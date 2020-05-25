command -v protoc >/dev/null 2>&1 || {
	if [ -f xaa ]; then
		cat xa* | tar xvfj - -C /workspace/protobuf
	fi
	echo export PATH='"/workspace/protobuf/bin:$PATH"' >> .bashrc
	source .bashrc
}
