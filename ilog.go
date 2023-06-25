package golog

type ILog interface {
    // 设置日志等级
    SetLogLevel(int)
    // 更新元素。
    // 将日志中的对应的元素用新值替换掉
    UpdateElement(newelement ILogElement)
}
