FROM golang:1.19 as builder

COPY . /teacup
WORKDIR /teacup/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN go get github.com/AdaLogics/go-fuzz-headers
RUN apt update && apt install -y clang

RUN cd fuzz_greatape_activitypub && go-fuzz-build -libfuzzer -o fuzz_greatape_activitypub.a && \
    clang -fsanitize=fuzzer fuzz_greatape_activitypub.a -o fuzz_greatape_activitypub.libfuzzer

RUN cd fuzz_greatape_validators && go-fuzz-build -libfuzzer -o fuzz_greatape_validators.a && \
    clang -fsanitize=fuzzer fuzz_greatape_validators.a -o fuzz_greatape_validators.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /teacup/mayhem/fuzz_greatape_activitypub/fuzz_greatape_activitypub.libfuzzer /
COPY --from=builder /teacup/mayhem/fuzz_greatape_validators/fuzz_greatape_validators.libfuzzer /