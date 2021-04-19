

프로토콜 버퍼 컴파일러 설치
```
go install google.golang.org/protobuf/cmd/protoc-gen-go
```

프로토콜 버퍼파일 빌드
```
protoc -I proto grpc.proto \
                --go_out=. \    // message 컴파일 (모델클래스)
                --go-grpc_out=. // client/server 컴파일 (서비스클래스)
```

