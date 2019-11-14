FROM dkzhang007/grpc-python:latest

RUN python -m pip install numpy
RUN python -m pip install scipy
RUN python -m pip install matplotlib

WORKDIR /myApp

RUN git clone https://github.com/dkzhang/Robot2019.git #20191114-004

CMD python /myApp/Robot2019/dataServer/thermalImagingRendering/server/server.py
