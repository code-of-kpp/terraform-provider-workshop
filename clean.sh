#!/usr/bin/env sh

find \
  \( -path '*-solution/terraform-provider-playground' \
  -or -path '*-solution/examples/*/terraform.tfstate*' \) \
  -print -delete
