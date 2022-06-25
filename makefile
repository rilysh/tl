all:
	rm -rf ./build
	mkdir build
	make compile
	make tarpkg

compile:
	go build -ldflags "-w" tl.go
	mv tl ./build

cleanup:
	rm -rf ./build

tarpkg:
	tar cf ./build/tl.tar ./build/tl
	xz -z ./build/tl.tar
