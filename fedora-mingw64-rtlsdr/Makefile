SHELL=bash
.PHONY: debug container push all

debug:
	docker build . -t voltagex/fedora-mingw64-rtlsdr
	docker run -it voltagex/fedora-mingw64-rtlsdr

#container:
#	docker build -t voltagex/fedora-mingw64-qemu .

#push:
#	docker push voltagex/fedora-mingw64-qemu
#	docker push voltagex/fedora-mingw64-qemu:debug

#all: debug container push
