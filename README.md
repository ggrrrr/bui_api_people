# bui-api-login

## Project setup
```
go mod tidy -compat=1.17
```

`.env.local`
```bash
LISTEN_ADDR=:8100


```

### Compiles and hot-reloads for development
```
export $(xargs <.env.local)
go run main.go
```

## TESTS
```
export T=`curl -s -X POST -d '{"email":"asd@asd.com","password":"asd"}' localhost:8000/userLogin | jq -r '.token'`

curl -v  -H "Authorization: Bearer $T" http://localhost:8000/tokenVerify

```