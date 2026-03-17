package cmd

import (
	"fmt"
	"math"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	MaxSingleUploadSize = 20 * 1024 * 1024 // 20MB
	ChunkSizeMB         = 10
	ChunkSizeBytes      = ChunkSizeMB * 1024 * 1024
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload [file_path]",
	Short: "Upload a file to FastClaw Storage (Auto-detects Large Files)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		fileInfo, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			fmt.Printf("Error: File '%s' not found.\n", filePath)
			return
		}

		fileSize := fileInfo.Size()
		if fileSize <= MaxSingleUploadSize {
			uploadSingle(filePath)
		} else {
			uploadMultipart(filePath, fileSize)
		}
	},
}

// 기존 단일 파일 업로드 로직
func uploadSingle(filePath string) {
	fmt.Printf("Uploading small file (<20MB): %s ...\n", filepath.Base(filePath))
	client := GetClient()
	resp, err := client.R().
		SetFile("file", filePath).
		SetResult(&UploadResponse{}).
		Post("/api/skill/storage/r2/upload")

	if err != nil || resp.IsError() {
		fmt.Printf("Error: Single upload failed: %v %s\n", err, resp.String())
		return
	}

	result := resp.Result().(*UploadResponse)
	printUploadSuccess(result.PublicUrl, result.FileKey, result.Size, result.Cost)
}

// 대용량 멀티파트 업로드 로직
func uploadMultipart(filePath string, fileSize int64) {
	filename := filepath.Base(filePath)
	chunkCount := int(math.Ceil(float64(fileSize) / float64(ChunkSizeBytes)))
	fmt.Printf("🚀 Large file detected (%.2f MB). Starting Multipart Upload (%d chunks)...\n", float64(fileSize)/1024/1024, chunkCount)

	client := GetClient()

	// 1. Init
	fmt.Println("1️⃣ Initializing upload (paying deposit)...")
	var initResult struct {
		UploadId    string  `json:"uploadId"`
		FileKey     string  `json:"fileKey"`
		DepositPaid float64 `json:"depositPaid"`
	}
	resp, err := client.R().
		SetBody(map[string]interface{}{
			"filename":       filename,
			"chunkCount":     chunkCount,
			"maxChunkSizeMB": ChunkSizeMB,
			"contentType":    "application/octet-stream",
		}).
		SetResult(&initResult).
		Post("/api/skill/storage/r2/multipart/init")

	if err != nil || resp.IsError() {
		fmt.Printf("❌ Init failed: %v %s\n", err, resp.String())
		return
	}
	fmt.Printf("✔️ Init complete. Deposit Paid: %f Credits\n", initResult.DepositPaid)

	// 2. Get Presigned URLs
	fmt.Println("2️⃣ Requesting Presigned URLs...")
	partNumbers := make([]int, chunkCount)
	for i := 0; i < chunkCount; i++ {
		partNumbers[i] = i + 1
	}

	var urlsResult struct {
		Urls []struct {
			PartNumber int    `json:"partNumber"`
			Url        string `json:"url"`
		} `json:"urls"`
	}
	resp, err = client.R().
		SetBody(map[string]interface{}{
			"uploadId":    initResult.UploadId,
			"fileKey":     initResult.FileKey,
			"partNumbers": partNumbers,
		}).
		SetResult(&urlsResult).
		Post("/api/skill/storage/r2/multipart/urls")

	if err != nil || resp.IsError() {
		fmt.Printf("❌ Failed to get URLs: %v %s\n", err, resp.String())
		return
	}

	// 3. Upload Chunks (Direct to R2)
	fmt.Println("3️⃣ Uploading chunks directly to R2...")
	file, _ := os.Open(filePath)
	defer file.Close()

	completedParts := make([]map[string]interface{}, 0)
	for _, item := range urlsResult.Urls {
		fmt.Printf("   - [Chunk %d/%d] Sending...", item.PartNumber, chunkCount)
		
		buffer := make([]byte, ChunkSizeBytes)
		file.Seek(int64((item.PartNumber-1)*ChunkSizeBytes), 0)
		n, _ := file.Read(buffer)
		
		// PUT request to R2
		putResp, putErr := client.R().
			SetBody(buffer[:n]).
			Put(item.Url)

		if putErr != nil || putResp.IsError() {
			fmt.Printf(" ❌ Failed: %v\n", putErr)
			return
		}

		etag := putResp.Header().Get("ETag")
		completedParts = append(completedParts, map[string]interface{}{
			"PartNumber": item.PartNumber,
			"ETag":       etag,
		})
		fmt.Println(" ✔️")
	}

	// 4. Complete
	fmt.Println("4️⃣ Finalizing upload...")
	var finalResult struct {
		PublicUrl  string  `json:"publicUrl"`
		TotalCost  float64 `json:"totalCost"`
		ActualSize int64   `json:"actualSize"`
	}
	resp, err = client.R().
		SetBody(map[string]interface{}{
			"uploadId":    initResult.UploadId,
			"fileKey":     initResult.FileKey,
			"parts":       completedParts,
			"depositPaid": initResult.DepositPaid,
		}).
		SetResult(&finalResult).
		Post("/api/skill/storage/r2/multipart/complete")

	if err != nil || resp.IsError() {
		fmt.Printf("❌ Finalize failed: %v %s\n", err, resp.String())
		return
	}

	printUploadSuccess(finalResult.PublicUrl, initResult.FileKey, finalResult.ActualSize, finalResult.TotalCost)
}

func printUploadSuccess(url, key string, size int64, cost float64) {
	fmt.Println("-------------------------------------------")
	fmt.Printf("✨ Success! Public URL: %s\n", url)
	fmt.Printf("File Key: %s\n", key)
	fmt.Printf("Size: %.2f MB\n", float64(size)/1024/1024)
	fmt.Printf("Cost: %f Credits\n", cost)
	fmt.Println("-------------------------------------------")
}

// UploadResponse represents the response from upload API
type UploadResponse struct {
	PublicUrl string  `json:"publicUrl"`
	FileKey   string  `json:"fileKey"`
	Size      int64   `json:"size"`
	Cost      float64 `json:"cost"`
}

func init() {
	RootCmd.AddCommand(uploadCmd)
}
