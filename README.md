# find-n-unique-integers-sum-up-to-zero

## 題目解讀：

### 題目來源:
[find-n-unique-integers-sum-up-to-zero](https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/)

### 原文:

Given an integer n, return any array containing n unique integers such that they add up to 0.
### 解讀：

給定一個正整數n, 回傳一個整數陣列 包含n 個整數 符合以下條件:

1 每個整數數值都不同

2 所有整數相加總和為0

## 初步解法:
### 初步觀察:

首先已知總和要是0

代表 陣列中正整數和 跟負整數和 一樣多

最簡單的作法是 

如果 n是偶數 放入 1,...,n/2, -1,...,-n/2

如果是 奇數 則取   M= floor(n/2), 放入 0, 1, ..., M, -1, ..., -M

### 初步設計:

Given an integer n,

Step 0: let M = floor(n/2), integer array result = make([]int, n)

Step 1: if n % 2 == 0, result = {1,..,M, -1, ..., -M}

Step 2: if n % 2 == 1, result = {1,...,M,0,-1,...,-M}

Step 3: return result
## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
```golang
package sum_zero

func sumZero(n int) []int {
	result := make([]int, n)
	M := n / 2
	if n%2 == 0 {
		for idx := range result {
			if idx < M {
				result[idx] = idx + 1
			} else {
				result[idx] = M - idx - 1
			}
		}
	} else {
		for idx := range result {
			if idx < M {
				result[idx] = idx + 1
			} else if idx > M {
				result[idx] = M - idx
			} else {
				result[idx] = 0
			}
		}
	}
	return result
}
```
## 測資的撰寫
```golang
package sum_zero

import (
	"reflect"
	"testing"
)

func Test_sumZero(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				n: 5,
			},
			want: []int{1, 2, 0, -1, -2},
		},
		{
			name: "Example2",
			args: args{
				n: 3,
			},
			want: []int{1, 0, -1},
		},
		{
			name: "Example3",
			args: args{
				n: 1,
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumZero(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

```
## my self record
[golang leetcode 30day 29th day](https://hackmd.io/tiwlpPgjQ4yUhxl9F3cmog?view)
## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)