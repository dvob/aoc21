#!/bin/sh

awk '
$1 == "forward" {
  h += $2
  depth += aim * $2
}
$1 == "down" { aim += $2 }
$1 == "up" { aim -= $2 }

END {
  printf("horizontal=%d, depth=%d, aim=%d\n", h, depth, aim)
  printf("solution horizontal * depth = %d\n", h * depth)
}
' input
