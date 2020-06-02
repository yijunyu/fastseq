mkdir -p /workspace/bin
cat > /workspace/bin/push <<EOF
git commit -am "$@*" && git push
EOF
cat > /home/gitpod/.inputrc <<EOF
"\e[A": history-search-backward
"\e[B": history-search-forward
"\e[C": forward-char
"\e[D": backward-char
"\M-o": "\C-p\C-a\M-f "
$if Bash
  Space: magic-space
$endif
set visible-stats on
set show-all-if-ambiguous on
EOF
chmod +x /workspace/bin/push
export PATH="/home/gitpod/bin:$PATH"
export PATH="/home/gitpod/nucleus/usr/local/clif/bin:$PATH"
export PATH="/workspace/bin:$PATH"
export LD_LIBRARY_PATH="/home/gitpod/nucleus/usr/local/lib"
