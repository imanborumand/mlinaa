FROM golang:1.13.8

USER root

#define app env
ENV APP_NAME mlinaa
ENV PORT 8080

COPY . /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}
RUN go get .
RUN go build -o main .
RUN export GIN_MODE=release
EXPOSE 8080
# Run the executable
CMD ["./main"]
