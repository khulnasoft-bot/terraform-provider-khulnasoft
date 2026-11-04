#Listing of single API Key details
data "khulnasoft_khulnasoft_api_keys" "single" {
  id = khulnasoft_khulnasoft_api_key.terraform_api_key.id
}

output "api_key_single_id" {
  value = data.khulnasoft_khulnasoft_api_keys.single.apikeys[0].id
}
output "group_id" {
  value = data.khulnasoft_khulnasoft_api_keys.single.apikeys[0].group_id
}
output "expiration" {
  value = data.khulnasoft_khulnasoft_api_keys.single.apikeys[0].expiration
}
output "permission_ids" {
  value = data.khulnasoft_khulnasoft_api_keys.single.apikeys[0].permission_ids
}
output "roles" {
  value = data.khulnasoft_khulnasoft_api_keys.single.apikeys[0].roles
}

output "secret" {
  value     = khulnasoft_khulnasoft_api_key.terraform_api_key.secret //resource.resource_name.secret
  sensitive = true
}

#Listing of all API Keys

data "khulnasoft_khulnasoft_api_keys" "list" {
  limit  = 10
  offset = 0
}

output "api_key_list" {
  value     = data.khulnasoft_khulnasoft_api_keys.list
  sensitive = true
}
