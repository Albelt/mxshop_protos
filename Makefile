.PHONY: good
good:
	protoc -I /Users/albelt/Gitlab/common/protos -I . \
		--go_out=albelt/good_srv \
		--go-grpc_out=albelt/good_srv \
		--validate_out="lang=go:albelt/good_srv" \
		albelt/good_srv/msg.proto albelt/good_srv/svc.proto


.PHONY: order
order:
	protoc -I /Users/albelt/Gitlab/common/protos -I . \
		--go_out=albelt/order_srv \
		--go-grpc_out=albelt/order_srv \
		--validate_out="lang=go:albelt/order_srv" \
		albelt/order_srv/msg.proto albelt/order_srv/svc.proto albelt/order_srv/const.proto


.PHONY: stock
stock:
	protoc -I /Users/albelt/Gitlab/common/protos -I . \
		--go_out=albelt/stock_srv \
		--go-grpc_out=albelt/stock_srv \
		--validate_out="lang=go:albelt/stock_srv" \
		albelt/stock_srv/msg.proto albelt/stock_srv/svc.proto albelt/stock_srv/const.proto


.PHONY: user
user:
	protoc -I /Users/albelt/Gitlab/common/protos -I . \
		--go_out=albelt/user_srv \
		--go-grpc_out=albelt/user_srv \
		--validate_out="lang=go:albelt/user_srv" \
		albelt/user_srv/msg.proto albelt/user_srv/svc.proto albelt/user_srv/const.proto


.PHONY: all
all:
	make good
	make order
	make stock
	make user

.PHONY: clean
clean:
	rm -rf ./albelt/**/go
