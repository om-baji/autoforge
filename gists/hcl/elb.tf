resource "aws_lb" "__ALB_NAME__" {
  name               = "__ALB_ACTUAL_NAME__"
  internal           = __IS_INTERNAL__
  load_balancer_type = "application"
  security_groups    = ["__SECURITY_GROUP_ID__"]
  subnets            = ["__SUBNET_ID_1__", "__SUBNET_ID_2__"]

  tags = {
    Environment = "__ENVIRONMENT__"
  }
}
