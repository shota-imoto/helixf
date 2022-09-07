FROM golang:1.19 as dev
ENV HELIXF_ROOT=/go/src/app
COPY docker/conf/air.toml $HELIXF_ROOT/air.toml
COPY src/server $HELIXF_ROOT/src/server
COPY lib $HELIXF_ROOT/lib
WORKDIR $HELIXF_ROOT/src/server
RUN go mod download
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", "../../air.toml"]
# CMD ["go", "run", "main.go"]
# CMD ["sleep", "7200"]
