FROM dkzhang007/python_grpc-arm32v7

RUN pip3 install adafruit-circuitpython-amg88xx

WORKDIR /

RUN git clone https://github.com/dkzhang/Robot2019.git

CMD python3 /Robot2019/applicationDriverForRobot/thermalImagingDataCollect/server/server.py