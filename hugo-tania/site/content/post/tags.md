---
title: tags
servicename: tags
labels: 
- Readme
- Headless CMS
---
# Tag Service

Tag any resource by saving a tag associated with their respective ID

## cURL


### Tags Add
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/tags/Tags/Add' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "resource_created": 1,
  "resource_id": "string",
  "title": "string",
  "type": "string"
};
# Response
{}
```


### Tags List
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
ListRequest: list either by resource id or type.
 Optionally filter by min or max count.
```shell
> curl 'https://api.m3o.com/tags/Tags/List' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "max_count": 1,
  "min_count": 1,
  "resource_id": "string",
  "type": "string"
};
# Response
{
  "tags": [
    {
      "count": 1,
      "description": "string",
      "slug": "string",
      "title": "string",
      "type": "Type is useful for namespacing and listing across resources,. ie. list tags for posts, customers etc."
    }
  ]
}
```


### Tags Remove
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/tags/Tags/Remove' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "resource_id": "string",
  "title": "string",
  "type": "string"
};
# Response
{}
```


### Tags Update
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/tags/Tags/Update' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "description": "string",
  "title": "string",
  "type": "string"
};
# Response
{}
```


