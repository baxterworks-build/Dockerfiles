FROM fedora:latest
RUN echo fastestmirror=1 >> /etc/dnf/dnf.conf
RUN dnf -v install --assumeyes make mingw64-gcc mingw64-binutils && dnf clean all
