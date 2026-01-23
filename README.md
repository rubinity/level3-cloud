# OpenStack
```
# creating a user
sudo useradd -s /bin/bash -d /opt/stack -m stack
#make it executable
sudo chmod +x /opt/stack
# add it to sudoers
echo "stack ALL=(ALL) NOPASSWD: ALL" | sudo tee /etc/sudoers.d/stack
#switch to user
sudo -u stack -i
#cloning devstack repo
git clone https://opendev.org/openstack/devstack
cd devstack
```
### creating local.conf

[local.conf](local.conf)

```
#running the script to install openstack
./stack.sh
```
### Some other files I used
[port forwarding script](forward.sh)

### Diagram
![diagram](diagram_simple.png)