import json
def filter_to_string(full_output): #string
    data = json.loads(full_output)
    keys_to_skip = [
        "success", "linkedInIdentifier", "photoUrl", "companyLogo", 
        "linkedInUrl", "linkedinUrl", "linkedinId", "websiteUrl","credits_left","rate_limit_left","schoolLogo",
        "logo","websiteUrl", "companyLogo"
    ]
    filtered_data = recursive_filter(data,keys_to_skip).get('person')
    return json.dumps(filtered_data) #string

def recursive_filter(d, keys_to_skip):
    if isinstance(d, dict):
        return {k: recursive_filter(v, keys_to_skip) for k, v in d.items() if k not in keys_to_skip}
    elif isinstance(d, list):
        return [recursive_filter(i, keys_to_skip) for i in d]
    else:
        return d
def update_provider(full_output): #string
    data = json.loads(full_output)
    credits_left = data.get("credits_left")
    rate_limit_left = data.get("rate_limit_left")
    return [credits_left,rate_limit_left]