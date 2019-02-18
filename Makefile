PHONY: build generate run

working_dir = $(shell basename `pwd`)
generate: 
	docker run -it -v $(PWD):/go/src/github.com/bdna/$(shell basename `pwd`) -w /go/src/github.com/bdna/$(shell basename `pwd`) docker-goa goa gen github.com/bdna/$(shell basename `pwd`)/design

lint:
	docker-compose -f docker-compose.yml run proxy /bin/sh -c "golangci-lint run"
