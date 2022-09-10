terraform {
  required_providers {
    timeouts = {
      source = "bendbennett/timeouts"
    }
  }
}

resource "timeouts_example" "example" {
  configurable_attribute = "some-value"

  timeouts {
    create = "60m"
  }
}
