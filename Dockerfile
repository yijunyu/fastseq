FROM gitpod/workspace-full
USER root
RUN apt-get update -y
RUN apt-get install -y graphviz
RUN apt-get install -y python3-dev python3-pip python3-tk python3-lxml python3-six python3-numpy python3-h5py
RUN usermod -a -G root gitpod
USER gitpod
