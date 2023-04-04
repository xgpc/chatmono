create table checkup_info
(
    id                int unsigned auto_increment
        primary key,
    created_user_id   int unsigned  not null comment '体检人',
    created_at        int unsigned  not null comment '下单时间',
    sign              int default 0 not null comment '签到状态:[0未签到, 1已签到]',
    sign_user_id      int unsigned  null comment '签到确认人',
    appointment_at    int unsigned  not null comment '预约时间',
    user_name         varchar(200)  null comment '体检人姓名',
    user_mobile       varchar(11)   not null comment '用户手机号',
    user_id_card_data blob          not null comment '身份证数据',

    constraint id
        unique (id)
)
    comment '体检信息';

create table checkup_package
(
    checkup_package_id   int unsigned auto_increment
        primary key,
    checkup_package_name varchar(200)                not null comment '套餐名称',
    price                decimal(10, 2) default 0.00 not null comment '套餐价格',
    zlxm_list            json                        null,
    detail               text                        null comment '套餐描述',
    created_at           int unsigned                not null,
    created_user_id      int unsigned                not null,
    constraint checkup_package_id
        unique (checkup_package_id)
)
    comment '体检套餐';

create table dept
(
    dept_id      int unsigned auto_increment
        primary key,
    dept_name    varchar(200)  not null,
    dept_address varchar(2000) null comment '科室地址',
    dept_user_id int unsigned  null comment '科室负责人',
    constraint dept_id
        unique (dept_id)
)
    comment '科室';

create table zlxm
(
    zlxm_id    int unsigned auto_increment
        primary key,
    sn         varchar(200)                not null,
    zlxm_name  varchar(200)                not null,
    unit       varchar(20)                 null comment '计价单位',
    price      decimal(10, 2) default 0.00 not null comment '价格',
    child_list json                        null comment '捆绑包子集',
    detail     text                        null,
    constraint id
        unique (zlxm_id),
    constraint sn
        unique (sn)
)
    comment '诊疗项目';

