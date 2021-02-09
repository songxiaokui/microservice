
1. 创建数据库和数据表
    CREATE DATABASE user;
    create table if not exists user.user
        (
             id         bigint auto_increment
                 primary key,
             username   varchar(100)                        not null,
             password   varchar(100)                        not null,
             email      varchar(100)                        not null,
             created_at timestamp default CURRENT_TIMESTAMP not null
        );
2. 修改configs下yaml文件为自己的数据库配置

3. cd cmd/user
4. go build
5. ./user