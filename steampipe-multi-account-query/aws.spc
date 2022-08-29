connection "aws" {
  plugin = "aws"
  regions = ["eu-central-1"]
}

connection "megaproaktiv" {
 type        = "aggregator"
 plugin      = "aws"
 connections = ["megaproaktiv*"]
}

connection "megaproaktiv_prod" {
  plugin = "aws"
  regions = ["eu-central-1"]
  profile = "megaproaktiv_prod"
}

connection "megaproaktiv_test" {
  plugin = "aws"
  regions = ["eu-central-1"]
  profile = "megaproaktiv_test"
}

connection "megaproaktiv_dev" {
  plugin = "aws"
  regions = ["eu-central-1"]
  profile = "megaproaktiv_dev"
}