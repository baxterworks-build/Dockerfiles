FROM fedora:latest
RUN echo fastestmirror=1 >> /etc/dnf/dnf.conf
RUN dnf -v install --assumeyes make mingw64-gcc mingw64-binutils binutils findutils flex bison mingw64-pkg-config perl-podlators texinfo  \
mingw64-glib2 mingw64-pixman mingw64-SDL2 mingw64-gettext \
mingw64-curl mingw64-libpng mingw64-libjpeg-turbo \
mingw64-libgcrypt mingw64-gnutls mingw64-bzip2 mingw64-libssh2 mingw64-libxml2 p7zip p7zip-plugins gcc mingw32-nsis mingw32-nsiswrapper bzip2 wget && dnf clean all

COPY headers/* /usr/x86_64-w64-mingw32/sys-root/mingw/include/
