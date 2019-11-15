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


class ThermalImagingDataCollectService(tir_pb2_grpc.ThermalImagingRenderingServiceServicer):

    def CollectThermalImagingData(self, request, context):

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
