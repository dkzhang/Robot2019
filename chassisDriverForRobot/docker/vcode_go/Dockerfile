FROM dkzhang007/golang_protobuf_grpc-x86:my1.0

RUN go get github.com/gomodule/redigo/redis

WORKDIR /go/src/Robot2019/chassisDriverForRobot/

ENV theApp "./docker/vcode_go/main.go"
CMD go run $theApp