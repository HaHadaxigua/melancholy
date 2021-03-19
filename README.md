# melancholy

melancholy /ˈmel.əŋ.kɒl.i/ 忧愁的

这是一首很好听的纯音乐，同时也以此来纪念我那不成熟的岁月。

这是一个在线网盘系统， 支持web / android / pc

- [ ] **后端**
    - [ ] 文件系统
        - [ ] 文件管理
            - [ ] 上传文件
            - [ ] 下载文件(逻辑删除)
            - [ ] 文件命名
            - [ ] 创建文件
            - [ ] 文件删除
            - [ ] 文件属性管理
                - [ ] 弹幕
                - [ ] 评论
                - [ ] 点赞
                - [ ] 加密
    - [ ] 用户系统
        - [x] 注册
        - [x] 登录
        - [x] 登出
    - [ ] 后台管理

## FileModule

In this module, we use id to represent a path, we can find a file's parent. Then we can return a tree struct. So when
cur-path is requested, we need a cur-pid.At first, we set everyone's initial root path is zero, and user can see a full
struct.

### About upload File

由于我们将文件上传到阿里云的对象存储中，因此我们并不需要亲自去处理文件。

核心思路： 将请求转发到阿里云的公共api

解决办法：

1. 直接在请求中进行转发
2. 新开一个消费队列去消费，这种复杂度太高了， 不太合适

# Git Commit Rule

分类 | 说明
---- | ----
normal | 普通提交，无意义，改动的地方很小
feat | 新功能
fix | 修复了错误

## gitignore

*.log # 忽略.log后缀的文件 bin/ # 过滤bin文件夹的内容 /idea/ # 过滤整个文件夹

### usage

1. git rm -r --cached .
2. git add .
3. git commit -m "fix: ..."
4. git push origin <branch name>

# Docker 构建镜像

创建一个Dockerfile文件, 在此文件的目录下输入：

```shell
docker build -t TagName:version .
```

# CI/CD 持续集成

使用GithubActions来持续发布与部署