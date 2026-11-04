
data "khulnasoft_response_policies" "all" {}

output "response_policy_id" {
  value = data.khulnasoft_response_policies.all.data
}