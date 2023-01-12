# goblog

## 配置

代码本身不提供配置,需要定义配置文件供程序使用

应用配置名需为`.gbconf.yaml`, 应用会尝试查找下列路径下的配置文件

- 应用同级文件夹
- /etc/gb_blog

### 配置项

```yaml
debug:
  pprof: false
mysql:
  dataSourceName: "{user}:{password}@/{database}"
  maxOpenConns: 200
  maxIdleConns: 100
```

