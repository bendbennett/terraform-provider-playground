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

  map_nested_attribute = {
    "one" = {
      int64_attribute = 9223372036854775807
      list_attribute  = ["list-element", "list-element"]
    },
    "two" = {
      int64_attribute = 9223372036854775807
      list_attribute  = ["list-element", "list-element"]
    }
  }

  map_nested_attribute_custom = {
    "one" = {
      int64_attribute = 9223372036854775807
      list_attribute  = ["list-element", "list-element"]
    },
    "two" = {
      int64_attribute = 9223372036854775807
      list_attribute  = ["list-element", "list-element"]
    }
  }

  set_nested_attribute = [
    {
      int64_attribute = 9223372036854775807
      list_attribute  = ["list-element", "list-element"]
    },
    {
      int64_attribute = 9223372036854775807
      list_attribute  = ["list-element", "list-element"]
    }
  ]

  set_nested_attribute_custom = [
    {
      int64_attribute = 9223372036854775807
      list_attribute  = ["list-element", "list-element"]
    },
    {
      int64_attribute = 9223372036854775807
      list_attribute  = ["list-element", "list-element"]
    }
  ]

  single_nested_attribute = {
    int64_attribute = 9223372036854775807
    list_attribute  = ["list-element", "list-element"]
  }

  single_nested_attribute_custom = {
    int64_attribute = 9223372036854775807
    list_attribute  = ["list-element", "list-element"]
  }

  list_nested_block {
    int64_attribute = 9223372036854775807
    list_attribute  = ["list-element", "list-element"]
  }
  list_nested_block {
    int64_attribute = 9223372036854775807
    list_attribute  = ["list-element", "list-element"]
  }

#  list_nested_block_custom {
#    int64_attribute = 9223372036854775807
##    list_attribute  = ["list-element", "list-element"]
#  }
#  list_nested_block_custom {
#    int64_attribute = 9223372036854775807
##    list_attribute  = ["list-element", "list-element"]
#  }

  single_nested_block {
    int64_attribute = 9223372036854775807
    list_attribute  = ["list-element", "list-element"]
  }

  single_nested_block_custom {
    int64_attribute = 9223372036854775807
    list_attribute  = ["list-element", "list-element"]
  }
}
