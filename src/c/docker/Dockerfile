FROM yijun/nucleus
RUN apt -y install antlr4
RUN curl -O https://www.antlr.org/download/antlr-4.8-complete.jar
RUN mv antlr-4.8-complete.jar /usr/share/java/antlr4.jar
RUN rm /usr/share/java/antlr4-runtime.jar
RUN git clone https://github.com/google/flatbuffers.git
RUN apt -y install cmake
RUN cd flatbuffers \
 && cmake -G "Unix Makefiles" \
 && make \
 && make install
RUN apt-get -y install software-properties-common
RUN add-apt-repository ppa:longsleep/golang-backports
RUN apt-get update -y
RUN apt-get -y install golang
RUN go get github.com/antlr/antlr4/runtime/Go/antlr
RUN go get github.com/google/flatbuffers/go
RUN go get gopkg.in/yaml.v2
RUN go get github.com/bachvtuan/mime2extension
RUN apt-get -y install file
