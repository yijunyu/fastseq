FROM ubuntu:16.04
RUN apt-get -y update
RUN apt-get install -y sudo git
RUN apt -y install --no-install-recommends python2.7-minimal python2.7  
RUN apt-get install -y python-pip python-numpy
RUN pip install future
ENV PATH /root/bin:$PATH
ENV CLIF_FLAGS "--copt=-msse4.1 --copt=-msse4.2 --copt=-mavx --copt=-O3"
RUN git clone https://github.com/google/nucleus /opt/nucleus \
 && cd /opt/nucleus && ./install.sh --prereqs_only \
 && mv /opt /workspace
RUN cd /workspace/nucleus \
 && bazel build -c opt ${CLIF_FLAGS} nucleus/... \
 && bazel build :licenses_zip
CMD /bin/bash
