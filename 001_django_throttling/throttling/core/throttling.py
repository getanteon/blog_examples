from rest_framework.throttling import SimpleRateThrottle

class ConcurrencyThrottleApiKey(SimpleRateThrottle):
    rate = "1/s"

    def get_cache_key(self, request, view):
        return request.query_params['api_key']
