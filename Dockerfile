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
RUN cd /root \
 && git clone https://github.com/tensorflow/tensorflow \
 && cd tensorflow \
 && git checkout v${NUCLEUS_TENSORFLOW_VERSION} \
 && echo | ./configure
ENV NUCLEUS_BAZEL_VERSION "0.26.1"
RUN cd /root \
 && curl -L -O https://github.com/bazelbuild/bazel/releases/download/"${NUCLEUS_BAZEL_VERSION}"/bazel-"${NUCLEUS_BAZEL_VERSION}"-installer-linux-x86_64.sh \
 && chmod +x bazel-*.sh \
 && ./bazel-"${NUCLEUS_BAZEL_VERSION}"-installer-linux-x86_64.sh \
 && rm bazel-"${NUCLEUS_BAZEL_VERSION}"-installer-linux-x86_64.sh
RUN pip3 install numpy
RUN cd /root \
 && git clone https://github.com/google/nucleus \
 && cd nucleus \
 && bazel build -c opt "--copt=-msse4.1 --copt=-msse4.2 --copt=-mavx --copt=-O3" nucleus/...
RUN cd /root/nucleus \
 && RUN bazel build :licenses_zip
RUN chmod -R gitpod:gitpod /root/nucleus
RUN apt-get -y install graphviz
USER gitpod
