setup:
	go get github.com/sasha-s/goimpl/cmd/goimpl

watch:
	ginkgo watch -r -notify -cover -succinct
