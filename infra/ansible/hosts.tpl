[masters]
master ansible_host=${master_ip} ansible_user=duongha

[workers]
worker1 ansible_host=${worker1_ip} ansible_user=duongha

[all:vars]
ansible_python_interpreter=/usr/bin/python3
