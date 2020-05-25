#!/bin/sh
printf '%s\n%s\n%s\n%s' "$AWS_ACCESS_KEY_ID" "$AWS_SECRET_ACCESS_KEY" "$AWS_REGION" "json"| docker run --rm -i amazon/aws-cli configure