# Consul

https://learn.hashicorp.com/consul/datacenter-deploy/reference-architecture


### Запуск сервера в режиме разработки   

```bash
consul agent -dev -node maschine
```

### Запись в базу Consul

```bash
curl -X PUT -d 'value1' localhost:8500/v1/kv/group1/key1
```
### Чтение из базы Consul

```bash
curl -X GET -d localhost:8500/v1/kv/key1
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


## Add Data
First, insert or "put" some values into the KV store with the consul kv put command. The first entry after the command is the key, and the second entry is the value.

```sh
$ consul kv put redis/config/minconns 1
Success! Data written to: redis/config/minconns

$ consul kv put redis/config/maxconns 25
Success! Data written to: redis/config/maxconns

$ consul kv put -flags=42 redis/config/users/admin abcd1234
Success! Data written to: redis/config/users/admin
```

