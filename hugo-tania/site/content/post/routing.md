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

## Features

- **Directions** - Turn by turn driving directions
- **ETAs** - Estimated time of arrival from origin to destination
- **Routes** - GPS point based routes with distance and duration

## Config

The following config must be set

- **routing.mode**: "google" or "osrm"
- **routing.address"**: location of osrm api if used
- **google.apikey**: google maps api key if using google

## cURL


### Routing Directions
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/routing/Routing/Directions' \
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
  "directions": [
    {
      "distance": 1,
      "duration": 1,
      "instruction": "human readable instruction",
      "intersections": [
        {
          "bearings": [
            1
          ],
          "location": {
            "latitude": 1,
            "longitude": 1
          }
        }
      ],
      "maneuver": {
        "action": "string",
        "bearing_after": 1,
        "bearing_before": 1,
        "direction": "string",
        "location": {
          "latitude": 1,
          "longitude": 1
        }
      },
      "name": "street name or location",
      "reference": "alternative reference"
    }
  ],
  "distance": 1,
  "duration": 1,
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
  "distance": 1,
  "duration": 1,
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


