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


variable "datasources_domain_zone" {
  type        = string
  description = "Datasources domain zone name"
}

variable "datasources_domain_zone_pending_delegation" {
  type        = string
  description = "Datasources domain zone pending delegation name"
}

variable "datasources_arpa_zone" {
  type        = string
  description = "Datasources ARPA zone name"
}

variable "datasources_arpa_domain_example" {
  type        = string
  description = "Datasources ARPA example domain"
}

variable "datasources_network_prefix" {
  type        = string
  description = "Datasources network prefix name"
}

variable "resources_domain_zone" {
  type        = string
  description = "Resource domain zone name"
}

variable "resources_arpa_zone" {
  type        = string
  description = "Resource ARPA zone name"
}

variable "config_file" {
  type        = string
  description = "Config file path"

  default = "../files/test_config.json"
}