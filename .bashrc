mkdir -p /workspace/bin
cat > /workspace/bin/push <<EOF
git commit -am "$@*" && git push
EOF
chmod +x /workspace/bin/push
export PATH="/workspace/bin:$PATH"
export PATH="/workspace/nucleus/usr/local/clif/bin:$PATH"
export LD_LIBRARY_PATH="/workspace/nucleus/usr/local/lib"
