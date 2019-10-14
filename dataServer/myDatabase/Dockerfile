FROM dkzhang007/my-grpc-golang:my1.0

WORKDIR /go/src
RUN git clone https://github.com/dkzhang/Robot2019.git
WORKDIR /go/src/Robot2019/dataServer/myDatabase

CMD go test -v -run "TestInitializeDB"
