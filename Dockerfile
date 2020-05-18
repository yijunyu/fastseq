FROM gitpod/workspace-full
USER root
RUN apt-get update -y
RUN apt-get install -y git-core subversion git-svn
RUN usermod -a -G root gitpod
USER gitpod
