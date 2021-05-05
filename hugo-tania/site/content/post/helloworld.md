---
title: helloworld
servicename: helloworld
labels: 
- Readme
- Backend
---
A simple hello world service

# Helloworld Service

Helloworld simply returns a personalized message in response to a name.

## cURL


### Helloworld Call
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Call returns a personalised "Hello $name" response
```shell
> curl 'https://api.m3o.com/helloworld/Helloworld/Call' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "name": "string"
};
# Response
{
  "msg": "string"
}
```


