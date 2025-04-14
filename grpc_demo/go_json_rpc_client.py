import requests

request = {
    "id":0,
    "params":["boby"],
    "method": "HelloService.Hello"
}

rsp = requests.post("http://localhost:1234/jsonrpc", json=request)
print(rsp.text)