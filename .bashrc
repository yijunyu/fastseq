export PATH="/home/gitpod/.pyenv/versions/3.8.2/bin:/usr/lib/llvm-11/bin:/workspace/clang/bin:/workspace/bin:$PATH"
export LD_LIBRARY_PATH=/home/gitpod/.pyenv/versions/3.8.2/lib:$(llvm-config-11 --libdir)
export PKG_CONFIG_PATH=/home/gitpod/.pyenv/versions/3.8.2/lib/pkgconfig
mkdir -p /workspace/bin
cat > /workspace/bin/push <<EOF
git commit -am "$@*" && git push
EOF
chmod +x /workspace/bin/push
pyenv global 3.8.2
export PATH="/workspace/flex/bin:$PATH"
export PATH="/workspace/graphviz/bin:$PATH"
