from rest_framework import status
from rest_framework.generics import GenericAPIView
from rest_framework.response import Response

from core import throttling


class DjangoThrottlingAPIView(GenericAPIView):
    lookup_field = 'id'
    throttle_classes = [throttling.ConcurrencyThrottleApiKey]

    def get(self, request):
        return Response("okk", status=status.HTTP_200_OK)
