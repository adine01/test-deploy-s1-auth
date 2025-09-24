GET
/health
Health check endpoint


Returns the health status of the auth service (public endpoint)

Parameters
Cancel
No parameters

ExecuteClear
Responses
Curl

curl -X 'GET' \
  'https://b18a02cf-5dea-45c5-a25e-d69ff8710855-dev.e1-us-east-azure.choreoapis.dev/sts-test/auth-service/v1.0/health' \
  -H 'accept: application/json' \
  -H 'Test-Key: '
Request URL
https://b18a02cf-5dea-45c5-a25e-d69ff8710855-dev.e1-us-east-azure.choreoapis.dev/sts-test/auth-service/v1.0/health
Server response
Code	Details
401
Undocumented
Error: Unauthorized

Response body
Download
{"error_message":"Invalid Credentials","code":"900901","error_description":"Make sure you have provided the correct security credentials"}
Response headers
 content-length: 138 
 content-type: text/plain 