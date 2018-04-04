resource "aws_api_gateway_rest_api" "PParrotAPI" {
  name = "PParrotAPI"
  description = "public endpoint for partyparrot api"
}

resource "aws_api_gateway_resource" "pparrot_convert" {
  rest_api_id = "${aws_api_gateway_rest_api.PParrotAPI.id}"
  parent_id = "${aws_api_gateway_rest_api.PParrotAPI.root_resource_id}"
  path_part = "convert"
}

resource "aws_api_gateway_method" "pparrot_convert_proxy" {
  rest_api_id = "${aws_api_gateway_rest_api.PParrotAPI.id}"
  resource_id = "${aws_api_gateway_resource.pparrot_convert.id}"
  http_method = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "pplambda" {
  rest_api_id = "${aws_api_gateway_rest_api.PParrotAPI.id}"
  resource_id = "${aws_api_gateway_method.pparrot_convert_proxy.resource_id}"
  http_method = "${aws_api_gateway_method.pparrot_convert_proxy.http_method}"

  integration_http_method = "POST"
  type = "AWS_PROXY"
  uri = "${aws_lambda_function.pplambda.invoke_arn}"
}

resource "aws_api_gateway_deployment" "PParrotAPI_prod" {
  depends_on = [
    "aws_api_gateway_integration.pplambda",
  ]

  rest_api_id = "${aws_api_gateway_rest_api.PParrotAPI.id}"
  stage_name  = "prod"
}

resource "aws_api_gateway_deployment" "PParrotAPI_test" {
  depends_on = [
    "aws_api_gateway_integration.pplambda",
  ]

  rest_api_id = "${aws_api_gateway_rest_api.PParrotAPI.id}"
  stage_name  = "test"
}

resource "aws_api_gateway_domain_name" "api_whobe_us" {
  domain_name = "api.whobe.us"
  certificate_arn = "${var.cert_arn}"
}

resource "aws_api_gateway_base_path_mapping" "api_whobe_us_pparrot" {
  api_id = "${aws_api_gateway_rest_api.PParrotAPI.id}"
  stage_name = "${aws_api_gateway_deployment.PParrotAPI_prod.stage_name}"
  domain_name = "${aws_api_gateway_domain_name.api_whobe_us.domain_name}"
  base_path = "partyparrot"
}

output "base_url" {
  value = "${aws_api_gateway_deployment.PParrotAPI_test.invoke_url}"
}