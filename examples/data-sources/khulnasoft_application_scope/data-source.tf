data "khulnasoft_application_scope" "default" {
  name = "Global"
}

output "scopes" {
  value = data.khulnasoft_application_scope.default
}

 output "codebuild_config" {
   value = [
     for category in data.khulnasoft_application_scope.default.categories : [
       for artifact in category.artifacts : artifact.codebuild if artifact.codebuild != null
     ] if category.artifacts != null
   ][0][0]
}
