# FROM gitpod/workspace-full
# USER root
FROM ubuntu:16.04
RUN apt-get -y update
RUN apt-get -y install pkg-config zip g++ zlib1g-dev unzip curl git lsb-release
RUN apt-get -y install libssl-dev libcurl4-openssl-dev liblz-dev libbz2-dev liblzma-dev
RUN apt-get -y install python3-dev python3-pip python3-wheel python3-setuptools python3-numpy python3-h5py
RUN curl "https://storage.googleapis.com/deepvariant/packages/oss_clif/oss_clif.ubuntu-16.latest.tgz" | tar xzf -
RUN mkdir -p /workspace \
 && cd /workspace \
 && git clone https://github.com/tensorflow/tensorflow \
 && cd tensorflow \
 && git checkout v2.0.0 \
 && echo | ./configure
RUN cd /workspace \
 && curl -L -O https://github.com/bazelbuild/bazel/releases/download/0.26.1/bazel-0.26.1-installer-linux-x86_64.sh \
 && chmod +x bazel-*.sh \
 && ./bazel-0.26.1-installer-linux-x86_64.sh --prefix=/usr/local \
 && rm bazel-0.26.1-installer-linux-x86_64.sh
RUN pip3 install numpy
RUN apt-get -y install graphviz
# USER gitpod