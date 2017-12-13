pkgname=zabbix-agent-extension-php-fpm
pkgver=20170919.4_5f2bca6
pkgrel=1
pkgdesc="Extension for zabbix-agent to monitoring PHP-FPM metrics"
arch=('any')
license=('GPL')
makedepends=('go' 'make')
install='install.sh'
source=("https://github.com/zarplata/$pkgname.git#branch=master")
md5sums=('SKIP')

pkgver() {
    cd "$srcdir/$pkgname"
    make ver
}
    
build() {
    cd "$srcdir/$pkgname"
    make
}

package() {
    cd "$srcdir/$pkgname"

    install -Dm 0755 .out/"${pkgname}" "${pkgdir}/usr/bin/${pkgname}"
    install -Dm 0644 "${pkgname}.conf" "${pkgdir}/etc/zabbix/zabbix_agentd.conf.d/${pkgname}.conf"
}
