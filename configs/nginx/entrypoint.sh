#!/bin/sh

envsubst '${PORT},${WEB_PORT},${CLIENT_PORT}' < /etc/nginx/nginx.template.conf > /etc/nginx/nginx.conf

nginx -g 'daemon off;'
