#!/bin/bash

files=$(find input.*)

for file in $files
do
    go run lets.go < $file
done
