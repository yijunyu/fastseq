command -v flex >/dev/null 2>&1 || {
	if [ ! -f flex.tar.bz2 ]; then
	else
		tar xvfj /workspace/fastseq/flex.tar.bz2 -C /workspace 
	fi
	echo export PATH='"/workspace/flex/bin:$PATH"' >> .bashrc
	source .bashrc
}
