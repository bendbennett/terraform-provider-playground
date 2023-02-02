terraform {
  required_providers {
    playground = {
      source = "bendbennett/playground"
    }
  }
}

resource "playground_example" "example" {
}
