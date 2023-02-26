
CREATE DATABASE IF NOT EXISTS `master_db` CHARACTER SET 'utf8';
GRANT SElECT ON *.* TO 'viewer'@'%' IDENTIFIED BY "1234567";
STOP SLAVE;
-- RESET SLAVE;
change master to master_host='mysql_master',master_user='rep',master_password='1234567',master_auto_position=1;
START SLAVE;
SHOW slave status\G;

-- mysql -h 127.0.0.1 -P 3307 -u viewer -p'1234567'
-- SHOW VARIABLES LIKE '%_query%';
