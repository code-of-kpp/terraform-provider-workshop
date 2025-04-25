terraform {
  required_providers {
    playground = {
      source = "registry.opentofu.org/example/playground"
    }
  }
}

provider "playground" {}

output "what-we-are-here-for" {
  value = provider::playground::meaning_of_life()
}
