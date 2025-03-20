terraform {
  required_version = ">= 1.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 4.67.0"
    }
  }

}

provider "aws" {
  region = "eu-west-1"
}

locals {

  secret_data = {
    one = {
      name  = "secret_one"
      value = "one"
    }
    two = {
      name  = "secret_two"
      value = "two"
    }
    three = {
      name  = "secret_three"
      value = "three"
    }
  }
}

resource "aws_secretsmanager_secret" "pk_secrets" {
  for_each                = local.secret_data
  name                    = "mysecret-${each.key}"
  recovery_window_in_days = 0
}

resource "aws_secretsmanager_secret_version" "git_secrets_version" {
  for_each  = aws_secretsmanager_secret.pk_secrets
  secret_id = each.value.id
  secret_string = jsonencode({
    name  = "${local.secret_data[each.key].name}"
    value = "${local.secret_data[each.key].value}"
  })
}
