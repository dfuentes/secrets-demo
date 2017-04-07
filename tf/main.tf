provider "aws" {}

module "credstash" {
  source = "github.com/dfuentes/tf-credstash"
}
