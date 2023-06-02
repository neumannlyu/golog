# golog

本日志模块自定义程度高，搭配出自己喜欢的样式。

## 使用方法

- 新建log对象

  ```
  log := NewDefaultLog()
  log.Level = 7
  ```

- 自定义data样式

  ```
  var mytime StTime
  mytime.Fgcolor = color.FgRed
  mytime.Format = "[[2006-01-02 15:04:05]]"
  log.UpdateElement(mytime)
  ```

- 自定义标签

  ```
  var tag StLevelTag
  tag.Bgcolor = color.BgBlack
  tag.tag = "INFOF"
  log.UpdateElement(tag)
  ```

- 打印日志

  ```
  log.Println(color.CyanString("msg"))
  log.Fatal("fatal\n")
  log.Error("error\n")
  log.Warn("warn\n")
  log.Info("info\n")
  log.Debug("debug\n")
  log.Trace("trace\n")
  ```

  

