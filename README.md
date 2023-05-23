# Go + onion architecture example

TODO app with onion architecture.

## module

### db

use sqlite3.

```
go get github.com/mattn/go-sqlite3
```

### web

use echo.

```
go get github.com/labstack/echo/v4
```

## api

### todo

#### create

request

```
curl -X POST http://localhost:8080/1/todos -H "Content-Type: application/json" -d '{"title":"test","description":"test"}'
```

response 

```
{
  "id": 9
}
```

#### list

request

```
curl -X GET http://localhost:8080/1/todos
```

response 

```
{
  "todos": [
    {
      "id": 6,
      "title": "test",
      "description": "test",
      "created_at": "2023-05-20T07:46:49.112686+09:00",
      "delete_at": "0001-01-01T00:00:00Z"
    },
    {
      "id": 7,
      "title": "test",
      "description": "test",
      "created_at": "2023-05-21T14:58:50.918357+09:00",
      "delete_at": "0001-01-01T00:00:00Z"
    },
    {
      "id": 8,
      "title": "test",
      "description": "test",
      "created_at": "2023-05-21T17:31:32.128915+09:00",
      "delete_at": "0001-01-01T00:00:00Z"
    },
    {
      "id": 9,
      "title": "test",
      "description": "test",
      "created_at": "2023-05-21T17:31:45.145144+09:00",
      "delete_at": "0001-01-01T00:00:00Z"
    }
  ]
}
```

#### delete

request

```
curl -X DELETE localhost:8080/todos/9
```

response

```
HTTP/1.1 204 No Content
```

### user

#### create

request

```
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"name":"yamanaka"}'
```

response 

```
{
  "id": 2
}
```

#### Get

request

```
curl -X GET http://localhost:8080/user/2
```

response 

```
{
  "user": {
    "id": 2,
    "name": "yamanaka",
    "created_at": "2023-05-21T17:39:00.358339+09:00"
  }
}
```

## test

### testify

```
go get github.com/stretchr/testify
```

### gomock

```
go get github.com/golang/mock/gomock
```

create mock

```
go generate ./...
```

### dockertest

```
go get -u github.com/ory/dockertest/v3
```

**but, this sample use sqlite3. so, not use dockertest.**

### golden test

```
go get -u github.com/sebdah/goldie/v2
```