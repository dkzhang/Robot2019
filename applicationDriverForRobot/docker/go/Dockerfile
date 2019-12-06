FROM dkzhang007/golang_protobuf_grpc-arm32v7-arm32v7

RUN go get github.com/stianeikeland/go-rpio

WORKDIR /go/src
RUN git clone https://github.com/dkzhang/Robot2019.git #20191206-001

WORKDIR /go/src/Robot2019/applicationDriverForRobot/

ENV theApp "./docker/go/main.go"

CMD go run $theApp