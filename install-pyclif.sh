command -v pyclif >/dev/null 2>&1 || {
	if [ -f pyclif.tar.bz2 ]; then
		tar xvfj /workspace/fastseq/pyclif.tar.bz2 -C /workspace 
		patch /workspace/.pip-modules/lib/python3.8/site-packages/clif/protos/ast_pb2.py < /workspace/fastseq/ast_pb2.py.patch 
		patch /workspace/.pip-modules/lib/python3.8/site-packages/clif/python/gen.py < /workspace/fastseq/gen.py.patch 
	fi
	echo export PATH='"/workspace/.pipmodules/bin:$PATH"' >> .bashrc
	source .bashrc
}
