#!/bin/sh
if [ -f "$FILE" ]; then
    echo "exist"
else
    echo "does not exist"
fi