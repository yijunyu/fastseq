export PATH="/home/gitpod/.pyenv/versions/3.6.0/bin:/usr/lib/llvm-11/bin:/workspace/clang/bin:$PATH"
export LD_LIBRARY_PATH=/home/gitpod/.pyenv/versions/3.6.0/lib:$(llvm-config-11 --libdir)
export PKG_CONFIG_PATH=/home/gitpod/.pyenv/versions/3.6.0/lib/pkgconfig
alias push="git commit -am update && git push"
pyenv global 3.6.0
