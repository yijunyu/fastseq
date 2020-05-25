mkdir -p /workspace/bin
cat > /workspace/bin/push <<EOF
git commit -am "$@*" && git push
EOF
chmod +x /workspace/bin/push
pyenv global 3.8.2
export PATH="/workspace/bin:$PATH"
export PATH="/workspace/clang/bin:$PATH"
export PATH="/workspace/flex/bin:$PATH"
export PATH="/workspace/graphviz/bin:$PATH"
export PATH="/workspace/clang/bin:$PATH"
