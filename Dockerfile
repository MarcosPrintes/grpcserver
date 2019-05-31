# FROM golang:1.12.4 as builder
# # ADD . ${GOPATH}/src/
# # WORKDIR ${GOPATH}/src/
# # RUN mkdir -pgithub.com.br/MarcosPrintes && cdgithub.com.br/MarcosPrintes
# WORKDIR ${GOPATH}/src/github.com.br/MarcosPrintes/grpc

# COPY . .

# RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && dep ensure
# #&& glide install
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main

# # FROM iron/go
# # RUN ln -s /app/env/env .env
# RUN adduser -S -D -H -h /app appuser
# USER appuser
# WORKDIR /app
# ENTRYPOINT ["grpc"]
# # COPY --from=builder ${GOPATH}/src/github.com.br/MarcosPrintes/grpc/main .
# #ADD env/env /app/env/env
# # CMD ["./"]
# EXPOSE 8080

FROM golang

LABEL maintener="marcosprintes"

WORKDIR $GOPATH/src/github.com.br/MarcosPrintes/grpc 

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

ENTRYPOINT [ "grpc" ]