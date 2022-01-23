#!/bin/sh

awk '
{ 
  s[NR] += $1
  s[NR-1] += $1
  s[NR-2] += $1
}
END {
  for (i = 1; i <= NR - 3; i++) {
    if ( s[i] > s[i-1] ) {
      inc++
    }
  }
  print inc
}

' input
