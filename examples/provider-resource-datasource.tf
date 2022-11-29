terraform {
  required_providers {
    example = {
      source = "bendbennett/playground"
    }
  }
}

provider "example" {
  configurable_attribute = "some-value"
}

resource "example_resource" "example" {
  configurable_attribute = "some-value"
}

data "example_datasource" "example" {
  configurable_attribute = "some-value"
}
