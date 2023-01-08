FROM golang as builder

WORKDIR /sekura

COPY go.mod go.sum ./

RUN go mod download

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/agent-fake ./main.go


FROM scratch as agent-fake

WORKDIR /sekura

COPY --from=builder /sekura/build/agent-fake /sekura/agent-fake

ENTRYPOINT [ "./agent-fake" ]
CMD [ "-h" ]
