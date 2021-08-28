# IaC-teste
IaC-Teste Lab to deploy a AWS EC2 with Terraform, Ansible, Apache2,  Docker and Golang


# Ansible


### Create a host file in ansible folder with target machines

### community.general.apache2_module – Enables/disables a module of the Apache2 webserver.
This plugin is part of the community.general collection (version 3.5.0).

To install it use: ansible-galaxy collection install community.general.

To use it in a playbook, specify: community.general.apache2_module.

### ansible.posix.synchronize – A wrapper around rsync to make common tasks in your playbooks quick and easy

To install it use: ansible-galaxy collection install ansible.posix.

To use it in a playbook, specify: ansible.posix.synchronize.
