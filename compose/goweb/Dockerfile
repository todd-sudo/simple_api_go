FROM golang:1.17.5-alpine

WORKDIR /go/src/app
COPY . .

RUN apk add build-base &&\
    go get -d -v ./... &&\
    go install -v ./...

RUN addgroup --system todd &&\
    adduser --system --ingroup todd todd &&\
    chown todd:todd /go/src/app

USER todd

CMD ["web"]
