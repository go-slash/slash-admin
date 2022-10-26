## 使用swagger

> 环境配置

[go-swagger](https://goswagger.io/install.html)

> 在项目根目录运行 slash-admin/

```shell
swagger generate spec --output=./core.yml --scan-models

swagger serve --no-open -F=swagger --port 36666 core.yaml
```

![pic](../../assets/swagger.png)

> 获取 Token 
> 
> 先登录系统，在任意请求中复制 authorization

![pic](../../assets/get_token.png)

> 粘贴到 swagger

![pic](../../assets/swagger_authority.png)


