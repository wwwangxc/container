test:
	go test -v -covermode=count -coverprofile=coverage_unit.out -coverpkg=./... -gcflags="-N -l" ./...


bench:
	go test -benchmem -bench . github.com/wwwangxc/container/...
