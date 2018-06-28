#!/bin/bash

cd tests 
./get-expected-out.sh
cd ..

./Scrambler

cd php-src

./buildconf

make install

cd /polyscripted-php
find . -type d -empty -delete
