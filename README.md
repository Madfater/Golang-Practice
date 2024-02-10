# Golang-Practice

## 安裝環境

```bash
sudo apt-get install golang-go
```

## module管理

- 專案初始化
```bash
go mod init <projectName>
```
- 維護專案的函示庫引入
```bash
go mod tidy
```

## 賦值

```go
//方式一
var example string = "This is example"
//方式二
var exaple = "This is example"
//方式三
example := "this is example"
```

## 函式

- 大小區分是否為public or private

```go
//private
func example(test string)int{

}
//public
func Example(test string)int{

}
```

## Array

- 靜態的Array
- size固定
- 在宣告時要固定長度
- 可使用...讓其自動判定宣告array的長度

### Array的宣告

```go
//方式一
var arr = [5]int
arr := [5]int{1,2,3,4,5}

//方式二 使用...自動判定長度
var arr = [...]int{1,2,3} //len=3
```

## Slice

- 動態的Array
- Size不固定
- 可使用append新增
- 空間不夠時會將原有空間翻倍
- 可使用Python的用法將Slice分割
- 在Slice中有capacity和length兩種長度
	- length: Slice的元素數量
	- capacity: Slice的空間長度

### Slice的宣告

```go
//方式一 使用make
var arr = make([]int,3,3)//make(宣告的型別,length,capacity)
//方式二
var arr = []int
var arr = []int{1,2,3}
```

### Slice的操作

```go
//新增
arr.append()

//擷取 可用來刪除
arr[1:3:4]//[擷取的頭:擷取的尾+1:擷取的capacity]
```

## Map

- 與Python的dictionary差不多

### Map的宣告
```go
//方法一 使用make
var m = make([int]string)

//方法二
var m = [int]string{}
```

### Map的操作
```go
//新增
m[1]="tom"

//刪除
delete(m,1)
```

## Struct

- 與C++的差不多
- 首字需大寫

```go
type Person struct{
	Name string
	Age number
}
```

### struct的宣告

```go
//使用指標比較方便
var p = &Person{
	Name:"Tom",
	Age:6,
}

//存取不須改為箭頭
p.name
```

## Json的轉換

- 可以將Map與Struct轉成Json
- Struct需要上標籤

```go
struct Person struct{
	Name string `json"name"`
	Age int `json"age"`
}
```
### 轉換方式

- obj皆為指標
```go
//obj to Json
json,err = json.Marshel(obj)
//json to obj
err = json.unMarshel(json,obj)//struct
err = json.unMarshel([]byte(jsons),obj)//map
```
## 常用函式

```go
//取元素數
len()
//取空間長度
cap()
```

## 協程

- 使用關鍵字go開始一個新線程
- 使用chan可以讓多個線程進行安全通訊
- 關鍵字defer可以確保線程結束後執行 如垃圾清理、關閉連線等