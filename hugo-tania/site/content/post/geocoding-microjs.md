---
title: geocoding Micro.js
servicename: geocoding
labels: 
- Micro.js
- Logistics
---

## Micro.js


### Geocoding Lookup
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Lookup returns a geocoded address including normalized address and gps coordinates
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/geocoding/Geocoding/Lookup",
        "micro",
        {
          "address": "string",
          "city": "string",
          "country": "string",
          "postcode": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Geocoding Reverse
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Reverse lookup an address from gps coordinates
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/geocoding/Geocoding/Reverse",
        "micro",
        {
          "latitude": 1,
          "longitude": 1
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


