#FROM golang:1.13.8
FROM golang:1.13.8-alpine

WORKDIR /m4d/
COPY go.mod go.mod
COPY pkg/connectors/protobuf pkg/connectors/protobuf

WORKDIR /m4d/pkg/policy-compiler/
COPY pkg/policy-compiler/policy-compiler policy-compiler
COPY pkg/policy-compiler/pc-bl pc-bl

WORKDIR /m4d/test/services/pilot/policy-manager-mock
COPY test/services/pilot/policy-manager-mock .

RUN go get -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o policy-manager-mock mockup_pilot.go

EXPOSE 50070

CMD [ "./policy-manager-mock"]
