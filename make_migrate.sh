# Se necessário acesse o mysql manualmente e:
# GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'root';
# FLUSH PRIVILEGES;


migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up