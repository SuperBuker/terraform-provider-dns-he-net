variable "dhn_username" {
  description = "Username used to authenticate against dns.he.net"
  default     = ""
}

variable "dhn_password" {
  description = "Password used to authenticate against dns.he.net"
  default     = ""
  sensitive   = true
}

variable "dhn_2fa" {
  description = "2FA seed used to authenticate against dns.he.net"
  default     = ""
  sensitive   = true
}

variable "dhn_store_type" {
  description = "Cookie authentication store type"
  default     = "dummy"
}