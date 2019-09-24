#!/bin/bash

go build logviewer.go
rm /var/www/logviewer/logviewer
cp -r conf.json public/ /var/www/logviewer/
cp -f logviewer /usr/local/bin/