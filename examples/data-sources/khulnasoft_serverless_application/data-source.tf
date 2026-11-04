data "khulnasoft_serverless_applications" "serverless_application" {}

output "acknowledges" {
  value = data.khulnasoft_serverless_applications.serverless_application
}