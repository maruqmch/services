---
title: chats Micro.js
servicename: chats
labels: 
- Micro.js
---

## Micro.js


### Chats CreateChat
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Create a new chat between mulitple users
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/chats/Chats/CreateChat",
        "micro",
        {
          "id": "The chat ID",
          "user_ids": [
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


### Chats ListMessages
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
List messages within a chat
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/chats/Chats/ListMessages",
        "micro",
        {
          "chat_id": "unique id of the chat",
          "limit": 1,
          "offset": 1,
          "order": "order \"asc\" or \"desc\" (defaults to reverse chronological)"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Chats SendMessage
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Send a message to a chat room
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/chats/Chats/SendMessage",
        "micro",
        {
          "author_id": "string",
          "chat_id": "string",
          "id": "string",
          "text": "string"
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


