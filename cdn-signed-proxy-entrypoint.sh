#!/bin/sh

export KEY

envsubst '${KEY}' < /etc/nginx/sites-available/pre-signed-url-reverse-proxy.conf.template > /etc/nginx/sites-available/pre-signed-url-reverse-proxy.conf

rm /etc/nginx/sites-enabled/* && ln -s '/etc/nginx/sites-available/pre-signed-url-reverse-proxy.conf' /etc/nginx/sites-enabled/pre-signed-url-reverse-proxy.conf

exec "$@"
