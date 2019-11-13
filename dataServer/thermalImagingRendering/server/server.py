from concurrent import futures
import logging
import os

import grpc

import tir_pb2
import tir_pb2_grpc


class Greeter(helloworld_pb2_grpc.GreeterServicer):
    def ThermalImagingRender(self, request, context):
        if not os.path.exists(request.filepath):
            os.makedirs(request.filepath)

        print(string(request.dataArray))
        print(request.height)
        print(request.width)
        print(request.filepath)
        print(request.filename)
        
        return tir_pb2.ThermalImagingRenderingReply(errorMessage="")


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    tir_pb2_grpc.add_ThermalImagingRenderingServiceServicer_to_server(ThermalImagingRender(), server)

    server.add_insecure_port('[::]:50061')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
