curl -X POST http://localhost:5000/login -F 'username=test' -F 'password=test'

curl -X POST http://localhost:5000/login -H "Content-Type: application/x-www-form-urlencoded" -d "username=test&password=test"

curl http://localhost:5000/api/videos -H "Accept: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzdCIsImFkbWluIjp0cnVlLCJleHAiOjE2NDg0ODE1MjcsImlhdCI6MTY0ODIyMjMyNywiaXNzIjoiZXhhbXBsZS5jb20ifQ.B2N6i1VaUQI2Ql53hFxfSyorGRTUbTp8wz4zmAcNypE"

