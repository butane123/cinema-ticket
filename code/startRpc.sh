#rpc
go run service/film/rpc/film.go -f service/film/rpc/etc/film.yaml &
go run service/order/rpc/order.go -f service/order/rpc/etc/order.yaml
