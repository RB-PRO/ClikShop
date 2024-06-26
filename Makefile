all: run

run:
	go run cmd/main/main.go

push:
	git push git@github.com:RB-PRO/ClikShop.git

pull:
	git pull git@github.com:RB-PRO/ClikShop.git

pushW:
	git push https://github.com/RB-PRO/ClikShop.git

pullW:
	git pull https://github.com/RB-PRO/ClikShop.git

pushCar:
	scp main root@194.87.107.129:go/ClikShop/

build-config:
	go env GOOS GOARCH

build-windows-to-linux:
	set GOARCH=amd64 set GOOS=linux go build cmd/main/main.go  

build-linux-to-windows:
	export GOARCH=amd64 export GOOS=windows go build cmd/main/main.go 

build-car:
	set GOARCH=amd64
	set GOOS=linux
	set CGO_ENABLED=0
	go env GOOS GOARCH
	go build -o updator cmd/main/main.go
	scp main Delivery.xlsx config_rb.json config_wp.json root@194.87.107.129:go/ClikShop/

build-carW:
	set GOARCH=amd64
	set GOOS=windows 
	go env GOOS GOARCH
	go build -o main.exe cmd/main/main.go
	scp main Delivery.xlsx config_rb.json config_wp.json root@194.87.107.129:go/ClikShop/


pushBaraki:
	scp main notification_updator.json sender.json root@185.154.192.111:updator/

pushTrudeks:
	scp main notification_updator.json sender.json root@193.124.117.19:go/clikshopupd/