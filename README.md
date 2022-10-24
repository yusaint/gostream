1. stream.of(1,2,3).filter(xxx).sum()
2. op1->op2->op3--->opN
3. opN是terminalOP，terminalOP中执行evaluate方法，驱动stream执行
4. evaluate方法中，开始头结点，启动head sink
   1. headSink.begin()
   2. headSink.Foreach(headSink)
      1. headSink会链式调用下一级的sink
      2. sink区分stateless和stateful
         1. stateless的是不需要保持状态的，例如做print，来一个输出一个
         2. stateful是需要保持状态的，例如输出到S3，需要等所有数据ready才能输出
         3. 每次调用sink的begin方法驱动一次，end方法结尾一次
   3. headSink.end()