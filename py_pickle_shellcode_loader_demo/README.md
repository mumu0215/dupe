### 思路

python中提供了模块pickle帮助完成序列化

pickle.dumps()将对象obj对象序列化并返回一个byte对象（序列化）

pickle.loads(),从字节对象中读取被封装的对象（反序列化）

序列化封装shellcode：利用类似php的__wakeup__，python中的__reduce__来实现对象被pickle时调用

### 具体使用

使用generater生成序列化后的载荷

使用demo_run.py运行反序列化载荷并达到命令执行的效果

### 本demo只提供基本思路，后续利用可将python shellcode loader代码进行序列化操作，而后反序列化利用执行加载shellcode，具体代码请执行完成

