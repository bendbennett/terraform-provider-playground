terraform {
  required_providers {
    playground = {
      source = "bendbennett/playground"
    }
  }
}

resource "playground_resource" "example" {
  configurable_attribute = provider::playground::function("some-value")
  #  configurable_attribute = "some-value"
}

#output "test" {
#  value = provider::example::function("some-value")
#}
