#!/bin/bash

OUT_FILE=./tests/output.txt
EXPECTED=./tests/expected-output.txt
PHP_TEST=./tests/php-tests.php

if [[ -s $EXPECTED  &&  -s $PHP_TEST ]]
then

touch $OUT_FILE;

diff <(/polyscripted-php/bin/php $PHP_TEST) $EXPECTED > $OUT_FILE

SUCCESS='### SUCCESS ###'$'\n'
FAIL='### TEST OUTCOME NOT WHAT EXPECTED###'$'\n'
FAIL_OUT="$(cat $OUT_FILE)"

if [ -s $OUT_FILE ]
then  
	echo "$FAIL""$FAIL_OUT"
else 
	echo $SUCCESS
	rm  $OUT_FILE
fi

else
	echo Missing required test files
fi 
