FROM golang:1.19 as dev
ENV HELIXF_ROOT=/go/src/app
COPY docker/conf/air.toml $HELIXF_ROOT/air.toml
COPY src/server/go.mod src/server/go.sum $HELIXF_ROOT/src/server/
COPY lib/go.mod lib/go.sum $HELIXF_ROOT/lib/
WORKDIR $HELIXF_ROOT/src/server
RUN go mod download
RUN go install github.com/cosmtrek/air@latest
CMD ["air", "-c", "../../air.toml"]
# CMD ["sleep", "7200"]
