# FastClaw CLI 🚀

FastClaw CLI는 AI 기능을 터미널에서 즉시 실행하고, 클라우드 스토리지 및 다양한 워크스페이스(Google, Notion 등)를 관리할 수 있는 **AI 기반 다목적 CLI 도구**입니다.

Go 언어로 작성되어 빠르고 가벼우며, 단일 실행 파일로 어디서든 편리하게 사용할 수 있습니다.

## ✨ 주요 기능

- 📁 **스마트 파일 업로드**: 20MB 이하는 단일 업로드, 대용량 파일은 자동으로 분할(Multipart) 업로드 수행.
- 👁️ **NVIDIA Vision 분석**: 이미지 URL 분석은 물론, 로컬 이미지를 자동으로 업로드하여 OCR 및 내용 분석 수행.
- 🎨 **AI 이미지 생성**: Gemini 3.1 모델 및 초고속 Riverflow 모델을 통한 이미지 생성.
- 🎵 **AI 음악 생성**: Suno AI를 활용한 맞춤 스타일, 가사, 보컬 음악 생성.
- 🔍 **특화 검색**: 구글 뉴스, 이미지, 지도(장소 정보) 등 유형별 맞춤 결과 출력.
- 🏢 **워크스페이스 통합**: Gmail, Calendar, Drive, Sheets, Tasks, Notion 완벽 연동.
- 📱 **소셜 자동화**: Instagram 비즈니스 포스팅 및 Reddit 게시글 관리.

## 🛠 설치 방법

### 1. 소스 코드 빌드 (Go 설치 필요)
```powershell
git clone https://github.com/your-username/fast-claw-cli.git
cd fast-claw-cli
go build -o fastclaw.exe main.go
```

### 2. 환경 변수 설정 (선택 사항)
매번 API 키를 입력하지 않으려면 환경 변수를 설정하세요.
```powershell
$env:FASTCLAW_API_KEY="your_api_key_here"
```

## 🔑 인증 및 연결

- **API Key**: [FastClaw 대시보드](https://fast-claw.xyz)에서 발급받은 키가 필요합니다.
- **서비스 연동**: Google이나 소셜 계정 연동은 다음 매직 링크를 통해 최초 1회 인증이 필요합니다.
  - Gmail: [인증하기](https://fast-claw.xyz/api/skill/tool/auth-link?appName=gmail)
  - Google Drive: [인증하기](https://fast-claw.xyz/api/skill/tool/auth-link?appName=drive)
  - Notion: [인증하기](https://fast-claw.xyz/api/skill/tool/auth-link?appName=notion)

## 📖 사용 가이드

### 1. 파일 업로드 & 이미지 분석
```powershell
# 대용량 파일 업로드 (자동 멀티파트 처리)
./fastclaw upload ./my_large_video.mp4

# 로컬 이미지 분석 (자동 업로드 후 분석)
./fastclaw vision ./receipt.jpg --prompt "이 영수증의 합계 금액을 알려줘"
```

### 2. AI 이미지 생성
```powershell
# 고품질 이미지 생성 (Gemini)
./fastclaw image "A futuristic Seoul city with flying cars" --ratio 16:9

# 초고속 미리보기 이미지 생성 (Riverflow)
./fastclaw riverflow "Cute robot cat"
```

### 3. AI 음악 생성 (Suno)
```powershell
# 음악 생성 (자동 가사 생성)
./fastclaw music "여름 바다로 떠나는 신나는 여행" --style "K-Pop, Dance, Upbeat" --title "Summer Wave" --email user@example.com --vocal Female

# Suno API 상태 확인
./fastclaw music health
```

### 4. 구글 검색 (유형별 특화)
```powershell
# 일반 검색
./fastclaw search "Fast-Claw API"

# 지도 검색 (장소 정보 및 평점 출력)
./fastclaw search maps "Best coffee shops in Gangnam" --num 3

# 뉴스 검색
./fastclaw search news "latest AI trends 2026"
```

### 5. 워크스페이스 & 소셜
```powershell
# 구글 캘린더 일정 추가 (자연어 지원)
./fastclaw google calendar add "Tomorrow at 2pm meeting with team"

# Gmail 전송
./fastclaw gmail send --to "example@gmail.com" --subject "Hello" --body "Sent from CLI"

# Reddit 포스팅
./fastclaw social reddit post "test" "Hello World" "This is a post from FastClaw CLI"
```

## 📜 커맨드 목록

| 명령어 | 설명 |
| :--- | :--- |
| `upload` | 파일 업로드 (20MB 초과 시 자동 멀티파트) |
| `vision` | 이미지 내용 분석 및 OCR |
| `search` | 구글 검색 (통합, 뉴스, 이미지, 지도) |
| `gmail` | 메일 목록 조회 및 발송 |
| `google` | 캘린더, 드라이브, 시트, 태스크 관리 |
| `image` | Gemini 기반 고품질 이미지 생성 |
| `riverflow` | 초고속 이미지 생성 모델 |
| `music` | Suno AI 기반 음악 생성 |
| `scrape` | 웹 페이지 텍스트 데이터 추출 |
| `notion` | Notion 페이지 생성 및 검색 |
| `social` | Instagram 및 Reddit 연동 |

## 🏗 빌드 & 릴리즈

이 프로젝트는 [GoReleaser](https://goreleaser.com)를 사용하여 자동 크로스플랫폼 빌드 및 GitHub 릴리즈를 수행합니다.

### 로컬 빌드
```powershell
go build -o fastclaw.exe main.go
```

### 릴리즈 절차
1. `main` 브랜치에 변경사항 커밋
2. 버전 태그 생성 (예: `v1.1.0`)
3. GoReleaser 실행하여 빌드 및 배포

```powershell
git add .
git commit -m "feat: 설명"
git tag v1.1.0
goreleaser release --clean
git push origin main --tags
```

GoReleaser가 **linux**, **darwin**, **windows** × **amd64**, **arm64** 바이너리를 빌드하여 [GitHub Releases](https://github.com/nadmzda/fast-claw-cli/releases)에 배포합니다.

> 배포를 위해 `GITHUB_TOKEN` 환경 변수가 설정되어 있어야 합니다.

## 🤝 기여하기
이슈 제보나 기능 제안은 언제든지 환영합니다! PR을 통해 프로젝트 발전에 기여해 주세요.

---
© 2026 Fast-Claw CLI. Built with Go and AI.
