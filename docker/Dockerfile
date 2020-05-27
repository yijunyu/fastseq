FROM ubuntu:16.04

RUN apt-get -y update
RUN apt-get install -y sudo git

# Note: changes to install-prereqs.sh will invalidate the Docker
# cache, which is the behavior we want.
# Note: This will put the tensorflow sources at /opt/tensorflow.
RUN git clone https://github.com/google/nucleus /workspace/nucleus
RUN cd /workspace/nucleus && ./install.sh --prereqs_only

ENV PATH /root/bin:$PATH

RUN	curl -L -O https://github.com/bazelbuild/bazel/releases/download/0.26.1/bazel-0.26.1-installer-linux-x86_64.sh \
 &&	chmod +x bazel-*.sh \
 && ./bazel-0.26.1-installer-linux-x86_64.sh --prefix=/usr/local \
 && rm bazel-0.26.1-installer-linux-x86_64.sh

RUN pip3 install numpy
RUN git clone https://github.com/tensorflow/tensorflow /workspace/tensorflow \
 && cd /workspace/tensorflow \
 && git checkout v2.0.0 \
 && echo | ./configure
RUN pip install future
ENV CLIF_FLAGS "--copt=-msse4.1 --copt=-msse4.2 --copt=-mavx --copt=-O3"
RUN cd /workspace/nucleus \
 && bazel build -c opt ${CLIF_FLAGS} nucleus/... \
 && bazel build :licenses_zip
CMD /bin/bash

