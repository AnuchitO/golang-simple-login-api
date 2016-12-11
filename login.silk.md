# Login

## POST /login
* Content-Type: `application/json`

### Example body for request

```json
{
  "user": "kob@gmail.com",
  "password": "aobaob"
}
```

===

### Example response

* Status: `200`
* Content-Type: `application/json; charset=utf-8`
* Body: /"token":".*"/


## GET /users
* Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJrb2JAZ21haWwuY29tIiwiaXNzIjoiYXBwIGtvYiJ9.ZjrECc-Y_EjVQ-Ui6-_Qos4Yl5ZJ9Tx-K_IRtj5croQ`

===

### Example response

* Status: `200`
* Content-Type: `application/json; charset=utf-8`
* Body: /"user":"kob@gmail.com"/


## GET /users
* Authorization: `Bearer bad token`

===

### Example response

* Status: `401`
* Content-Type: `application/json; charset=utf-8`
* Body: /"error":"not a compact JWS"/
