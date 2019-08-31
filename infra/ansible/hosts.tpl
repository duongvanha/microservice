[masters]
master ansible_host=${master_ip} ansible_user=haduong

[workers]
worker1 ansible_host=${worker1_ip} ansible_user=haduong

[all:vars]
ansible_python_interpreter=/usr/bin/python3
