data "khulnasoft_response_policy_config" "config" {
}

output "response_policy_config_id" {
  value = data.khulnasoft_response_policy_config.config.triggers
}