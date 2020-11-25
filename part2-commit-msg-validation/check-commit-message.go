package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/tidwall/gjson"
)

// Split on these characters
func Split(r rune) bool {
	return r == ':' || r == ' '
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func QueryJira(ticketId string, credentials string) string {

	queryUrl := fmt.Sprintf("https://kevinhartwig.atlassian.net/rest/api/2/issue/%s", ticketId)
	credentialsHeader := fmt.Sprintf("Basic %s", credentials)

	client := &http.Client{}

	req, err := http.NewRequest("GET", queryUrl, nil)
	// ...
	if err != nil {
		exitErrorf("Unable to create new http request: ", err)
	}
	req.Header.Add("Authorization", credentialsHeader)
	req.Header.Add("Accept", "application/json")

	result, err := client.Do(req)

	if err != nil {
		exitErrorf("Unable to complete request: ", err)
	}

	if result.StatusCode > 399 {
		exitErrorf("Error retrieving ticketId: %s statusCode: %d", ticketId, result.StatusCode)
	}

	body, err := ioutil.ReadAll(result.Body)

	if err != nil {
		exitErrorf("Unable to ready body of the response: ", err)
	}

	return string(body)
}

func ValidateTicket(ticketJson string) {
	ticketId := gjson.Get(ticketJson, "key").String()
	status := gjson.Get(ticketJson, "fields.status.name").String()
	assignee := gjson.Get(ticketJson, "fields.assignee").String()

	if status != "Approved" {
		exitErrorf("Status is not in approved state: %s", status)
	}

	if assignee == "" {
		exitErrorf("Ticket is unassigned")
	}

	fmt.Printf("Commit message for ticket id %s is valid\n", ticketId)
}

func main() {
	commitPtr := flag.String("commit-msg", "", "The commit message you would like to verify")

	flag.Parse()

	// Verify required fields exist

	if *commitPtr == "" {
		flag.Usage()
		os.Exit(1)
	}

	commitMsg := *commitPtr

	ticketId := strings.FieldsFunc(commitMsg, Split)[0]
	userId := os.Getenv("JIRA_USER_NAME")
	apiKey := os.Getenv("JIRA_API_KEY")

	if userId == "" ||
		apiKey == "" {
		exitErrorf("User ID or API Key undefined.")
	}

	credentialsString := fmt.Sprintf("%s:%s", userId, apiKey)
	credentialsEnc := base64.StdEncoding.EncodeToString([]byte(credentialsString))

	responseJson := QueryJira(ticketId, credentialsEnc)

	ValidateTicket(responseJson)
}
