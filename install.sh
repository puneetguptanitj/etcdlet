# install docker
cat > /etc/yum.repos.d/docker.repo <<EOS
[dockerrepo]
name=Docker CE Stable
baseurl=https://download.docker.com/linux/centos/7/\$basearch/stable
enabled=1
gpgcheck=0

EOS
yum makecache fast

yum install -y "docker-ce-17.09.1.ce"
mkdir -p /etc/docker
echo '{"storage-driver": "devicemapper"}' > /etc/docker/daemon.json
systemctl enable docker
systemctl start docker
mkdir -p /tmp/inventory


