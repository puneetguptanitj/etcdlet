# Build Etcdlet

From project root directory

```
pushd service
make image
popd
```
Make image builds the service binary and container image

# Start 3 node centos cluster
```
cp service/etcdlet.tar . 
vagrant up
```
# Bootstrap etcd cluster

Create a file /tmp/inventory/clusterspec with the following content to bootstrap etcd cluster

```
clusterspec:
  members:
  - name: etcd0
    address: 192.168.60.10 
  - name: etcd1
    address: 192.168.60.11 
  - name: etcd2
    address: 192.168.60.12 
```
