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


variable "datasources_zone" {
  type        = string
  description = "Datasources zone name"

  #validation {
  # TODO: Must be a valid domain
  #  condition     = length(var.image_id) > 4 && substr(var.image_id, 0, 4) == "ami-"
  #  error_message = "The image_id value must be a valid AMI id, starting with \"ami-\"."
  #}
}

variable "resources_zone" {
  type        = string
  description = "Resource zone name"

  #validation {
  # TODO: Must be a valid domain
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
