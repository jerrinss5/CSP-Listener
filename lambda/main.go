package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// HandleRequest ...
func HandleRequest(ctx context.Context, cspstruct CSPstruct) (string, error) {
	cspReport := cspstruct.CspReport
	log.Printf("line-number:%d, source-file:%s, script-sample:%s, Document-uri:%s, referrer:%s, violated-directive:%s, effective-directive:%s, original-policy:%s, blocked-uri:%s, status-code:%d", cspReport.LineNumber, cspReport.SourceFile, cspReport.ScriptSample, cspReport.DocumentURI, cspReport.Referrer, cspReport.ViolatedDirective, cspReport.EffectiveDirective, cspReport.OriginalPolicy, cspReport.BlockedURI, cspReport.StatusCode)
	return fmt.Sprintf("Report Processed!"), nil
}

func main() {
	lambda.Start(HandleRequest)
}
