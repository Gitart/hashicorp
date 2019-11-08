### Запуск сервера в режиме разработки   

```bash
consul agent -dev -node maschine
```


### Запись в базу Consul


```bash
curl -X PUT -d 'value1' localhost:8500/v1/kv/group1/key1
```


