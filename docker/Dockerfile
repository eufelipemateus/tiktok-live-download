FROM golang:1.22


RUN apt update
RUN  apt install openssl  libssl-dev -y

WORKDIR /app

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 && go build  -o /calculadora-back-go -tags=release -ldflags "-s -w"

CMD [ "/go-video" ]