This repository was added to support the POC or the CASE-STUDY given for an interview
1) Please pull the whole target into a common directory
2) Then build and install the postgre_package and mongo_package
3) Then build the targetPoc.go(contains the main function) to generate the executable
4) By defau:lt the application will be running in port 3000
5) Hit http://localhost:3000 to see the available API's

REQUEST Examples:

API 1) CREATE A PRODUCT:
curl -X POST \
  http://localhost:3000/createproduct \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -F productid=155556 \
  -F name=mmmmmm \
  -F current_price=20 \
  -F currency=INR \
  -F description=testinggggg

API 1) UPDTAE A PRODUCT:
curl -X PUT \
  http://localhost:3000/updateproduct \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
	"productid": "155556",
	"title": "vijay",
	"current_price": 20.12,
	"currency": "INR",
	"description": "testing json"
}'

API 3) GET THE PRODUCT USING ID: Here productid is the unique ID
curl -X POST \
  http://localhost:3000/findproduct \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -F productid=15555

API 4) DELETE A PRODUCT USING ID: Here productid is the unique ID
curl -X POST \
  http://localhost:3000/deleteproduct \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -F productid=155555
