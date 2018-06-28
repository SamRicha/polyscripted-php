#!/bin/bash
  
./resetPhp/reset-php.sh
./tests/remove-expected-out.sh

cd tests
./get-expected-out.sh
cd ..

./scrambler


./transformer -f /php/php-src/ext/phar/phar/phar.php -replace=true
./transformer -f /php/php-src/ext/phar/build_precommand.php -replace=true

cd php-src
make install -k

cd /polyscripted-php
find . -type d -empty -delete

