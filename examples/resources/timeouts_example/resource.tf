terraform {
  required_providers {
    timeouts = {
      source = "bendbennett/timeouts"
    }
  }
}

resource "timeouts_example" "example" {
  configurable_attribute = "some-value"

#  Block
  timeouts {
    create = "60m"
    read = "30m"
  }

#  Attributes
#  timeouts = {
#    create = "60m"
#    read = "30m"
#  }
}
