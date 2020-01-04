FROM fedora:latest
RUN echo fastestmirror=1 >> /etc/dnf/dnf.conf
# https://bgstack15.wordpress.com/2017/05/13/dnf-ignore-weak-dependencies/
RUN dnf install --setopt=install_weak_deps=False --assumeyes  git mingw64-gcc mingw64-binutils && dnf clean all
RUN dnf search mingw64

RUN mkdir src
WORKDIR /src

RUN git clone https://github.com/osmocom/rtl-sdr
WORKDIR /src/rtl-sdr
