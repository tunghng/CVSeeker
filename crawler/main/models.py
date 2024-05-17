from django.db import models
import requests
import json
from .filterdata import relevance, scrapin

class Scrapin(models.Model):
    email_name = models.CharField(default="test", max_length=50)
    remain_credits = models.IntegerField()
    api_key = models.CharField(max_length=100)
    def __str__(self) -> str:
        return self.email_name
    def get_profile_fulltext_jsonstring(self, linkedinUrl):
        url = "https://api.scrapin.io/enrichment/profile"
        params = {
            "linkedinUrl": linkedinUrl,
            "apikey": self.api_key
        }
        try:
            response = requests.get(url, params=params)
            code = response.status_code
            if(code == 402):
                return 'Dont have enough credits'
            if(code == 403 or code == 401):
                return 'API wrong'
            if(code == 500):
                return 'LinkedIn URL wrong'
            if(code == 400):
                return 'Missing parameters'
            credits_and_limit = scrapin.update_provider(response.text)
            self.remain_credits = credits_and_limit[0] 
            return scrapin.filter_to_string(response.text)
        except:
            return 'failed'

class RelevanceAI(models.Model):
    email_name = models.CharField(default="test", max_length=50)
    remain_credits = models.IntegerField()
    project_id = models.CharField(max_length=100)
    api_key = models.CharField(max_length=100)
    endpoint = models.CharField(max_length=200)
    def get_profile_fulltext_jsonstring(self, linked_url):
        headers = {
            "Content-Type": "application/json",
            "Authorization": f"{self.project_id}:{self.api_key}",    
        }
        data = {
            "params": {
                "url": linked_url
            },
            "project": f"{self.project_id}"
        }
        try:
            response = requests.post(self.endpoint, json=data, headers=headers)
            self.remain_credits = self.remain_credits - 4
            response_json = response.json()
            if(response.status_code == 200):
                status = response_json.get('status')
                errors = response_json.get('errors')
                if(status == 'failed' or len(errors) > 0):
                    return 'failed'
            
            return relevance.filter_to_string(response.text)
        except:
            return 'failed'
        
    def __str__(self) -> str:
        return self.email_name

class Phantombuster(models.Model):
    email = models.CharField(max_length=100)
    password = models.CharField(max_length=100)
    cookie  = models.CharField(max_length=300)
    link_to_setup = models.CharField(max_length=300)
    link_to_launch = models.CharField(max_length=300)
    def __str__(self) -> str:
        return self.email

class ProviderManagement(models.Model):
    name = models.CharField(max_length=50)
    enable = models.BooleanField(default=True)
    number_account = models.IntegerField()
    error_limit = models.IntegerField(default=5)
    def __str__(self) -> str:
        return self.name
