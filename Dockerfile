FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /go_test/
COPY . /go_test/
CMD ["./rpcx_service"]

#wayne

#900edabdcd5e

#docker exec -it 900edabdcd5e /bin/bash

#apt-get update
#apt-get install vim

#docker run -p 9100:80 --name my-custom-nginx-container --mount type=bind,source=/usr/local/nginx/nginx_data/nginx.conf,destination  /etc/nginx/nginx.conf:ro -d nginx

#docker run -p 9101:80 --name my-first-nginx-container -v /usr/local/nginx/nginx_data/nginx.conf:/etc/nginx/nginx.conf -d nginx

