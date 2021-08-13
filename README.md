# TUTORIAL

## Create project and dependencies
```bash
mkdir rest-go-demo && cd rest-go-demo
go mod init rest-go-demo
go get gitlab.com/avarf/getenvs
go get github.com/gorilla/mux
go get github.com/jinzhu/gorm
go get github.com/go-sql-driver/mysql
```

## Initialize MySQL database

### Use init scripts (optional)

`-v db/scripts:/docker-entrypoint-initdb.d/`

```bash
docker network create rest-go-demo-net
mkdir -p db/data
mkdir -p db/conf
cat >> db/conf/mysql-custom.cnf <<EOL
[mysqld]
max_connections=250
EOL
docker run -d --rm --name mysql -p 3306:3306 --network=rest-go-demo-net -e MYSQL_USER=demo -e MYSQL_PASSWORD=demo -e MYSQL_DATABASE=demo -e MYSQL_ROOT_PASSWORD=demo -v $(pwd)/db/data:/var/lib/mysql/ -v $(pwd)/db/conf/mysql-custom.cnf:/etc/mysql/conf.d -d mysql/mysql-server
```

## Run the project

```
go run main.go
```

## Compile the executable

```
go build -a -o demo main.go
```

## Run the app

```
./demo
```

## Build Docker image

```bash
docker build -t rest-go-demo .
docker run --rm -it -p 8080:8080 --network=rest-go-demo-net -e MYSQL_HOST=mysql rest-go-demo
```
