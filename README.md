# Drop
Put/Remove your links and access them from multiple devices as they're updated in real-time.

## Quicker demo
[https://farnasirim.ir/drop](https://farnasirim.ir/drop)

## Quick deployment

Create a docker bridge network to contain the following three containers:
```
docker network create drop
```

Create a redis server with the name `redis`:
```
docker run -v $(pwd)/data:/data -d --net drop -p 127.0.0.1:6379:6379 --name redis redis --appendonly yes
```

Run the envoy to proxy incoming http requests to the grpc server:
```
docker run --net drop -v /path/to/cert-chain.pem:/etc/fullchain.pem -v /path/to/privkey.pem:/etc/privkey.pem -d -p 0.0.0.0:18082:8080 --name envoy-sec quay.io/farnasirim/drop-envoy-sec:v0.0.1
```
If you're not willing to use tls (!) use the container that is named `quay.io/farnasirim/drop-envoy:v0.0.1` and remove the `-v` directives mounting the credentials.

Run drop grpc server:
```
docker run -d --name drop-server --net drop -p 127.0.0.1:12345:12345 quay.io/farnasirim/drop-server:v0.1.2
```

The frontend is still not served by the drop server, so you need to manually build a `bundle.js` from `frontend/client.js` and somehow serve it alongside `frontend/index.html`.

## Development
To create any of the above artifacts please refer to the `Makefile`.

## License
MIT
