# artist-store 

## setting up backend for dev mode

### 1. MySql setup
```
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql
```

### 2. Sqlc setup
Refer to the link below for setting up Sqlc  
link: https://github.com/SeokbyeongChae/go-sqlc-example

### 3. Skaffold setup
Refer to the link below for setting up Skaffold  
link: https://github.com/SeokbyeongChae/nodejs-skaffold-example

### 4. Start Skaffold
```
skaffold dev
```
