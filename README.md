# ELK POC

This is a proof of concept to do log centralize with elk stack:

- elasticsearh: for searching and collect the log
- logstash: for log pre-processing before insert log to elasticsearch
- filebeat: detect and collect the log from application then send to logstash
- kibana: dashboard visualization

## Get Started

1. Set STACK_VERSION of ELK in .env file
2. Run docker compose

```bash
docker compose up -d
```

3. Run service-go to produce log

```bash
cd service-go && go run main.go && cd ..
```

4. Open Kibina dashboard to see the result
   Go to: `http://localhost:5601`

## Basic command for elasticsearch

References:

- https://www.notion.so/ribbinpo/CRUD-Command-935694d016244c4eb9d03e9e88adf2e1
- https://www.elastic.co/guide/en/elasticsearch/reference/current/release-highlights.html

### Test tokenizer

```
POST _analyze
{
  "tokenizer": "thai",
  "text": "สวัสดีครับคุณครู"
}
```

### Update custom default analyzer

```
PUT myindex
{
  "settings": {
    "index": {
      "analysis": {
        "analyzer": {
          "analyzer_mythai": {
            "tokenizer": "thai",
            "filter": [
              "filter_shingle"
            ]
          }
        },
        "filter": {
          "filter_shingle": {
            "type": "shingle",
            "max_shingle_size": 3,
            "min_shingle_size": 2,
            "output_unigrams": "true"
          }
        }
      }
    }
  },
  "mappings": {
    "properties": {
      "content": {
        "analyzer": "analyzer_mythai",
        "type": "text"
      }
    }
  }
}
```

### Test create

```
POST myindex/_doc/1
{
  "content": "สวัสดีปีใหม่",
  "type": "hello"
}

POST myindex/_doc/2
{
  "content": "สวัสดีคุณครู",
  "type": "hello"
}
```

### Search

```
GET myindex/_search
{
  "query" :
    {"match" :
      {"content" : "สวัสดี"}
    }
}
```
