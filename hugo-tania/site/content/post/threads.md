---
title: threads
servicename: threads
labels: 
- Readme
---
Threaded conversations

# Threads Service

Threads provides threaded conversations as a service grouped by topics.


## cURL


### Threads CreateMessage
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/threads/Threads/CreateMessage' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "author_id": "string",
  "id": "string",
  "text": "string",
  "thread_id": "string"
};
# Response
{
  "message": {
    "author_id": "string",
    "id": "string",
    "sent_at": "string",
    "text": "string",
    "thread_id": "string"
  }
}
```


### Threads CreateThread
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/threads/Threads/CreateThread' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "group_id": "string",
  "topic": "string"
};
# Response
{
  "thread": {
    "created_at": "string",
    "group_id": "string",
    "id": "string",
    "topic": "string"
  }
}
```


### Threads DeleteThread
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/threads/Threads/DeleteThread' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "string"
};
# Response
{}
```


### Threads ListMessages
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/threads/Threads/ListMessages' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "limit": 1,
  "offset": 1,
  "order": "string",
  "thread_id": "string"
};
# Response
{
  "messages": [
    {
      "author_id": "string",
      "id": "string",
      "sent_at": "string",
      "text": "string",
      "thread_id": "string"
    }
  ]
}
```


### Threads ListThreads
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/threads/Threads/ListThreads' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "group_id": "string"
};
# Response
{
  "threads": [
    {
      "created_at": "string",
      "group_id": "string",
      "id": "string",
      "topic": "string"
    }
  ]
}
```


### Threads ReadThread
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/threads/Threads/ReadThread' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "group_id": "string",
  "id": "string"
};
# Response
{
  "thread": {
    "created_at": "string",
    "group_id": "string",
    "id": "string",
    "topic": "string"
  }
}
```


### Threads RecentMessages
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/threads/Threads/RecentMessages' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "group_id": "string",
  "limit_per_thread": 1,
  "thread_ids": [
    "string"
  ]
};
# Response
{
  "messages": [
    {
      "author_id": "string",
      "id": "string",
      "sent_at": "string",
      "text": "string",
      "thread_id": "string"
    }
  ]
}
```


### Threads UpdateThread
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/threads/Threads/UpdateThread' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "string",
  "topic": "string"
};
# Response
{
  "thread": {
    "created_at": "string",
    "group_id": "string",
    "id": "string",
    "topic": "string"
  }
}
```


