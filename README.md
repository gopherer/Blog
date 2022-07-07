# Blog 

```
git clone https://github.com/gopherer/Blog.git
```
注意事项：检查Config/Blog.json配置信息与本机配置信息是否相符

## 项目结构：

- **Controller**:接收浏览器(前端)的请求并调用*模板引擎*返回处理的HTML页面或处理的结果，目前内置两个中间件1、访问权限的限制，2、重新向恶意多次的请求
- **Service**:承接Controller传递的参数进行一定的处理并返回处理的结果
- **DataAccess**:承接Service传递的参数并结合*Xorm框架*进行对数据库的增删改查，并返回处理结构
- **Model**:用于结构体的定义、Blog、User、Config
- **Config**:内置Blog.json配置文件
- **Html**:用于保存Html页面、保存有Blog、User等一系列页面
- **Tools**:定义了连接**Mysql**,**Redis**数据库、初始化**Session**、解析配置文件等一系列函数
- **Upload**:用于保存用户头像和发布博客时上传的图片


![image ](https://github.com/gopherer/Blog/blob/main/MDPhoto/code.png)

![image ](https://github.com/gopherer/Blog/blob/main/MDPhoto/home.png)

## 注意

**问题**：使用Navicat设计表时，在字符集为utf8 排序规则为utf8_general_ci时添加注释认为乱码 

**解决**：win+i 打开设置-时间和语言-语言-管理语言设置-更改系统区域设置-打勾✔ Beta版：使用Unicode UTF-8提供全球语言支持 。重启即可。

**问题**：git 报错：error: failed to push some refs to 'github.com:gopherer/Blog.git' 

**解决**：原因是直接在GitHub上修改代码或文件导致本地仓库和远程仓库数据不一致。方法：**git push origin master -f**

## 还有许多不足，不完美的地方目前仍在优化中