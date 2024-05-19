from django.shortcuts import render
from django.http import HttpResponse
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from concurrent.futures import ThreadPoolExecutor, as_completed
from .models import Scrapin, RelevanceAI, Phantombuster, ProviderManagement

class GetFulltext(APIView):
    def get(self, request):
        # Get parameters list_url from request
        list_url = request.query_params.get('list_url', None)
        
        if list_url is None:
            return Response({'error': 'list_url parameter is required'}, status=status.HTTP_400_BAD_REQUEST)

        # Split and store in list 
        urls = list_url.split(',')
        number_links = len(urls)

        # Checking number urls
        if(number_links > 5):
            return Response({'error': 'Too many urls (max = 5)'}, status=status.HTTP_400_BAD_REQUEST)
        
        # Init useable providers each type
        rele_providers = RelevanceAI.objects.filter(remain_credits__gt=0) # greater than 0 credits
        scrapin_providers = Scrapin.objects.filter(remain_credits__gt=0) # greater than 0 credits
        phantom_providers = Phantombuster.objects.filter(remain_time__gt=10) # greater than 10 minutes

        # Get the number of providers each type
        number_rele = len(rele_providers)
        number_scrapin = len(scrapin_providers)
        number_phantom = len(phantom_providers)

        # Checking providers's responsiveness to urls
        providers = [None] * number_links
        if(number_rele >= number_links):
            providers = rele_providers
        else:
            if(number_rele + number_scrapin >= number_links):
                providers = rele_providers + scrapin_providers[:(number_links - number_rele)]
            else:
                if(number_rele + number_scrapin + number_phantom >= number_links):
                    providers = rele_providers + scrapin_providers + phantom_providers[:(number_links - number_rele - number_scrapin)]
                else:
                    return Response({'error': 'Dont enough scraper objects, please get less urls'}, status=status.HTTP_400_BAD_REQUEST)
        
        profiles = [] # Store resutls
        failed_providers = [] # Store failed provider for handling then

        # Multithread for crawling
        with ThreadPoolExecutor(max_workers=5) as executor:
            futures = []
            for i in range(len(urls)):
                future = executor.submit(providers[i].get_profile_fulltext_jsonstring, urls[i])
                futures.append((i, future)) # Store result with order each thread
            
            # Get results
            for i, future in futures:
                result = future.result()
                if(result[0] == 200 or result[0] == 500 or result[0] == 402):
                    # Save successed providers
                    profile = {
                        "content" : result[1],
                        "fileBytes" : urls[i],
                    }
                    profiles.append(profile)
                else:
                    # Save failed providers for handling then
                    failed_providers.append(i)
        response_data = {'resumes': profiles}

        #Response to client
        return Response(response_data, status=status.HTTP_200_OK)
    


def index(response):
    return HttpResponse("<h1> \"Death is like the wind, always by my side\" - Yasuo (15p gg) </h1>")



# Create your views here.
