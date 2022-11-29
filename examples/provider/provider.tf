terraform {
  required_providers {
    example = {
      source = "bendbennett/playground"
    }
  }
}

provider "example" {
  configurable_attribute = "some-value"
}
