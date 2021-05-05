---
title: invites
servicename: invites
labels: 
- Readme
---
Manage and send user and group invites

# Invites Service

The invites services allows you to create and manage invites for users and groups by providing invite codes.


## cURL


### Invites Create
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/invites/Invites/Create' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "email": "string",
  "group_id": "string"
};
# Response
{
  "invite": {
    "code": "string",
    "email": "string",
    "group_id": "string",
    "id": "string"
  }
}
```


### Invites Delete
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/invites/Invites/Delete' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "string"
};
# Response
{}
```


### Invites List
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/invites/Invites/List' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "email": "string",
  "group_id": "string"
};
# Response
{
  "invites": [
    {
      "code": "string",
      "email": "string",
      "group_id": "string",
      "id": "string"
    }
  ]
}
```


### Invites Read
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/invites/Invites/Read' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "code": "string",
  "id": "string"
};
# Response
{
  "invite": {
    "code": "string",
    "email": "string",
    "group_id": "string",
    "id": "string"
  }
}
```


