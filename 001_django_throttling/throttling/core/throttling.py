from rest_framework.throttling import SimpleRateThrottle
from django.core.cache import caches

class ConcurrencyThrottleApiKey(SimpleRateThrottle):
    cache = caches['alternate'] # use redis cache. Delete if you want to use default Django cache
    rate = "1/s"

    def get_cache_key(self, request, view):
        return request.query_params['api_key']
