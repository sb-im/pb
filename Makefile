.PHONY: all
all: gen-signal

.PHONY: gen-signal
gen-signal:
	@cd signal && ./protogen.sh
