FROM gitpod/workspace-full
RUN git clone https://github.com/google/nucleus \
 && cd nucleus \
 && sed -i -e "s/16/18/g" ./install.sh \
 && ./install.sh
