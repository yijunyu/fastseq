FROM gitpod/workspace-full
USER root
RUN apt-get -y update
RUN apt-get install -y python3-pip python3-numpy
ENV HOME /home/gitpod
ENV PATH ${HOME}/bin:$PATH
RUN git clone https://github.com/google/nucleus ${HOME}/nucleus \
 && cd ${HOME}/nucleus && ./install.sh --prereqs_only
ENV COPT_FLAGS "--copt=-msse4.1 --copt=-msse4.2 --copt=-mavx --copt=-O3"
#RUN cd ${HOME}/nucleus \
# && bazel build -c opt ${COPT_FLAGS} nucleus/... \
# && bazel build :licenses_zip
CMD /bin/bash
