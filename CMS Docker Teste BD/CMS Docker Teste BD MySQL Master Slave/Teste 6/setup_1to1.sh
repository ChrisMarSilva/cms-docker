#!/bin/bash
# MySQL主从同步详解与配置 https://zhuanlan.zhihu.com/p/335142300

echo -e "\n--- Setp1: on master server, create sync user ---"
mysql -h 127.0.0.1 -P 4406 -u root -p'123' < setup_master.sql > ../tmp/mysql-`date +%F_%H-%M-%S`.out

echo -e "\n--- Step2: backup master initial db ---"
mysqldump -h 127.0.0.1 -P 4406 -u root -p'123' master_db --single-transaction --master-data=1 --flush-logs > ../tmp/mysql_init.sql

echo -e "\n--- Setp3: restore master db to slave db ---"
mysql -h 127.0.0.1 -P 5506 -u root -p'456' < ../tmp/mysql_init.sql

echo -e "\n--- Setp4: on slave server, set master db ---"
mysql -h 127.0.0.1 -P 5506 -u root -p'456' < setup_slave.sql > ../tmp/mysql-`date +%F_%H-%M-%S`.out

echo -e "\n--- Setp5: 查看同步状态，Slave_IO_Running 和 Slave_SQL_Running 都要是 Yes ---"
mysql -h 127.0.0.1 -P 5506 -u root -p'456' -e "SHOW slave status\G;" | grep Running
echo -e "\nAll done."
