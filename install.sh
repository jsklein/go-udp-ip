wget https://github.com/kingtuna/go-udp-ip/releases/download/v1.1/go_ip_linux_port_1024
mv -f go_ip_linux_port_1024 /usr/local/bin/
chmod +x /usr/local/bin/go_ip_linux_port_1024

curl https://raw.githubusercontent.com/kingtuna/go-udp-ip/master/go_ip.upstart > /etc/init.d/go_ip
chmod +x /etc/init.d/go_ip

update-rc.d go_ip defaults
