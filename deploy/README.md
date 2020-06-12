# Deploy CSP to lambda via AWS SDK for go
This program can be used to deploy lambda function to your environment.

# Command
go run main.go -z function -f csplistener -h main -a arn:aws:iam::849229154245:role/lambda -r go1.x\

where

z => The name of the ZIP file (without the .zip extension)

f => The name of the Lambda function

h => The name of the package.class handling the call

a => The ARN of the role that calls the function

r => The runtime for the function
