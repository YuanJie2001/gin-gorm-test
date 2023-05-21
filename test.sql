drop database if exists test_gorm_db;
create database test_gorm_db default character set utf8mb4 collate utf8mb4_0900_ai_ci;

use test_gorm_db;

drop table if exists `user`;
create table `user`(
                       `id` bigint(20) unsigned auto_increment comment '主键',
                       `name` varchar(255) not null comment '姓名',
                       `age` int(11) not null comment '年龄',
                       primary key (`id`)
) engine=InnoDB default charset=utf8mb4 comment='用户表';

insert into `user`(`name`, `age`) values('张三', 18);
insert into `user`(`name`, `age`) values('李四', 20);
insert into `user`(`name`, `age`) values('王五', 22);
insert into `user`(`name`, `age`) values('赵六', 24);
insert into `user`(`name`, `age`) values('田七', 26);
insert into `user`(`name`, `age`) values('周八', 28);
insert into `user`(`name`, `age`) values('吴九', 30);
insert into `user`(`name`, `age`) values('郑十', 32);
