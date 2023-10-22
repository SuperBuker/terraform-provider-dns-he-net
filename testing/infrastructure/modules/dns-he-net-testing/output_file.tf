# Outputs config file to some location

resource "local_sensitive_file" "config_file" {
  content  = jsonencode(local.test_config)
  filename = "${path.root}/${var.config_file}"
}
