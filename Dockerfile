FROM golang:1.19 AS backend-build

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
ENV GOFLAGS=-buildvcs=false

# Create a symbolic link from .env.dev to .env
RUN ln -sf .env.dev .env

CMD ["air"]
