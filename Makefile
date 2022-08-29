CMD_DIR=./cmd/goenum

GO_FILES:=$(shell find . -name '*.go')
MOD_FILES:=go.mod go.sum

TARGET_DIR=bin
TARGET=$(TARGET_DIR)/goenum

.PHONY: default clean install test

default: $(TARGET)

$(TARGET_DIR):
	mkdir -p $(TARGET_DIR)

$(TARGET): $(GO_FILES) $(MOD_FILES) $(TARGET_DIR)
	go build -o $(TARGET) $(CMD_DIR)

install: $(TARGET)
	go install $(CMD_DIR)

clean:
	rm -r $(TARGET_DIR)

test:
	go test ./... -cover -race -count=1
