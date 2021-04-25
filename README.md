# Blue/Green deployment with Ansible
# Description
In this solution, there are three virtual machines that hostnames can be seen in the inventory file.

To solve the problem of setting up an environment to deploy an application without downtime I've used Nginx reverse proxy which exposes healthy apps to port 3000.

I also set my blue and green environment mapped to 3001, 3002 port numbers to migrate from blue env to the green env and vice-versa easily by these two vars.

this playbook runs on 3 VMs and moves between blue and green env in new deployment without knowing the env of other VMs. 

(there is no problem if one of container start in green and the other start from the blue state in cases we add a new VM our infra deployment VMs)


please follow this steps by order:

**1.An ansible playbook for setting up and configuring the VMs**

#ansible-playbook presetup.yaml -i inventory/sample/hosts.yml   

**2.An ansible playbook for building and pushing docker images**

#ansible-playbook buildpush.yaml -i inventory/sample/hosts.yml 

please note because of my poor internet connection I couldn't push my images to the docker hub and I used the local docker registry instead. 

**3.An ansible playbook which deploys the new tag with blue/green technique.**

#ansible-playbook --extra-vars 'tag=3.0' bluegreen.yaml -i inventory/sample/hosts.yml

by tag variable we specify version of deploymen, if the newer version of deployment is healthy, application will be updated without any down time else 
the older healthy version will stay as our stable deployment.
also it's possible to rollback to the previous versions by specifying the tag.
