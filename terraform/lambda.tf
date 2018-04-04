resource "aws_lambda_function" "pplambda" {
  function_name = "pplambda"
  filename = "${var.deploy_file}"

  handler = "pplambda"

  role = "${aws_iam_role.lambda_exec.arn}"
  description = "Shitposts parrots to slack"
  runtime = "go1.x"
  environment = {
    variables = {
      SLACK_TEAM_TOKENS = "${slack_tokens}"
    }
  }
}

resource "aws_iam_role" "lambda_exec" {
  name = "lambda_role"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Effect": "Allow",
      "Principal": {
        "Service": [
          "lambda.amazonaws.com",
          "edgelambda.amazonaws.com"
        ]
      },
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_permission" "apigw_prod" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.pplambda.arn}"
  principal     = "apigateway.amazonaws.com"

  # The /*/* portion grants access from any method on any resource
  # within the API Gateway "REST API".
  source_arn = "${aws_api_gateway_deployment.PParrotAPI_prod.execution_arn}/*/*"
}