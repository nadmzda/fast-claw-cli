# FastClaw CLI Auto Installer for Windows
$repo = "nadmzda/fast-claw-cli"
$ErrorActionPreference = "Stop"

Write-Host "[INFO] Fetching latest version info from GitHub..." -ForegroundColor Cyan
$latestRelease = Invoke-RestMethod -Uri "https://api.github.com/repos/$repo/releases/latest"
$version = $latestRelease.tag_name
$versionNum = $version.TrimStart("v")

$arch = if ($env:PROCESSOR_ARCHITECTURE -eq "ARM64") { "arm64" } else { "amd64" }
$assetName = "fast-claw-cli_$( $versionNum )_windows_$arch.zip"
$asset = $latestRelease.assets | Where-Object { $_.name -eq $assetName }
$url = $asset.browser_download_url

$destDir = "$HOME\bin\fastclaw"
if (!(Test-Path $destDir)) { 
    New-Item -ItemType Directory -Path $destDir -Force | Out-Null 
}

Write-Host "[INFO] Downloading FastClaw CLI $version ($arch)..." -ForegroundColor Cyan
Invoke-WebRequest -Uri $url -OutFile "$destDir\fastclaw.zip"

Write-Host "[INFO] Extracting files..." -ForegroundColor Cyan
Expand-Archive -Path "$destDir\fastclaw.zip" -DestinationPath $destDir -Force
Remove-Item "$destDir\fastclaw.zip"

$exePath = "$destDir\fastclaw.exe"
if (Test-Path $exePath) {
    $currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
    if ($currentPath -notlike "*$destDir*") {
        [Environment]::SetEnvironmentVariable("Path", "$currentPath;$destDir", "User")
        $env:Path += ";$destDir"
        Write-Host "[OK] Added $destDir to User PATH." -ForegroundColor Green
    }
    Write-Host "[OK] Successfully installed FastClaw CLI!" -ForegroundColor Green
    Write-Host "[INFO] Please restart your terminal or run: `$env:Path = [System.Environment]::GetEnvironmentVariable('Path','User')`" -ForegroundColor Yellow
} else {
    Write-Error "[FAILED] Installation failed: fastclaw.exe not found."
}
