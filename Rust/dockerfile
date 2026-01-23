FROM docker.io/library/rust:1.92.0-alpine3.20 AS builder

ADD . .

RUN cargo build --release

FROM docker.io/library/alpine:3.20 AS final

COPY --from=builder /target/release/terminal-vpet /terminal-vpet

CMD ["./terminal-vpet"]
