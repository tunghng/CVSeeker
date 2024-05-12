from django.urls import path

from . import views
urlpatterns = [ 
    path("home", view=views.index, name="index"),
    path("", view=views.getURL, name="getURL")
]