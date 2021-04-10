package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	xflag:=1//记录符号
	sB:=[]byte(str)//转字节数组好处理
	l:=0//前有效位：left

	//' ' ==> 32 '+' ==> 43 ‘-’==>45 '0'==>48 '9'==>57
	for i:=0;i<len(sB)&&sB[i]==32;i++{//跳过空格字符
		l++
	}

	if l==len(sB){//字符串全是空格的请款
		return 0
	}

	if sB[l]==43{//判断首位符号
		l++
	}else if sB[l]==45{
		xflag=-1
		l++
	}

	r:=l
	//找右界
	for i:=l;i<len(sB)&&sB[r]>=48&&sB[r]<=57;i++{
		r++
	}

	num:=make(map[byte]int)//字符与数字的映射
	var k byte=48
	for i:=0;i<10;i++{
		num[k]=i
		k++
	}

	result:=0
	for _,v:=range sB[l:r]{
		result=result*10+num[v]//构建新的数
		//溢出情况处理
		if result>math.MaxInt32&&xflag==1{
			return math.MaxInt32
		}else if result>math.MaxInt32&&xflag==-1{
			return math.MinInt32
		}
	}
	result*=xflag//符号位决定正负
	return result
}


func main() {
	fmt.Println(myAtoi("-50"))
}
