# artist-store 

## setting up backend for dev mode

### 1. Setting up MySql
```
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql
```

### 1. Setting up Sqlc
Refer to the link below for setting up Sqlc  
link: https://github.com/SeokbyeongChae/go-sqlc-example

### 2. Setting up Skaffold
Refer to the link below for setting up Skaffold  
link: https://github.com/SeokbyeongChae/nodejs-skaffold-example

### 3. Start Skaffold
```
skaffold dev
```
