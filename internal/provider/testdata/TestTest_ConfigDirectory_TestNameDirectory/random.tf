# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

terraform {
  required_providers {
    random = {
      source = "registry.terraform.io/hashicorp/random"
      version = "3.6.1"
    }
  }
}

provider "random" {}

resource "random_string" "test" {
  length = 8

  numeric = false
}

resource "playground_resource" "example" {
  configurable_attribute = random_string.test.result
}
