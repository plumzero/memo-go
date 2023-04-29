
Go Web 服务器请求和响应的流程如下:

1.客户端发送请求
2.服务器端的多路复用器收到请求
3.多路复用器根据请求的 URL 找到注册的处理器，将请求交由处理器处理
4.处理器执行程序逻辑，如果必要，则与数据库进行交互，得到处理结果
5.处理器调用模板引擎将指定的模板和上一步得到的结果渲染成客户端可识别的数据格式(通常是 HTML 格式)
6.服务器端将数据通过 HTTP 响应返回给客户端
7.客户端收到数据，执行对应的操作(例如渲染出来呈现给用户)