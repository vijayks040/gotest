# 

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

This Go code was added to support the POC or the CASE-STUDY given for an interview.

# Pre-requisites!

  - Hoping that the GO environment already present,run the following command,
  
        -go get github.com/labstack/echo gopkg.in/mgo.v2/bson github.com/lib/pq
        
  - Install postgre sql and MongoDb
        
        -Create a database table in postgre sql with following sql command,
        
            -CREATE TABLE public.product(
                 "Productname" text COLLATE pg_catalog."default" NOT NULL,
    
                  "currentPrice" double precision NOT NULL,
    
                  "productId" bigint NOT NULL,
    
                  currency text COLLATE pg_catalog."default" NOT NULL,
    
                  CONSTRAINT product_pkey PRIMARY KEY ("productId")
    
            )

**Here public is the database schema and product is the database table name.
        
        -Create a collection in MongoDB to enter the details of the products created.**

### API's
**Create Product**
```sh
curl -X POST 
http://localhost:3000/createproduct 
-H 'cache-control: no-cache' 
-H 'content-type: multipart/form-data 
-F productid=155556 
-F name=shirt 
-F current_price=20 
-F currency=INR 
-F description=full and half sleeves
```

**Update Product**: JSON input Format

```sh
curl -X PUT 
http://localhost:3000/updateproduct 
-H 'cache-control: no-cache' 
-H 'content-type: application/json' 
-d '{ "productid": "155556", "title": "vijay", "current_price": 20.12, "currency": "INR", "description": "testing json" }'
```

**Get Product**: Using Product ID( Here product is the unique product ID )
```sh
http://localhost:3000/findproduct 
-H 'cache-control: no-cache' 
-H 'content-type: multipart/form-data
-F productid=15555
```

**Delete Product**: Using Product ID( Here product is the unique product ID )
```sh
curl -X POST 
http://localhost:3000/deleteproduct 
-H 'cache-control: no-cache' 
-H 'content-type: multipart/form-data
-F productid=155555
```
### Executables

Both the windows and linux platform executables have been provided along with source code.

| Platform | Executable |
| ------ | ------ |
| Windows7 and higher version | target.exe (Run using admin permissions) |
| Linux | target_linux (with sudo on run ./target_linux command in terminal) |
| Docker Image | Please request the developer |

### Development Environment

Want to contribute? Great!

**GO version : 1.10.3 windows/amd64**

**Postgres : 9.6**

**MongoDB : 4.0**

**-This piece of code is developed in windows environment.**

**-Anyhow Go is such a platform freindly coding language, you can run this code in any other platforms such as linux,mac etc**




### Todos

 - Write MORE Tests
 - Manking parameters configurable
 - Using dependancy management **dep**

License
----

@vijay


**Free Software, Hell Yeah!**
