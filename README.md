<h1 align="center">IaC-teste</h1>

Lab to deploy a AWS EC2 with Terraform, Ansible, Apache2, Docker and container with Golang

Table of contents
=================

<!--ts-->
   
   * [Prerequisites](#prerequisites)
   * [Installation](#installation)
   * [Usage](#usage)
      * [Terraform](#terraform)
      * [Ansible](#ansible)
   * [Accessing the application](#accessing-the-application)
   * [Update the application](#update-the-application)


<!--te-->

## Prerequisites

#### Create a AWS user with Programmatic access.
Link: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html#id_users_create_console

#### Create the KeyPars in aws
Link: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html

#### Install Terraform
Link: https://learn.hashicorp.com/tutorials/terraform/install-cli

#### Install Ansible

Link: https://docs.ansible.com/ansible/latest/installation_guide/index.html

#### Install Ansible Plugins

community.general.apache2_module – Enables/disables a module of the Apache2 webserver.
```
ansible-galaxy collection install community.general
```
ansible.posix.synchronize – A wrapper around rsync to make common tasks in your playbooks quick and easy
```
ansible-galaxy collection install ansible.posix
```


## Installation

Clone this repo

```bash
git clone https://github.com/marcosvlt/IaC-teste.git
```

# Usage

# Terraform

cd into terraform folder
```bash
cd terraform
```

Change the key name in main.tf for the same name that created in aws keypars

```bash
vim main.tf
```

```bash
key_name               = "terraform_teste"
```

### Build Infrastructure:

Initialize the directory
```bash
terraform init
```

Format and validate the configuration
```bash
terraform fmt
```

Create infrastructure
```bash
terraform apply
```

Access AWS and copy the EC2hostnames


# Ansible

Create the ansible host files with target machines
```bash
cd ansible
```
```bash
vim hosts
```
Copy and paste the code below with target machine
```bash
[marcos-teste]
foo.example.com
```
Run ansible with private key configured in AWS KeyPars
```bash
ansible-playbook -i hosts mytask.yml -u ubuntu --private-key PRIVATE_KEY_LOCATION
```

After that ansible will run the playbook and install Apache2, docker and create the container listening on port 8080


# Accessing the application

In your browser access with EC2 ip. Apache will proxy to the container in the port 8080 and display the application

http://EC2_IP

# Update the application

The container is mounted in the /home/marcos/ with persistent files, just update the files and will reflect in the container




