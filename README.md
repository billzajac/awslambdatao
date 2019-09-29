* https://www.alexedwards.net/blog/serverless-api-with-go-and-aws-lambda

* Initial setup of role
```
aws iam create-role --role-name lambda-tao-executor --assume-role-policy-document file://./trust-policy.json --profile admin

aws iam attach-role-policy --role-name lambda-tao-executor --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole --profile admin
```

* ENV VARIABLES
```
ACCOUNT_ID=....
ZONE=us-east-1
```

* Deploy with Ansible
```
ansible-playbook deploy.yml
```




* Deploy
```
aws lambda create-function --function-name tao --runtime go1.x \
--role arn:aws:iam::${ACCOUNT_ID}:role/lambda-tao-executor \
--handler main --zip-file fileb:///tmp/main.zip
```

* Test
```
# aws lambda invoke --function-name tao /tmp/output.json && cat /tmp/output.json
aws lambda invoke --function-name tao /dev/stdout
```

* Rebuild and redeploy
```
env GOOS=linux GOARCH=amd64 go build -o /tmp/main
zip -j /tmp/main.zip /tmp/main
aws lambda update-function-code --function-name tao --zip-file fileb:///tmp/main.zip
```

* Create API gateway
```
aws apigateway create-rest-api --name taopassages
REST_API_ID=....
```

* Get the id of the root API resource ("/")
```
aws apigateway get-resources --rest-api-id ${REST_API_ID}
ROOT_PATH_ID=....
```

* Create /tao
```
aws apigateway create-resource --rest-api-id ${REST_API_ID} --parent-id ${ROOT_PATH_ID} --path-part tao --profile admin
RESOURCE_ID=....
```

* Configure API gateway to respond to ANY HTTP method
```
aws apigateway put-method --rest-api-id ${REST_API_ID} \
--resource-id ${RESOURCE_ID} --http-method ANY \
--authorization-type NONE
```

* Connect the gateway to proxy with POST to the lambda function
```
aws apigateway put-integration --rest-api-id ${REST_API_ID} \
--resource-id ${RESOURCE_ID} --http-method ANY --type AWS_PROXY \
--integration-http-method POST \
--uri arn:aws:apigateway:${ZONE}:lambda:path/2015-03-31/functions/arn:aws:lambda:${ZONE}:${ACCOUNT_ID}:function:tao/invocations
```

* Test API gateway
```
aws apigateway test-invoke-method --rest-api-id ${REST_API_ID} --resource-id ${RESOURCE_ID} --http-method "GET"
```

* Add/fix permissions on 
    * First build a GUID: https://www.guidgenerator.com/
        GUID=....
```
aws lambda add-permission --function-name tao --statement-id ${GUID} \
--action lambda:InvokeFunction --principal apigateway.amazonaws.com \
--source-arn arn:aws:execute-api:${ZONE}:${ACCOUNT_ID}:${REST_API_ID}/*/*/*
```

* Test API gateway
```
aws apigateway test-invoke-method --rest-api-id ${REST_API_ID} --resource-id ${RESOURCE_ID} --http-method "GET"
```

* Now fix the response for API gateway from lambda (in the Go code)
```
go get github.com/aws/aws-lambda-go/events

# Update the code to use the new structs
vi tao/main.go

cd tao && \
env GOOS=linux GOARCH=amd64 go build -o /tmp/main && \
zip -j /tmp/main.zip /tmp/main && \
aws lambda update-function-code --function-name tao --zip-file fileb:///tmp/main.zip
```

* Test API gateway
```
aws apigateway test-invoke-method --rest-api-id ${REST_API_ID} --resource-id ${RESOURCE_ID} --http-method "GET" --path-with-query-string "/tao"
```

* Create Deployment
```
aws apigateway create-deployment --rest-api-id ${REST_API_ID} \
--stage-name staging
```

* Test GET with curl
```
curl https://${REST_API_ID}.execute-api.us-east-1.amazonaws.com/staging/tao
```

* Tail CloudWatch logs
    * https://github.com/lucagrulla/cw
```
# aws logs filter-log-events --log-group-name /aws/lambda/tao \
# --filter-pattern "ERROR"
cw tail -f /aws/lambda/tao
```
