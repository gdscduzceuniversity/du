PREFIX := /usr
BINDIR := $(PREFIX)/bin
BUILD := build
GOC := go

all: build install

build:
	mkdir -pv $(BUILD) $(BINDIR)
	$(GOC) mod download && go mod verify
	$(GOC) build -v -o $(BUILD)/duapp

install:
	install -vm 755 $(BUILD)/duapp $(BINDIR)

clean:
	rm -rv $(BUILD)

.PHONY: build install clean
