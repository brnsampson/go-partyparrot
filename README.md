# go-partyparrot
Defines an aws lambda function to convert slack text into party parrots

Inspired by https://github.com/Shrugs/partyparrot, we're doing basically the same thing except in golang and running as an aws lambda function.

I've also moved to a 4x4 grid for each character as personal preference, as friends that use mobile are destroyed by this and every little bit helps.

However, the main motivation was to save $$$, as paying for even a micro ec2 instance is waaaay more expensive for my use case.

To build:

1. Build the code (I use windows and you'll need to setup your go environment and navigate to your go root before this):
    go get github.com/brnsampson/go-partyparrot/pplambda
    $env:GOOS="linux"
    $env:GOARCH="amd64"
    go build -o pplambda src\github.com\brnsampson\go-partyparrot\pplambda
    (or just `GOOS=linux GOARCH=amd64 go build -o pplambda src\github.com\brnsampson\go-partyparrot\pplambda` if you're running linux)
    
2. Download the lambda packages
    go get github.com/aws/aws-lambda-go/events
    go get github.com/aws/aws-lambda-go/lambda

3.a. (only if you're on windows like me :'(  ) get the lambda zip creater
    go.exe get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
    ~\go\bin\build-lambda-zip.exe -o .\pplambda.zip .\pplambda

3.b. (for linux)
    chmod +x pplambda
    zip pplambda.zip pplambda


To deploy:

1. Get an AWS account. See https://aws.amazon.com/
2. Create the lambda function by going to Services > compute/lambda and clicking "Create function"
3. Add the API Gateway as a trigger for your function
4. Upload your zip file as the function code
5. Create an environment variable named `SLACK_TEAM_TOKENS` and add all of the slack tokens you wish to use delimited by a colon (:)
6. Save your function
7. Add a test that looks like this:
    {
      "Body": "token=<INSERT YOUR TOKEN HERE>&team_id=WTFBBQ&team_domain=dangerouspeople&channel_id=NUMBERS&channel_name=directmessage&user_id=DANGEROUSPERSON&user_name=cooldude&command=%2Fparty&text=test&response_url=https%3A%2F%2Fhooks.slack.com%2Fcommands%2FT0BKPG7MM%2F320588249252%2FpVJ4B4U3ndfluZaY4UMkWvZ2"
    }
8. Your test _should_ work at this point.
9. You should have been prompted to set up a minimal api gateway resource/method if you didn't already have one. If you already had one, I assume you should know how to add a resource for your new function.
10. To get the URL you can use, go to API Gateway > APIs > <The api created for this> > Stages and then click on whatever stage you created or added your lambda function to.
11. The URL to use will be the "Invoke URL" at the top of the pane. The URL to put into slack is <That URL>/<Whatever the resource is under that stage>


Whew! Too much detail. But hopefully it'll help someone out some day.
