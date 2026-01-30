# copy the initial script to the openstack server
file=scripts/init.sh
# file1=scripts/post-init.sh
# downloads=downloads/*
# copy files from local
# scp -P 2003 -i /Users/mariia.rubina13/.ssh/cluster $file ubuntu@188.34.101.189:~/terraform
# scp -P 2003 -i /Users/mariia.rubina13/.ssh/cluster $file ubuntu@188.34.101.189:~/
scp $file stack@188.34.101.189:~/terraform