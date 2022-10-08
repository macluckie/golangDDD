FROM golang:1.17

WORKDIR /go/src/app
COPY . /go/src/app
RUN ls
RUN go get  github.com/swaggo/files
RUN go get  github.com/swaggo/gin-swagger
RUN go get  github.com/swaggo/swag/cmd/swag
RUN go get -u github.com/gin-gonic/gin
RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/sqlite
RUN go get github.com/hamba/avro
EXPOSE 8080
# RUN go build -o ./out/go-sample-app .
# CMD ["./out/go-sample-app"]
CMD go run cmd/main.go
