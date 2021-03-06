from concurrent import futures
import logging
import os

import grpc

import tir_pb2
import tir_pb2_grpc

import numpy as np
from scipy import interpolate
import pylab as pl
import matplotlib as mpl
from mpl_toolkits.mplot3d import Axes3D

class ThermalImagingRenderingService(tir_pb2_grpc.ThermalImagingRenderingServiceServicer):

    def ThermalImagingRender(self, request, context):
        if not os.path.exists(request.filepath):
            os.makedirs(request.filepath)

        x = np.linspace(-1, 1, request.width)  # x
        y = np.linspace(-1, 1, request.height)  # y
        z = np.array(request.dataArray).reshape((request.height, request.width))

        # 三次样条二维插值
        newfunc = interpolate.interp2d(x, y, z, kind='cubic')

        # 原来是计算100*100的网格上的插值，先改为放大20倍的网格
        xnew = np.linspace(-1, 1, request.width * 10)  # x
        ynew = np.linspace(-1, 1, request.height * 10)  # y
        fnew = newfunc(xnew, ynew)

        # 绘图
        # 为了更明显地比较插值前后的区别，使用关键字参数interpolation='nearest'
        # 关闭imshow()内置的插值运算。
        # cm = pl.cm.get_cmap('Spectral_r')
        pl.subplot(311)
        im1 = pl.imshow(z, extent=[-1 * (request.width / request.height), 1 * (request.width / request.height), -1, 1],
                        cmap=mpl.cm.rainbow, interpolation='nearest', origin="lower")
        #               cmap = mpl.cm.hot,
        pl.colorbar(im1)

        pl.subplot(312)
        im2 = pl.imshow(fnew, extent=[-1 * (request.width / request.height), 1 * (request.width / request.height), -1, 1],
                        cmap=mpl.cm.rainbow, interpolation='nearest', origin="lower")
        pl.colorbar(im2)

        xn, yn = np.meshgrid(xnew, ynew)
        ax = pl.subplot(313,projection='3d')
        surf = ax.plot_surface(xn, yn, fnew, rstride=2, cstride=2, cmap=mpl.cm.rainbow,linewidth=0.5, antialiased=True)
        pl.colorbar(surf)

        pl.savefig(request.filepath + request.filename + '.png', dpi=256)
        pl.clf()

        print(request.dataArray)
        print(request.height)
        print(request.width)
        print(request.filepath)
        print(request.filename)

        return tir_pb2.ThermalImagingRenderingReply(errorMessage="")


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    tir_pb2_grpc.add_ThermalImagingRenderingServiceServicer_to_server(ThermalImagingRenderingService(), server)

    server.add_insecure_port('[::]:50061')
    print("Server is running:")
    server.start()

    # Notice the version of grpc, wait_for_termination func not exists in early version.
    # To update grpc use this command.
    # python3 -m pip install --upgrade grpcio
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
