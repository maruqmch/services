---
title: places
servicename: places
labels: 
- Readme
- Logistics
---
Store and search for points of interest

# Places Service

The places API stores points of interest and enables you to search for places nearby or last visited.


## cURL


### Places Last
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/places/Places/Last' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "ids": [
    "string"
  ]
};
# Response
Schema related to #/components/responses/PlacesLastResponse not found
```


### Places Near
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/places/Places/Near' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "latitude": {},
  "longitude": {},
  "radius": {}
};
# Response
Schema related to #/components/responses/PlacesNearResponse not found
```


### Places Read
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/places/Places/Read' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "after": "string",
  "before": "string",
  "ids": [
    "string"
  ]
};
# Response
Schema related to #/components/responses/PlacesReadResponse not found
```


### Places Save
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/places/Places/Save' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "places": [
    {
      "id": "string",
      "latitude": {},
      "longitude": {},
      "metadata": [
        {
          "key": "string",
          "value": "string"
        }
      ],
      "name": "string",
      "timestamp": "string"
    }
  ]
};
# Response
{}
```


