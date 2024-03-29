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

### Database schema

```text
id UUID primary_key not_null
status string not_null
url string not null
stored_path string
created_at datetime index
```

### Docs

![](./extras/docs.png)

After running docker compose,

Visit: http://localhost:9010/docs

### Used Libraries

- [Fiber Framework](https://docs.gofiber.io/)
- [Ent](https://entgo.io/)
- [Test Container](https://golang.testcontainers.org/)
- [Testify](https://github.com/stretchr/testify)
- Viper
- Logrus
- Rod

### Support

<a href="https://www.buymeacoffee.com/sakib.ninja" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" style="height: 60px !important;width: 217px !important;" ></a>

### License

[MIT](./LICENSE)
