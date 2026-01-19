resource "aws_subnet" "__SUBNET_NAME__" {
  vpc_id                  = "__VPC_ID__"
  cidr_block              = "__SUBNET_CIDR__"
  availability_zone       = "__AVAILABILITY_ZONE__"
  map_public_ip_on_launch = __MAP_PUBLIC_IP__

  tags = {
    Name = "__SUBNET_TAG_NAME__"
  }
}
