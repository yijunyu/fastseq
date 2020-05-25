command -v /workspace/clang/bin/clif-matcher >/dev/null 2>&1 || {
	if [ -f clif.tar.bz2 ]; then
		tar xvfj /workspace/fastseq/bin/clif.tar.bz2 -C /workspace 
	fi
	echo export PATH='"/workspace/clang/bin:$PATH"' >> .bashrc
	source .bashrc
}
