# Snapify
Screenshot as a service of web pages.

#### Architecture
- Screen Shot Creation Flow
![](./extras/creation_flow.png)

#### Endpoints
1. `POST /v1/screenshots`
```json
{
    "urls": [
        "www.sakib.ninja",
        "www.codersgarage.com",
        "www.google.com"
    ]
}
```

Result `202`
```json
{
    "title": "Screenshot creation is in progress",
    "data": [
        {
            "id": "c0a10b26-ab7e-48c8-8e58-6bbeb8afe8a7",
            "status": "queued",
            "website": "www.sakib.ninja",
            "created_at": "2019-10-29T13:05:26.1003577Z",
            "updated_at": "2019-10-29T13:05:26.1003726Z"
        },
        {
            "id": "ce47935c-2afa-4c03-bcba-a211250dedae",
            "status": "queued",
            "website": "www.codersgarage.com",
            "created_at": "2019-10-29T13:05:26.1320367Z",
            "updated_at": "2019-10-29T13:05:26.13205Z"
        },
        {
            "id": "ae2742db-a3cc-4e1d-999b-f46355695c5b",
            "status": "queued",
            "website": "www.google.com",
            "created_at": "2019-10-29T13:05:26.1334529Z",
            "updated_at": "2019-10-29T13:05:26.1334668Z"
        }
    ]
}
```

2. `GET /v1/screenshots?page={page}&limit={limit}`

Result `200`
```json
{
    "data": {
        "meta": {
            "current_page": 1,
            "page_limit": 10,
            "total": 15,
            "total_pages": 2
        },
        "screenshots": [
            {
                "id": "b1305a82-996c-492b-a932-db5a2cc16dfd",
                "status": "done",
                "website": "www.google.com",
                "created_at": "2019-10-29T09:00:38.750841Z",
                "updated_at": "2019-10-29T09:30:35.370862Z"
            },
            {
                "id": "70cb187b-f6f0-44a3-ba78-a43197b1bf4c",
                "status": "done",
                "website": "www.codersgarage.com",
                "created_at": "2019-10-29T09:00:38.734694Z",
                "updated_at": "2019-10-29T09:30:35.609259Z"
            },
            {
                "id": "29140d04-f35a-46ff-866e-aee1f9879ecb",
                "status": "done",
                "website": "www.sakib.ninja",
                "created_at": "2019-10-29T09:00:38.686813Z",
                "updated_at": "2019-10-29T09:30:55.145758Z"
            },
            {
                "id": "2c8dbe74-8a57-4393-a27f-1755ae9ac0c1",
                "status": "queued",
                "website": "www.google.com",
                "created_at": "2019-10-29T08:51:29.009733Z",
                "updated_at": "2019-10-29T08:51:29.009788Z"
            },
            {
                "id": "b3c78432-357c-4c57-ae39-c3e17d4b90be",
                "status": "queued",
                "website": "www.codersgarage.com",
                "created_at": "2019-10-29T08:51:29.008021Z",
                "updated_at": "2019-10-29T08:51:29.008129Z"
            },
            {
                "id": "b7232608-41fd-45a3-9989-dd6c8299598f",
                "status": "queued",
                "website": "www.sakib.ninja",
                "created_at": "2019-10-29T08:51:28.972604Z",
                "updated_at": "2019-10-29T08:51:28.972618Z"
            },
            {
                "id": "c13f6651-5028-4820-956d-f5682ae63dd7",
                "status": "queued",
                "website": "www.google.com",
                "created_at": "2019-10-29T08:44:00.68891Z",
                "updated_at": "2019-10-29T08:44:00.688925Z"
            },
            {
                "id": "220a3675-eaa0-45bc-b88f-5a37902e6b5a",
                "status": "failed",
                "website": "www.codersgarage.co",
                "created_at": "2019-10-29T08:44:00.687227Z",
                "updated_at": "2019-10-29T08:44:00.687357Z"
            },
            {
                "id": "bbcc7bbe-ed10-4826-8657-b9e5bb0d38f4",
                "status": "queued",
                "website": "www.sakib.ninja",
                "created_at": "2019-10-29T08:44:00.624945Z",
                "updated_at": "2019-10-29T08:44:00.624959Z"
            },
            {
                "id": "dc123a81-cfc8-44bb-8169-76036422ac4a",
                "status": "done",
                "website": "www.google.com",
                "created_at": "2019-10-29T08:04:03.476972Z",
                "updated_at": "2019-10-29T08:04:05.746757Z"
            }
        ]
    }
}
```

3. `/v1/screenshots/{screenshot_id}`

Result `200`
```text
The screen shot will be shown
```

#### Dependencies

- PostgreSQL
- RabbitMQ
- Minio
- ChromeHeadless browser via chromedp

#### App Build
```bash
./build.sh build
```

#### Docker Image Build
```bash
./build.sh docker {{docker_version}}
```

#### App Run
```bash
./build.sh run
```

#### Run complete system
```bash
docker-compose.up -d
```
