#!/bin/sh

awk '
$1 == "forward" { x += $2 }
$1 == "down" { y += $2 }
$1 == "up" { y -= $2 }

END {
  printf("x=%d, depth=%d\n", x, y)
  printf("solution x * depth = %d\n", x * y)
}
' input
