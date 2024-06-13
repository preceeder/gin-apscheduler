# gin-apscheduler


```shell
// swag
http://0.0.0.0:8080/swagger/index.html
![img.png](https://github.com/preceeder/gin-apscheduler/blob/main/static/img.png)
```

```shell
# swagger 有引用外部类型的时候需要加参数   --parseDependency --parseInternal
  swag init  --parseDependency --parseInternal
```

```
// add_job
{
  "args": {
    "additionalProp1": "string"
  },
  "func_name": "task.test:Test",
  "id": "test002",
  "interval_trigger": {
    "Jitter": 0,
    "end_time": "",
    "interval": 7000,
    "start_time": "",
    "utc_time_zone": "UTC+8"
  },
  "max_instance": 0,
  "name": "test002",
  "replace": true,
  "store_name": "default",
  "timeout": 0
}



```