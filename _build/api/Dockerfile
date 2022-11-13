FROM golang:1.17.1-alpine as build

#SET TIMEZONE
ENV TZ=Asia/Bangkok

#UPDATES THE REPOSITORY
RUN apk update && \
    apk add --no-cache bash && \
    apk add --no-cache vim && \
    apk add --no-cache nano && \
    apk add --no-cache tzdata && \
    apk add --no-cache git && \
    apk add --no-cache curl && \
    apk add --no-cache ca-certificates && \
    apk add --no-cache postgresql-client && \
    apk add --no-cache busybox-extras && \
    rm -rf /var/lib/apt/lists/*

#SET WORKING DIRECTORY
WORKDIR /go/src/resume

#COPY CODE INTO WORKSPACE
# COPY go.mod .
# COPY go.sum .
COPY .env .
COPY assets .
COPY views .

COPY . .

#BUILD APP
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -v -mod=readonly -o ./api .


FROM alpine:latest as Dev

#SET TIMEZONE
ENV TZ=Asia/Bangkok
#UPDATES THE REPOSITORY
RUN apk update && \
    apk add --no-cache bash && \
    apk add --no-cache vim && \
    apk add --no-cache nano && \
    apk add --no-cache tzdata && \
    apk add --no-cache git && \
    apk add --no-cache curl && \
    apk add --no-cache ca-certificates && \
    apk add --no-cache postgresql-client && \
    apk add --no-cache busybox-extras && \
    rm -rf /var/lib/apt/lists/*

# COPIED THE BINARY FILE FROM BUILD CONTAINER TO APP-SERVER
COPY --from=build /go/src/resume/api /resume/api
COPY --from=build /go/src/resume/.env /resume/.env
COPY --from=build /go/src/resume/assets /resume/assets
COPY --from=build /go/src/resume/views /resume/views

# SWITCHES WORKING DIRECTORY TO APP-SERVER
WORKDIR "/resume"

# EXPOSES THE 8080 PORT FROM THE CONTAINER
EXPOSE 8080

# RUN THE BINARY ONCE THE CONTAINER START
ENTRYPOINT ["/resume/api"]