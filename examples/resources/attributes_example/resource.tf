// https://www.terraform.io/plugin/framework/types

terraform {
  required_providers {
    attributes = {
      source = "bendbennett/attributes"
    }
  }
}

provider "attributes" {
  // Simple/primitive attributes.
  bool_attribute = true

  float64_attribute = 1234.5

  int64_attribute = 9223372036854775807

  list_attribute = ["list-element", "list-element"]

  map_attribute = { "map-key-1" : "map-value-1" }

  number_attribute = 1.79769313486231570814527423731704356798070e+1000

  // Treated as an atomic unit. Cannot configure required, optional, computed,
  // PlanModifiers, Validators etc on individual fields.
  object_attribute = {
    bool_attribute    = true
    float64_attribute = 1234.5
    int64_attribute   = 9223372036854775807
    list_attribute    = ["obj-list-element", "obj-list-element"]
    map_attribute     = { "obj-map-key-1" : "obj-map-value-1" }
    number_attribute  = 1.79769313486231570814527423731704356798070e+1000
    set_attribute     = ["obj-set-element-1", "obj-set-element-2"]
    string_attribute  = "obj-string"
  }

  set_attribute = ["set-element-1", "set-element-2"]

  string_attribute = "string"

  // Nested attributes - can contain simple/primitive attributes or other
  // nested attributes.
  // See https://www.terraform.io/plugin/framework/schemas
  single_nested_attribute = {
    bool_attribute    = true
    float64_attribute = 1234.5
  }

  list_nested_attribute = [
    {
      int64_attribute  = 9223372036854775807
      list_attribute   = ["list-element", "list-element"]
      map_attribute    = { "map-key-1" : "map-value-1" }
      number_attribute = 1.79769313486231570814527423731704356798070e+1000
    },
    {
      int64_attribute  = 9223372036854775807
      list_attribute   = ["list-element", "list-element"]
      map_attribute    = { "map-key-1" : "map-value-1" }
      number_attribute = 1.79769313486231570814527423731704356798070e+1000
    }
  ]

  map_nested_attribute = {
    "obj_1" = {
      map_attribute    = { "map-key-1" : "map-value-1" }
      number_attribute = 1.79769313486231570814527423731704356798070e+1000
    },
    "obj_2" = {
      map_attribute    = { "map-key-1" : "map-value-1" }
      number_attribute = 1.79769313486231570814527423731704356798070e+1000
    }
  }

  // Each element in the set must be unique hence true/false for bool_attribute.
  set_nested_attribute = [
    {
      object_attribute = {
        bool_attribute    = true
        float64_attribute = 1234.5
        int64_attribute   = 9223372036854775807
        list_attribute    = ["obj-list-element", "obj-list-element"]
        map_attribute     = { "obj-map-key-1" : "obj-map-value-1" }
        number_attribute  = 1.79769313486231570814527423731704356798070e+1000
        set_attribute     = ["obj-set-element-1", "obj-set-element-2"]
        string_attribute  = "obj-string"
      }
      set_attribute    = ["set-element-1", "set-element-2"]
      string_attribute = "string"
    },
    {
      object_attribute = {
        bool_attribute    = false
        float64_attribute = 1234.5
        int64_attribute   = 9223372036854775807
        list_attribute    = ["obj-list-element", "obj-list-element"]
        map_attribute     = { "obj-map-key-1" : "obj-map-value-1" }
        number_attribute  = 1.79769313486231570814527423731704356798070e+1000
        set_attribute     = ["obj-set-element-1", "obj-set-element-2"]
        string_attribute  = "obj-string"
      }
      set_attribute    = ["set-element-1", "set-element-2"]
      string_attribute = "string"
    }
  ]

  // Blocks
  block {
    bool_attribute = true
    float64_attribute = 1234.5
    int64_attribute = 9223372036854775807
    list_attribute = ["list-element", "list-element"]
    nested_block {
      bool_attribute = true
    }
  }
}

resource "attributes_example" "example" {
  configurable_attribute = "some-value"
}
