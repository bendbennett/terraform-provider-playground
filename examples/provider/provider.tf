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
