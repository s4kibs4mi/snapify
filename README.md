# Snapify

A RESTful API service to take screenshot of any webpage.

### Dependencies

* Go
* Docker (needed to be installed on your machine)
* Postgresql
* Minio (S3 compatible storage)
* Redis (Queue)
* Rod (Headless Browser library for golang)

### Run application

```shell
make docker-up
```

Note: schema will be populated automatically.

### Cleanup

```shell
make docker-down
```

#### Database schema

```text
id UUID primary_key not_null
status string not_null
url string not null
stored_path string
created_at datetime index
```

#### APIs

* Create screenshot

`POST /v1/screenshots`

```json
{
  "url": "https://sakib.ninja"
}
```

Result `202`

```json
{
  "data": {
    "id": "9773afb9-7beb-4d6d-84dd-d421919b785d",
    "url": "https://sakib.ninja",
    "status": "queued",
    "created_at": "2023-02-18T20:23:14Z"
  }
}
```

* List screenshots

`GET /v1/screenshots?page={page}&limit={limit}`

Result `200`

```json
{
  "data": [
    {
      "id": "9773afb9-7beb-4d6d-84dd-d421919b785d",
      "url": "https://sakib.ninja",
      "status": "completed",
      "created_at": "2023-02-18T20:23:14Z"
    }
  ]
}
```

* Get screenshot

`GET /v1/screenshots/{screenshot_id}`

Result `200`

```text
{
    "data": {
        "id": "9773afb9-7beb-4d6d-84dd-d421919b785d",
        "url": "https://sakib.ninja",
        "status": "completed",
        "created_at": "2023-02-18T20:23:14Z",
        "screenshot_url": "http://minio:9000/snapify/10fa94b2-890a-43aa-91f2-e3ed7c8aeced.png?X-Amz-Algorithm=AWS4-HMAC-SHA256\u0026X-Amz-Credential=MINIO_ACCESS_KEY%2F20230218%2Fus-east-1%2Fs3%2Faws4_request\u0026X-Amz-Date=20230218T202357Z\u0026X-Amz-Expires=300\u0026X-Amz-SignedHeaders=host\u0026X-Amz-Signature=a3c22c5847b6400034cda7d7d3201da25deebf699288b780afcacbc1d5f814ec"
    }
}
```

### Used Libraries

- [Fiber Framework](https://docs.gofiber.io/)
- [Ent](https://entgo.io/)
- [Test Container](https://golang.testcontainers.org/)
- [Testify](https://github.com/stretchr/testify)
- Viper
- Logrus
- Rod

### License

[MIT](./LICENSE)
