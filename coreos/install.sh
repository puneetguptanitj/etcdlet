# install docker
cat > /etc/systemd/system/etcdlet.service <<EOS
[Unit]
Description=Ectdlet
After=docker.service

[Service]
ExecStart=/usr/bin/docker run --net=host -v /data/etcd:/var/etcd/data  -v /tmp/inventory:/tmp/inventory -v /var/run/docker.sock:/var/run/docker.sock --name etcdlet etcdlet /service
ExecStop=/usr/bin/docker stop etcdlet
ExecStopPost=/usr/bin/docker rm etcdlet

[Install]
WantedBy=multi-user.target
EOS
mkdir -p /tmp/inventory
docker load  < /home/core/share/etcdlet.tar

systemctl dameon-reload
systemctl enable etcdlet.service
systemctl restart etcdlet.service
