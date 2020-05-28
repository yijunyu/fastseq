FROM ubuntu:16.04
RUN apt-get -y update
RUN apt-get install -y sudo git
RUN apt -y install --no-install-recommends python2.7-minimal python2.7  
RUN apt-get install -y python-pip python-numpy
RUN pip install future
ENV HOME /home/gitpod
ENV PATH ${HOME}/bin:$PATH
RUN git clone https://github.com/google/nucleus ${HOME}/nucleus \
 && cd ${HOME}/nucleus && ./install.sh --prereqs_only
ENV COPT_FLAGS "--copt=-msse4.1 --copt=-msse4.2 --copt=-mavx --copt=-O3"
RUN sed -i -e 's/"--py3output"/#"--py3output"/g' ${HOME}/nucleus/third_party/clif.bzl
RUN cd ${HOME}/nucleus \
 && bazel build -c opt ${COPT_FLAGS} nucleus/... \
 && bazel build :licenses_zip
CMD /bin/bash
