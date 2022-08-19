from django.urls import path
import core.views as core_views

urlpatterns = [
    path('', core_views.DjangoThrottlingAPIView.as_view(), name="throtling"),
]
