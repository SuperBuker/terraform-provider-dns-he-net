variable "account" {
  type = object({
    username = string
    password = string
    otp_secret = optional(string, "")
    store_type = optional(string, "encrypted")
    mask_creds = optional(bool, true)
  })
  description = "Credentials for dns.he.net"
  sensitive = true
}


variable "datasources_zone" {
  type        = string
  description = "Datasources zone name"
}

variable "resources_zone" {
  type        = string
  description = "Resource zone name"
}

variable "config_file" {
  type        = string
  description = "Config file path"

  default = "../files/test_config.json"
}