from django.shortcuts import render
from django.http import HttpResponse
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from concurrent.futures import ThreadPoolExecutor, as_completed
from .models import Scrapin, RelevanceAI, Phantombuster, ProviderManagement
class GetFulltext(APIView):
    def get(self, request):
        # Lấy tham số 'list_url' từ request
        list_url = request.query_params.get('list_url', None)
        
        if list_url is None:
            return Response({'error': 'list_url parameter is required'}, status=status.HTTP_400_BAD_REQUEST)

        # Tách các URL từ chuỗi 'list_url' (giả sử các URL được phân cách bằng dấu phẩy)
        urls = list_url.split(',')
        providers = RelevanceAI.objects.all()
        data = []
        with ThreadPoolExecutor(max_workers=5) as executor: 
            futures = []
            for i in range(3):
                future = executor.submit(providers[i].get_profile_fulltext_jsonstring, urls[i])
                futures.append(future)
            for future in as_completed(futures):
                data.append(future.result())
        # Thực hiện logic xử lý các URL (ở đây chỉ đơn giản là trả về chúng)
        response_data = {'list profile': data}
        
        return Response(response_data, status=status.HTTP_200_OK)
def index(response):
    return HttpResponse("<h1> Home Index </h1>")



# Create your views here.
