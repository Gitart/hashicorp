### Запуск сервера в режиме разработки   

```bash
consul agent -dev -node maschine
```

### Запись в базу Consul

```bash
curl -X PUT -d 'value1' localhost:8500/v1/kv/group1/key1
```

## Просмотр подключений
consul members -detailed

## Просомтр узлов
curl localhost:8500/v1/catalog/nodes

```json
[
{
ID: "d5d9ed5f-c557-ebf8-7bcb-319acaa04f46",
Node: "airpc-ubuntu",
Address: "127.0.0.1",
Datacenter: "dc1",
TaggedAddresses: {
lan: "127.0.0.1",
wan: "127.0.0.1"
},
Meta: {
consul-network-segment: ""
},
CreateIndex: 9,
ModifyIndex: 10
}
]
```
 ## Просомтор информации о серврере
 dig @127.0.0.1 -p 8600 machine.node.consul
   
 ## Конфигурация
 создать дирректорию 
 makedir 
 
 
 ## Конфигурация
 $ mkdir ./consul.d
 
 ```
 $ echo '{"service":
  {"name": "web",
   "tags": ["rails"],
   "port": 80
  }
}' > ./consul.d/web.json
```

## Запуск
```sh
consul agent -dev -enable-script-checks -config-dir=./consul.d
```
