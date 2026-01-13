if command -v apt >/dev/null 2>&1; then
    sudo apt update
    sudo apt install -y python3 python3-venv python3-pip
elif command -v dnf >/dev/null 2>&1; then
    sudo dnf install -y python3 python3-venv python3-pip
elif command -v yum >/dev/null 2>&1; then
    sudo yum install -y python3 python3-venv python3-pip
elif command -v pacman >/dev/null 2>&1; then
    sudo pacman -Sy --noconfirm python python-pip
else
    echo "Unsupported distro. Install Python manually."
    exit 1
fi

python3 --version
pip3 --version
