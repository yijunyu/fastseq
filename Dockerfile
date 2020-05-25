FROM gitpod/workspace-full
USER root
RUN apt-get update -y
# RUN apt-get install -y subversion
# RUN apt-get install -y llvm clang
# RUN apt-get install -y ocaml opam
# RUN apt-get install -y python3-dev python3-pip python3-tk python3-lxml python3-six python3-numpy python3-h5py
RUN apt-get install -y graphviz
# RUN apt-get install -y texlive
RUN usermod -a -G root gitpod
USER gitpod
