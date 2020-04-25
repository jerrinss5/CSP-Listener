# CSP

Content Security Policy is the name of the security response header that is used to provide defense in depth measure to enhance the security of the webpage.

The CSP header allows to restrict how resources such as Javascript, CSS or pretty much anything that browser loads.

# CSP-Listener

CSP Listener is a server which logs all the CSP violations

Currently the logged information would be of the format

Timestamp   Content

2020/04/13 01:13:09 {http://localhost:4000/  script-src-elem script-src-elem script-src 'self'; report-uri https://localhost:900/report report inline 1 http://localhost:4000/ 200 }

2020/04/13 01:13:09 line-number:1, source-file:http://localhost:4000/, script-sample:, Document-uri:http://localhost:4000/, referrer:, violated-directive:script-src-elem, effective-directive:script-src-elem, original-policy:script-src 'self'; report-uri http://localhost/report, blocked-uri:inline, status-code:200

## Purpose

Before we can start enforcing CSPs - we need to make sure it does not break existing applications.

So to test any new policy - we can apply the policy in report only mode and observe what parts of the website are being affected by the policy and apply different measures to be compliant to the policy.

It also give an added benefit of violations reported when an actual attack is performed.

### Development

Listener service is written in GO and deployed as an executable on scratch container image

The service is listening on port 9000 by default and can be changed via the environment variable PORT

### To test locally without executable

go run *.go

### Building the docker image

docker build -t csplistener:0.1 .

### Running the docker image

docker run -d -p 9000:9000 --name csplistener --restart always csplistener:0.1

### Running the executable (by default listening on port 9000)

For linux: ./cspListenerLinux

For Mac: ./cspListenerMac

### Ensuring the server has started
You would see the message

CSP Report logger service running on 9000 (by default)

If you need to change the port and host it can be done by setting at the environment variables

EXPORT PORT=9000
EXPORT HOST=0.0.0.0

## Testing

Additionally you can test with the following payload by send a POST request to the endpoint

http://localhost:9000/report

{"csp-report":{"document-uri":"http://localhost:4000/","referrer":"","violated-directive":"script-src-elem","effective-directive":"script-src-elem","original-policy":"script-src 'self'; report-uri http://localhost/report","disposition":"report","blocked-uri":"inline","line-number":1,"source-file":"http://localhost:4000/","status-code":200,"script-sample":""}}

## Health Check

Send a GET request to /health endpoint to get a 204 (Status - No Content)

http://localhost:9000/health