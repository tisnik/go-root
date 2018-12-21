#!/bin/bash

KEYWORDS="
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var"

for keyword in $KEYWORDS
do
    grep -h -v "//" `find . -name *.go` | grep "\b$keyword\b" | wc -l | tr '\n' ' '
    echo $keyword
done | sort -n
