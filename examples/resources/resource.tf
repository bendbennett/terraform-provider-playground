terraform {
  required_providers {
    playground = {
      source = "bendbennett/playground"
    }
  }
}

resource "playground_resource" "example" {
    list_nested_attribute = [
      {
        int64_attribute = 9223372036854775807
        list_attribute  = ["list-element", "list-element"]
      },
      {
        int64_attribute = 9223372036854775807
        list_attribute  = ["list-element", "list-element"]
      }
    ]

  list_nested_attribute_custom = [
    {
      int64_attribute = 9223372036854775807
      list_attribute  = ["list-element", "list-element"]
    },
    {
      int64_attribute = 9223372036854775807
      list_attribute  = ["list-element", "list-element"]
    }
  ]
}
