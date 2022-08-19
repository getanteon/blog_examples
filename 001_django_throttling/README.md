# Django Throttling Example with API Key and Make a Load Test with Ddosify

## Create Django Environment

```bash
python3 -m venv env
```

## Activate Django Environment

```bash
source env/bin/activate
```

## Create requirements.txt

```txt
Django==4.0.7
djangorestframework==3.13.1
gunicorn==20.1.0
redis==4.3.4
```

## Install Django Dependencies

```bash
pip3 install -r requirements.txt
```

## Create Django Project and Application

```bash
django-admin startproject throttling
python3 manage.py startapp core
```

## Update Django Settings

Settings file: `throttling/settings.py`

📌 Add `rest_framework` into `INSTALLED_APPS` list.

📌 Add `REST_FRAMEWORK` dictionary:

```python
REST_FRAMEWORK = {
    'DEFAULT_RENDERER_CLASSES': (
        'rest_framework.renderers.JSONRenderer',
    ),
}
```

## Add a basic endpoint

📌 Add a basic endpoint into `core/views.py`

```python
from rest_framework import status
from rest_framework.generics import GenericAPIView
from rest_framework.response import Response

from core import throttling


class DjangoThrottlingAPIView(GenericAPIView):
    throttle_classes = [throttling.ConcurrencyThrottleApiKey]

    def get(self, request):
        return Response(status=status.HTTP_200_OK)

```

📌 Create a custom throttling class: `core/throttling.py`

```python
from rest_framework.throttling import SimpleRateThrottle

class ConcurrencyThrottleApiKey(SimpleRateThrottle):
    rate = "1/s"

    def get_cache_key(self, request, view):
        return request.query_params['api_key']
```

📌 Add URL into: `core/urls.py`

```python
from django.urls import path
import core.views as core_views

urlpatterns = [
    path('', core_views.DjangoThrottlingAPIView.as_view(), name="throtling"),
]
```

📌 Update root urls: `throttling/urls.py`

```python
from django.urls import path

urlpatterns = [
    path('', include('core.urls')),
]
```

## Run gunicorn webserver

```bash
gunicorn --workers 1 --bind 0.0.0.0:9018 throttling.wsgi
```
