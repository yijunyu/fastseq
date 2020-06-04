FROM yijun/nucleus:py3
RUN apt-get -y update
RUN apt-get -y install vim
RUN apt-get -y install ipython
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
