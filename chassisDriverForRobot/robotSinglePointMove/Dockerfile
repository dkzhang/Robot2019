FROM dkzhang007/golang_protobuf_grpc-x86:my1.0

WORKDIR /go/src

RUN git clone https://github.com/dkzhang/Robot2019.git #20191204-001

WORKDIR /go/src/Robot2019/chassisDriverForRobot/robotSinglePointMove

CMD go run ./serverMain.go

#50071
