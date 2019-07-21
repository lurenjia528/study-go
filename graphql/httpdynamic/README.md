# 动态加载数据文件

![run](image/run.png)

```json
[
  {
    "id": "1",
    "name": "Dan"
  },
  {
    "id": "2",
    "name": "Lee"
  },
  {
    "id": "3",
    "name": "Nick"
  }
]
```

![altair](image/altair.png)

![no-signal](image/no-signal.png)

修改data.json

```json
[
  {
    "id": "1",
    "name": "Dan",
    "surname": "Jones"
  },
  {
    "id": "2",
    "name": "Lee"
  },
  {
    "id": "3",
    "name": "Nick"
  }
]
```

发送信号
`kill -SIGUSR1 27454 `

![signal](image/signal.png)

重新查询

![sig-reload](image/sig-reload.png)