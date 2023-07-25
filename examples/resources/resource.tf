terraform {
  required_providers {
    example = {
      source = "bendbennett/playground"
    }
  }
}

resource "example_resource" "example" {
#  configurable_attribute = "some-value"
}
