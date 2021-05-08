---
title: feeds
servicename: feeds
labels: 
- Readme
- Headless CMS
---
A single uniform API for crawling and indexing RSS feeds

# Feeds Service

Designed to populate the posts service with RSS feeds from other blogs. Useful for migration or just to get outside content into the posts service.


## cURL


### Feeds Add
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/feeds/Feeds/Add' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "category": "category to add",
  "name": "rss feed name. eg. a16z",
  "url": "rss feed url. eg. http://a16z.com/feed/"
};
# Response
{}
```


### Feeds Entries
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/feeds/Feeds/Entries' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "url": "rss feed url. eg. http://a16z.com/feed/"
};
# Response
{
  "entries": [
    {
      "category": "string",
      "content": "string",
      "date": 1,
      "domain": "string",
      "id": "string",
      "title": "string",
      "url": "string"
    }
  ]
}
```


### Feeds List
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/feeds/Feeds/List' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {};
# Response
{
  "feeds": [
    {
      "category": "category of the feed",
      "name": "rss feed name. eg. a16z",
      "url": "rss feed url. eg. http://a16z.com/feed/"
    }
  ]
}
```


### Feeds Remove
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/feeds/Feeds/Remove' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "name": "rss feed name. eg. a16z"
};
# Response
{}
```


