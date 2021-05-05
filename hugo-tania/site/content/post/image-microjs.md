---
title: image Micro.js
servicename: image
labels: 
- Micro.js
---

## Micro.js


### Image Convert
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Convert an image from one format (jpeg, png etc.) to an other either on the fly (from base64 to base64),
 or by uploading the conversion result.
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/image/Image/Convert",
        "micro",
        {
          "base64": "base64 encoded image to resize,. ie. \"data:image/png;base64, iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg==\"",
          "image_id": "output name of the image including extension, ie. \"cat.png\"",
          "output_url": true,
          "url": "url of the image to resize"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Image Resize
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Resize an image on the fly without storing it (by sending and receiving a base64 encoded image), or resize and upload depending on parameters.
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/image/Image/Resize",
        "micro",
        {
          "base64": "base64 encoded image to resize,. ie. \"data:image/png;base64, iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg==\"",
          "height": 1,
          "image_id": "output name of the image including extension, ie. \"cat.png\"",
          "output_url": true,
          "url": "url of the image to resize",
          "width": 1
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Image Upload
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Upload an image by either sending a base64 encoded image to this endpoint or a URL.
 To resize an image before uploading, see the Resize endpoint.
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/image/Image/Upload",
        "micro",
        {
          "base64": "Base64 encoded image to upload,. ie. \"data:image/png;base64, iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg==\"",
          "image_id": "Output name of the image including extension, ie. \"cat.png\"",
          "url": "URL of the image to upload"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


