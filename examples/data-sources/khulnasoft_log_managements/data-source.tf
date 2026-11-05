data "khulnasoft_log_managements" "logmanagement" {
  name = "CloudWatch"
}

output "log_managements_data_source_name" {
  value = data.khulnasoft_log_managements.logmanagement.name
}

output "log_managements_data_source_enable" {
  value = data.khulnasoft_log_managements.logmanagement.enable
}