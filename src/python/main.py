import grpc
from proto.sample_pb2 import pinPonRequest
from proto.sample_pb2_grpc import PinPonServiceStub

import sys
import pprint
import time

def run():
    # go server の起動を待つ
    time.sleep(10)
    c = 'protoc-go:1234'
    try:
        with grpc.insecure_channel(c) as channel:
            stub = PinPonServiceStub(channel)
            response = stub.send(pinPonRequest(word='ping'))
            print("PinPon client received: " + response.word)
    except grpc.RpcError as e:
        print(f"An RPC error occurred: {e.code()} {e.details()}")
    except Exception as e:
        print(f"An error occurred: {e}")
    finally:
        pprint.pprint(sys.path)

if __name__ == '__main__':
    run()
