FROM gitpod/workspace-full
USER root
RUN apt-get -y install graphviz
RUN apt-get -y install python-numpy
USER gitpod
