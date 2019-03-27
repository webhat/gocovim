

install:
	go install

gen-cover:
	go test -coverprofile cover.out
