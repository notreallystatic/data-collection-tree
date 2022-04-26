FROM golang:latest

RUN mkdir /dct
WORKDIR /dct
ADD . /dct

RUN go mod download

CMD ["go", "run", "main.go"]

