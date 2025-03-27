FROM golang:1.23.2 AS builder

# Instala OpenCV e dependências
RUN apt-get update && apt-get install -y libopencv-dev pkg-config && rm -rf /var/lib/apt/lists/*

# Define os caminhos para CGO
ENV CGO_ENABLED=1
ENV CGO_CFLAGS="-I/usr/include/opencv4"
ENV CGO_LDFLAGS="-L/usr/lib/x86_64-linux-gnu -lopencv_core -lopencv_imgproc -lopencv_highgui -lopencv_videoio"
ENV LD_LIBRARY_PATH="/usr/lib/x86_64-linux-gnu:/usr/local/lib"

WORKDIR /app
COPY . .

# Cria links simbólicos para as bibliotecas necessárias
RUN ln -s /usr/lib/x86_64-linux-gnu/libopencv_videoio.so.4.5 /usr/lib/x86_64-linux-gnu/libopencv_videoio.so.4.5d
RUN ln -s /usr/lib/x86_64-linux-gnu/libopencv_imgproc.so.4.5 /usr/lib/x86_64-linux-gnu/libopencv_imgproc.so.4.5d
RUN ln -s /usr/lib/x86_64-linux-gnu/libopencv_core.so.4.5 /usr/lib/x86_64-linux-gnu/libopencv_core.so.4.5d

# Compila o projeto
RUN go build -o ./go-video -tags=release

# Imagem final mais leve
FROM debian:bullseye-slim
COPY --from=builder /usr/lib/x86_64-linux-gnu /usr/lib/x86_64-linux-gnu
COPY --from=builder /app/go-video /go-video
COPY --from=builder /app/lib /lib

# Define o caminho das libs
ENV LD_LIBRARY_PATH="/usr/lib/x86_64-linux-gnu:/usr/local/lib"

CMD ["/go-video"]