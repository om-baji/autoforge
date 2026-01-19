resource "aws_vpc" "__VPC_NAME__" {
  cidr_block           = "__VPC_CIDR__"
  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = {
    Name = "__VPC_TAG_NAME__"
  }
}
