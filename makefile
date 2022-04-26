PROJECT_NAME:="Go Boiler Plate Service"

PROJECT_ROOT?=$(shell pwd)
PROJECT_MAIN_PKG=cmd

GO_CMD:=go
GO_BUILD:=${GO_CMD} build

BINARY_NAME:=go-boiler-plate

## Run local go
.PHONY: run
run: 
	@echo " "
	@echo " > Choose a command run in "${PROJECT_NAME}":"
	@echo " > Build -> "${PROJECT_ROOT}
	@${GO_BUILD} -o  ${BINARY_NAME} ${PROJECT_MAIN_PKG}/*.${GO_CMD}
	@echo " > Build -> Done "
	@echo " > "run ${PROJECT_MAIN_PKG}/*.${GO_CMD}
	@echo " > Starting Server...\n"
	@./${BINARY_NAME}
