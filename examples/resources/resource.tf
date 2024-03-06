terraform {
  required_providers {
    example = {
      source = "bendbennett/playground"
    }
  }
}

resource "example_resource" "example" {
  configurable_attribute = provider::example::function("some-value")
}

#output "test" {
#  value = provider::example::function("some-value")
#}
