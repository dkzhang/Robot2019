FROM dkzhang007/golang_protobuf_grpc-arm32v7-arm32v7

WORKDIR /go/src

RUN git clone https://github.com/dkzhang/Robot2019.git #20191204-001

WORKDIR /go/src/Robot2019/applicationDriverForRobot/playAudio/server

CMD go run ./server.go

#docker run --privileged -p 50051:50051
