FROM ubuntu:16.04

RUN apt-get -y update
RUN apt-get install -y sudo git

# Note: changes to install-prereqs.sh will invalidate the Docker
# cache, which is the behavior we want.
# Note: This will put the tensorflow sources at /opt/tensorflow.
RUN git clone https://github.com/google/nucleus /opt/nucleus
RUN cd /opt/nucleus && ./install.sh --prereqs_only
RUN rm -rf /opt/nucleus

ENV PATH /root/bin:$PATH

RUN	curl -L -O https://github.com/bazelbuild/bazel/releases/download/0.26.1/bazel-0.26.1-installer-linux-x86_64.sh \
 &&	chmod +x bazel-*.sh \
 && ./bazel-0.26.1-installer-linux-x86_64.sh --prefix=/usr/local \
 && rm bazel-0.26.1-installer-linux-x86_64.sh

RUN pip install numpy
CMD /bin/bash
