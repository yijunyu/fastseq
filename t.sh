#!/bin/bash
function gdownload {
	if [ ! -f $1 ]; then
		gsutil cp gs://deepvariant/case-study-testdata/$1 $1
	fi
}
export -f gdownload
function wdownload {
	if [ ! -f NA12878_calls.vcf.gz ]; then
		wget ftp://ftp-trace.ncbi.nlm.nih.gov/giab/ftp/release/NA12878_HG001/latest/GRCh37/HG001_GRCh37_GIAB_highconf_CG-IllFB-IllGATKHC-Ion-10X-SOLID_CHROM1-X_v.3.3.2_highconf_PGandRTGphasetransfer.vcf.gz -O NA12878_calls.vcf.gz
	fi
	if [ ! -f NA12878_calls.vcf.gz.tbi ]; then
		wget ftp://ftp-trace.ncbi.nlm.nih.gov/giab/ftp/release/NA12878_HG001/latest/GRCh37/HG001_GRCh37_GIAB_highconf_CG-IllFB-IllGATKHC-Ion-10X-SOLID_CHROM1-X_v.3.3.2_highconf_PGandRTGphasetransfer.vcf.gz.tbi -O NA12878_calls.vcf.gz.tbi
	fi
	if [ ! -f test_samples.vcf ]; then
		wget https://raw.githubusercontent.com/google/nucleus/master/nucleus/testdata/test_samples.vcf
	fi
}
export -f wdownload
cd /host
gdownload NA12878_sliced.bam
gdownload NA12878_sliced.bam.bai
gdownload hs37d5.fa.gz
gdownload hs37d5.fa.gz.fai
gdownload hs37d5.fa.gz.gzi
wdownload

