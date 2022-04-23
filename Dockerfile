FROM golang:1.18

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN go get github.com/spf13/cobra@latest && \
    go install github.com/golang/mock/mockgen@v1.5.0

RUN apt-get update && apt-get install sqlite3 -y

RUN usermod -u 1000 www-data

USER www-data

CMD ["tail", "-f", "/dev/null"]