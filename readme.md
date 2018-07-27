## 前后端分离的简单尝试——Hotspot

后端：

- 爬虫：goquery
- Server：gin

前端：

- vue
- Vuetify

TODO:

- pwa

## example

![](http://7xi3ed.com1.z0.glb.clouddn.com/5b6bf2cd8a1326775f17381a424bfca6.png)

## 定时执行 spider

crontab -e

```
1 *  *   *   * cd /var/hotspots && ./hotspots -m spider >> /tmp/main.log 2>&1
```