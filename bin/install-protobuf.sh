command -v protoc >/dev/null 2>&1 || {
	git clone https://github.com/google/protobuf
	cd protobuf
	git checkout v3.6.0
	./autogen.sh
	./configure --prefix=/workplace/fastseq/nucleus-0.5.1/usr/local
	make -j 16
	make install
	grep "nucleus-0.5.1/usr/local/bin" .bashrc > /dev/null 2>& 1 || {
		echo export PATH='"/workspace/fastseq/nucleus-0.5.1/usr/local/bin:$PATH"' >> .bashrc
	}
}
