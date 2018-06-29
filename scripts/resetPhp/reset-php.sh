#!/bin/bash

cp /php/resetPhp/zend_language_parser.y /php/php-src/Zend/
cp /php/resetPhp/zend_language_scanner.l /php/php-src/Zend/
cp /php/resetPhp/phar.php /php/php-src/ext/phar/phar/phar.php
cp /php/resetPhp/build_precommand.php /php/php-src/ext/phar/build_precommand.php

if [[ $1 == "-revert" ]]; then
        cd /php/php-src/
        make install
fi
