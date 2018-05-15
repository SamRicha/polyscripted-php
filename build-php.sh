git clone --depth 1 https://github.com/php/php-src.git
cd php-src

./buildconf

sed 's/PHP_EXTRA_VERSION=.*/PHP_EXTRA_VERSION=-polyscripted/' configure > configure.polyscripted
yes | cp configure.polyscripted configure

./configure \
  --without-pear \
  --exec-prefix=/polyscripted-php \
  --prefix=/polyscripted-php

make install

