# ldappasswd

ldappasswd是一个简单的ldap自服务密码管理工具，当前仅支持389端口，不支持636端口，使用docker运行

```shell
$ docker run -itd-e LDAP_SERVER=$your_ldap_server -e LDAP_PORT=$your_ldap_port -p 8000:8389 cherryleo/ldappasswd
```

 

运行截图

![](https://raw.githubusercontent.com/cherryleo/ldappasswd/master/screenshot/1.png)

![](https://raw.githubusercontent.com/cherryleo/ldappasswd/master/screenshot/4.png)