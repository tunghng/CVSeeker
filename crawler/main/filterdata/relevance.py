import json
def filter_to_string(full_output):
    data = json.loads(full_output)
    keys_to_skip = [
        "company_linkedin_url", "company_logo_url", "school_linkedin_url", "company_public_url",
        "linkedin_url", "profile_image_url","company_website"
    ]
    filtered_data = recursive_filter(data,keys_to_skip).get('output').get('linkedin_full_data')
    return json.dumps(filtered_data)


def recursive_filter(d, keys_to_skip):
    if isinstance(d, dict):
        return {k: recursive_filter(v, keys_to_skip) for k, v in d.items() if k not in keys_to_skip}
    elif isinstance(d, list):
        return [recursive_filter(i, keys_to_skip) for i in d]
    else:
        return d