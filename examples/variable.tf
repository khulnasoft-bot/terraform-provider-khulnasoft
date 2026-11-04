variable "khulnasoft_username" {
  description = "Khulnasoft Username"
  type        = string
  sensitive   = true
}

variable "khulnasoft_url" {
  description = "Khulnasoft Console URL"
  type        = string
}

variable "khulnasoft_password" {
  description = "Khulnasoft password"
  type        = string
  sensitive   = true
}

variable "aws_access_key" {
  description = "AWS Access Key"
  type        = string
  sensitive   = true
}

variable "aws_secret_key" {
  description = "AWS Secret Key"
  type        = string
  sensitive   = true
}

variable "aws_region" {
  description = "AWS Region"
  type        = string
}

variable "aws_log_group" {
  description = "AWS Log Group"
  type        = string
}

variable "enable_log_management" {
  description = "Enable Log Management"
  type        = bool
  default     = true
}