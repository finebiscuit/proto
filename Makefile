.PHONY: all
all: biscuit biscuitdata

.PHONY: biscuit
biscuit:
	buf generate biscuit --template biscuit/buf.gen.yaml -o biscuit

.PHONY: biscuitdata
biscuitdata:
	buf generate biscuitdata --template biscuitdata/buf.gen.yaml -o biscuitdata

