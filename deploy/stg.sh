#!/usr/bin/env bash
APP_NAME=food_delivery
DEPLOY_CONNECT=root@146.190.98.185
cd .. 
echo "load docker images..."
docker load -i ${APP_NAME}.tar
echo "remove old image..."
docker rm -f ${APP_NAME}
echo "run docker..."
docker run -d --name ${APP_NAME} \
  --network buy_me \
  -e S3BUCKETNAME="food-delivery-tb" \
  -e S3REGION="ap-southeast-1" \
  -e S3APIKEY="AKIA46IFCIM4KQVEV5WF" \
  -e S3SECRETKEY="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" \
  -e S3DOMAIN="http://food-delivery-tb.s3-website-ap-southeast-1.amazonaws.com" \
  -e SYSTEM_SECRET="thaibinh123" \
  -e MYSQL_CONNECTION="food_delivery:thaibinh123@tcp(mysql:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local" \
  -p 8080:8080 \
  food_delivery