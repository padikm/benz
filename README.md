# Language Used Golang
## setup svc1
1. clone the repo
2. cd to svc1
3. run go build
4. it will create a binary svc1.exe(windows).
5. run svc1.exe it will start http server
**Note: Run svc1 with environment GODEBUG=x509ignoreCN=0 eg: set GODEBUG=x509ignoreCN=0 in windows. Some problem i am facing with ca certs. 
## setup svc2 
1. clone the repo
2. cd to svc2
3. run go build
4. it will create a binary svc2.exe(windows).
5. run svc2.exe it will start grpc server

## Api endpoints
1. Create:
   endpoint -> http://localhost:8000/employees
	 Method -> Post
	 Body ->
	 {
    "name":"name",
    "age" : 25,
    "address": "bangalore"
		}
2. read:
   endpoint -> http://localhost:8000/employees
	 Method -> Get
3. Edit:
   endpoint -> http://localhost:8000/employees/{id}
	 Method -> Post
	 Body ->
	 {
    "name":"name",
    "age" : 25,
    "address": "bangalore"
		}
		Note: id here is employee id which we get from read or create api
		
## Deployment
