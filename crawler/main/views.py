from django.shortcuts import render
from django.http import HttpResponse
from .scrapin.scrape import ScrapinCrawler
import json

def index(response):
    return HttpResponse("<h1> Home Index </h1>")
def getURL(request):
    result = None
    if request.method == 'POST':
        user_input = request.POST.get('URL_LinkedIn', '')
        crawler = ScrapinCrawler()
        crawler.sendData(LinkedInURL=user_input)
        result = crawler.getData()
        with open("profile.json", "w") as json_file:
            json.dump(result, json_file)
    return render(request, 'main/form.html')


# Create your views here.
