terraform {
  required_providers {
    timeouts = {
      source = "bendbennett/timeouts"
    }
  }
}

resource "timeouts_example" "example" {
  configurable_attribute = "some-values"

#  #  Block
#  timeouts {
#    create = "60m"
#    read   = "30m"
#    update = "12m"
#  }

  #  Attributes
    timeouts = {
      create = "60m"
      read = "30m"
    }
}
