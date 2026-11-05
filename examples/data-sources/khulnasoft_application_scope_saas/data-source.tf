data "khulnasoft_application_scope_saas" "saas" {
  name = "Global"
}

output "scopes" {
  value = data.khulnasoft_application_scope_saas.saas
}

output "name" {
  value = data.khulnasoft_application_scope_saas.saas.name
}

output "categories" {
  value = data.khulnasoft_application_scope_saas.saas.categories
}

