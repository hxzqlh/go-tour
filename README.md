## Go 指南练习

在 [Go 指南](https://tour.go-zh.org/list) 的指导下，完成下列练习：

* [循环和函数](https://tour.go-zh.org/flowcontrol/8)
* [slice](https://tour.go-zh.org/moretypes/18)
* [map](https://tour.go-zh.org/moretypes/23)
* [斐波纳契闭包](https://tour.go-zh.org/moretypes/26)
* [Stringers](https://tour.go-zh.org/methods/18)
* [错误](https://tour.go-zh.org/methods/20)
* [Reader](https://tour.go-zh.org/methods/22)
* [rot13Reader](https://tour.go-zh.org/methods/23)
* [图片](https://tour.go-zh.org/methods/25)
* [等价二叉树](https://tour.go-zh.org/concurrency/8)
* [Web爬虫](https://tour.go-zh.org/concurrency/10)

## Go 高级练习

1. 有四个协程1、2、3、4, 协程1的功能就是输出1，协程2的功能就是输出2，以此类推……现在有四个文件A、B、C、D，初始都为空，编程实现让这四个文件呈现如下格式：

    ```
    A：1 2 3 4 1 2...
    B：2 3 4 1 2 3...
    C：3 4 1 2 3 4...
    D：4 1 2 3 4 1...
    ```

2. 交替打印数字和字母，使得最终效果如下：

    ```
    12AB34CD56EF78GH910IJ...
    ```