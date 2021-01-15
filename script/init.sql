use
melancholy;

drop table casbin_rule if exists;
create table casbin_rule
(
    p_type varchar(100) default '' not null,
    v0     varchar(100) default '' not null,
    v1     varchar(100) default '' not null,
    v2     varchar(100) default '' not null,
    v3     varchar(100) default '' not null,
    v4     varchar(100) default '' not null,
    v5     varchar(100) default '' not null
) charset = utf8;

create
index IDX_casbin_rule_p_type
    on casbin_rule (p_type);

create
index IDX_casbin_rule_v0
    on casbin_rule (v0);

create
index IDX_casbin_rule_v1
    on casbin_rule (v1);

create
index IDX_casbin_rule_v2
    on casbin_rule (v2);

create
index IDX_casbin_rule_v3
    on casbin_rule (v3);

create
index IDX_casbin_rule_v4
    on casbin_rule (v4);

create
index IDX_casbin_rule_v5
    on casbin_rule (v5);


create table exit_logs
(
    id      int auto_increment
        primary key,
    user_id int          not null,
    token   varchar(255) not null,
    date    timestamp null
);


create table folders
(
    id         int auto_increment
        primary key,
    parent     int null,
    name       varchar(255) null,
    owner      int null,
    size       int default 0 null,
    status     int default 0 null,
    created_at timestamp null,
    updated_at timestamp null,
    deleted_at timestamp null
);


create table m_files
(
    id         int auto_increment
        primary key,
    parent     int not null,
    name       varchar(255) null,
    author     int null,
    md5        int null,
    size       int default 0 null,
    m_type     int default 0 null,
    `desc`     varchar(255) null,
    status     int default 0 null,
    updated_at timestamp null,
    created_at timestamp null,
    deleted_at timestamp null
);


create table roles
(
    id         int auto_increment
        primary key,
    name       varchar(255) null,
    status     int default 0 null,
    created_at timestamp null,
    updated_at timestamp null,
    deleted_at timestamp null,
    constraint role_name_uindex
        unique (name)
);

create table users
(
    id         int auto_increment
        primary key,
    username   varchar(255) null,
    password   varchar(255) null,
    phone      int null,
    email      varchar(255) null,
    state      int default 0 null,
    salt       varchar(255) null,
    created_at timestamp null,
    updated_at timestamp null,
    deleted_at timestamp null
);

