if command -v apt >/dev/null 2>&1; then
    sudo apt update
    sudo apt install -y nginx
elif command -v dnf >/dev/null 2>&1; then
    sudo dnf install -y nginx
elif command -v yum >/dev/null 2>&1; then
    sudo yum install -y nginx
elif command -v pacman >/dev/null 2>&1; then
    sudo pacman -Sy --noconfirm nginx
else
    echo "Unsupported distro. Install Nginx manually."
    exit 1
fi

sudo systemctl enable --now nginx
nginx -v
systemctl status nginx
