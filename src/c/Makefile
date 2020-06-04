flatc = flatc
targets += c.fbs c.out

c.fbs c.out: c
	./c examples/BinaryDigit.c c.fbs
	./c c.fbs > c.out

targets += fast_generated.go fast_generated.go-e

fast_generated.go: fast.fbs
	$(flatc) -g --gen-onefile --gen-mutable fast.fbs
	sed -i -e 's/package/package main/' $@
	sed -i -e '/fast_ "fast_"/d' $@
	sed -i -e 's/fast___Element_\.//g' $@
	sed -i -e 's/fast___Data_\.//g' $@
	sed -i -e 's/fast___Bugs_\.//g' $@
	sed -i -e 's/fast___Bugs___Bug_\.//g' $@
	sed -i -e 's/fast___Log_\.//g' $@
	sed -i -e 's/fast___Log___Commit_\.//g' $@
	sed -i -e 's/fast___Log___Commit___LogDiff_\.//g' $@
	sed -i -e 's/fast___Log___Commit___LogDiff___Hunk_\.//g' $@
	sed -i -e 's/fast___Delta_\.//g' $@
	sed -i -e 's/fast___Delta___Diff_\.//g' $@
	sed -i -e 's/fast___Pairs_\.//g' $@
	sed -i -e 's/fast___Pairs___Pair_\.//g' $@
	sed -i -e 's/fast___Slices_\.//g' $@
	sed -i -e 's/fast___Slices___Slice_\.//g' $@
	sed -i -e 's/fast___Slices___Slice___SourceFile_\.//g' $@
	sed -i -e 's/fast___Slices___Slice___SourceFile___Function_\.//g' $@
	sed -i -e 's/fast___Slices___Slice___SourceFile___Function___Variable_\.//g' $@
	sed -i -e 's/fast___Graph_\.//g' $@
	sed -i -e 's/fast_\.//g' $@
	sed -i -e 's/fast___Element___Unit_.LanguageType/LanguageType/g' $@
	sed -i -e 's/fast___Element___Literal_.LiteralType/LiteralType/g' $@

targets += fast.fbs fast.fbs-e
fast.fbs: fast.proto
	$(flatc) --proto fast.proto
	sed -i -e 's/namespace fast/namespace fast_/g' $@
	sed -i -e 's/fast\./fast_./g' $@
	sed -i -e 's/\[ubyte\]/string/g' $@

targets += c
c: main.go parser/c_parser.go fast_generated.go
	go build

targets += parser
parser/c_parser.go: C.g4
	antlr4 -Dlanguage=Go -o parser C.g4

clean:
	rm -rf $(targets)
