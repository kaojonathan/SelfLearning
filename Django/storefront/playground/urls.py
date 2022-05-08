# the urls to which the view functions get mapped to
from django.urls import path
from . import views # from the curr folder import the view functions

# URLConf
# array of url pattern objects
# use path function to create them
urlpatterns = [
    # url of playground/hello mapped to say_hello view function
    path('hello/', views.say_hello)
]
# need to import the URLConf into the main app urls.py