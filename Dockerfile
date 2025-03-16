FROM golang:1.4-alpine
LABEL authors="ulr9661"

ENTRYPOINT ["top", "-b"]