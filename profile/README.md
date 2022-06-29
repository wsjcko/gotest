1 安装graphviz(https://graphviz.org/download/)         sudo apt-get install graphviz
2 编辑dot文件，例如ex.dot, 产生图片：dot -Tsvg ex.dot -o ex.svg


Go语言工具链中的 go pprof 可以帮助开发者快速分析及定位各种性能问题，如 CPU 消耗、内存分配及阻塞分析。

性能分析首先需要使用 runtime.pprof 包嵌入到待分析程序的入口和结束处。runtime.pprof 包在运行时对程序进行每秒 100 次的采样，最少采样 1 秒。然后将生成的数据输出，让开发者写入文件或者其他媒介上进行分析。

go pprof 工具链配合 Graphviz 图形化工具可以将 runtime.pprof 包生成的数据转换为 PDF 格式，以图片的方式展示程序的性能分析结果。
安装第三方图形化显式分析数据工具（Graphviz）
Graphviz 是一套通过文本描述的方法生成图形的工具包。描述文本的语言叫做 DOT。

在 www.graphviz.org 网站可以获取到最新的 Graphviz 各平台的安装包。


go build -o cpu cpu.go        

./cpu
2022/06/01 11:20:58 profile: cpu profiling enabled, cpu.pprof
2022/06/01 11:20:59 profile: cpu profiling disabled, cpu.pprof

go tool pprof --pdf cpu cpu.pprof > cpu.pdf