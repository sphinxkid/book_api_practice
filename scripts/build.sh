#!/bin/bash

. config.sh

cd $GO_SRC_DIR
go build -v -o $BIN_DIR all