# Vagrantfile
Vagrant.configure("2") do |config|
  # Define the virtual machine settings
  config.vm.box = "debian/buster64"
  config.vm.network "private_network", type: "dhcp"
  config.vm.provider "virtualbox" do |vb|
    vb.memory = "2048" # Adjust the memory as needed
    vb.cpus = 2       # Adjust the CPU count as needed
  end

  # Script to provision the virtual machine
  config.vm.provision "shell", inline: <<-SHELL
    # Update the package repository and install dependencies
    sudo apt update
    sudo apt install -y apt-transport-https ca-certificates curl gnupg2 software-properties-common

    # Install Docker
    curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -
    sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs) stable"
    sudo apt update
    sudo apt install -y docker-ce

    # Install Docker Compose
    sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    sudo chmod +x /usr/local/bin/docker-compose

    # Install Kubernetes
    curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
    echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
    sudo apt update
    sudo apt install -y kubeadm kubelet kubectl

    # Install GCC (GNU Compiler Collection)
    sudo apt install -y gcc

    # Initialize Kubernetes (you may need to adjust this step as needed)
    # sudo kubeadm init --pod-network-cidr=10.244.0.0/16 --apiserver-advertise-address=192.168.33.10

    # Add your user to the Docker group
    sudo usermod -aG docker ${USER}

    # Start Docker service
    sudo systemctl start docker
    sudo systemctl enable docker

    # Optionally, you can set up a Kubernetes cluster using kubeadm here

    echo "All installations completed."
  SHELL
end
