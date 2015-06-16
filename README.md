# beeblog
无闻老师的web编程基础，完成基本的blog功能

对models的方法进行了分类，不同数据表对应不同文件，各自保存其数据操作的方法

把数据库连接相关的操作，集成到了models中base.go中

国际化的代码，放到了controllers的base.go中

路由用单独的文件routers.go分配