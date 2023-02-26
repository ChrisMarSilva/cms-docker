
CREATE DATABASE `master_db` CHARACTER SET 'utf8';
CREATE user 'rep'@'%' IDENTIFIED BY '1234567';
GRANT replication slave,replication client on *.* TO 'rep'@'%';

-- mysql -h 127.0.0.1 -P 3306 -u rep -p'1234567'
