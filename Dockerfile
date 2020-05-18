FROM gitpod/workspace-full
USER root
RUN apt-get update -y
RUN apt-get install -y subversion
RUN usermod -a -G root gitpod
USER gitpod
