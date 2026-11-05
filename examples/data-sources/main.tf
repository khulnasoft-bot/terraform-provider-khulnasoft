terraform {
  required_providers {
    khulnasoft = {
      //      version = "0.8.30"
      source = "khulnasoft/khulnasoft"
    }
  }
}

provider "khulnasoft" {
  username = "admin"
  khulnasoft_url = "https://khulnasoft.com"
  password = "@password"
}

data "khulnasoft_users" "testusers" {
}

output "name" {
  value = data.khulnasoft_users.testusers
}

data "khulnasoft_integration_registry" "testregistries" {
  name = "samplename"
}

output "registries" {
  value = data.khulnasoft_integration_registry.testregistries
}

data "khulnasoft_service" "test-svc" {
  name = "test-svc"
}

output "service" {
  value = data.khulnasoft_service.test-svc
}

data "khulnasoft_enforcer_groups" "testegdata" {
  group_id = "default"
}

output "enforcergroups" {
  value = data.khulnasoft_enforcer_groups.testegdata
}

data "khulnasoft_image" "test" {
  registry   = "Docker Hub"
  repository = "elasticsearch"
  tag        = "7.10.1"
}

output "image" {
  value = data.khulnasoft_image.test
}

data "khulnasoft_container_runtime_policy" "test" {
  name = "test-container-runtime-policy"
}

output "test-crp" {
  value = data.khulnasoft_container_runtime_policy.test
}

data "khulnasoft_function_runtime_policy" "test" {
  name = "test-function-runtime-policy"
}

output "test-frp" {
  value = data.khulnasoft_function_runtime_policy.test
}

data "khulnasoft_host_runtime_policy" "test" {
  name = "test-host-runtime-policy"
}

output "test-hrp" {
  value = data.khulnasoft_host_runtime_policy.test
}


data "khulnasoft_gateways" "testgateways" {
}

output "gateways" {
  value = data.khulnasoft_gateways.testgateways
}

data "khulnasoft_image_assurance_policy" "default-iap" {
  name = "DTA"
}

output "image-assurance" {
  value = data.khulnasoft_image_assurance_policy.default-iap
}

data "khulnasoft_permissions_sets" "testpermissionsset" {}

output "permissions_sets" {
  value = data.khulnasoft_permissions_sets.testpermissionsset
}

output "permissions_sets_names" {
  value = data.khulnasoft_permissions_sets.testpermissionsset[*].permissions_sets[*].name
}


data "khulnasoft_host_assurance_policy" "default-hap" {
  name = "Default"
}

output "host-assurance" {
  value = data.khulnasoft_host_assurance_policy.default-hap
}

data "khulnasoft_function_assurance_policy" "default-fap" {
  name = "Default"
}

output "function-assurance" {
  value = data.khulnasoft_function_assurance_policy.default-fap
}

data "khulnasoft_kubernetes_assurance_policy" "default-kap" {
  name = "Default"
}

output "kubernetes-assurance" {
  value = data.khulnasoft_kubernetes_assurance_policy.default-kap
}

data "khulnasoft_application_scope" "default" {
  name = "Global"
}

output "scopes" {
  value = data.khulnasoft_application_scope.default
}