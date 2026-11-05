data "khulnasoft_assurance_custom_script" "example" {
  script_id = "ID of the custom script"
}

output "name" {
  value = data.khulnasoft_assurance_custom_script.example.name
}