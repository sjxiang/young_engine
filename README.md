# YoungEngine

本项目实现了一个简单的规则引擎。 
- 引擎自定义了一套词法、语法。
- 在自定义词法语法的基础上实现了一个典型的编译器前端，能够生成表达式对应的抽象语法树。
- 基于编译构建的抽象语法树实现了go版本的虚拟机。通过注入参数可以获得执行结果。

## 词法
引擎支持指定的运算符和数据类型

**运算符**
- 一元计算符 : `!` `-` `+`
- 二元计算符 : `+` `-` `/` `*` `%`
- 二元比较符 : `>` `>=` `<` `<=`  `==` `!=`
- 逻辑操作符 : `||` `&&`
- 括号 : `(` `)`

**数据类型**
- 字符串 `"abc"` `'def'`
- 十进制int `123`
- 十进制float `123.4`
- bool `true`
- 变量 `id`

**表达式词法**
- 表达式以换行结束、不支持多行表达式。形如`a + 7 > 100`
- 支持字面量 (上述数据类型的常量)、变量和运算符(上述运算符)
- 变量：由字母数字下划线构成且必须以字母开头，形如：`_id`、`foo`
- 关键字：系统内置部分关键字 
  - `true`: bool类型常量
  - `false`: bool类型常量

## 语法
支持简单的表达式语法 
- 一元运算: `!true`
- 二元运算: `a + b > c`
- 逻辑运算: `a || b == 100`
- 括号: `(a + b) * c`

运算符的优先级

| 优先级 | 运算符                         |
|-----|-----------------------------|
| 0   | `or`                        |
| 1   | `&&`                        |
| 2   | `!` `-` `+`                 |
| 3   | `>` `>=` `<` `<=` `==` `!=` |
| 4   | `+` `-`                     |
| 5   | `*` `/`                     |

## 项目结构
``` shell
.
├── README.md
├── compiler.go
├── compiler_test.go
├── compiler
│   ├── lexical.go 
│   ├── parser.go   # 语法分析
│   ├── parser_test.go
│   ├── planner.go  # 构建语法树
│   ├── scanner.go  # 词法分析
│   └── scanner_test.go
├── executor
│   ├── ast.go      # 抽象语法树定义
│   ├── operator.go # 语法树执行
│   ├── svg.go      # 可视化打印语法树 - 辅助工具
│   ├── symbol.go   # 符号定义
│   ├── type.go     # 类型定义
│   └── type_checker.go # 类型检查
└── token
    ├── kind.go      # token类型
    ├── kind_test.go
    ├── lexer.go     # 词法定义
    └── token.go     # token定义
```

![](image/node.svg)

# 项目运行

## 检测是否安装项目依赖
```shell
chmod a+x ./setup.sh
./setup.sh
```

## 启动 DB
```shell
docker-compose up
```

## 运行项目
```shell
go run ./main.go
```



## 规则引擎设计与实现

业务逻辑

抖音商城要搞活动啦~
活动期间，用户购买相应的产品会获得商城积分！




输入：计算规则、商品价格、用户标签、商品属性 ..
输出：积分

规则简单，容易配置、可扩展


组成
  数据输入
  规则理解
  规则执行





四、课后作业
4.1 实现一个在线规则引擎
课上我们重点讲了规则引擎的设计和实现，结合前面课程的内容课后实现一个在线版本的规则引擎
4.1.1 项目要求
使用Hertz框架开发一个HTTP服务，服务使用mysql，支持表达式的增删查改和编译执行。
并实现以下接口
直接表达式执行：
请求参数为待执行的表达式和表达式中参数的值，并输出编译结果

实时编译并执行结果，不需要写入DB中


POST api/engine/run
Request

{
    "exp": "uid == 12345 && did > 0",
    "params": {
        "uid": 123456,
        "did": 0
    }
}
复制代码

Response

{
    "code": 0,
    "message": "success",
    "data": {  // 执行结果
        "result": true
    }
}
复制代码
新增表达式：
新增一条表达式到DB中，并返回表达式在DB中的ID
需要检测表达式是否已经存在，如果已经存在，直接返回表达式的ID
需要检测表达式是否合法(编译是否通过) ，如果编译失败，返回错误码 20001和编译错误

POST api/engine/exp/new
Request

{
    "exp": "uid == 12345 && did > 0",
}
复制代码

Response

{
    "code": 0,
    "message": "success",
    "data": {  // 表达式ID
        "id": 1
    }
}

// 编译失败时
{
    "code": -1,
    "message": "compile error: xxxxx", // 编译失败的信息
    "data": {  // 表达式ID
        "id": 0
    }
}
复制代码
查询表达式：
查询数据库中所有的表达式

GET api/engine/exp/list
Response

{
    "code": 0,
    "message": "success",
    "data": [  
        {
            "id": 1,
            "exp": "uid > 0"
        }
    ]
}
复制代码
删除表达式：
根据ID删除表达式，表达式不存在时返回错误码20002 , 和错误信息
删除成功返回被删除的表达式信息

DELETE api/engine/exp/:id
Response

// 删除成功时
{
    "code": 0,
    "message": "success",
    "data": {  // 表达式ID
        "id": 1,
        "exp": "uid > 0"
    }
}

// 删除失败时
{
    "code": -1,
    "message": "exp id 1 not exist", //查询失败的信息
    "data": {}
}
复制代码
执行表达式
根据表达式的ID，查询出表达式内容，并编译执行。表达式不存在时返回错误码20002 , 和错误信息

POST api/engine/exp/run
Request

{
    "exp_id": 1,
    "parmas": {
        "uid": 123456,
        "did": 0
    }
}
复制代码

Response

{
    "code": 0,
    "message": "success",
    "data": {  // 执行结果
        "result": true
    }
}

// 表达式不存在时
{
    "code": -1,
    "message": "exp id 1 not exist", //查询失败的信息
    "data": {}
}
复制代码
