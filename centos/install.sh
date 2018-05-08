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

#install etcdlet.service
cat > /etc/systemd/system/etcdlet.service <<EOS
[Unit]
Description=Ectdlet
After=docker.service

[Service]
ExecStart=/usr/bin/docker run --net=host -v /tmp/inventory:/tmp/inventory -v /var/run/docker.sock:/var/run/docker.sock --name etcdlet etcdlet /service
ExecStop=/usr/bin/docker stop etcdlet
ExecStopPost=/usr/bin/docker rm etcdlet

[Install]
WantedBy=multi-user.target
EOS
mkdir -p /tmp/inventory
docker load  < /vagrant/etcdlet.tar

systemctl system-reload
systemctl enable etcdlet.service
systemctl restart etcdlet.service
