from concurrent import futures
import logging
import os

import grpc

import tidc_pb2
import tidc_pb2_grpc

class ThermalImagingDataCollectService(tidc_pb2_grpc.ThermalImagingDataCollectServiceServicer):

    def CollectThermalImagingData(self, request, context):

        return tidc_pb2.ThermalImagingDataCollectReply(errorMessage="")


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    tidc_pb2_grpc.add_ThermalImagingDataCollectServiceServicer_to_server(ThermalImagingDataCollectService(), server)

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
