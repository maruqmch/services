---
title: chats
servicename: chats
labels: 
- Readme
---
Chats is a service for direct messaging

# Chats Service

The chats service enables direct messaging between one or more parties.

## cURL


### Chats CreateChat
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Create a new chat between mulitple users
```shell
> curl 'https://api.m3o.com/chats/Chats/CreateChat' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "The chat ID",
  "user_ids": [
    "string"
  ]
};
# Response
{
  "chat": {
    "created_at": 1,
    "id": "unique id of the chat",
    "user_ids": [
      "string"
    ]
  }
}
```


### Chats ListMessages
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
List messages within a chat
```shell
> curl 'https://api.m3o.com/chats/Chats/ListMessages' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "chat_id": "unique id of the chat",
  "limit": 1,
  "offset": 1,
  "order": "order \"asc\" or \"desc\" (defaults to reverse chronological)"
};
# Response
{
  "messages": [
    {
      "author_id": "user id of the message",
      "chat_id": "chat id the message belongs to",
      "id": "unique id of the message",
      "sent_at": 1,
      "text": "text within the message"
    }
  ]
}
```


### Chats SendMessage
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Send a message to a chat room
```shell
> curl 'https://api.m3o.com/chats/Chats/SendMessage' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "author_id": "string",
  "chat_id": "string",
  "id": "string",
  "text": "string"
};
# Response
{
  "message": {
    "author_id": "user id of the message",
    "chat_id": "chat id the message belongs to",
    "id": "unique id of the message",
    "sent_at": 1,
    "text": "text within the message"
  }
}
```


