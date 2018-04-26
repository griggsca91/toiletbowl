go-files = ./rename/db.go \
			main.go

build: get-libraries main frontend

run: main
	ulimit -n 65536
	nohup ./main &

get-libraries:
	dep ensure

watch:
	./node_modules/.bin/webpack -d --watch

run-dev: main
	go run -x main.go --port=3000 

frontend:
	./node_modules/.bin/webpack

init-db:
	createuser --createdb --createrole --superuser --replication renameuser;

main: $(go-files)
	go build -o main

clean:
	rm main
