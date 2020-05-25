command -v pyclif >/dev/null 2>&1 || {
	if [ -f bin/pyclif.tar.bz2 ]; then
		tar xvfj /workspace/fastseq/bin/pyclif.tar.bz2 -C /workspace 
		patch /workspace/.pip-modules/lib/python3.8/site-packages/clif/protos/ast_pb2.py < /workspace/fastseq/bin/ast_pb2.py.patch 
		patch /workspace/.pip-modules/lib/python3.8/site-packages/clif/python/gen.py < /workspace/fastseq/bin/gen.py.patch 
	fi
	echo export PATH='"/workspace/.pipmodules/bin:$PATH"' >> .bashrc
	source .bashrc
}
