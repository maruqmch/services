---
title: routing Micro.js
servicename: routing
labels: 
- Micro.js
- Logistics
---

## Micro.js


### Routing Eta
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/routing/Routing/Eta",
        "micro",
        {
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
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Routing Route
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/routing/Routing/Route",
        "micro",
        {
          "destination": {
                    "latitude": 1,
                    "longitude": 1
          },
          "origin": {
                    "latitude": 1,
                    "longitude": 1
          }
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


