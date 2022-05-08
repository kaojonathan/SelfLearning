from django.shortcuts import render
from django.http import HttpResponse

# Create your views here.

# debug practice with breakpoints and vscode
def calculate():
    x = 1
    y = 2
    return x

# view function: request -> response
#   request handler

def say_hello(request):
    # Ex. Pull data from db, transform data, send an email...
    # return HttpResponse('Hello World')
    # render a html page instead
    x = calculate()
    return render(request, 'hello.html', {'name': 'Justin'}) 

# map the view function to an url
# when a request is called at that url, the view function is called





