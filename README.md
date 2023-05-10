# server
Simple HTTP server for logging all GET and POST requests


```bash
docker run -d --name fake-server -p 10000:8080 futuretea/server
```

```bash
curl http://127.0.0.1:10000
```