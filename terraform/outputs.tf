output "instance_ip" {
  value = aws_lightsail_static_ip.portfolio-leosama.ip_address
}
output "ecr_pull_access_key_id" {
  value = aws_iam_access_key.ecr_pull.id
}

output "ecr_pull_secret_access_key" {
  value     = aws_iam_access_key.ecr_pull.secret
  sensitive = true
}
