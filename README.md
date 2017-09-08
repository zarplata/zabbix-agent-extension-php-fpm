# zabbix-agent-extension-php-fpm

zabbix-agent-extension-php-fpm - this extension for monitoring PHP-FPM.

### Installation

#### Manual build

```sh
# Building
git clone https://github.com/zarplata/zabbix-agent-extension-php-fpm.git
cd zabbix-agent-extension-php-fpm
make

#Installing
make install

# By default, binary installs into /usr/bin/ and zabbix config in /etc/zabbix/zabbix_agentd.conf.d/ but,
# you may manually copy binary to your executable path and zabbix config to specific include directory
```

### Dependencies

zabbix-agent-extension-php-fpm requires [zabbix-agent](http://www.zabbix.com/download) v2.4+ to run.
