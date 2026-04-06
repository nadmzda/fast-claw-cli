# FastClaw CLI - Gemini Context

이 파일은 FastClaw CLI 프로젝트의 구조, 기술 스택, 개발 컨벤션 및 주요 기능을 설명하여 AI 에이전트가 프로젝트를 이해하고 보조하는 데 도움을 주기 위한 지침서입니다.

## 🚀 프로젝트 개요

FastClaw CLI는 AI 기능을 터미널에서 즉시 실행하고, 클라우드 스토리지 및 다양한 워크스페이스(Google, Notion 등)와 소셜 미디어를 관리할 수 있는 **AI 기반 다목적 CLI 도구**입니다.

- **언어:** Go (1.24.11 이상 권장)
- **주요 라이브러리:**
  - `spf13/cobra`: CLI 커맨드 및 플래그 관리
  - `go-resty/resty/v2`: HTTP 클라이언트 요청 처리
  - `creativeprojects/go-selfupdate`: 자가 업데이트 기능
- **백엔드 API:** `https://fast-claw.xyz`
- **인증 방식:** API Key (X-API-KEY 헤더 사용)

## 📂 프로젝트 구조

```text
D:\work\web\fast-claw\fast-claw-cli\
├── main.go             # 프로그램 진입점
├── cmd/                # Cobra 커맨드 정의 (핵심 로직)
│   ├── root.go         # 루트 커맨드 및 공통 플래그/설정
│   ├── config.go       # API 키 설정 관리 (~/.fastclaw_config)
│   ├── upload.go       # 파일 업로드 (20MB 초과 시 자동 멀티파트)
│   ├── google.go       # Google Workspace (Calendar, Drive, Sheets, Tasks)
│   ├── vision.go       # NVIDIA Vision 분석 및 OCR
│   └── ...             # 기타 기능별 커맨드 (image, music, search, social 등)
├── go.mod / go.sum     # 의존성 관리
└── .goreleaser.yaml    # 배포 및 릴리즈 설정
```

## 🛠 빌드 및 실행

### 로컬 빌드
```powershell
go build -o fastclaw.exe main.go
```

### 실행 및 테스트
```powershell
# API 키 설정
./fastclaw config set-key YOUR_API_KEY

# 도움말 확인
./fastclaw --help
```

### 배포 (GoReleaser)
```powershell
git tag vX.Y.Z
goreleaser release --clean
```

## 💡 개발 컨벤션 및 원칙

1.  **Cobra 패턴 준수:** 모든 커맨드는 `cmd/` 디렉토리 내 개별 파일로 관리하며, `init()` 함수를 통해 `RootCmd`에 등록합니다.
2.  **공통 실행 로직:** 대부분의 기능은 `cmd/root.go`의 `ExecuteToolAction` 함수를 사용하여 `/api/skill/tool/execute` 엔드포인트를 호출합니다.
3.  **에러 처리:** 사용자에게 명확한 에러 메시지를 출력하고, 필요한 경우 `--verbose` 플래그를 통해 상세 정보를 제공합니다.
4.  **API 키 우선순위:** Flag (`-k`) > Env (`FASTCLAW_API_KEY`) > Config File (`~/.fastclaw_config`) 순으로 적용됩니다.
5.  **보안:** API 키는 로그나 출력 시 `maskKey` 함수를 사용하여 일부 마스킹 처리합니다.

## 🎯 주요 기능 및 명령어

| 구분 | 명령어 | 설명 |
| :--- | :--- | :--- |
| **핵심** | `upload` | 단일/멀티파트 업로드 (R2 스토리지) |
| **AI** | `vision`, `image`, `music`, `riverflow` | 이미지 분석, 생성, 음악 생성 |
| **검색** | `search`, `scrape` | 특화 검색 (지도, 뉴스) 및 웹 크롤링 |
| **생산성** | `google`, `gmail`, `notion` | 워크스페이스 연동 및 관리 |
| **소셜** | `social` | Instagram 및 Reddit 포스팅 |
| **관리** | `config`, `update` | 설정 관리 및 셀프 업데이트 |

## 📝 향후 작업 및 주의사항
- **인증 통합:** 현재 `google_tokens` 테이블을 폐기하고 `ai_users`의 `apikey`로 인증 체계를 통합하는 작업이 예정되어 있습니다.
- **Suno 연동:** 음악 생성 기능은 `headless: false` 옵션이 필요할 수 있습니다.
- **Composio:** 내부적으로 Composio를 사용하지만, 사용자에게는 'Tool' 또는 Kong API 고유 기능으로 브랜딩합니다.
