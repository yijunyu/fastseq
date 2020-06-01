from nucleus.io import fastq
r = fastq.FastqReader('test_reads.fastq')
for v in r:
  print(v)
