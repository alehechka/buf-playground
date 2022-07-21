buf_generate:
	cd proto ; rm -rf gen ; buf generate


# START SERVER

start_rest:
	PORT=3000 go run gateways/rest/main.go

start_inventory:
	go run services/inventory-api/main.go

start_session:
	go run services/session-api/main.go


# BUILD SERVER

build_rest:
	go build -o out/rest-server gateways/rest/main.go

build_inventory:
	go build -o out/inventory-api services/inventory-api/main.go

build_session:
	go build -o out/session-api services/session-api/main.go

build_all: build_rest build_inventory build_session
	