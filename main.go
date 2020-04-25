package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func logReport(w http.ResponseWriter, r *http.Request) {
	var cspstruct CSPstruct
	// EOF at 5 MB body payload - we don't anticipate the payload to be more than 5MB at any point
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 5048576))
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &cspstruct); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	cspReport := cspstruct.CspReport
	log.Println(cspReport)
	log.Printf("line-number:%d, source-file:%s, script-sample:%s, Document-uri:%s, referrer:%s, violated-directive:%s, effective-directive:%s, original-policy:%s, blocked-uri:%s, status-code:%d", cspReport.LineNumber, cspReport.SourceFile, cspReport.ScriptSample, cspReport.DocumentURI, cspReport.Referrer, cspReport.ViolatedDirective, cspReport.EffectiveDirective, cspReport.OriginalPolicy, cspReport.BlockedURI, cspReport.StatusCode)
	fmt.Fprintf(w, "Report processed")
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// fetching host and port from environment if present
func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func main() {
	http.HandleFunc("/report", logReport)
	http.HandleFunc("/health", healthz)

	listenerHost := getEnv("HOST", "")
	listenerPort := getEnv("PORT", "9000")

	fmt.Printf("CSP Report logger service running on %s", listenerPort)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", listenerHost, listenerPort), nil); err != nil {
		log.Fatal(err)
	}
}
