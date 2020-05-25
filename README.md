# golang-cli-tool

[![Build Status](https://travis-ci.org/Manuhmutua/golang-cli-tool.svg?branch=master)](https://travis-ci.org/Manuhmutua/golang-cli-tool)

The question required me to create a CLI that allows the user to list all the resources below:
- IAM users
- EC2 instances
- RDS instances
- S3 buckets
- ECR repos
- ECS clusters

The credentials are to be stored in a file on the user’s machine. The file is a simple text file in the following format:
```sh
AWS_ACCESS_KEY_ID = AKIAXXXXXXX
AWS_SECRET_ACCESS_KEY = XXXXXXXXXXX
AWS_REGION = xx_xxxxx
```

The application should take an input flag corresponding to the location of the credentials file, handle user input errors gracefully and provide a help menu.

## Running the application