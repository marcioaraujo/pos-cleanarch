# Primeira vez pode ser necessário
# ./make_migrate.sh 



cd cmd/ordersystem
go run main.go wire_gen.go
cd ..
cd ..