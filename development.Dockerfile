FROM golang:latest

RUN mkdir /dct
WORKDIR /dct
ADD . /dct

RUN go mod download
RUN go install github.com/githubnemo/CompileDaemon@latest

ENTRYPOINT CompileDaemon --build="go build main.go" --command="./main"