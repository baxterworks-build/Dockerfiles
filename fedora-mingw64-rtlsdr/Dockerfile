FROM fedora:latest

#as close to apt update as I can get
#RUN dnf check-update &> /dev/null || true

# https://bgstack15.wordpress.com/2017/05/13/dnf-ignore-weak-dependencies/
# https://stackoverflow.com/questions/33124261/usr-bin-mingw32-cmake-line-84-fg-no-job-control-midori-cross-compiling-erro
RUN echo fastestmirror=1 >> /etc/dnf/dnf.conf && \ 
    dnf install --setopt=install_weak_deps=False --assumeyes git cmake make mingw64-gcc mingw64-binutils mingw64-libusbx \ 
    && dnf clean all

RUN mkdir src
WORKDIR /src

RUN git clone https://github.com/osmocom/rtl-sdr
RUN mkdir /src/rtl-sdr/build
WORKDIR /src/rtl-sdr/build

RUN mingw64-cmake .. && mingw64-make install
