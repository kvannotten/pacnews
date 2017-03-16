# Maintainer: Kristof Vannotten <kristof@vannotten.be>
pkgname=pacnews
pkgver=20170316.2_f74a11c
pkgrel=1
pkgdesc="Arch news reader"
arch=('i686' 'x86_64')
license=('GPL')
depends=(
)
makedepends=(
	'go'
	'git'
)

source=(
	"pacnews::git://github.com/kvannotten/pacnews#branch=${BRANCH:-master}"
	"pacnews.go"
	"pacnews.hook"
	"utility.go"
)

md5sums=(
	'SKIP'
	'1f188b5fd15ebeeb5d7bbd3229e4e333'
	'97e6eb7ff7a657632971ac37135754b8'
	'1d8027b32f5159b71dfb2ffd198df19f'
)

backup=(
)

pkgver() {
	if [[ "$PKGVER" ]]; then
		echo "$PKGVER"
		return
	fi

	cd "$srcdir/$pkgname"
	local date=$(git log -1 --format="%cd" --date=short | sed s/-//g)
	local count=$(git rev-list --count HEAD)
	local commit=$(git rev-parse --short HEAD)
	echo "$date.${count}_$commit"
}

build() {
	cd "$srcdir/$pkgname"

	if [ -L "$srcdir/$pkgname" ]; then
		rm "$srcdir/$pkgname" -rf
		mv "$srcdir/.go/src/$pkgname/" "$srcdir/$pkgname"
	fi

	rm -rf "$srcdir/.go/src"

	mkdir -p "$srcdir/.go/src"

	export GOPATH="$srcdir/.go"

	mv "$srcdir/$pkgname" "$srcdir/.go/src/"

	cd "$srcdir/.go/src/$pkgname/"
	ln -sf "$srcdir/.go/src/$pkgname/" "$srcdir/$pkgname"

	git submodule update --init

	go get -v \
		-gcflags "-trimpath $GOPATH/src"
}

package() {
	find "$srcdir/.go/bin/" -type f -executable | while read filename; do
		install -DT "$filename" "$pkgdir/usr/bin/$(basename $filename)"
	done
	install -DT -m0755 "$srcdir/pacnews.hook" "$pkgdir/etc/pacman.d/hooks/pacnews.hook"
}
