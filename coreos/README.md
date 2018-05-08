# Build Etcdlet

From project root directory

```
pushd service
make image
popd
```
Make image builds the service binary and container image

# Start 3 node centos cluster
## Get upstream vagrant repo
```
git clone https://github.com/coreos/coreos-vagrant.git
```
## Update Vagrant file
```
set number of instances to 3
uncomment NFS share 
```
## Copy image tar
```
cp service/etcdlet.tar . 
vagrant up
```
# Bootstrap a etcd cluster

Create a file /tmp/inventor/clusterspec with the following content to bootstrap etcd cluster

```
clusterspec:
  members:
  - name: etcd1
    address: 172.17.8.101
  - name: etcd2
    address: 172.17.8.102
  - name: etcd3
    address: 172.17.8.103
```

