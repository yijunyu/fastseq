FROM gitpod/workspace-full
USER root
RUN apt-get update -y
RUN apt-get install curl
RUN curl -s https://packagecloud.io/install/repositories/github/git-lfs/script.deb.sh | bash
RUN apt-get update -y
RUN apt-get install git-lfs -y
RUN usermod -a -G root gitpod
USER gitpod
