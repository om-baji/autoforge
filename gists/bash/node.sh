curl -fsSL https://deb.nodesource.com/setup_20.x | sudo -E bash -
sudo apt-get install -y nodejs || sudo dnf install -y nodejs || sudo pacman -Sy --noconfirm nodejs npm

node -v
npm -v
