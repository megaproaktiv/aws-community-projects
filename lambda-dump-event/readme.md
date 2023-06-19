# Swap Lambda Function with dump event code

Replace running Lambda function with dump event code.
Then events are written in CloudWatch Logs.

1) save old Lambda function code
`task save`

2) replace Lambda function code with dump event code
`task deploy`

3) run Lambda function with events

4) Restore Lambda function code
`task restore`
