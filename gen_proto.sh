#!/bin/zsh

#cd ..
#clang-format mall-demo/proto/order/order.proto -i
#clang-format mall-demo/proto/product/product.proto -i
#clang-format mall-demo/proto/product/spu.proto -i
#clang-format mall-demo/proto/coupon/coupon.proto -i
#clang-format mall-demo/proto/coupon/spu.proto -i
#clang-format mall-demo/proto/member/member.proto -i
#clang-format mall-demo/proto/ware/ware.proto -i
#clang-format mall-demo/proto/thirdparty/thirdparty.proto -i
#
#
#protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  mall-demo/proto/coupon/*.proto
#
#protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  mall-demo/proto/product/*.proto
#
#protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  mall-demo/proto/member/*.proto
#
#protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative mall-demo/proto/order/*.proto
#
#protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative mall-demo/proto/ware/*.proto
#
#protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative mall-demo/proto/thirdparty/*.proto
#
#cd mall-demo
#

cd ..
services=(api coupon member order product thirdparty ware search auth-server cart)

# shellcheck disable=SC2068
for service in ${services[@]}
do
  # shellcheck disable=SC2045
  for dir in $(ls "mall-demo/micro-mall-${service}/proto")
  do
    for file in $(ls "mall-demo/micro-mall-${service}/proto/${dir}")
    do
      if echo "${file}" | grep -q -E '\.proto$'
      then
        clang-format "mall-demo/micro-mall-${service}/proto/${dir}/${file}" -i
        protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative "mall-demo/micro-mall-${service}/proto/${dir}/${file}"
      fi
    done
  done
done

cd mall-demo

