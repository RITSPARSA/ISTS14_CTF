#!/bin/bash

docker build -t flask-app .
docker run -it -e WERKZEUG_DEBUG_PIN="569-122-328" -p 8080:8080 flask-app

