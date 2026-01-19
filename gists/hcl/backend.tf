terraform {
  backend "__PLACEHOLDER__" {
    bucket       = "__BUCKET_PLACEHOLDER__"
    key          = "__BUCKET_KEY__"
    region       = "__AWS_REGION__"
    profile      = "__AWS_PROFILE__"
    use_lockfile = true
    encrypt      = true
  }
}
