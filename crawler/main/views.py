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
            return Response({'error': 'Too many urls (max = 6)'}, status=status.HTTP_400_BAD_REQUEST)
        
        # Init useable providers each type
        providers = RelevanceAI.objects.all() # greater than 4 credits

        number_providers = len(providers) # default = 6 only for demo
        threads_per_link = int(number_providers/number_links)
        profiles = [] # Store resutls

        # Multithread for crawling
        with ThreadPoolExecutor(max_workers=6) as executor:
            futures = []
            for i in range(number_links):
                for j in range(threads_per_link):
                    future = executor.submit(providers[i*threads_per_link + j].get_profile_fulltext_jsonstring, urls[i])
                    futures.append((i, future)) # Store order of each thread and link for tracking then
    
            successed = []
            for i, future in futures:
                result = future.result()
                if(result[0] == 200 or result[0] == 500):
                    # Kiểm tra link này đã chạy thành công và được lưu chưa, nếu chưa thì lưu, không thì skip để kết quả trả về không bị lặp
                    if i not in successed: 
                        profile = {
                            "content" : result[1],
                            "link" : urls[i],
                        }
                        successed.append(i)
                        profiles.append(profile)

        response_data = {'profiles': profiles}
        #Response to client
        return Response(response_data, status=status.HTTP_200_OK)

def index(response):
    return HttpResponse("<h1> \"Death is like the wind, always by my side\" - Yasuo (15p gg) </h1>")

