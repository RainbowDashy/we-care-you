FROM node:slim as frontend

RUN npm install -g pnpm

WORKDIR /frontend

COPY ./frontend/package.json ./frontend/pnpm-lock.yaml ./
RUN pnpm install

COPY ./frontend .
RUN pnpm build


FROM golang:1.17 as backend

WORKDIR /backend
COPY go.mod go.sum ./
# Downloading go dependencies is too slow in China. So set a proxy.
ENV GOPROXY="https://goproxy.cn/"
RUN go mod download

COPY ./store ./store
COPY  *.go ./
RUN go build -a -ldflags '-linkmode external -extldflags "-static"' -o we-care-you .

FROM scratch
WORKDIR /we-care-you
COPY --from=frontend /frontend/dist ./frontend/dist
COPY --from=backend /backend/we-care-you .

ENV GIN_MODE=release
ENV PORT=8080
ENTRYPOINT ["./we-care-you"]
