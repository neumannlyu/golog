# golog

本日志模块自定义程度高，搭配出自己喜欢的样式。

## 使用方法

- 设置level

  ```
  SetLogLevel(LOGLEVEL_ALL)
  ```

- 简单日志样式

  ```
    simplelog := NewSimpleLog()
    simplelog.Logln("这是一个测试的日志。")
  ```

- 设置日志样式

  ```
    Trace("hello ", "pop\n")
    GetCommonLog().TraceTag.Tag = "NewTag"
    Trace("New Hello.\n")
  ```

- 打印日志

  ```
    Fatal("fatal\n")
    Error("error\n")
    Warn("warn\n")
    Info("info\n")
    Debug("debug\n")
    Trace("trace\n")
  ```

  

