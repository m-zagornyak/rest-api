FROM golang:1.15-alpine3.12 AS builder

COPY . /github.com/m-zagornyak/rest-api
WORKDIR /github.com/m-zagornyak/rest-api

RUN go mod download
RUN go build -o /bin/main cmd/app/main.go

FROM apline:latest

WORKDIR /root/

COPY --from=0 /github.com/m-zagornyak/rest-api/bin/app .
COPY --from=0 /github.com/m-zagornyak/rest-api/bin/configs configs/

EXPOSE 80

CMD ["./main"]
