#!/bin/bash

for file in *.php
do
	/polyscripted-php/bin/php $file > "expected_${file}"
done
