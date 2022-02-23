clear

echo "\nWelcome to XENA botnet kit!\n"

echo "Checking Python3 installation."
if ! python3 --version; then
  echo "Python3 is required, proceeding with installation..."
  sudo apt install python3
fi

echo "Checking Docker installation."
if ! docker -v; then
  echo "Docker is required, proceeding with installation..."
  sudo apt install docker-ce docker-ce-cli containerd.io
fi

# Generate a key pair.
echo ""
echo "Do not set the password for the private key, it is not yet supported, just hit enter."
ssh-keygen -t rsa -b 4096 -m PEM -f xena.private.key
openssl rsa -in xena.private.key -pubout -outform PEM -out xena.public.key

python3 setup-helper.py