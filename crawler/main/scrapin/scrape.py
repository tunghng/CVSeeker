import requests
import json

def get_Profile_Fulltext_JSON(linkedinUrl):
    url = "https://api.scrapin.io/enrichment/profile"
    apikey = "sk_live_663f41863a4e8207e2e32013_key_8pf359rr5s8"
    params = {
        "linkedinUrl": linkedinUrl,
        "apikey": apikey
    }
    response = requests.get(url, params=params)
    return json.loads(response.text)
