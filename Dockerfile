FROM golang:1.10.1-alpine3.7
copy main.go /go/src/

ENV LogsPerSecond 1
ENV TotalLogs 5
ENV MyProgramName myprog
CMD go run /go/src/main.go
