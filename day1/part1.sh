#!/bin/sh

awk 'prev && $1 > prev { inc++ } { prev = $1 } END{ print inc }' input
