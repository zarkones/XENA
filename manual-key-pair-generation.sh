ssh-keygen -t rsa -b 4096 -m PEM -f xena.private.key
openssl rsa -in xena.private.key -pubout -outform PEM -out xena.public.key