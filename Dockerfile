FROM gitpod/workspace-full
USER root
RUN apt-get -y update
RUN apt-get -y install pkg-config zip g++ zlib1g-dev unzip curl git lsb-release
RUN apt-get -y install libssl-dev libcurl4-openssl-dev liblz-dev libbz2-dev liblzma-dev
RUN apt-get -y install python3-dev python3-pip python3-wheel python3-setuptools
RUN curl "https://storage.googleapis.com/deepvariant/packages/oss_clif/oss_clif.ubuntu-18.latest.tgz" > /tmp/oss_clif.tgz \
 && cd / \
 && tar xzf /tmp/oss_clif.tgz \
 && rm -f /tmp/oss_clif.tgz
ENV NUCLEUS_TENSORFLOW_VERSION "2.0.0"
RUN git clone https://github.com/tensorflow/tensorflow \
 && cd tensorflow \
 && git checkout v${NUCLEUS_TENSORFLOW_VERSION} \
 && echo | ./configure
RUN git clone https://github.com/google/nucleus
RUN cd /root/nucleus \
 && bazel build -c opt "--copt=-msse4.1 --copt=-msse4.2 --copt=-mavx --copt=-O3" nucleus/...
RUN cd /root/nucleus \
 && RUN bazel build :licenses_zip
RUN chmod -R gitpod:gitpod /root/nucleus
RUN apt-get -y install graphviz
USER gitpod
