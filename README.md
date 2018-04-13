"# swagger_sample" 

To generate server side code execute the following command:
swagger generate server -f api\swagger.yml -A swaggertest


To build the server
cd cmd\swaggertest-server
go build -o ..\..\server.exe


To run the server.
cd ..\..\
server.exe -port=8080

Note: Since the default is to start with a random Port, port has been specified in the argument.


To access the endpoint:
curl "http://localhost:8080/search?user_id=10"


To generate client side code execute the following command:
swagger generate client -f api\swagger.yml -A swaggertest

