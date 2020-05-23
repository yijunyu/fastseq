export PATH="$PATH:/home/gitpod/.pyenv/versions/3.8.2/bin:/usr/lib/llvm-11/bin:/workspace/clang/bin"
export LD_LIBRARY_PATH=/home/gitpod/.pyenv/versions/3.8.2/lib:$(llvm-config-11 --libdir)
export PKG_CONFIG_PATH=/home/gitpod/.pyenv/versions/3.8.2/lib/pkgconfig
alias push="git commit -am update && git push"
pyenv global 3.8.2
