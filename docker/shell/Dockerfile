FROM golang:1.19 as dev
ENV HELIXF_ROOT=/go/src/app
ENV HELIXF_ENV=test
RUN apt update
RUN apt install -y default-mysql-client-core/stable
WORKDIR /go/src/app
CMD ["bash"]
