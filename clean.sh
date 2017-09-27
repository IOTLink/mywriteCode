#!/bin/bash
for f in `find ./   -path ".git" -prune -o -type f`
do 
    if [ $f[3] == $'.git' ]
	    echo $f
    fi
    /usr/bin/find $f
done
