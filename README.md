# 📍 Coupon Service

![Docker Version](https://img.shields.io/badge/docker-v25.0.2-2496ED.svg?style=for-the-badge&logo=docker)
![Golang Version](https://img.shields.io/badge/go-v1.21.6-00ADD8.svg?style=for-the-badge&logo=go) ![Grafana](https://img.shields.io/badge/Grafana-F46800.svg?style=for-the-badge&logo=grafana&logoColor=white) ![Prometheus](https://img.shields.io/badge/Prometheus-E6522C.svg?style=for-the-badge&logo=prometheus&logoColor=white) ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1.svg?style=for-the-badge&logo=postgresql&logoColor=white)

---

## Steps to run the application with Docker Compose

1. Make sure you have Docker installed on your machine.
```sh
$ docker version && docker compose version
```

2. Run Docker Compose
```sh
$ docker compose -f docker-compose.yaml up --build -d
```
---
## 📕  API Documentation

> ❗ Prices are represented in cents.


### Create a coupon
```sh
method: POST 
path: /api/coupon/create
```

Request:
```sh
curl --request POST \
  --url http://localhost:8080/api/coupon/create \
  --header 'content-type: application/json' \
  --data '{"code": "A000A","discount": 50,"min_basket_value": 100}'
````

Response:
```json
{
  "data": null,
  "method": "POST",
  "status_code": 201,
  "url": "/api/coupon/create"
}
```
---
### Get coupons
```sh
method: GET 
path: /api/coupon/
```

Request:
```sh
curl --request GET --url http://localhost:8080/api/coupon/
````

Response:
```json
{
  "data": [
    "A000A"
  ],
  "method": "GET",
  "status_code": 200,
  "url": "/api/coupon/"
}
```

### Apply coupon
```sh
method: POST 
path: /api/coupon/apply
```

Request:
```sh
curl --request POST \
  --url http://localhost:8080/api/coupon/apply \
  --header 'content-type: application/json' \
  --data '{"code": "A000A","value": 100}'
````

Response:
```json
{
  "data": {
    "value": 100,
    "applied_discount": 50,
    "application_successful": true,
    "final_value": 50
  },
  "method": "POST",
  "status_code": 200,
  "url": "/api/coupon/apply"
}
```
---
## Ports where other applications are running

- Prometheus: http://localhost:9090
- Grafana: http://localhost:3000
- PostgreSQL: http://localhost:5432
- API: http://localhost:8080