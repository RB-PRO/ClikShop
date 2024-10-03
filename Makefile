all: run

run:
	go run cmd/main/main.go

push:
	git push git@github.com:RB-PRO/ClikShop.git

pull:
	git pull git@github.com:RB-PRO/ClikShop.git

pushW:
	git push https://ClikShop.git

pullW:
	git pull https://ClikShop.git

pushCar:
	scp main root@194.87.107.129:go/ClikShop/

build-config:
	go env GOOS GOARCH

build-windows-to-linux:
	set GOARCH=amd64 set GOOS=linux go build cmd/main/main.go  

build-linux-to-windows:
	export GOARCH=amd64 export GOOS=windows go build cmd/main/main.go 

build-updator:
	set GOARCH=amd64
	set GOOS=linux
	set CGO_ENABLED=0
	go env GOOS GOARCH
	go build -o bin/updator ./updator/cmd/main.go

build-actualizer:
	set GOARCH=amd64
	set GOOS=linux
	set CGO_ENABLED=0
	go env GOOS GOARCH
	go build -o bin/actualizer ./actualizer/cmd/main.go

scp-car:
	set GOARCH=amd64
	set GOOS=linux
	set CGO_ENABLED=0
	go env GOOS GOARCH
	go build -o updator cmd/main/main.go
	scp bin/updator config.json root@194.87.107.129:go/clikshop/

pushBaraki:
	scp updator config.json root@185.154.192.111:updator/

pushTrudeks:
	scp updator config.json root@193.124.117.19:go/clikshopupd/