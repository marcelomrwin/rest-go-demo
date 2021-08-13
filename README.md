# TUTORIAL

## Create project and dependencies
```bash
go mod init rest-go-demo
go get github.com/gorilla/mux
go get github.com/jinzhu/gorm
go get github.com/go-sql-driver/mysql
```

## Initialize MySQL database

### Use init scripts (optional)

`-v db/scripts:/docker-entrypoint-initdb.d/`

```bash
docker network create rest-go-demo-net
docker run -d --rm --name mysql -p 3306:3306 --network=rest-go-demo-net -e MYSQL_USER=demo -e MYSQL_PASSWORD=demo -e MYSQL_DATABASE=demo -e MYSQL_ROOT_PASSWORD=demo -v $(pwd)/db/data:/var/lib/mysql/ -v $(pwd)/db/conf/mysql-custom.cnf:/etc/mysql/conf.d -d mysql/mysql-server
```
