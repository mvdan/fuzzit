FROM golang:1.12-stretch

LABEL maintainer="Fuzzit.dev, inc."

RUN apt-get -qqy update && apt-get install -y wget gnupg2

RUN echo "deb http://apt.llvm.org/stretch/ llvm-toolchain-stretch main" >> /etc/apt/sources.list
RUN echo "deb-src http://apt.llvm.org/stretch/ llvm-toolchain-stretch main" >> /etc/apt/sources.list
RUN wget -O - https://apt.llvm.org/llvm-snapshot.gpg.key| apt-key add -
RUN apt-get update && apt-get install -y libllvm-9-ocaml-dev libllvm9 llvm-9 llvm-9-dev llvm-9-doc llvm-9-examples llvm-9-runtime clang-9 lldb-9 lld-9
RUN ln -s /usr/lib/llvm-9/bin/llvm-symbolizer /bin/llvm-symbolizer

WORKDIR /app
