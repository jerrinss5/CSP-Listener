package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"

	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	zipFilePtr := flag.String("z", "", "The name of the ZIP file (without the .zip extension)")
	functionPtr := flag.String("f", "", "The name of the Lambda function")
	handlerPtr := flag.String("h", "", "The name of the package.class handling the call")
	resourcePtr := flag.String("a", "", "The ARN of the role that calls the function")
	runtimePtr := flag.String("r", "", "The runtime for the function.")

	flag.Parse()

	if *zipFilePtr == "" || *functionPtr == "" || *handlerPtr == "" || *resourcePtr == "" || *runtimePtr == "" {
		fmt.Println("You must supply a zip file name, bucket name, function name, handler, ARN, and runtime.")
		os.Exit(0)
	}

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-1"),
		Credentials: credentials.NewSharedCredentials("", "default"),
	})

	svc := lambda.New(sess)

	deleteArgs := &lambda.DeleteFunctionInput{
		FunctionName: functionPtr,
	}

	deleteresult, err := svc.DeleteFunction(deleteArgs)
	if err != nil {
		fmt.Println("Cannot delete function: " + err.Error())
	} else {
		fmt.Println(deleteresult)
	}

	contents, err := ioutil.ReadFile(*zipFilePtr + ".zip")
	if err != nil {
		fmt.Println("Could not read " + *zipFilePtr + ".zip")
		os.Exit(0)
	}

	createCode := &lambda.FunctionCode{
		ZipFile: contents,
	}

	createArgs := &lambda.CreateFunctionInput{
		Code:         createCode,
		FunctionName: functionPtr,
		Handler:      handlerPtr,
		Role:         resourcePtr,
		Runtime:      runtimePtr,
	}

	createResult, err := svc.CreateFunction(createArgs)
	if err != nil {
		fmt.Println("Cannot create function: " + err.Error())
	} else {
		fmt.Println(createResult)
	}
}

