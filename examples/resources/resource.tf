terraform {
  required_providers {
    playground = {
      source = "bendbennett/playground"
    }
  }
}

resource "random_string" "example" {
  length = 8
}

resource "playground_resource" "example" {
  configurable_attribute = random_string.example.result
}
