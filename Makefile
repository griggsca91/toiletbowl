go-files = ./toiletbowl/db.go \
			main.go

build: get-libraries main frontend

run: main
	ulimit -n 65536
	nohup ./main &

get-libraries:
	dep ensure

watch: install-frontend
	./node_modules/.bin/webpack -d --watch

run-dev: get-libraries frontend main
	go run -x main.go --port=3000 

frontend: install-frontend
	./node_modules/.bin/webpack

init-db:
	CREATE USER toiletbowl
		WITH
			SUPERUSER
			CREATEDB
			CREATEROLE
			LOGIN
			PASSWORD 'password';

	createdb toiletbowl;
	alter user toiletbowl with encrypted password 'password';

install-frontend: package.json
	npm i

main:  $(go-files)
	go build -o main

clean:
	rm main
