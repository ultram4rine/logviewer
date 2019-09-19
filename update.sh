#!/bin/bash

go build main.go
rm /var/www/logviewer/main
cp -r main conf.json public/ /var/www/logviewer/
