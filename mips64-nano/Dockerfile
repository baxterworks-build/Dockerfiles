FROM debian:buster-slim
WORKDIR /opt
RUN apt update
RUN apt -qy install make bzip2 xz-utils wget
RUN wget -qO- https://toolchains.bootlin.com/downloads/releases/toolchains/mips64-n32/tarballs/mips64-n32--musl--stable-2018.11-1.tar.bz2 | tar -jxf -
RUN wget -qO- https://toolchains.bootlin.com/downloads/releases/toolchains/x86-64-core-i7/tarballs/x86-64-core-i7--musl--stable-2018.11-1.tar.bz2 | tar -jxf -
ENV PATH="/opt/mips64-n32--musl--stable-2018.11-1/bin:/opt/x86-64-core-i7--musl--stable-2018.11-1/bin:${PATH}"
WORKDIR /usr/src

RUN wget -O- https://ftpmirror.gnu.org/ncurses/ncurses-6.1.tar.gz | tar -zxf -
WORKDIR /usr/src/ncurses-6.1
#ncurses has to be one of the strangest cross-compiles I've encountered
#--host and --target don't make sense until you realise it wants a native *and* a cross compiling toolchain
#Thank FSM for http://www.linuxfromscratch.org/lfs/view/stable/chapter06/ncurses.html

ENV CFLAGS="-static -I/opt/mips64/include/ncursesw -I/opt/mips64/include"
ENV CPPFLAGS="-static -I/opt/mips64/include/ncursesw -I/opt/mips64/include"
ENV LDFLAGS="-static -L/opt/mips64/lib/"
#Hoo boy did this take a while - see that -static flag for the build cc? (used to build tools like make_hash)
#because I'm compiling on a glibc based distro, make_hash was failing to launch with a 'not found'
#I can probably fix it by launching it with something out of the musl toolchain or
#just giving up and installing Debian's GCC, but this works
#with any other combination of host/target/build/build cc, I got an x86_64 copy of ncursesw.a
RUN ./configure --disable-stripping --disable-shared --enable-static \
	--host=mips64-buildroot-linux-musl --with-build-cc="x86_64-buildroot-linux-musl-gcc -static" --prefix=/opt/mips64 \
	--enable-pc-files --enable-widec
RUN make -j
RUN make install 1> /dev/null
WORKDIR /usr/src
ENV PKG_CONFIG_PATH=/opt/mips64-n32--musl--stable-2018.11-1/lib/pkgconfig

RUN wget -qO- https://ftpmirror.gnu.org/nano/nano-4.3.tar.xz | tar -Jxf -
WORKDIR /usr/src/nano-4.3
RUN ./configure --host=mips64-buildroot-linux-musl --prefix=/opt/mips64
RUN make -j
RUN make install

RUN mkdir /output
RUN cp /opt/mips64/bin/* /output
RUN mkdir -p /output/usr/share/terminfo/x
RUN cp /opt/mips64/share/terminfo/x/xterm /output/usr/share/terminfo/x
ENTRYPOINT tar --create /output

