from django.shortcuts import render
from django.http import HttpResponse
from .scrapin import scrape as sr
import json
def index(response):
    return HttpResponse("<h1> Home Index </h1>")
def getURL(request):
    result = None
    if request.method == 'POST':
        user_input = request.POST.get('URL_LinkedIn', '')
        result = sr.get_Profile_Fulltext_JSON(user_input)
        with open("profile.json", "w") as json_file:
            json.dump(result, json_file)
    return render(request, 'form.html')


# Create your views here.
