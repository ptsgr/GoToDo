FROM golang:1.15-buster AS build

ENV GOPATH=/

WORKDIR /src/
COPY ./ /src/

RUN go mod download
RUN CGO_ENABLED=0 go build -o /bin/todo ./cmd/main.go


FROM scratch
COPY --from=build /bin/todo /bin/todo

COPY --from=build /src/configs/config.yaml /configs/config.yaml
ENTRYPOINT ["/bin/todo"]