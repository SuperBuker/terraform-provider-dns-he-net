variable "account" {
  type = object({
    username   = string
    password   = string
    otp_secret = optional(string, "")
    store_type = optional(string, "encrypted")
    mask_creds = optional(bool, true)
  })
  description = "Credentials for dns.he.net"
  #sensitive = true
}


variable "datasources_domain_zone" {
  type        = string
  description = "Datasources domain zone name"

  #validation {
  # TODO: Must be a valid domain zone
  #  condition     = length(var.image_id) > 4 && substr(var.image_id, 0, 4) == "ami-"
  #  error_message = "The image_id value must be a valid AMI id, starting with \"ami-\"."
  #}
}

variable "datasources_domain_zone_pending_delegation" {
  type        = string
  description = "Datasources domain zone pending delegation name"

  #validation {
  # TODO: Must be a valid domain zone
  #  condition     = length(var.image_id) > 4 && substr(var.image_id, 0, 4) == "ami-"
  #  error_message = "The image_id value must be a valid AMI id, starting with \"ami-\"."
  #}
}

variable "datasources_arpa_zone" {
  type        = string
  description = "Datasources ARPA zone name"

  #validation {
  # TODO: Must be a valid ARPA zone
  #  condition     = length(var.image_id) > 4 && substr(var.image_id, 0, 4) == "ami-"
  #  error_message = "The image_id value must be a valid AMI id, starting with \"ami-\"."
  #}
}

variable "datasources_arpa_domain_example" {
  type        = string
  description = "Datasources ARPA example domain"

  #validation {
  # TODO: Must be a valid domain within ARPA zone
  #  condition     = length(var.image_id) > 4 && substr(var.image_id, 0, 4) == "ami-"
  #  error_message = "The image_id value must be a valid AMI id, starting with \"ami-\"."
  #}
}

variable "datasources_network_prefix" {
  type        = string
  description = "Datasources network prefix name"

  #validation {
  # TODO: Must be a valid Network prefix
  #  condition     = length(var.image_id) > 4 && substr(var.image_id, 0, 4) == "ami-"
  #  error_message = "The image_id value must be a valid AMI id, starting with \"ami-\"."
  #}
}

variable "resources_domain_zone" {
  type        = string
  description = "Resource domain zone name"

  #validation {
  # TODO: Must be a valid domain zone
  #  condition     = length(var.image_id) > 4 && substr(var.image_id, 0, 4) == "ami-"
  #  error_message = "The image_id value must be a valid AMI id, starting with \"ami-\"."
  #}
}

variable "resources_arpa_zone" {
  type        = string
  description = "Resource ARPA zone name"

  #validation {
  # TODO: Must be a valid ARPA zone
  #  condition     = length(var.image_id) > 4 && substr(var.image_id, 0, 4) == "ami-"
  #  error_message = "The image_id value must be a valid AMI id, starting with \"ami-\"."
  #}
}


variable "config_file" {
  type        = string
  description = "Config file path"

  #validation {
  # TODO: Must be a valid path
  #  condition     = length(var.image_id) > 4 && substr(var.image_id, 0, 4) == "ami-"
  #  error_message = "The image_id value must be a valid AMI id, starting with \"ami-\"."
  #}
}
