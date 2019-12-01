# =grafana-telegram-proxy=
sending Grafana alarm messages to Telegram via the Webhook channel

[![Build Status](https://api.travis-ci.org/maslick/grafana-telegram-proxy.svg)](https://travis-ci.org/maslick/grafana-telegram-proxy)
[![Dockerhub](https://img.shields.io/badge/image%20size-2.5MB-blue.svg)](https://hub.docker.com/r/maslick/grafana-telegram-proxy)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)


## Motivation
Sending Alert notifications via Teleram fails when Grafana instance runs on a server without direct access to Telegram API (e.g. Russian ISPs block Telegram servers). One way to solve this is to leverage the webhook notification channel and run an HTTP proxy that listens to Grafana webhooks, parses the payload and sends the notifications to Telegram API from a DMZ (or better say ``de-Russia-nized zone`` - any cloud provider, i.e. Heroku, GAE, EC2, foreign VPS, baremetal, etc.)


## Features
* Grafana alert notifications to Telegram via the webhook notification channel
* Lightweight static binary: ~2.5 MB zipped
* Cloud-native friendly: Docker + k8s
* Secure: Basic authentication (optional)

![Grafana](screenshot.png)

## Installation
```zsh
$ go test
$ go build -ldflags="-s -w"
$ go build -ldflags="-s -w" && upx grafana-telegram-proxy
```

## Usage
* Without authentication:
```zsh
$ export BOT_TOKEN=1234567890abcdef
$ export CHAT_ID=-12345
$ ./grafana-telegram-proxy
Starting server on port 8080 ...
```

* With Basic authentication:
```zsh
$ export BOT_TOKEN=1234567890abcdef
$ export CHAT_ID=-12345
$ export USERNAME=maslick
$ export PASSWORD=12345
$ export PORT=4000
$ ./grafana-telegram-proxy
Starting server on port 4000 ...
```

## Docker
```zsh
$ docker build -t maslick/grafana-telegram-proxy .
$ docker run -d \
   -e BOT_TOKEN=1234567890abcdef \
   -e CHAT_ID=-12345 \
   -p 8081:8080 \
   maslick/grafana-telegram-proxy

$ docker run -d \
   -e BOT_TOKEN=1234567890abcdef \
   -e CHAT_ID=-12345 \
   -e USERNAME=maslick \
   -e PASSWORD=12345 \
   -p 8082:8080 \
   maslick/grafana-telegram-proxy
```

## Kubernetes
```zsh
$ kubectl apply -f k8s
$ kubectl set env deploy grafana-telegram-proxy \
   BOT_TOKEN=1234567890abcdef \
   CHAT_ID=-12345 \
   USERNAME=maslick \
   PASSWORD=12345
```

## Heroku
```zsh
$ git clone https://github.com/maslick/grafana-telegram-proxy.git
$ cd grafana-telegram-proxy

$ export HEROKU_APP_NAME=hello-world-app
$ heroku login
$ heroku create $HEROKU_APP_NAME
$ git push heroku master
$ heroku config:set BOT_TOKEN=$BOT_TOKEN
$ heroku config:set CHAT_ID=$CHAT_ID
$ heroku config:set USERNAME=$USERNAME
$ heroku config:set PASSWORD=$PASSWORD
$ open https://$HEROKU_APP_NAME.herokuapp.com/health
```
