# requires 64-bit version of tdm-gcc: https://jmeubank.github.io/tdm-gcc/download/
tomcat.dll:	main.go
	go build -buildmode=c-shared -o $@ && cp -v $@ ../plugins

clean:	
	rm -fv *.dll