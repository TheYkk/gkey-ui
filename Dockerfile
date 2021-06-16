FROM golang:1.16-alpine as build
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY main.go ./

RUN GOOS=js GOARCH=wasm go build -o test.wasm main.go

FROM theykk/docker-nginx-static:v1.19.8
COPY --from=build /app/test.wasm /static/test.wasm
COPY wasm.conf /etc/nginx/conf.d/wasm.conf
COPY index.html /static/index.html
COPY wasm_exec.js /static/wasm_exec.js