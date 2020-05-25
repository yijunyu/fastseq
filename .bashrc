mkdir -p /workspace/bin
cat > /workspace/bin/push <<EOF
git commit -am "$@*" && git push
EOF
chmod +x /workspace/bin/push
pyenv global 3.8.2
export PATH="/workspace/bin:$PATH"
export PATH="/workspace/flex/bin:$PATH"
export PATH="/workspace/graphviz/bin:$PATH"
export PATH="/workspace/protobuf/bin:$PATH"
export PATH="/workspace/clang/bin:$PATH"
export PATH="/workspace/.pipmodules/bin:$PATH"
export PATH="/workspace/protobuf/bin:$PATH"
export PATH="/workspace/fastseq/nucleus-0.5.1/usr/local/clif/bin:$PATH"
export LD_LIBRARY_PATH="/workspace/fastseq/nucleus-0.5.1/usr/local/lib"
export PATH="/workspace/graphviz/bin:$PATH"
export PATH="/workspace/protobuf/bin:$PATH"
export PATH="/workspace/clang/bin:$PATH"
export PATH="/workspace/.pipmodules/bin:$PATH"
export PATH="/workspace/fastseq/nucleus-0.5.1/usr/local/clif/bin:$PATH"
export LD_LIBRARY_PATH="/workspace/fastseq/nucleus-0.5.1/usr/local/lib"
export PATH="/workspace/graphviz/bin:$PATH"
export PATH="/workspace/protobuf/bin:$PATH"
export PATH="/workspace/clang/bin:$PATH"
export PATH="/workspace/fastseq/nucleus-0.5.1/usr/local/clif/bin:$PATH"
export LD_LIBRARY_PATH="/workspace/fastseq/nucleus-0.5.1/usr/local/lib"
export PATH="/workspace/clang/bin:$PATH"
export PATH="/workspace/fastseq/nucleus-0.5.1/usr/local/clif/bin:$PATH"
export LD_LIBRARY_PATH="/workspace/fastseq/nucleus-0.5.1/usr/local/lib"
export PATH="/workspace/clang/bin:$PATH"
export PATH="/workspace/fastseq/nucleus-0.5.1/usr/local/clif/bin:$PATH"
export LD_LIBRARY_PATH="/workspace/fastseq/nucleus-0.5.1/usr/local/lib"
