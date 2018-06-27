#!/bin/bash

./polyscript_transformer

OUT_FILE=./tests/outputTransform.txt
TRANSFORMED=./tests/transformed.php
EXPECTED=./tests/php-tests.php

if [[ -s $EXPECTED && -s $TRANSFORMED ]]
then

touch $OUT_FILE;

diff  $TRANSFORMED $EXPECTED > $OUT_FILE

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

