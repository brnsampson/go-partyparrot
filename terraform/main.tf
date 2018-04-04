provider "aws" {
  region = "${var.region}"
}

resource "aws_s3_bucket" "terraform_state" {
  bucket = "pplambda"

  versioning {
    enabled = true
  }

  lifecycle {
    prevent_destroy = true
  }
}

terraform {
  backend "s3" {
    bucket = "pplambda"
    key = "terraform/${var.env}/terraform.tfstate"
    region = "${var.region}"
  }
}

data "terraform_remote_state" 