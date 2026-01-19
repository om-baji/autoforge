resource "aws_security_group" "__SG_NAME__" {
  name        = "__SG_ACTUAL_NAME__"
  description = "__SG_DESCRIPTION__"
  vpc_id      = "__VPC_ID__"

  ingress {
    description = "__INGRESS_DESC__"
    from_port   = __FROM_PORT__
    to_port     = __TO_PORT__
    protocol    = "__PROTOCOL__"
    cidr_blocks = ["__CIDR_BLOCK__"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "__SG_TAG_NAME__"
  }
}
