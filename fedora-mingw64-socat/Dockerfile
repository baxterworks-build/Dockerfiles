FROM fedora:latest
RUN echo fastestmirror=1 >> /etc/dnf/dnf.conf
RUN dnf -v install --assumeyes make mingw64-gcc mingw64-binutils automake autoconf git && dnf clean all
WORKDIR /usr/src
RUN git clone --depth=1 https://repo.or.cz/socat.git
WORKDIR /usr/src/socat
COPY patches/*.patch .
RUN git apply *.patch
RUN autoconf && ./configure --host=x86_64-w64-mingw32 && make

