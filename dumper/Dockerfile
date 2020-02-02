FROM golang:alpine AS BUILD
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /server
COPY . .
RUN go build -a -o server .

FROM scratch
ARG PORT=80
ENV PORT $PORT
COPY --from=BUILD /server/server /server
ENTRYPOINT ["/server"]