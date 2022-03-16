# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "generic/alpine315"

  # temporal port
  config.vm.network "forwarded_port",
    guest: 7233,
    host: 7233,
    host_ip: "127.0.0.1"
  # temporal-web port
  config.vm.network "forwarded_port",
    guest: 8088,
    host: 8088,
    host_ip: "127.0.0.1"

  config.vm.synced_folder ".", "/vagrant"

  config.vm.provider "virtualbox" do |vb|
    vb.cpus = 8
    vb.memory = "8192"
  end

  config.vm.provision "shell", inline: <<-SHELL
    sudo apk update
    sudo apk add docker docker-compose
    sudo rc-update add docker
    sudo service docker start

    cd /vagrant && docker-compose up --detach # --wait
  SHELL
end
