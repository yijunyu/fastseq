if [ ! -f nucleus-0.5.1/bazel-bin/licenses.zip ]; then
	bin/install-protobuf.sh
	bin/install-clif.sh
	bin/install-pyclif.sh
	if [ -f bin/nucleusaa ]; then
		if [ ! -d nucleus-0.5.1 ]; then
			wget https://github.com/google/nucleus/archive/0.5.1.tar.gz
			tar xvfz 0.5.1.tar.gz
		fi
		cat bin/nucleusa* | tar xvfj - -C nucleus-0.5.1
	fi
	echo export PATH='"/workspace/fastseq/nucleus-0.5.1/usr/local/clif/bin:$PATH"' >> .bashrc
	echo export LD_LIBRARY_PATH='"/workspace/fastseq/nucleus-0.5.1/usr/local/lib"' >> .bashrc
fi
cd nucleus-0.5.1
