FROM golang:1.17-buster as build

WORKDIR /app
ADD . /app

ARG APP_NAME
ENV APP_NAME ${APP_NAME}

RUN go get -d -v ./...
RUN mkdir build
RUN CGO_ENABLED=0 GOOS=linux go build -o build/${APP_NAME} cmd/${APP_NAME}/main.go

# ---------------------------------------------
FROM gcr.io/distroless/base-debian10

ARG APP_NAME
ENV APP_NAME ${APP_NAME}

COPY --from=build /app/build/${APP_NAME} /
