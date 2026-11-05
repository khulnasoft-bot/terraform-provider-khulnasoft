---
page_title: "khulnasoft_permission_set_saas Resource - terraform-provider-khulnasoft"
subcategory: ""
description: |-
  The khulnasoft_permission_set_saas resource manages your Permission Set within Khulnasoft SaaS platform.
---

# khulnasoft_permission_set_saas (Resource)

The `khulnasoft_permission_set_saas` resource manages your Permission Set within Khulnasoft SaaS platform.

## Example Usage

```terraform
resource "khulnasoft_permission_set_saas" "example" {
  name        = "my_saas_perm_set"
  description = "Test Permissions Sets for SaaS"
  actions = [
    "account_mgmt.groups.read",
    "cspm.cloud_accounts.read",
    "cnapp.inventory.read",
    "cnapp.insights.read",
    "cnapp.dashboards.read"
  ]
}
```

## Schema
### Required

- `name` (String) Name of the permission set
- `actions` (List of String) List of allowed actions for the permission set

## Optional

- `description` (String) Description of the permission set

## Read-Only

- `id` (String) The ID of this resource