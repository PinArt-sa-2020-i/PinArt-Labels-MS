FROM golang:alpine
RUN apk update && apk add --no-cache git
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN go get -d -v
RUN go build -o main .
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ["./main"]