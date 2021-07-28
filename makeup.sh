# Install nginx
sudo amazon-linux-extras install nginx1
# [TODO] Copy nginx.conf
sudo systemctl enable nginx
sudo systemctl start nginx

# Install certbot
sudo wget -r --no-parent -A 'epel-release-*.rpm' http://dl.fedoraproject.org/pub/epel/7/x86_64/Packages/e/
sudo rpm -Uvh dl.fedoraproject.org/pub/epel/7/x86_64/Packages/e/epel-release-*.rpm
sudo yum-config-manager --enable epel*
sudo yum install certbot-nginx
sudo certbot --nginx
# 1. Enter email address
# 2. Agree with term of service (a)
# 3. Share email address or not (N)
# 4. Choose domain name (server_name in nginx.conf)
# Auto update by using crontab
# 39      1,13    *       *       *       root    certbot renew --no-self-upgrade && systemctl restart nginx

# Install docker, docker-compose
sudo yum istall -y docker
sudo service docker start
sudo usermod -a -G docker ec2-usern

sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
