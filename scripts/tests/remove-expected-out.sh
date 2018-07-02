#!/bin/bash

for file in *
do  
	if [[ $file == expected* ]]
	then
		rm $file
	fi

done
