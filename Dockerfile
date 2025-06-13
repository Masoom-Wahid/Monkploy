# FROM node:20.17-slim AS front-build

# ENV VITE_BASE_API_URL=
# ENV VITE_BASE_WEBSOCKET_API_URL=

# WORKDIR /app/views/

# # Force rebuild by copying a dummy file with a timestamp
# # This ensures `npm install` and `npm run build` are always re-executed
# ARG CACHE_BREAKER=1
# COPY ./views/.npmrc ./
# COPY ./views/package.json ./
# RUN echo $CACHE_BREAKER && npm install --frozen-lockfile

# COPY ./views .
# RUN echo $CACHE_BREAKER && npm run build

FROM golang:1.19 AS backend-build

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

# COPY --from=front-build /app/views/dist ./views/dist

COPY . .
ENV GOFLAGS=-buildvcs=false

CMD ["air"]