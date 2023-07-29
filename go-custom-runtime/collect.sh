#!/bin/bash
#  Name;Memory;Init;Cold;Billed
FILE=speed.csv

echo "Name;Memory;Init;Cold;Billed" > $FILE
echo "Call functions ========================================="
for f in hello-node hello-py311 hello-runtime-al2 hello-runtime-go
do
    echo "Function:  $f ==============="
    ./dist/coldcalls --lambda $f --times 10 --memory "1024"
done
echo "Collect results ========================================="
for f in hello-node hello-py311 hello-runtime-al2 hello-runtime-go
do
    echo "Function:  $f ==============="
    ./dist/fetchreport --lambda $f >>$FILE
done