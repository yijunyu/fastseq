cd /workspace \
 && git clone https://github.com/google/nucleus
export CLIF_FLAGS="--copt=-msse4.1 --copt=-msse4.2 --copt=-mavx --copt=-O3"
cd /workspace/nucleus \
  && bazel build -c opt ${CLIF_FLAGS} nucleus/...
cd /workspace/nucleus \
  && bazel build :licenses_zip
