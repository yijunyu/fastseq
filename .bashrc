export PATH="$PATH:/workspace/fastseq/bin:/home/gitpod/.pyenv/versions/3.8.2/bin:/usr/lib/llvm-11/bin"
export LD_LIBRARY_PATH=/workspace/fastseq/lib:$(llvm-config-11 --libdir)
ln -sf /workspace/fastseq/bin ~/bin
ln -sf /workspace/fastseq/lib ~/lib
ln -sf /workspace/fastseq/include ~/include
