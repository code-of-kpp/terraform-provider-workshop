terraform {
  required_providers {
    // Provider installation:
    playground = {
      source = "registry.opentofu.org/example/playground"
    }
  }
}

provider "playground" {
  folder = var.playground_path
}

variable "playground_path" {
  type        = string
  description = "Local path for our playground!"
  default     = "."
  validation {
    condition     = provider::playground::valid_path(var.playground_path)
    error_message = "The playground path is not valid."
  }
}
