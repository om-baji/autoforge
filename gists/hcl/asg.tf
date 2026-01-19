resource "aws_autoscaling_group" "__ASG_NAME__" {
  name                 = "__ASG_ACTUAL_NAME__"
  max_size             = __MAX_SIZE__
  min_size             = __MIN_SIZE__
  desired_capacity     = __DESIRED_CAPACITY__
  vpc_zone_identifier  = ["__SUBNET_ID__"]
  launch_configuration = "__LAUNCH_CONFIG_NAME__"

  tag {
    key                 = "Name"
    value               = "__EC2_TAG_NAME__"
    propagate_at_launch = true
  }
}
