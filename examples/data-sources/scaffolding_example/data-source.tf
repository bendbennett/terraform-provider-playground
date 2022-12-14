terraform {
  required_providers {
    timeouts = {
      source = "bendbennett/timeouts"
    }
  }
}

data "timeouts_example" "example" {
  configurable_attribute = "some-value"

#    #  Block
#    timeouts {
#      create = "10m"
#      read   = "30m"
#    }

  #  Attributes
  timeouts = {
    read = "30m"
  }
}