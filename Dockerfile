FROM golang:1.23.5 as build
WORKDIR /app
COPY . .
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun ./cmd/server/main.go 

FROM scratch
WORKDIR /app
COPY --from=build /app/cloudrun .
ENTRYPOINT ["./cloudrun"]