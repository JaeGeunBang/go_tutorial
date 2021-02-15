# go_tutorial

### 목차
  - 설치 및 간단한 프로그램 작성
    - go를 설치하면 Workspace 폴더를 생성하는데, 3개의 서브 디렉토리 (bin, pkg, src)가 있다.
    - 해당 Workspace 폴더는 GOPATH로 추가할수 있다. (default $GOPATH: $HOME/go)
    - 중요한 환경 변수는 GOROOT, GOPATH가 있다.
      - GOROOT: go가 설치된 디렉토리로 실행파일은 bin, 패키지들은 pkg에 저장된다.
      - GOPATH: go 표준 패키지 외 3rd party 패키지, 사용자 정의 패키지들이 저장된다.
      - `go env` 를 통해 조회할 수 있다.
    - 외부에서 패키지 다운방법
      - `go get github.com/mattn/go-sqlite3` 과 같이 입력하면 GOPATH로 지정한 곳에 관련 패키지들을 다운받는다.
      - 내가 작업할 공간도 $GOPATH/src에서 하면되며, Github에 업로드 해두면 언제든 패키지를 다운받을수 있다.
        - 프로젝트 초기 생성 `go mod init github.com/.../...`
        - 이후 작업하면 됨
  - 패키지
    - 패키지는 코드의 모듈화, 코드의 재사용 기능을 제공한다.
    - "main" 이라 명명된 패키지는 Go Compiler에 의해 특별히 인식되며, executable 프로그램으로 만들어준다.
      - 즉, main 함수안엔 `package main`을 기재해준다.
    - 패키지를 import를 하면 기본 GOROOT, GOPATH 횐경변수를 검색한다.
      - `GOROOT/pkg`, `GOPATH/pkg` 또는 `../src` 경로를 찾는다.
    - `go install`을 수행하면 /pkg, 또는 /bin 폴더에 생성된다.
      - 패키지 하려는게 main 함수 라면, /bin 폴더에 `실행파일`이 생성된다. 
      - 패키지 하려는게 main 함수가 아니라면, /pkg 파일에 `.a 파일`이 생성된다.