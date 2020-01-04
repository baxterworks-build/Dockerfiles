FROM fedora:latest
RUN echo fastestmirror=1 >> /etc/dnf/dnf.conf
# https://bgstack15.wordpress.com/2017/05/13/dnf-ignore-weak-dependencies/
RUN dnf install --setopt=install_weak_deps=False --assumeyes  git mingw64-gcc mingw64-binutils cmake && dnf clean all

RUN mkdir src
WORKDIR /src

RUN git clone https://github.com/osmocom/rtl-sdr
RUN mkdir /src/rtl-sdr/build
WORKDIR /src/rtl-sdr/build

# https://stackoverflow.com/questions/51509173/how-can-i-make-cmake-use-mingw-w64-gcc-g
RUN cmake -G "MinGW Makefiles"
