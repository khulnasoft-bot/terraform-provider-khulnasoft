## 0.15.2 (Unreleased)

BUG FIXES:

* `khulnasoft_container_runtime_policy`: Preserve `allowed_executables` while upgrading legacy state layouts.
* `khulnasoft_permission_set_saas` and `khulnasoft_role_mapping_saas`: Use current SaaS access-management endpoints and normalize server-provided action names.
* `khulnasoft_service` and `khulnasoft_firewall_policy`: Canonicalize omitted `anywhere` firewall resources as `0.0.0.0/0` to prevent persistent plan differences.
* SaaS development authentication: Route API-key provisioning through the development endpoint.

## 0.15.1 (June 24, 2026)

BUG FIXES:

* `khulnasoft_application_scope_saas`: Fix "unexpected end of JSON input" error on all CRUD operations. The SaaS code path was routing to an incorrect API endpoint. Now uses the auto-resolved CSP URL with the correct `/api/v2/access_management/scopes` path. No user configuration change required. ([#388](https://github.com/khulnasoft/terraform-provider-khulnasoft/pull/388))

## 0.1.0 (Unreleased)

BACKWARDS INCOMPATIBILITIES / NOTES:
