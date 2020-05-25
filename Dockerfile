FROM gitpod/workspace-full
USER root
RUN apt-get install -y graphviz
RUN cd /root \
 && git clone https://github.com/google/nucleus
RUN sed -i -e "s/16/18/g" /root/nucleus/install.sh
RUN cd /root/nucleus \
 && ./install.sh
RUN chmod -R gitpod:gitpod /root/nucleus
#RUN cd /root \
# && git clone https://github.com/google/nucleus \
# && cd nucleus \
# && sed -i -e "s/16/18/g" ./install.sh \
# && ./install.sh \
# && chmod -R gitpod:gitpod /root/nucleus
USER gitpod
