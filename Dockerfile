FROM golang:latest AS build

WORKDIR /random_work_dir

# first download dependencies
COPY go.mod /random_work_dir
COPY go.sum /random_work_dir
RUN go mod download

# then copy source code
COPY / /random_work_dir


RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /mailbear ./cmd/mailbear


FROM alpine:latest

WORKDIR /

# Copy executable
COPY --from=build /mailbear /bin/mailbear
RUN chmod +x /bin/mailbear

# Copy config
RUN mkdir /mailbear/
COPY config_sample.yml /mailbear/config.yml

WORKDIR /mailbear



EXPOSE 1234
EXPOSE 9090
VOLUME ["/mailbear"]

CMD ["/bin/mailbear"]