FROM golang:1.22 as build

RUN go install golang.org/x/tools/cmd/present@latest

FROM debian:bookworm-slim

COPY --from=build /go/bin/present /usr/local/bin/present

RUN groupadd -r present && useradd --no-log-init -r -g present present
USER present

WORKDIR /talks

ADD fosdem22 /talks/fosdem22
ADD gcil-2023 /talks/gcil-2023
ADD go-august-penguin /talks/go-august-penguin
ADD go-il-ifaces /talks/go-il-ifaces
ADD go-intro /talks/go-intro
ADD go-quiz /talks/go-quiz
ADD go-std-iface /talks/go-std-iface
ADD go-var /talks/go-var
ADD import-c-go-il /talks/import-c-go-il
ADD json-the-fine-print /talks/json-the-fine-print
ADD nrsc /talks/nrsc
ADD rang-func /talks/range-func



CMD ["present", "-http", ":8080", "-play=false"]
