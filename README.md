# gostream

`gostream` 项目受到 [Go-Stream](https://github.com/reugn/go-streams) 项目的启发，在此感谢该项目作者的工作成果。
`gostream` 参考了JDK中stream标准库的思路，尽量提供一个贴近JDK stream库的使用体验,当然整体实现和JDK有很大的差距。 

## 使用介绍
gostream整体基于Go泛型实现，所以如果想使用的话，需要Go的版本不低于1.8。项目整体结构分为三大块
* 基本的数据结构支持库 - 主要是`arrays`,`set`,`generic`三个文件夹，提供了基本的泛型数据结构，当然还在不断扩展中。
* Stream库 - 提供了常用的操作算子，例如`filter`,`groupby`,`parallel`,`sort`, `map`等等。
* 扩展库 - 封装了数据源的操作工具类库，例如支持将文件流转为一个stream数据源，将kafka消息转为一个stream数据源等，这里也是支持扩展的。只需要实现了`Splittable`接口即可。


## Examples
Usage samples are available in the examples directory.

## License
Licensed under the MIT License.