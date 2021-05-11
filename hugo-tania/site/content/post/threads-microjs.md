---
title: threads Micro.js
servicename: threads
labels: 
- Micro.js
---

## Micro.js


### Threads CreateMessage
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/threads/Threads/CreateMessage",
        "micro",
        {
          "author_id": "string",
          "id": "string",
          "text": "string",
          "thread_id": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Threads CreateThread
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/threads/Threads/CreateThread",
        "micro",
        {
          "group_id": "string",
          "topic": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Threads DeleteThread
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/threads/Threads/DeleteThread",
        "micro",
        {
          "id": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Threads ListMessages
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/threads/Threads/ListMessages",
        "micro",
        {
          "limit": 1,
          "offset": 1,
          "order": "string",
          "thread_id": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Threads ListThreads
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/threads/Threads/ListThreads",
        "micro",
        {
          "group_id": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Threads ReadThread
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/threads/Threads/ReadThread",
        "micro",
        {
          "group_id": "string",
          "id": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Threads RecentMessages
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/threads/Threads/RecentMessages",
        "micro",
        {
          "limit_per_thread": 1,
          "thread_ids": [
                    "string"
          ]
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Threads UpdateThread
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/threads/Threads/UpdateThread",
        "micro",
        {
          "id": "string",
          "topic": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


