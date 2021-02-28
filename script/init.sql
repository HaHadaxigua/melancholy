use
melancholy;

drop table users if exists
create table users
(
    id         int auto_increment
        primary key,
    username   varchar(255) not null,
    password   varchar(255) not null,
    mobile     varchar(255) null,
    email      varchar(255) not null,
    status     int default 0 null,
    salt       varchar(255) not null,
    created_at timestamp null,
    updated_at timestamp null,
    deleted_at timestamp null,
    constraint users_email_uindex
        unique (email)
);

drop table user_role if exists
create table user_role
(
    id      int auto_increment
        primary key,
    user_id int not null,
    role_id int not null
);

drop table roles if exists
create table roles
(
    id         int auto_increment
        primary key,
    name       varchar(255) not null,
    created_at timestamp null,
    updated_at timestamp null,
    deleted_at timestamp null,
    constraint role_name_uindex
        unique (name)
);

drop table role_permission if exists
create table role_permission
(
    id            int auto_increment
        primary key,
    role_id       int not null,
    permission_id int not null
);


drop table permissions if exists
create table permissions
(
    id         int auto_increment
        primary key,
    name       varchar(255) not null,
    created_at timestamp null,
    updated_at timestamp null,
    deleted_at timestamp null,
    constraint permissions_name_uindex
        unique (name)
);