FROM golang:1.13.0-buster

# install kubectl
RUN apt-get update \
    && apt-get install -y wget \
    && rm -rf /var/lib/apt/lists/* \
    && wget https://storage.googleapis.com/kubernetes-release/release/v1.13.0/bin/linux/amd64/kubectl \
    && mv kubectl /usr/local/bin/kubectl \
    && chmod +x /usr/local/bin/kubectl

WORKDIR /simframe
COPY . .
RUN go build -o ./build

CMD [ "bash" ]