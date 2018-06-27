#!/bin/bash

for file in *.php
do
	touch "expected_${file}"
	#/polyscripted-php/bin/php $file > "${file}_expected"
done
