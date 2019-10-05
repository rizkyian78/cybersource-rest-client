#! /bin/bash

echo Delete the previously generated SDK code

rm -rf ../client
rm -rf ../models

echo Command to generate SDK

swagger generate client -A Cybersource -f cybersource-rest-spec.json -t ../
