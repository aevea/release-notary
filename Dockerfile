FROM golang:1.17.8-alpine as builder
RUN mkdir /app
WORKDIR /app

RUN apk add --no-cache make git gcc musl-dev

# This step is done separately than `COPY . /app/` in order to
# cache dependencies.
COPY go.mod go.sum /app/
RUN go mod download

COPY . /app/
RUN CGO_ENABLED=0 go build -a -tags "osusergo netgo" --ldflags "-linkmode external -extldflags '-static'" -o build/release-notary .

FROM alpine:3.15.0
RUN  apk add --no-cache --virtual=.run-deps ca-certificates git &&\
    mkdir /app

RUN addgroup --gid 10001 --system nonroot \
    && adduser  --uid 10000 --system --ingroup nonroot --home /home/nonroot nonroot

RUN apk add --no-cache tini

WORKDIR /app
COPY --from=builder /app/build/release-notary ./release-notary

RUN ln -s $PWD/release-notary /usr/local/bin

USER nonroot

ENTRYPOINT ["/sbin/tini", "--", "release-notary" ]

CMD [ "publish" ]