FROM gitpod/workspace-full
USER root
RUN cd /root \
 && git clone https://github.com/google/nucleus \
 && cd nucleus \
 && sed -i -e "s/16/18/g" ./install.sh \
 && ./install.sh \
 && chmod -R gitpod:gitpod /root/nucleus
USER gitpod
