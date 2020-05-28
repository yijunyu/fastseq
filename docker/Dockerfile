#FROM ubuntu:16.04
#RUN apt-get -y update
#RUN apt-get -y install git
#RUN apt-get -y install --no-install-recommends python3.5-minimal python3.5  
#ENV HOME /home/gitpod
#ENV PATH ${HOME}/bin:$PATH
#RUN mkdir -p ${HOME}/fastseq
#RUN apt-get -y install sudo
#RUN git clone https://github.com/google/nucleus ${HOME}/fastseq/nucleus
#RUN apt-get -y install python3-numpy
#RUN cd ${HOME}/fastseq/nucleus && ./install.sh --prereqs_only
#ENV COPT_FLAGS "--copt=-msse4.1 --copt=-msse4.2 --copt=-mavx --copt=-O3"
#RUN cd ${HOME}/fastseq/nucleus \
# && bazel build :licenses_zip
#RUN pip3 install future
#RUN update-alternatives --remove python /usr/bin/python2
#RUN update-alternatives --install /usr/bin/python python /usr/bin/python3 10
#RUN cd ${HOME}/fastseq/nucleus \
# && bazel build -c opt ${COPT_FLAGS} nucleus/...
FROM yijun/nucleus
CMD /bin/bash
