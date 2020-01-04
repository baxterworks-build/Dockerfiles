FROM fedora:latest
RUN echo fastestmirror=1 >> /etc/dnf/dnf.conf
RUN dnf -v install --assumeyes git mingw64-gcc mingw64-binutils && dnf clean
RUN dnf search mingw64

RUN mkdir src
WORKDIR /src

RUN git clone https://github.com/osmocom/rtl-sdr
WORKDIR /src/rtl-sdr
