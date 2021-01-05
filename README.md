# melancholy
melancholy  /ˈmel.əŋ.kɒl.i/ 忧愁的

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
        - [ ] 添加好友
        - [ ] 删除好友
        - [ ] 私聊
        - [ ] 群聊 
    - [ ] 后台管理
    - [ ] 短路由
- [ ] **前端**    
    - [ ] web
        - [ ] 在线解析
            - [ ] 图片
            - [ ] 音频 
            - [ ] 视频
            - [ ] 弹幕
    - [ ] android 
    
    
## 2020.10.27.11:57

# Git Commit Rule
分类 | 说明
---- | ----
normal | 普通提交，无意义，说明没有说明改动
feat | 新功能
fix | 修复了错误

## gitignore
*.log   # 忽略.log后缀的文件
bin/    # 过滤bin文件夹的内容
/idea/  # 过滤整个文件夹

### usage
1. git rm -r --cached .
2. git add .
3. git commit -m "fix: ..."
4. git push origin <branch name>