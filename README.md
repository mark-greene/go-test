# go-test
Go-Test golang HTTP test server

Only runs on an AWS instance as it uses the AWS API `http://169.254.169.254/latest/dynamic/instance-identity/document`.
Must be built on the same platform as the AWS Instance.

## Routes
```
/
/headers
/status
/instance
```
### To Build
```
go install .
```
### To Run
```
export PORT=80
sudo -E ./go-test
```
### To Test
```
curl http://localhost/status
```
