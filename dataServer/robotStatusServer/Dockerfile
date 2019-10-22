FROM dkzhang007/golang_protobuf_grpc-x86:my1.0

WORKDIR $GOPATH/src/
RUN git clone https://github.com/dkzhang/Robot2019.git

RUN protoc -I Robot2019/dataServer/robotStatusServer/grpc/ Robot2019/dataServer/robotStatusServer/grpc/rs.proto --go_out=plugins=grpc:Robot2019/dataServer/robotStatusServer/grpc/

RUN go build Robot2019/dataServer/robotStatusServer/server.go

CMD ./server