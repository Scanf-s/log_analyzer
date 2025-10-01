# Day 1-2: Go 기초 & 로그 분석 CLI 과제 명세서

## 학습 목표
- Go 언어의 기본 문법, 타입 시스템, 모듈 관리 방식 이해
- 표준 라이브러리를 활용한 파일 입출력 및 에러 처리 구현
- CLI 플래그를 통한 사용자 입력 처리 및 JSON 기반 결과 출력 설계
- 테이블 기반 단위 테스트를 작성하여 핵심 로직 검증

## 사전 준비 사항
- Go 1.21 이상 설치 및 환경 변수 설정 (`go version` 확인)
- 기본 터미널 사용 능력 및 Git 사용 가능
- 샘플 로그 파일 확보 (없을 경우 Day 1에서 직접 생성)

## 프로젝트 개요
- **프로젝트명:** `log-analyzer`
- **프로젝트 구조:**
  - `cmd/log-analyzer/main.go` : CLI 엔트리 포인트
  - `internal/cli` : 플래그 파싱 및 실행 흐름 제어
  - `internal/analyzer` : 로그 통계 계산 로직
  - `testdata/` : 단위 테스트용 샘플 로그 파일 모음
- **기능 요약:**
  1. 지정된 로그 파일을 읽어 총 라인 수를 계산
  2. 로그 레벨/타임스탬프 등 주요 필드를 집계하여 JSON으로 출력
  3. 에러 상황(파일 없음, 잘못된 형식 등)에 대한 사용자 친화적 메시지 제공

## Day 1: Go 기초 다지기 & CLI 스켈레톤 구축
### 학습 포커스
- 패키지, 함수, 구조체, 포인터 등 핵심 문법 복습
- 표준 `flag` 패키지를 이용한 CLI 플래그 정의
- 파일 읽기(`os`, `bufio`), 에러 처리 패턴(`if err != nil`) 연습

### 진행 순서
1. **Go 모듈 초기화**
   - `mkdir -p cmd/log-analyzer internal/analyzer internal/cli testdata`
   - `go mod init <모듈 경로>` 실행 후 `.gitignore`에 `bin/` 등 추가
2. **엔트리 포인트 작성** (`cmd/log-analyzer/main.go`)
   - `--log-path` 플래그(기본값: `sample.log`) 정의
   - `internal/cli` 패키지로 위임하여 실행 흐름 제어
3. **CLI 계층 구현** (`internal/cli/runner.go` 권장)
   - 플래그 검증(파일 존재 여부)과 에러 메시지 출력 담당
   - `internal/analyzer` 호출 후 결과를 표준 출력에 전달
4. **로그 파일 스캔 로직 작성** (`internal/analyzer/readlogfiles.go`)
   - 버퍼 리더(`bufio.Scanner`)로 줄 단위로 이루어진 {}형태의 스트링을 한 줄  읽는 로직 구현
   - 빈 파일/대형 파일 처리 및 에러 전파
   - 반환 타입은 `LogEntry` 구조체(총 라인 수 필드 포함)
5. **간단한 로그 파일 분석 로직 구현** (`internal/analyzer/analyze.go`)
   - 총 라인 수 계산
   - 에러 발생 시 적절히 래핑하여 상위로 전달
6. **테스트 베이스라인 구축** (`internal/analyzer/linecount_test.go`)
   - 테이블 기반 테스트로 정상/빈 파일/없는 파일 시나리오 검증
   - `testdata/`에 최소 2개 이상의 샘플 로그 파일 저장
7. **샘플 로그 및 실행 확인**
   - `testdata/sample.log` 작성 (INFO/ERROR 혼합)
   - `go run ./cmd/log-analyzer --log-path testdata/sample.log`

### 완료 기준(Checklist)
- [x] `go test ./...`가 라인 카운트 로직에 대해 성공적으로 통과한다
- [x] CLI 실행 시 총 라인 수가 사람이 계산한 값과 일치한다
- [x] 잘못된 파일 경로 입력 시 사용자 친화적 에러 메시지를 출력한다

## Day 2: 통계 확장 & JSON 출력 고도화
### 학습 포커스
- 구조체, 맵, 슬라이스로 복합 데이터 모델링
- `encoding/json`을 활용한 직렬화/역직렬화
- 시간(`time`), 문자열 처리(`strings`) 등 표준 패키지 활용
- 더 정교한 단위 테스트 및 에러 핸들링 심화

### 진행 순서
1. **로그 파싱 모델 정의**
   - 로그 한 줄을 표현하는 `LogEntry` 구조체 설계 (예: 레벨, 타임스탬프, 메시지)
   - 레벨은 `string`, 타임스탬프는 `time.Time` 변환을 시도하고 실패 시 처리 전략 수립
2. **통계 로직 추가** (`internal/analyzer/stats.go`)
   - 레벨별 카운트, 시간대별 카운트(시간 또는 분 단위) 집계
   - `AnalyzerResult`에 `Totals`, `ByLevel`, `ByHour` 등 필드 확장
3. **JSON 출력 포맷 설계**
   - `internal/cli`에서 결과를 JSON Pretty Print로 출력 (`json.MarshalIndent`)
   - `--output` 플래그를 통해 `stdout` 혹은 파일 저장 옵션 제공 (기본: `stdout`)
4. **추가 플래그 지원**
   - `--top` (기본 5): 가장 많이 등장한 메시지/패턴 출력 (필요 시 문자열 정규화)
   - `--level` (다중 선택 가능): 특정 레벨만 필터링 후 분석
   - 플래그 충돌이나 잘못된 입력에 대한 검증 로직 강화
5. **테이블 기반 테스트 확대**
   - 정상 케이스 외에 잘못된 타임스탬프, 없는 레벨 등 에러 케이스 포함
   - `internal/cli`는 `testing`과 `bytes.Buffer`를 활용해 출력 검증
6. **도구 활용 & 문서화**
   - `go fmt ./...`, `go vet ./...`, `staticcheck`(가능하다면) 실행
   - README에 주요 명령어 및 사용 예시 추가

### 완료 기준(Checklist)
- [ ] JSON 출력 샘플이 요구한 필드를 모두 포함하고 포맷 규칙을 따른다
- [ ] 플래그 조합(`--level=ERROR --top=3`) 실행 시 예상대로 필터링된다
- [ ] 실패 케이스에 대한 단위 테스트가 존재하며 통과한다
- [ ] `go vet` 및 선택적 정적 분석 도구 실행 시 경고가 없다

## 확장 과제 (선택)
- 여러 로그 파일을 동시에 분석하는 `--inputs`(쉼표 구분) 플래그 추가
- goroutine과 channel을 활용하여 파일별 통계를 병렬로 수집하고 병합
- CLI 출력 결과를 파일로 저장할 때 JSON Schema를 명세하여 호환성 확보

## 제출물
- `README.md`에 실행 방법, 플래그 설명, 예시 출력, 테스트 명령어 정리
- 소스 코드 및 테스트 전체(`cmd/`, `internal/`, `testdata/`) 커밋
- 주요 설계 결정에 대한 간단한 회고(예: 에러 처리, 데이터 구조 선정) 문서화

## 학습 체크리스트
- [ ] 포인터와 값 타입의 차이를 코드 예제로 설명할 수 있다
- [ ] 사용자 정의 타입과 인터페이스를 활용해 파서를 확장할 계획을 세웠다
- [ ] 에러 래핑(`fmt.Errorf("...: %w", err)`)과 `errors.Is` 활용법을 이해했다
- [ ] `table-driven test` 패턴을 직접 작성하고 유지보수에 유리함을 설명할 수 있다

## 참고 자료
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Tour](https://go.dev/tour/welcome/1)
- [Go Blog: Table-Driven Tests](https://go.dev/blog/table-driven-tests)
- [Go CLI Applications with Cobra vs flag](https://go.dev/blog/crosscompile) *(flag 패키지를 우선 사용하되 확장 방향 참고)*
- [JSON and Go](https://go.dev/blog/json)
