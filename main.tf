# Retrieve the latest stack pack version
terraform {
  required_providers {
    ec = {
      source = "elastic/ec"
      version = "0.8.0"
    }
    elasticstack = {
      source = "elastic/elasticstack"
      version = "0.6.2"
    }
  }
}
