from django.shortcuts import render
from django.http import HttpResponse
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from concurrent.futures import ThreadPoolExecutor, as_completed
from .models import Scrapin, RelevanceAI, Phantombuster, ProviderManagement

def get_provider():
    # Get all providers not be ban
    enable_services = ProviderManagement.objects.filter(enable=True)
    useable_providers = [] # Store providers useable
    position_providers = [] # Store index of them in useable providers list
    for service in enable_services:
        if(service.id == 1):
            scrapin_providers = Scrapin.objects.filter(remain_credits__gt=0)
            position_providers.append(len(useable_providers) + len(scrapin_providers))
            useable_providers += scrapin_providers
        if(service.id == 2):
            rele_providers = RelevanceAI.objects.filter(remain_credits__gt=0)
            position_providers.append(len(useable_providers) + len(rele_providers))
            useable_providers += rele_providers
        if(service.id == 3):
            phantom_providers = Phantombuster.objects.filter(remain_time__gt=10) 
            position_providers.append(len(useable_providers) + len(phantom_providers))
            useable_providers += phantom_providers
    return [useable_providers, position_providers, enable_services]

def crawl_profiles(providers, failed_providers, successed_providers, times, urls):
    profiles = [] # Store result 
    # Multithread for crawling with maximum 6 thread
    with ThreadPoolExecutor(max_workers=6) as executor:
        futures = []
        running_providers = []
        for i in range(len(urls)):
            j = 0
            while j < len(providers):
                if (j not in failed_providers) and (j not in running_providers):
                    future = executor.submit(providers[j].get_profile_fulltext_jsonstring, urls[i])
                    futures.append((i, future)) # Store result with order each thread
                    running_providers.append(j)
                    break
                j += 1
                if j == len(providers):
                    return []
        # Get results
        for i, future in futures:
            result = future.result()
            if(result[0] == 200 or result[0] == 500):
                # Save successed providers
                profile = {
                    "content" : result[1],
                    "link" : urls[i],
                }
                profiles.append(profile)
                successed_providers.append(i)
            else:
                # Save failed providers for try again
                failed_providers.append(running_providers[i])
    if(len(failed_providers) == 0):
        return profiles
    else:
        # when trying times less than 3
        if(times <= 3):
            resend_urls = [urls[pos] for pos in failed_providers]
            profiles.append(crawl_profiles(providers, failed_providers, successed_providers, times+1, resend_urls))
        else:
            return profiles

            

def update_failed_providers(services, failed_list, successed_list, positions):
    for fail_index in failed_list:
        for i in range(len(positions)):
            if fail_index < positions[i]:
                services[i].number_errors += 1
                services[i].save()
                break
    for success_index in successed_list:
        for i in range(len(positions)):
            if success_index < positions[i]:
                if services[i].number_errors > 0:
                    services[i].number_errors -= 1
                    services[i].save()
                    break
    for service in services:
        if service.number_errors >= service.error_limit:
            service.enable = False
            service.save()


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
        if(number_links > 6):
            return Response({'error': 'Too many urls (max = 6)'}, status=status.HTTP_400_BAD_REQUEST)
        
        providers, position_providers, enable_services = get_provider()

        # Checking providers's responsiveness to urls
        if len(providers) < number_links:
            return Response({'error': 'Dont enough scraper objects, please get less urls'}, status=status.HTTP_400_BAD_REQUEST)
        # Store index of failed providers
        failed_providers = []
        # Store index of successed provider
        successed_providers = []

        profiles = crawl_profiles(providers, failed_providers, successed_providers, 0, urls)

        # Update enable of services after requests base on fail and success rate
        update_failed_providers(enable_services,failed_providers, successed_providers, position_providers)
        
        response_data = {'profiles': profiles}
        #Response to client
        return Response(response_data, status=status.HTTP_200_OK)
    


def index(response):
    return HttpResponse("<h1> Home </h1>")



# Create your views here.
