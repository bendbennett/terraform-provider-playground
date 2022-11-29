terraform {
  required_providers {
    example = {
      source = "bendbennett/playground"
    }
  }
}

data "example_datasource" "example" {
  configurable_attribute = "some-value"
}
