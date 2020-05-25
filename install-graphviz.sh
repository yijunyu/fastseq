command -v dot >/dev/null 2>&1 || {
	./install-flex.sh
	if [ ! -f graphviz.tar.bz2 ]; then
		if [ ! -d graphviz-2.44.0 ]; then
			wget https://gitlab.com/graphviz/graphviz/-/archive/2.44.0/graphviz-2.44.0.tar.bz2
			tar xvfj graphviz-2.44.0.tar.bz2
		fi
		cd graphviz-2.44.0
		./autogen.sh
		./configure --prefix=/workspace/graphviz
		make -j16
		make install
		cd /workspace && tar cvfj fastseq/graphviz.tar.bz2 graphviz && cd -
	else
		tar xvfj graphviz.tar.bz2 -C /workspace
	fi
	echo export PATH='"/workspace/graphviz/bin:$PATH"' >> .bashrc
	. .bashrc
}
