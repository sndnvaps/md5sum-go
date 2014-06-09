[![Build Status](https://travis-ci.org/sndnvaps/md5sum-golang.svg?branch=master)](http://travis-ci.org/sndnvaps/md5sum-golang)



# golang version md5sum 

```bash
Usage: Dirname/md5sum.exe -options=value ...
  -Type="": C = Check md5sum, M =  Make md5sum 
  -file="": File need to calculate md5sum
  -help="": print the help info
  -md5in="": Input md5sum file to check
  -md5out="": output md5sum to file
```

# example:

## calculate md5sum from file and write it the file 
	
	$ ./md5sum -Type="M" -file="README.md" -md5out="READMD.md.md5"
	
`output:`
	md5sum = 583a56e66624016487c236d738d5ad20 README.md
	
## Check md5sum from file 

	$ ./md5sum -Type="C" -file="README.md" -md5in="README.md.md5"
	
`output:`

	583a56e66624016487c236d738d5ad20 README.md
	verifiy success

## show help info 

```bash
$ ./md5sum -help=help
Usage: E:\golang_test\md5sum\md5sum.exe -options=value ...
  -Type="": C = Check md5sum, M =  Make md5sum
  -file="": File need to calculate md5sum
  -help="": print the help info
  -md5in="": Input md5sum file to check
  -md5out="": output md5sum to file
```

	
	
	


