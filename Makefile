GOBUILD=go build
GOTEST=go test

all:clean stop build_server
	./ecommerce-go
build_server:
	$(GOBUILD) -v .
clean:
	rm -f ecommerce-go
stop:
	pkill ecommerce-go || true
test:
	cd helper && $(GOTEST) -v .