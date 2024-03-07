terraform {
  required_providers {
    playground = {
      source = "bendbennett/playground"
    }
  }
}

provider "playground" {
  configurable_attribute = "some-value"
}

resource "playground_resource" "example" {
  configurable_attribute = "some-value"
}

data "playground_datasource" "example" {
  configurable_attribute = "some-value"
}
