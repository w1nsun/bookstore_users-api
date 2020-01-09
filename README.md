# bookstore_users-api
Users API

```
curl -X POST localhost:8080/users -d '{"id": 123, "first_name": "Fede", "email": "fedeleon@gmail.com"}' -v
```

Don't forget set ENV variables in IDE
```
mysql_users_username=something
mysql_users_password=something
mysql_users_host=127.0.0.1:3306
mysql_users_schema=users_db
```

run mysql for this service

```
docker run --name mysql57 --rm -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 mysql:5.7
```