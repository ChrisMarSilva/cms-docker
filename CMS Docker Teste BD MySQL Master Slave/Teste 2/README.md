# MySQL主从备份

> 1. master文件夹是主数据库的docker构建
> 2. slave文件夹是从数据库的docker构建
> 3. my.cnf是数据库的配置文件
> 4. docker-compose.yml构建docker
> 5. sources.list是Linux的国内源文件可以用于替换
> 6. dockerfiletmp是可以安装apt的docker文件

- 启动 docker-compose -f docker-compose.yml up -d
- 停止 docker-compose -f docker-compose.yml down

从数据库
    show master status;
    show  slave status;

    CHANGE MASTER TO
  MASTER_HOST='mysql',
  MASTER_USER='root',
  MASTER_PASSWORD='prisma',
  MASTER_LOG_FILE='replicas-mysql-bin.000003',
  MASTER_LOG_POS=154;
    stop slave;
    start slave;

    reset slave all;  #清除主从复制

主数据库
    show master status;

    show binlog events in 'replicas-mysql-bin.000001';

    #暂时不讲操作追加到Bin日志中   基于会话层面的  只针对当前连接的会话，其他会话不影响，但是如果设置的是（set global sql_log_bin=0)，这样会导致所有在Master数据库上执行的语句都不记录到binlog
    SET sql_log_bin=1;  #开启
    SET sql_log_bin=0;  #关闭

    show variables like "sql_log_bin";

<!-- set global general_log_file='/var/log/mysql/general_log.log' -->
<!-- SET GLOBAL general_log = 'ON'; -->
    SET FOREIGN_KEY_CHECKS=0;
    -- ----------------------------
    -- Table structure for course
    -- ----------------------------
    DROP TABLE IF EXISTS `course`;
    CREATE TABLE `course` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      `name` varchar(20) NOT NULL,
      `lesson_period` double(5,0) DEFAULT NULL,
      `score` double(10,0) DEFAULT NULL,
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
