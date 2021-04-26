---
title: routing
servicename: routing
labels: 
- Readme
- Logistics
---
Point to point routing directions

# Routing Service

The routing service provides point to point directions

## Config

The following config must be set

- **routing.mode**: "google" or "osrm"
- **google.api.key**: google maps api key if using google
- **osrm.api.address"**: location of osrm api

## cURL


### Routing Eta
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/routing/Routing/Eta' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "destination": {
    "latitude": 1,
    "longitude": 1
  },
  "origin": {
    "latitude": 1,
    "longitude": 1
  },
  "speed": 1,
  "type": "type of transport e.g car, foot, bicycle"
};
# Response
{
  "duration": 1
}
```


### Routing Route
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/routing/Routing/Route' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "destination": {
    "latitude": 1,
    "longitude": 1
  },
  "origin": {
    "latitude": 1,
    "longitude": 1
  }
};
# Response
{
  "waypoints": [
    {
      "location": {
        "latitude": 1,
        "longitude": 1
      },
      "name": "street name or related reference"
    }
  ]
}
```


