data "khulnasoft_permissions_sets_saas" "example" {}

output "permissions_sets" {
  value = data.khulnasoft_permissions_sets_saas.example
}

output "permissions_sets_names" {
  value = data.khulnasoft_permissions_sets_saas.example[*].permissions_sets[*].name
}

output "dashboard_permissions" {
  value = [
    for ps in data.khulnasoft_permissions_sets_saas.example.permissions_sets : ps.name
    if contains(ps.actions, "cnapp.dashboards.read")
  ]
}