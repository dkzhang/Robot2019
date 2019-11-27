FROM registry.cn-beijing.aliyuncs.com/dkzhang/python_grpc-arm32v7

RUN pip3 install -i https://pypi.tuna.tsinghua.edu.cn/simple/  adafruit-circuitpython-amg88xx

RUN pip3 install RPi.GPIO

RUN apt-get install i2c-tools

WORKDIR /

RUN git clone https://github.com/dkzhang/Robot2019.git

CMD python3 /Robot2019/applicationDriverForRobot/thermalImagingDataCollect/server/server.py