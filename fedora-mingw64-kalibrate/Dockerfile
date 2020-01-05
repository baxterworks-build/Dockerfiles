FROM voltagex/fedora-mingw64-rtlsdr
RUN dnf --assumeyes install mingw64-fftw automake autoconf libtool mingw64-gcc mingw64-gcc-c++
RUN mkdir -p /src /output

WORKDIR /src
RUN git clone --depth=1 https://github.com/baxterworks-build/kalibrate-rtl

WORKDIR /src/kalibrate-rtl

RUN ln -s /usr/x86_64-w64-mingw32/sys-root/mingw/include/windows.h /usr/x86_64-w64-mingw32/sys-root/mingw/include/Windows.h
RUN ./bootstrap
RUN sed -ie s/^LRT_FLAGS.*//gi Makefile.in && sed -ie s/^LRT_FLAGS.*//gi src/Makefile.in
RUN ./configure --host=x86_64-w64-mingw32 && make -j

RUN cp 	/usr/x86_64-w64-mingw32/sys-root/mingw/bin/libgcc_s_seh-1.dll \
	/usr/x86_64-w64-mingw32/sys-root/mingw/bin/libfftw3-3.dll \ 
	/usr/x86_64-w64-mingw32/sys-root/mingw/bin/librtlsdr.dll \
	/usr/x86_64-w64-mingw32/sys-root/mingw/bin/libusb-1.0.dll \
	/usr/x86_64-w64-mingw32/sys-root/mingw/bin/libwinpthread-1.dll \   
	/usr/x86_64-w64-mingw32/sys-root/mingw/bin/libstdc++-6.dll \
	/usr/x86_64-w64-mingw32/sys-root/mingw/bin/rtl*.exe \ 
	/src/kalibrate-rtl/src/kal.exe /output/

ENTRYPOINT tar --create /output
