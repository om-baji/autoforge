resource "aws_db_instance" "__RDS_NAME__" {
  identifier             = "__RDS_IDENTIFIER__"
  engine                 = "__DB_ENGINE__"
  instance_class         = "__DB_INSTANCE_CLASS__"
  allocated_storage      = __STORAGE_SIZE__
  username               = "__DB_USERNAME__"
  password               = "__DB_PASSWORD__"
  db_subnet_group_name   = "__DB_SUBNET_GROUP__"
  vpc_security_group_ids = ["__SECURITY_GROUP_ID__"]
  skip_final_snapshot    = true
}
