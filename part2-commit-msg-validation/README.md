## Commit Message Validation

This application verifies that a commit message is valid based on the following parameters. 

1. The first field of the commit message is the ticket number. (ex KT-1 or KT-1: are accepted)
2. The ticket associated with the commit message is in the status `Approved`
3. The ticket associated wtih the commit is assigned to someone.

If any of the above are not true then the check will fail.

This application uses the following environment variables for authenticating with Jira.

JIRA_USER_NAME: This should be set to your Jira username
JIRA_API_KEY: This should be set to your Jira API key

To print out the application usage you can run the application with the `-help` option.

```
kevinhartwig@kevins-mbp part2-commit-msg-validation % ./check-commit-message -help
Usage of ./check-commit-message:
  -commit-msg string
        The commit message you would like to verify
```

To check that a commit message is valid using the `-commit-msg` paramter and pass in your commit message as shown

```
kevinhartwig@kevins-mbp part2-commit-msg-validation % ./check-commit-message -commit-msg "KT-1: This is a test of a valid ticket"
Commit message for ticket id KT-1 is valid
```

An example of a ticket that is not assigned can be found below.

```
kevinhartwig@kevins-mbp part2-commit-msg-validation % ./check-commit-message -commit-msg "KT-1: This is a test of a valid ticket"
Commit message for ticket id KT-1 is valid
```

An example of a ticket that is not in the `Approved` status can be found below.
```
kevinhartwig@kevins-mbp part2-commit-msg-validation % ./check-commit-message -commit-msg "KT-1: This is a test of a valid ticket"
Commit message for ticket id KT-1 is valid
```