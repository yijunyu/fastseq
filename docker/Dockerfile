FROM ubuntu:16.04
RUN apt-get -y update
RUN apt-get install -y sudo git
RUN apt -y install --no-install-recommends python2.7-minimal python2.7  
RUN apt-get install -y python-pip python-numpy
RUN pip install future
ENV PATH /root/bin:$PATH
RUN git clone https://github.com/google/nucleus /home/gitpod/nucleus \
 && cd /home/gitpod/nucleus && ./install.sh --prereqs_only
ENV COPT_FLAGS "--copt=-msse4.1 --copt=-msse4.2 --copt=-mavx --copt=-O3"
#RUN cd /home/gitpod/nucleus \
# && bazel build -c opt ${COPT_FLAGS} nucleus/... \
# && bazel build :licenses_zip
CMD /bin/bash
