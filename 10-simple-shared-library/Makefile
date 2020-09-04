# requires 64-bit version of tdm-gcc: https://jmeubank.github.io/tdm-gcc/download/
shared.dll:	dll.go
	go build -buildmode=c-shared -o $@

clean:	
	rm -fv *.dll