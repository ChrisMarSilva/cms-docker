说明文档
=======

## Docker环境搭建  
对于 Ubuntu 18 以上系统，可执行 setup_on_ubuntu.sh 脚本搭建环境
* Docker v20.10.x
* docker-compose v1.28.x
* Mysql v5.6+ 使用 docker-compose 创建

## 配置主从数据库
* Step1: 创建和运行数据库，在 mysql 目录执行 docker-compose up，待容器跑起来后，确认 root 账号能登录
```bash
mysql -h 127.0.0.1 -P 3306 -u root -p'123'
```
* Step2: 创建数据库用户和配置主从同步，可执行脚本自动配置1主1从同步，脚本没有提醒错误的话继续下一步
```bash
cd mysql && ./setup_1to1.sh
```

* Step3: 确认从数据库的同步状态，在从数据库执行如下命令，
```bash
mysql -h 127.0.0.1 -P 3307 -u root -p'456' -e "SHOW slave status\G;" | grep Running
```
输出结果中 Slave_IO_Running 和 Slave_SQL_Running 都要是 Yes，示范结果如下
```bash
             Slave_IO_Running: Yes
            Slave_SQL_Running: Yes
      Slave_SQL_Running_State: Slave has read all relay log; waiting for more updates
```

* Step4: 确认数据会自动同步，在主数据库创建或修改表数据，确认马上有同步到从数据库

## 测试主从同步性能
* 可参考项目 [agilejzl/master-slave-db-demo](https://github.com/agilejzl/master-slave-db-demo)
