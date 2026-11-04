data "khulnasoft_suppression_rules" "all" {}

output "all_rule" {
  value = data.khulnasoft_suppression_rules.all
}