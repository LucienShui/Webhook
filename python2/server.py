#-*- coding:utf-8 -*-
import BaseHTTPServer
import os

class RequestHandler(BaseHTTPServer.BaseHTTPRequestHandler):
    '''Process & Accept requests'''

    # Process an GET requests
    def do_GET(self):
        os.system("bash -x main.sh")
        self.send_response(200)


if __name__ == '__main__':
    serverAddress = ('', 10086)
    server = BaseHTTPServer.HTTPServer(serverAddress, RequestHandler)
    server.serve_forever()

