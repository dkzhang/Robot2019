FROM dkzhang007/golang_protobuf_grpc-x86:my1.0

WORKDIR /go/src

RUN git clone https://github.com/dkzhang/Robot2019.git #20191203-003

WORKDIR /go/src/Robot2019/chassisDriverForRobot/robotMultiplePointsMove/server

CMD go run ./server.go

