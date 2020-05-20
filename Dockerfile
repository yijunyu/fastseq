FROM gitpod/workspace-full
USER root
RUN apt-get update -y
RUN apt-get install -y subversion
RUN apt-get install -y llvm clang
RUN apt-get install -y python3-dev python3-pip python3-tk python3-lxml python3-six
RUN apt-get install -y ocaml opam
RUN usermod -a -G root gitpod
USER gitpod
