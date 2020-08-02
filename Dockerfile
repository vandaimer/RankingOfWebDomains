FROM golang:alpine

RUN mkdir /app 

COPY . /app/

WORKDIR /app 

RUN go build

ENTRYPOINT ["./interview", "./customers.csv"]
