---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "playground Provider"
subcategory: ""
description: |-
  
---

# playground Provider



## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `configurable_attribute` (String)
