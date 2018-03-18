generate: cmd/bootstrap-egobook/main.ego.go

cmd/bootstrap-egobook/main.ego.go: cmd/bootstrap-egobook/main.ego
	go generate ./...

run: generate
	goo install ./cmd/bootstrap-egobook
	bootstrap-egobook

.PHONY: generate run