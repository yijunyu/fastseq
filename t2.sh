# NA12878_sliced.bam
samtools view -h \
ftp://ftp-trace.ncbi.nlm.nih.gov/giab/ftp/data/NA12878/NIST_NA12878_HG001_HiSeq_300x/RMNISTHS_30xdownsample.bam \
20:10,000,000-10,100,000 \
-o NA12878_sliced.bam


# NA12878_sliced.bam.bai`
samtools index NA12878_sliced.bam


# NA12878_calls.vcf.gz
wget \
ftp://ftp-trace.ncbi.nlm.nih.gov/giab/ftp/release/NA12878_HG001/latest/GRCh37/HG001_GRCh37_GIAB_highconf_CG-IllFB-IllGATKHC-Ion-10X-SOLID_CHROM1-X_v.3.3.2_highconf_PGandRTGphasetransfer.vcf.gz \
-O NA12878_calls.vcf.gz


# NA12878_calls.vcf.gz.tbi
wget \
ftp://ftp-trace.ncbi.nlm.nih.gov/giab/ftp/release/NA12878_HG001/latest/GRCh37/HG001_GRCh37_GIAB_highconf_CG-IllFB-IllGATKHC-Ion-10X-SOLID_CHROM1-X_v.3.3.2_highconf_PGandRTGphasetransfer.vcf.gz.tbi \
-O NA12878_calls.vcf.gz.tbi


# hs37d5.fa.gz
wget \
ftp://ftp.1000genomes.ebi.ac.uk/vol1/ftp/technical/reference/phase2_reference_assembly_sequence/hs37d5.fa.gz \
-O hs37d5.fa.gz


# hs37d5.fa.gz.fai
wget \
ftp://ftp.1000genomes.ebi.ac.uk/vol1/ftp/technical/reference/phase2_reference_assembly_sequence/hs37d5.fa.gz.fai \
-O hs37d5.fa.gz.fai


# hs37d5.fa.gz.gzi
wget \
ftp://ftp.1000genomes.ebi.ac.uk/vol1/ftp/technical/reference/phase2_reference_assembly_sequence/hs37d5.fa.gz.gzi \
-O hs37d5.fa.gz.gzi
