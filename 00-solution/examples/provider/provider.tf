terraform {
  required_providers {
    // Provider installation:
    playground = {
      source = "registry.opentofu.org/example/playground"
    }
  }
}

// Provider configuration:
provider "playground" {}
