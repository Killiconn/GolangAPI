## Running the OrthodoxAPI locally

# AWS RDS instance must be running
Enter the following command at a command prompt:
`mysql -h orthodoxcred.c9omrl6fspnx.eu-west-1.rds.amazonaws.com -P 3306 -u USERNAME -p`
Replace the `USERNAME` and enter password.

# Getting API running locally you must have Go installed on device
1. Locate this current directory
2. Run: `go run main.go`
3. Listening and serving HTTP on :8080
