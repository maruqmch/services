---
title: otp Micro.js
servicename: otp
labels: 
- Micro.js
---

## Micro.js


### Otp Generate
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Generate an OTP (one time pass) code
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/otp/Otp/Generate",
        "micro",
        {
          "id": "unique id, email or user to generate an OTP for"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Otp Validate
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Validate the code
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/otp/Otp/Validate",
        "micro",
        {
          "code": "6 digit one time pass code to validate",
          "id": "unique id, email or user for which the code was generated"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


