#rpc
go run service/film/rpc/film.go -f service/film/rpc/etc/film.yaml &
go run service/order/rpc/order.go -f service/order/rpc/etc/order.yaml &
#api
go run service/advert/api/advert.go -f service/advert/api/etc/advert.yaml &
go run service/comment/api/comment.go -f service/comment/api/etc/comment.yaml &
go run service/film/api/film.go -f service/film/api/etc/film.yaml &
go run service/order/api/order.go -f service/order/api/etc/order.yaml &
go run service/pay/api/pay.go -f service/pay/api/etc/pay.yaml &
go run service/user/api/user.go -f service/user/api/etc/user.yaml
