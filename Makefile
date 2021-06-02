BINARY=bin/app

build:
	go build -o ${BINARY} app/*.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t crypto_conversion .

docker-run:
	docker-compose up --build -d

run:
	go run app/main.go

stop:
	docker-compose down

.PHONY: clean build docker docker-run stop