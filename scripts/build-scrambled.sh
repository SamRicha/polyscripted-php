#!/bin/bash
 
 
./resetPhp/reset-php.sh

if [[ $1 = "-t" ]]
then
	cd tests
	./remove-expected-out.sh
	./get-expected-out.sh
cd ..
fi

./scrambler


./transformer -f /php/php-src/ext/phar/phar/phar.php -replace=true
./transformer -f /php/php-src/ext/phar/build_precommand.php -replace=true

cd php-src
./buildconf
make install -k

cd /polyscripted-php
find . -type d -empty -delete

