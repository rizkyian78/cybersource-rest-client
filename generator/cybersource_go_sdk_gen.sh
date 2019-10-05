#! /bin/bash

# cd %~dp0

echo Delete the previously generated SDK code

rm -rf ../client
rm -rf ../models

echo Command to generate SDK

swagger generate client -A CyberSource -f cybersource-rest-spec.json -t ../

# echo Batch script for changing accept type
# powershell -Command "(Get-Content ..\src\Api\SearchTransactionsApi.js) | ForEach-Object { $_ -replace 'accepts = \[''application/json;charset=utf-8', 'accepts = [''*/*'} | Set-Content ..\src\Api\SearchTransactionsApi.js"

# REM powershell -Command "(Get-Content ..\src\index.js) | ForEach-Object { $_ -replace \"require\('./api/Download([DTXS]{3})Api'\), \", \"\" } | Set-Content ..\src\index.js"

# REM powershell -Command "(Get-Content ..\src\index.js) | ForEach-Object { $_ -replace \"'api/Download([DTXS]{3})Api', \", \"\" } | Set-Content ..\src\index.js"

# REM powershell -Command "(Get-Content ..\src\index.js) | ForEach-Object { $_ -replace \"Download([DTXS]{3})Api, \", \"\" } | Set-Content ..\src\index.js"


# git checkout ..\README.md

# git checkout ..\package.json
