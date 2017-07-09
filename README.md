# go-test
Golang HTTP test server

Runs on an AWS instance as it uses the AWS API `http://169.254.169.254/latest/dynamic/instance-identity/document`.
Must be built for the same platform as the AWS Instance.

## Routes
```
/
/headers
/status
/instance
```
### To Build and Test
```
go get github.com/mark-greene/go-test
cd src/github.com/mark-greene/go-test/
go install .
PORT=8080 go-test &
curl http://localhost:8080/status
```
### To Run and Test on AWS instaqnce
```
export PORT=80
sudo -E ./go-test &
curl http://localhost/instance
```
