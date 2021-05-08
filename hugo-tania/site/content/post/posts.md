---
title: posts
servicename: posts
labels: 
- Readme
- Headless CMS
---
Build a blog or the foundations of a headless CMS with posts

# Post Service

Posts is the foundation of a headless CMS, storing blog posts with their metadata and enabling simple retrieval and querying.


## cURL


### Posts Delete
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/posts/Posts/Delete' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "string"
};
# Response
{}
```


### Posts Index
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/posts/Posts/Index' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "limit": 1,
  "offset": 1
};
# Response
{
  "posts": [
    {
      "author": "string",
      "content": "string",
      "created": 1,
      "id": "string",
      "image": "string",
      "metadata": [
        {
          "key": "string",
          "value": "string"
        }
      ],
      "slug": "string",
      "tags": [
        "string"
      ],
      "title": "string",
      "updated": 1
    }
  ]
}
```


### Posts Query
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
Query posts. Acts as a listing when no id or slug provided.
 Gets a single post by id or slug if any of them provided.
```shell
> curl 'https://api.m3o.com/posts/Posts/Query' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "id": "string",
  "limit": 1,
  "offset": 1,
  "slug": "string",
  "tag": "string"
};
# Response
{
  "posts": [
    {
      "author": "string",
      "content": "string",
      "created": 1,
      "id": "string",
      "image": "string",
      "metadata": [
        {
          "key": "string",
          "value": "string"
        }
      ],
      "slug": "string",
      "tags": [
        "string"
      ],
      "title": "string",
      "updated": 1
    }
  ]
}
```


### Posts Save
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/posts/Posts/Save' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "content": "string",
  "id": "string",
  "image": "string",
  "metadata": [
    {
      "key": "string",
      "value": "string"
    }
  ],
  "slug": "string",
  "tags": [
    "string"
  ],
  "timestamp": 1,
  "title": "string"
};
# Response
{
  "id": "string"
}
```


