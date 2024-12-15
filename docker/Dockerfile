FROM golang:1.23.2


RUN apt update
RUN  apt install openssl  libssl-dev -y
RUN apt-get install -y wget
RUN wget -q https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
RUN apt-get install -y ./google-chrome-stable_current_amd64.deb

WORKDIR /app

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 && go build  -o /go-video -tags=release -ldflags "-s -w"

CMD [ "/go-video" ]