# Maintainer: Kristof Vannotten <kristof@vannotten.be>

pkgname=pacnews
pkgver=0.0.1
pkgrel=1
pkgdesc="Arch news reader + pacman interruption"
arch=('x86_64' 'i686')
url="https://github.com/kvannotten/pacnews"
license=('GPLv2')
depends=('go')
makedepends=('git')
options=('!strip' '!emptydirs')
_gourl=github.com/kvannotten/github

build() {
  GOPATH="$srcdir" go get -fix -v -x ${_gourl}/...
}

check() {
  GOPATH="$GOPATH:$srcdir" go test -v -x ${_gourl}/...
}

package() {
  mkdir -p "$pkgdir/usr/bin"
  install -p -m755 "$srcdir/bin/"* "$pkgdir/usr/bin"

  mkdir -p "$pkgdir/$GOPATH"
  cp -Rv --preserve=timestamps "$srcdir/"{src,pkg} "$pkgdir/$GOPATH"

  if [ -e "$srcdir/src/$_gourl/pacnews.hook" ]; then
    install -Dm644 "$srcdir/src/$_gourl/pacnews.hook" \
      "$pkgdir/etc/pacman.d/hooks/pacnews.hook"
  fi

  touch /var/cache/pacnews.db

  # Package license (if available)
  for f in LICENSE COPYING LICENSE.* COPYING.*; do
    if [ -e "$srcdir/src/$_gourl/$f" ]; then
      install -Dm644 "$srcdir/src/$_gourl/$f" \
        "$pkgdir/usr/share/licenses/$pkgname/$f"
    fi
  done
}

# vim:set ts=2 sw=2 et:
