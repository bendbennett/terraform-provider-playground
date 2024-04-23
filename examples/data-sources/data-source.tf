terraform {
  required_providers {
    playground = {
      source = "bendbennett/playground"
    }
  }
}

data "playground_datasource" "example" {
  configurable_attribute = "some-value"
}
