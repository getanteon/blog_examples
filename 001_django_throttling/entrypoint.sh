#!/bin/bash

# cd throttling && gunicorn --workers 1 --bind 0.0.0.0:9018 throttling.wsgi
cd throttling && gunicorn --workers 6 --bind 0.0.0.0:9018 throttling.wsgi
