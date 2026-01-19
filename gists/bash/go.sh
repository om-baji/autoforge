GO_VERSION="1.21.2"
ARCH=$(uname -m)
OS="linux"
GO_TAR="go${GO_VERSION}.${OS}-${ARCH}.tar.gz"

wget -q https://go.dev/dl/${GO_TAR} -O /tmp/${GO_TAR}
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf /tmp/${GO_TAR}
rm /tmp/${GO_TAR}

export PATH=$PATH:/usr/local/go/bin
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc

go version
