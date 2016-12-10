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
* Body: /"token": ".*"/

