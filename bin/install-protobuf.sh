command -v protoc >/dev/null 2>&1 || {
	if [ -f bin/xaa ]; then
		mkdir -p /workspace/protobuf
		cat bin/xa* | tar xvfj - -C /workspace/protobuf
	fi
	echo export PATH='"/workspace/protobuf/bin:$PATH"' >> .bashrc
	source .bashrc
}
