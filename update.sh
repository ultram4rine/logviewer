#!/bin/bash

go build logviewer.go
mv -f logviewer /usr/local/bin/
cp -r conf.json public/ /var/www/logviewer/
systemctl restart logviewer