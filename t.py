from nucleus.io import vcf
r = vcf.VcfReader('test_samples.vcf')
for v in r:
    print(v.start)
