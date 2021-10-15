terraform {
  required_providers {
    sec = {
      source  = "mootpt/sec"
      version = "1.0.0"
    }
  }
}

variable "company_name" {
  type        = string
  description = "name of company to lookup s-1 status"
}

locals {
  message = data.sec_s1_filing.company.status ? "has filed their s-1" : "has not filed their s-1"
}

data "sec_s1_filing" "company" {
  name = var.company_name
}

output "s1_status" {
  value = "${data.sec_s1_filing.company.name} ${local.message}"
}
