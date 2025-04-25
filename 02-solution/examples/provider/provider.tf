terraform {
  required_providers {
    playground = {
      source = "registry.opentofu.org/example/playground"
    }
  }
}

provider "playground" {
  folder = null
}
provider "playground" {
  alias = "playagain"

  folder = var.playagain_path
}

variable "playagain_path" {
  type        = string
  description = "Local path for our playground!"
  default     = "."
  validation {
    condition     = provider::playground::valid_path(var.playagain_path)
    error_message = "The playground path is not valid."
  }
}
