data "khulnasoft_scanner_group" "all" {}

output "scanner_group_names" {
  value = [for sg in data.khulnasoft_scanner_group.all.scanner_groups : sg.name]
}

data "khulnasoft_scanner_group" "example" {
  name = "terraform-test"
}

output "scanner_group_desc" {
  value = data.khulnasoft_scanner_group.example.description
}