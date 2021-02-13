# go_tutorial

목차
  - 설치 및 간단한 프로그램 작성
    - Workspace 폴더를 생성하는데, 3개의 서브 디렉토리 (bin, pkg, src)가 있다.
    - 해당 Workspace 폴더는 GOPATH 경로에 추가할수 있으며, 하나 이상의 경로를 지정할수 있다.
    - 중요한 환경 변수는 GOROOT, GOPATH가 있다.
      - GOROOT: Go가 설치된 디렉토리로 설치파일은 bin, 패키지들은 pkg에 저장된다
      - GOPATH: go 표준 패키지 외 3rd party 패키지, 사용자 정의 패키지들이 저장된다.
      - `go env` 를 통해 조회할 수 있다
      - GOPATH는 ~/.zprofile에서 기존 경로를 바꿔줄수 있다.
    - 외부에서 패키지 다운방법
      - `go get github.com/mattn/go-sqlite3` 과 같이 입력하면 GOPATH로 지정한 곳에 관련 패키지들을 다운받는다.
    