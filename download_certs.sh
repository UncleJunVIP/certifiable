#!/bin/sh

curl https://ccadb.my.salesforce-sites.com/mozilla/IncludedRootsPEMTxt?TrustBitsInclude=Websites -o certificates.crt

printf "Downloaded!"
