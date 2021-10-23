COMPILER=go

test: 
	cd testfile/simple && \
	$(COMPILER) build -o $(shell pwd)/bin/$@ -race

server: 
	cd testfile/gin-test && \
	$(COMPILER) build -o $(shell pwd)/bin/$@ 


.PHONY: test server 
clean: 
	rm -rf $(shell pwd)/bin/test
	rm -rf $(shell pwd)/bin/server

all: test server 