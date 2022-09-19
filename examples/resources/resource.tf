terraform {
  required_providers {
    playground = {
      source = "bendbennett/playground"
    }
  }
}

resource "playground_resource" "example" {
  configurable_attribute = "some-value"
}
