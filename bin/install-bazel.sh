command -v bazel >/dev/null 2>&1 || {
	curl -L -O https://github.com/bazelbuild/bazel/releases/download/0.26.1/bazel-0.26.1-installer-linux-x86_64.sh
	chmod +x bazel-*.sh
	./bazel-0.26.1-installer-linux-x86_64.sh --prefix=/workspace
	rm bazel-0.26.1-installer-linux-x86_64.sh
}
