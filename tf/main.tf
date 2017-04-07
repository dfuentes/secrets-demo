provider "aws" {}

module "credstash" {
  source = "github.com/dfuentes/tf-credstash"
  key_alias = "alias/stealth-key-dev"
  table_name = "stealth-dev"
}
