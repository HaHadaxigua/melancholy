create table files
(
    id          varchar(255) not null
        primary key,
    owner_id    int          null comment '用户id',
    parent_id   varchar(255) null,
    name        varchar(255) null,
    suffix      varchar(255) null comment '文件后缀',
    hash        text         null,
    address     varchar(255) null comment 'oss地址',
    bucket_name varchar(255) null comment '存储桶名',
    object_name varchar(255) null comment '存储对象名字',
    size        int          null comment '文件大小',
    mode        int          null comment '文件模式',
    created_at  timestamp    null,
    updated_at  timestamp    null,
    deleted_at  timestamp    null
);

create table folder_sub
(
    id        int auto_increment
        primary key,
    folder_id varchar(255) not null,
    sub_id    varchar(255) not null
);

create table folders
(
    id         varchar(255) not null,
    parent_id  varchar(255) null,
    owner_id   int          not null,
    name       varchar(255) not null,
    created_at timestamp    null,
    updated_at timestamp    null,
    deleted_at timestamp    null,
    primary key (id, owner_id)
);

create table permissions
(
    id         int auto_increment
        primary key,
    name       varchar(255) not null,
    created_at timestamp    null,
    updated_at timestamp    null,
    deleted_at timestamp    null,
    constraint permissions_name_uindex
        unique (name)
);

create table role_permission
(
    id            int auto_increment
        primary key,
    role_id       int not null,
    permission_id int not null
);

create table roles
(
    id         int auto_increment
        primary key,
    name       varchar(255) not null,
    created_at timestamp    null,
    updated_at timestamp    null,
    deleted_at timestamp    null,
    constraint role_name_uindex
        unique (name)
);

create table user_role
(
    id      int auto_increment
        primary key,
    user_id int not null,
    role_id int not null
);

create table users
(
    id         int auto_increment
        primary key,
    username   varchar(255)  not null,
    password   varchar(255)  not null,
    mobile     varchar(255)  null,
    email      varchar(255)  not null,
    status     int default 0 null,
    salt       varchar(255)  not null,
    created_at timestamp     null,
    updated_at timestamp     null,
    deleted_at timestamp     null,
    constraint users_email_uindex
        unique (email)
);

