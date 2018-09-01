FROM golang:1.10
RUN mkdir -p /opt/gotest/src/gosam
RUN mkdir -p /opt/gotest/bin
RUN mkdir -p /opt/gotest/pkg
COPY ./ /opt/gotest/src/gosam
WORKDIR /opt/gotest/src/gosam
RUN export GOPATH=/opt/gotest
RUN go get github.com/gocarina/gocsv
RUN CGO_ENABLED=0 go build -v -o "dist"
FROM alpine:3.6
COPY --from=0 /opt/gotest/src/gosam/dist /usr/bin/gosam
RUN mkdir -p /opt/csvFile
COPY --from=0 /opt/gotest/src/gosam/csvFile /opt/csvFile
WORKDIR /opt
CMD ["gosam", "--help"]