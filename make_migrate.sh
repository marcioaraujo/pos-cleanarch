# Se necessário acesse o mysql manualmente e ajuste as permissões:
# docker exec -it mysql bash
# GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'root';
# FLUSH PRIVILEGES;

migrate -path=sql/migrations -database "mysql://root:root@tcp(127.0.0.1:3306)/orders" -verbose up
