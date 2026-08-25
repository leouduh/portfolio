resource "aws_lightsail_key_pair" "main" {
  name       = "portfolio-key"
  public_key = file("~/.ssh/id_ed25519.pub")
}

resource "aws_lightsail_instance" "portfolio-leosama" {
  name              = "portfolio-leosama-lightsail-instance"
  availability_zone = "eu-west-1a"
  blueprint_id      = "ubuntu_24_04"
  bundle_id         = "nano_3_0"
  key_pair_name     = aws_lightsail_key_pair.main.name
  user_data = <<-EOF
  #!/bin/bash
  apt-get update && apt-get upgrade -y
  apt-get install vim -y
  EOF
}

resource "aws_lightsail_static_ip" "portfolio-leosama" {
  name = "portfolio-leosama-static-ip"
}

resource "aws_lightsail_static_ip_attachment" "portfolio-leosama" {
  static_ip_name = aws_lightsail_static_ip.portfolio-leosama.name
  instance_name = aws_lightsail_instance.portfolio-leosama.name
}

resource "aws_lightsail_instance_public_ports" "portfolio-leosama" {
  instance_name = aws_lightsail_instance.portfolio-leosama.name

  port_info {
    protocol  = "tcp"
    from_port = 22
    to_port   = 22
  }

  port_info {
    protocol  = "tcp"
    from_port = 80
    to_port   = 80
  }

  port_info {
    protocol  = "tcp"
    from_port = 443
    to_port   = 443
  }
}
