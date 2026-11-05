data "khulnasoft_image_assurance_policy" "image_assurance_policy" {
  name = "ImageAssurancePolicy"
}

output "container_runtime_policy_details" {
  value = data.khulnasoft_image_assurance_policy.image_assurance_policy
}