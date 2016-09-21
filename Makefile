#
# Author	: 	Manuel A. Rodriguez 
# Email 	:	manuel.rdrs@gmail.com 
#
BIN_NAME 	:= go-proc-sentinel
BIN_DIR 	:= ./bin
SRC_DIR 	:= ./src

GOFLAGS ?= $(GOFLAGS:)

.PHONY: test bin cbin

test:
	cd $(SRC_DIR) && go test -v

bin:
	go build -o $(BIN_DIR)/$(BIN_NAME) $(SRC_DIR)/...

cbin:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cfgo $(GOFLAGS) -o $(BIN_DIR)/$(BIN_NAME) $(SRC_DIR)/...
