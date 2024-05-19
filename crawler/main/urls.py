from django.urls import path

from . import views
urlpatterns = [ 
    path("", view=views.index, name="index"),
    path("api/getfulltext/", view=views.GetFulltext.as_view(), name="Get Full Text")
]