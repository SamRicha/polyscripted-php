apt-get update && apt-get -y upgrade
apt-get install -y \
    git \
    make \
    autoconf \
    gcc \
    re2c \
    bison \
    libxml2-dev \
    vim

git clone --depth 1 https://github.com/php/php-src.git
cd php-src

sed 's/.*PHP_EXTRA_VERSION.*/PHP_EXTRA_VERSION = -polyscripted;/' Makefile > Makefile.polyscripted
yes | cp Makefile.polyscripted Makefile

./buildconf
./configure --without-pear 
make

