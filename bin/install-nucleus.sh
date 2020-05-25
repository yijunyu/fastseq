command -v nucleus >/dev/null 2>&1 || {
	bin/install-protobuf.sh
	bin/install-clif.sh
	bin/install-pyclif.sh
	if [ ! -f bin/nucleus.tar.bz2 ]; then
		if [ ! -d nucleus-0.5.1 ]; then
			wget https://github.com/google/nucleus/archive/0.5.1.tar.gz
			tar xvfz 0.5.1.tar.gz
		fi
		cd nucleus-0.5.1
        patch install.sh < /workspace/fastseq/bin/install.sh.patch
		./install.sh 
		cd /workspace && tar cvfj fastseq/bin/nucleus.tar.bz2 usr && cd -
	else
		tar xvfj bin/nucleus.tar.bz2 -C /workspace
	fi
	echo export PATH='"/workspace/.pipmodules/bin:$PATH"' >> .bashrc
	. .bashrc
}
