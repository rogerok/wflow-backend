# FROM golang:alpine
#
# RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base
#
# RUN mkdir /app
# WORKDIR /app
#
# COPY . .
# COPY .env .
#
# # RUN go get -d -v ./...
# # RUN go install -v ./...
# #
# # RUN go build -o /build
#
# RUN go mod download
#
# EXPOSE 5000
#
# CMD ["air"]