FROM golang:1.24.1-bullseye

RUN apt update -y \
    && apt install -y protobuf-compiler

WORKDIR /iterx

COPY ../../Makefile Makefile

RUN make init
