#!/bin/bash
echo "Call functions ========================================="
for f in hello-node hello-py311 hello-runtime-al2 hello-runtime-go
do
     echo "Function:  $f ==============="
     ./dist/coldcalls --lambda $f --times 10 --memory "128"
done
