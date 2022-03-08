FROM golang:1.17-alpine3.14 AS builder

RUN apk add --update gcc musl-dev
RUN mkdir -p /myapp
ADD . /myapp
WORKDIR /myapp

RUN adduser -u 10001 -D myapp

RUN go get -u github.com/swaggo/swag/cmd/swag  &&  go generate . && GOOS=linux GOARCH=amd64 CGO_ENABLED=1  go build -ldflags='-extldflags=-static'  -o myapp .


FROM scratch 

COPY --from=builder /etc/passwd /etc/passwd
USER myapp

WORKDIR /myapp

#COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/
COPY --from=builder /myapp/docs ./docs
COPY --from=builder /myapp/myapp ./myapp
COPY --from=builder /myapp/tests/postman ./tests/postman
VOLUME ./test/postman
EXPOSE 5007
CMD ["./myapp"]

