# 

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

This Go code was added to support the POC or the CASE-STUDY given for cache management.

# Pre-requisites!

  - Hoping that docker environment is setup.
  Docker Image @ **docker pull vijayks040/cachetool:final**
  - GO environment to run the test package
        
## TO TEST THIS TOOL
   -run command **go test** within the folder cachetool.
   
   
   -Postman Link: [https://www.getpostman.com/collections/92acdc069b27915b017f]

### API's
**Load the Cache**
```sh
curl -X GET 
http://localhost:8080/loadcache
**Output**: json response
example: 


Success-
{
    "item": [
        "name:vijay age:27 mobile:96209****",
        "name:vinay age:29 mobile:9999999999",
        "name:vikas age:28 mobile:7986767777",
        "name:ajey age:30 mobile:34324534556",
        "name:gani age:27 mobile:65436546546",
        "name:kris age:40 mobile:53456436543",
        "name:gopi age:70 mobile:65565467546",
        "name:pravi age:45 mobile:7868787665",
        "name:virat age:30 mobile:5654654667",
        "name:virat age:30 mobile:5654654667",
        "name:anushka age:31 mobile:43543564364"
    ]
}


Failed-Http response message "cache loading failed" along with error message
```

**Add item to Cache**: HTML form input

```sh
curl -X POST 
http://localhost:3000/putcache 
-H 'multipart/form-data' 
-F 'data=name:virat age:30 mobile:5654654667,name:anushka age:31 mobile:43543564364'

**Output**:
Example:
Success-"Data Added Successfully"

Failed-"cache adding failed"
```

**Get Cache Items**: Supports both GET and POST requests


**GET**
```sh
curl -X GET
http://localhost:8080/getcache

**Output**:
Example:
Success-json response
{
    "item": [
        "name:vijay age:27 mobile:96209*****",
        "name:vinay age:29 mobile:9999999999",
        "name:vikas age:28 mobile:7986767777",
        "name:ajey age:30 mobile:34324534556",
        "name:gani age:27 mobile:65436546546"
    ],
    "total_pages": 3
}
Failed-json response
{
    "description": "No items present in cache in now"
}

```


**POST**
```sh
curl -X POST 
http://localhost:8080/getcache
-H 'content-type: multipart/form-data
-F page=4

**Output**:
Example:
Success-json response
{
    "item": [
        "name:vijay age:27 mobile:96209*****",
        "name:vinay age:29 mobile:9999999999",
        "name:vikas age:28 mobile:7986767777",
        "name:ajey age:30 mobile:34324534556",
        "name:gani age:27 mobile:65436546546"
    ],
    "total_pages": 4
}

Failed-json response
{
    "total_pages": 3,
    "description": "Page not found"
}
```

**Clear Cache items**:
```sh
curl -X GET 
http://localhost:8080/clearcache 

**Output**:
Example:
Success-"Cache cleared Successfully"
Failed- "error truncating cache"
```
### Executables

Both the windows and linux platform executables have been provided along with source code.

| Platform | Executable |
| ------ | ------ |
| Docker Image | DockerHub image: [vijayks040/cachetool:final] |

### Development Environment

Want to contribute? Great!

**GO version : 1.10.3 windows/amd64**


**-This piece of code is developed in windows environment.**

**-Anyhow Go is such a platform freindly coding language, you can run this code in any other platforms such as linux,mac etc**




### Todos

 - Write MORE Tests
 - Making parameters configurable
 - Using dependancy management **dep**

License
----

@vijay


**Free Software, Hell Yeah!**
