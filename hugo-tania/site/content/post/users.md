---
title: users
servicename: users
labels: 
- Readme
- Backend
---
User management and authentication

# Users Service

The users service provides user management and authentication


## cURL


### Users Create
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/Create' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "email": "string",
  "id": "uuid",
  "password": "string",
  "username": "alphanumeric user or org"
};
# Response
{}
```


### Users Delete
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/Delete' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "string"
};
# Response
{}
```


### Users Login
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/Login' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "email": "string",
  "password": "string",
  "username": "string"
};
# Response
{
  "session": {
    "created": 1,
    "email": "string",
    "expires": 1,
    "id": "string",
    "username": "string"
  }
}
```


### Users Logout
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/Logout' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "session_id": "string"
};
# Response
{}
```


### Users Read
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/Read' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "string"
};
# Response
{
  "user": {
    "created": 1,
    "email": "string",
    "id": "uuid",
    "updated": 1,
    "username": "alphanumeric user or org"
  }
}
```


### Users ReadSession
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/ReadSession' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "session_id": "string"
};
# Response
{
  "session": {
    "created": 1,
    "email": "string",
    "expires": 1,
    "id": "string",
    "username": "string"
  }
}
```


### Users Search
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/Search' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "email": "string",
  "limit": 1,
  "offset": 1,
  "username": "string"
};
# Response
{
  "users": [
    {
      "created": 1,
      "email": "string",
      "id": "uuid",
      "updated": 1,
      "username": "alphanumeric user or org"
    }
  ]
}
```


### Users Update
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/Update' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "email": "string",
  "id": "uuid",
  "username": "alphanumeric user or org"
};
# Response
{}
```


### Users UpdatePassword
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/users/Users/UpdatePassword' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "confirm_password": "string",
  "new_password": "string",
  "old_password": "string",
  "user_id": "string"
};
# Response
{}
```


