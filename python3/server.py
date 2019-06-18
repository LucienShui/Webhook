from http.server import HTTPServer, BaseHTTPRequestHandler
import os

class RequestHandler(BaseHTTPRequestHandler):
    '''Process & Accept requests'''

    # Process an GET requests
    def do_GET(self):
        os.system("bash -x main.sh &")
        self.send_response(200)
        self.send_header("Content-Type", "text/html")
        self.send_header("Content-Length", str(len(self.Page)))
        self.end_headers()
        self.wfile.write(self.Page)


if __name__ == '__main__':
    serverAddress = ('', 10086)
    server = HTTPServer(serverAddress, RequestHandler)
    server.serve_forever()
