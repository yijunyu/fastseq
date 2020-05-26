# bin/install-bazel.sh
# source .bashrc
cd /workspace \
 && git clone https://github.com/tensorflow/tensorflow \
 && cd tensorflow \
 && git checkout v2.0.0 \
 && echo | ./configure
cd /workspace \
 && git clone https://github.com/google/nucleus
export CLIF_FLAGS="--copt=-msse4.1 --copt=-msse4.2 --copt=-mavx --copt=-O3"
cd /workspace/nucleus \
  && bazel build -c opt ${CLIF_FLAGS} nucleus/...
cd /workspace/nucleus \
  && bazel build :licenses_zip
