FROM dkzhang007/golang_protobuf_grpc-x86:my1.0

RUN go get github.com/gomodule/redigo/redis

WORKDIR /go/src

RUN protoc -I Robot2019/dataServer/robotStatusServer/grpc/ Robot2019/dataServer/robotStatusServer/grpc/rs.proto --go_out=plugins=grpc:Robot2019/dataServer/robotStatusServer/grpc/

WORKDIR /go/src/Robot2019/dataServer/

EXPOSE 50071

ENV theApp "./docker/vcode_go/main.go"
CMD go run $theApp


###################################################################