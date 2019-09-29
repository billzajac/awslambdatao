* https://www.alexedwards.net/blog/serverless-api-with-go-and-aws-lambda

* Environment variables
    * FIRST: Manually set "ACCOUNT_ID=...." in ENV.sh
```
echo "ZONE=us-east-1" > ENV.sh
source ENV.sh
```

* Initial setup of role
```
ROLE=lambda-tao-executor
POLICY_FILE=file://./trust-policy.json

aws iam create-role --role-name ${ROLE} --assume-role-policy-document ${POLICY_FILE} --profile admin
aws iam attach-role-policy --role-name ${ROLE} --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole --profile admin
```

* Deploy with Ansible
```
ansible-playbook deploy.yml
```

* Test
```
aws lambda invoke --function-name tao /dev/stdout
```

* Create API gateway
```
aws apigateway create-rest-api --name taopassages|python -c 'import sys, json; print "REST_API_ID=" + json.load(sys.stdin)["id"]' >> ENV.sh
source ENV.sh
```

* Get the id of the root API resource ("/")
```
aws apigateway get-resources --rest-api-id ${REST_API_ID}|python -c 'import sys, json; print "ROOT_PATH_ID=" + json.load(sys.stdin)["items"][0]["id"]' >> ENV.sh
source ENV.sh
```

* Create /tao
```
aws apigateway create-resource --rest-api-id ${REST_API_ID} --parent-id ${ROOT_PATH_ID} --path-part tao --profile admin | python -c 'import sys, json; print "RESOURCE_ID=" + json.load(sys.stdin)["id"]' >> ENV.sh
source ENV.sh
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

* Add/fix permissions, but first generate a GUID
```
python -c 'import sys,uuid; sys.stdout.write("GUID="+str(uuid.uuid4()))' >> ENV.sh
source ENV.sh
aws lambda add-permission --function-name tao --statement-id ${GUID} \
--action lambda:InvokeFunction --principal apigateway.amazonaws.com \
--source-arn arn:aws:execute-api:${ZONE}:${ACCOUNT_ID}:${REST_API_ID}/*/*/*
```

* Test API gateway
```
aws apigateway test-invoke-method --rest-api-id ${REST_API_ID} --resource-id ${RESOURCE_ID} --http-method "GET" --path-with-query-string "/tao"
```

* Create Deployment
```
aws apigateway create-deployment --rest-api-id ${REST_API_ID} --stage-name staging
```

* Test GET with curl
```
curl https://${REST_API_ID}.execute-api.us-east-1.amazonaws.com/staging/tao
```

* Tail CloudWatch logs
    * https://github.com/lucagrulla/cw
```
cw tail -f /aws/lambda/tao
```
