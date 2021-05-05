---
title: geocoding
servicename: geocoding
labels: 
- Readme
- Logistics
---
Geocode an address to gps location and the reverse.

# Geocoding Service

The geocoding service provides address to lat lng geocoding as well as the reverse. 
It's useful for building mapping or location based services.

## cURL


### Geocoding Lookup
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Lookup returns a geocoded address including normalized address and gps coordinates
```shell
> curl 'https://api.m3o.com/geocoding/Geocoding/Lookup' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "address": "string",
  "city": "string",
  "country": "string",
  "postcode": "string"
};
# Response
{
  "address": {
    "city": "string",
    "country": "string",
    "line_one": "string",
    "line_two": "string",
    "postcode": "string"
  },
  "location": {
    "latitude": 1,
    "longitude": 1
  }
}
```


### Geocoding Reverse
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Reverse lookup an address from gps coordinates
```shell
> curl 'https://api.m3o.com/geocoding/Geocoding/Reverse' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "latitude": 1,
  "longitude": 1
};
# Response
{
  "address": {
    "city": "string",
    "country": "string",
    "line_one": "string",
    "line_two": "string",
    "postcode": "string"
  },
  "location": {
    "latitude": 1,
    "longitude": 1
  }
}
```


