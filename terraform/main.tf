terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.27"
    }
  }

  required_version = ">= 0.14.9"
}


# Configurar o provedor AWS

provider "aws" {
  profile = "default"
  region  = "us-east-1"
}

resource "aws_instance" "marcos-teste" {
  ami                    = "ami-0747bdcabd34c712a"
  instance_type          = "t2.micro"
  key_name               = "terraform_teste"
  vpc_security_group_ids = ["${aws_security_group.marcos_teste.id}"]



  tags = {
    Name = "marcos-teste"
  }
}



