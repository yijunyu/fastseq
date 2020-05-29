# fastSeq - a faster genome sequencing toolkit

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/yijunyu/fastseq)

You can also run it using the following docker command:
```bash
docker run -it yijun/nucleus
```

## Dependencies
* https://github.com/google/protobuf
* https://github.com/google/clif
* https://github.com/google/nucleus
* https://github.com/google/deepvariant

## Usage
```bash
cd /home/gitpod/fastseq/nucleus
bazel build -c opt $COPT_FLAGS nucleus/examples:all
```
## References
[1] Yijun Yu. "[fAST: Flattening Abstract Syntax Trees for Efficiency](http://oro.open.ac.uk/59268/)". In: 41st ACM/IEEE International Conference on Software Engineering, 25-31 May 2019, Montreal, Canada, ACM and IEEE.
[demo](https://gitpod.io/#https://github.com/yijunyu/demo), [paper](http://oro.open.ac.uk/59268), [poster](doc/fast-poster-A0.pdf)
