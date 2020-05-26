export CLIF_FLAGS="--copt=-msse4.1 --copt=-msse4.2 --copt=-mavx --copt=-O3"
cd /workspace \
 && git clone https://github.com/tensorflow/tensorflow \
 && cd tensorflow \
 && git checkout v2.0.0 \
 && echo | ./configure \
 && cd .. \
 && git clone https://github.com/google/nucleus \
 && cd nucleus \
 && bazel build -c opt ${CLIF_FLAGS} nucleus/... \
 && bazel build :licenses_zip
