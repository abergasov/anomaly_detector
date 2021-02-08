### Dep
```shell
go mod vendor
```

### Dev env
up db/tarantool
```shell
bash dev.sh
```

### Test
```shell
wrk -s bench.lua -t12 -c400 -d30s http://localhost:31115/gather && curl http://localhost:29115/state
hey -m POST -d '{"id":123,"label":"view","value":5}' -z 10s http://localhost:31115/gather && curl http://localhost:29115/state
```
![result](data_gather.jpg)