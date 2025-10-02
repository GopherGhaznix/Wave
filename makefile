# Makefile

OUTPUT=index.html
SRC=./examples/wrapper-attributes/...

.PHONY: dev clean

dev:
	go run $(SRC) > $(OUTPUT)

clean:
	rm -f $(OUTPUT)