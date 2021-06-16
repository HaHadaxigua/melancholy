create table doc_files
(
    id      varchar(255) not null
        primary key,
    content text         null comment '文件内容'
);

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
    endpoint    varchar(255) null,
    size        int          null comment '文件大小',
    mode        int          null comment '文件模式',
    ftype       int          null comment '文件类型',
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

create table friends
(
    id         int auto_increment
        primary key,
    `from`     varchar(255) null comment '来自A用户的好友申请',
    `to`       varchar(255) null comment '来自B用户的好友同意',
    status     int          null,
    created_at timestamp    null comment '好友申请日期',
    updated_at timestamp    null comment '好友申请同意',
    deleted_at timestamp    null comment '好友申请拒绝'
);

create table music_files
(
    id        varchar(255) not null
        primary key,
    name      varchar(255) null comment '歌曲名',
    cover_url varchar(255) null comment '封面',
    duration  int          null comment '时长',
    singer    varchar(255) null comment '歌手',
    album     varchar(255) null comment '专辑',
    years     int          null comment '年份',
    species   varchar(255) null comment '分类',
    finished  tinyint(1)   null comment '是否上传完成',
    music_id  varchar(255) null comment '视频点播地址',
    region    varchar(255) null
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
    id                  int auto_increment
        primary key,
    username            varchar(255)  not null,
    password            varchar(255)  not null,
    mobile              varchar(255)  null,
    email               varchar(255)  not null,
    salt                varchar(255)  not null,
    avatar              varchar(255)  null comment '头像',
    status              int default 0 null,
    oss_end_point       varchar(255)  null,
    cloud_access_key    varchar(255)  null,
    oss_access_secret   varchar(255)  null,
    oss_access_key      varchar(255)  null,
    cloud_access_secret varchar(255)  null,
    created_at          timestamp     null,
    updated_at          timestamp     null,
    deleted_at          timestamp     null,
    constraint users_email_uindex
        unique (email)
);

create table video_files
(
    id                 varchar(255) not null
        primary key,
    title              varchar(255) null comment '视频标题',
    description        varchar(255) null comment '简介',
    cover_url          varchar(255) null comment '封面地址',
    area               varchar(255) null comment '地区',
    years              int          null comment '年份',
    production_company varchar(255) null comment '制作公司',
    species            varchar(255) null comment '视频类型',
    duration           int          null comment '时长(秒)',
    finished           tinyint(1)   null comment '是否上传完成',
    video_id           varchar(255) null comment '视频点播文件地址',
    region             varchar(255) null
)
    comment '视频文件信息';

