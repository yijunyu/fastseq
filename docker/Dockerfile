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
RUN apt -y install vim
RUN apt-get -y install ipython ipython-notebook
RUN pip3 install jupyter
RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
RUN apt-get -y install apt-transport-https ca-certificates gnupg
RUN curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -
RUN apt-get -y update && apt-get -y install google-cloud-sdk
RUN apt-get -y install wget
RUN pip3 install google-nucleus
RUN pip3 install tensorflow
EXPOSE 8080
WORKDIR /home/gitpod/fastseq/nucleus
CMD /bin/bash
