N = 3
Vagrant::configure("2") do |config|
    (0..N-1).each do |machine_id|
        config.vm.box = "centos/7"
        config.vm.define "etcd#{machine_id}" do |machine|
            machine.vm.hostname = "etcd#{machine_id}"
            machine.vm.network "private_network", ip: "192.168.60.#{10+machine_id}"
        end
        config.vm.provision "shell", inline: "yum install -y etcd"
        config.vm.synced_folder "." , "/vagrant"
    end
end
