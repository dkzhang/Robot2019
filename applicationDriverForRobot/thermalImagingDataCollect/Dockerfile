FROM registry.cn-beijing.aliyuncs.com/dkzhang/python_grpc-arm32v7

RUN pip3 install -i https://pypi.tuna.tsinghua.edu.cn/simple/  adafruit-circuitpython-amg88xx

RUN pip3 install RPi.GPIO

WORKDIR /

RUN git clone https://github.com/dkzhang/Robot2019.git  #20191029-001

CMD python3 /Robot2019/applicationDriverForRobot/thermalImagingDataCollect/server/server.py

#docker run --privileged -p 50061:50061