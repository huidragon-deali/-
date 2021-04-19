프로토콜버퍼 컴파일러 설치 (존재한다면 pass)
```
go install google.golang.org/protobuf/cmd/protoc-gen-go
```

goland 프로젝트 GOPATH 추가

프로토콜 버퍼파일 빌드
```
protoc -I proto \               // 컴파일 폴더
                grpc.proto \    // 컴파일 대상
                --go_out=. \    // message 컴파일 (모델클래스)
                --go-grpc_out=. // client/server 컴파일 (서비스클래스)

or

protoc --go-grpc_out=. \ 
       --go_out=.      \
        proto/*.proto  \  // 전체 빌드
```