PROG_NAME=simple-activemq
PRODUCER_PROG_NAME=${PROG_NAME}-producer
CONSUMER_PROG_NAME=${PROG_NAME}-consumer
SOURCE_PRODUCER=producer.go
SOURCE_CONSUMER=consumer.go
CMD=go
BIN_DIR=target

all: clear build

clear:
	if [ -e "${BIN_DIR}/${PRODUCER_PROG_NAME}" ]; then rm ${BIN_DIR}/${PRODUCER_PROG_NAME}; fi
	if [ -e "${BIN_DIR}/${PRODUCER_PROG_NAME}.exe" ]; then rm ${BIN_DIR}/${PRODUCER_PROG_NAME}.exe; fi
	if [ -e "${BIN_DIR}/${PRODUCER_PROG_NAME}-osx" ]; then rm ${BIN_DIR}/${PRODUCER_PROG_NAME}-osx; fi
	if [ -e "${BIN_DIR}/${CONSUMER_PROG_NAME}" ]; then rm ${BIN_DIR}/${CONSUMER_PROG_NAME}; fi
	if [ -e "${BIN_DIR}/${CONSUMER_PROG_NAME}.exe" ]; then rm ${BIN_DIR}/${CONSUMER_PROG_NAME}.exe; fi
	if [ -e "${BIN_DIR}/${CONSUMER_PROG_NAME}-osx" ]; then rm ${BIN_DIR}/${CONSUMER_PROG_NAME}-osx; fi
	if [ -d "${BIN_DIR}" ]; then rmdir -p ${BIN_DIR}; fi

build_all: build build_osx build_win

build:
	mkdir -p ${BIN_DIR}
	GOOS="linux" GOARCH="amd64" ${CMD} build -o ${BIN_DIR}/${PRODUCER_PROG_NAME} ${SOURCE_PRODUCER}
	GOOS="linux" GOARCH="amd64" ${CMD} build -o ${BIN_DIR}/${CONSUMER_PROG_NAME} ${SOURCE_CONSUMER}

build_win:
	mkdir -p ${BIN_DIR}
	GOOS="windows" GOARCH="amd64" ${CMD} build -o ${BIN_DIR}/${PRODUCER_PROG_NAME}.exe ${SOURCE_PRODUCER}
	GOOS="windows" GOARCH="amd64" ${CMD} build -o ${BIN_DIR}/${CONSUMER_PROG_NAME}.exe ${SOURCE_CONSUMER}

build_osx:
	mkdir -p ${BIN_DIR}
	GOOS="darwin" GOARCH="amd64" ${CMD} build -o ${BIN_DIR}/${PRODUCER_PROG_NAME}-osx ${SOURCE_PRODUCER}
	GOOS="darwin" GOARCH="amd64" ${CMD} build -o ${BIN_DIR}/${CONSUMER_PROG_NAME}-osx ${SOURCE_CONSUMER}
