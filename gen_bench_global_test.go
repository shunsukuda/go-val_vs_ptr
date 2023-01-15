package main

import "testing"

// type W1 [1]uint64
var rv1 = W1{}
var arg1vW1 = W1{}
var arg2vW1 = W1{}
var arg3vW1 = W1{}
var rp1 = new(W1)
var arg1pW1 = new(W1)
var arg2pW1 = new(W1)
var arg3pW1 = new(W1)

// type W2 [2]uint64
var rv2 = W2{}
var arg1vW2 = W2{}
var arg2vW2 = W2{}
var arg3vW2 = W2{}
var rp2 = new(W2)
var arg1pW2 = new(W2)
var arg2pW2 = new(W2)
var arg3pW2 = new(W2)

// type W3 [3]uint64
var rv3 = W3{}
var arg1vW3 = W3{}
var arg2vW3 = W3{}
var arg3vW3 = W3{}
var rp3 = new(W3)
var arg1pW3 = new(W3)
var arg2pW3 = new(W3)
var arg3pW3 = new(W3)

// type W4 [4]uint64
var rv4 = W4{}
var arg1vW4 = W4{}
var arg2vW4 = W4{}
var arg3vW4 = W4{}
var rp4 = new(W4)
var arg1pW4 = new(W4)
var arg2pW4 = new(W4)
var arg3pW4 = new(W4)

// type W5 [5]uint64
var rv5 = W5{}
var arg1vW5 = W5{}
var arg2vW5 = W5{}
var arg3vW5 = W5{}
var rp5 = new(W5)
var arg1pW5 = new(W5)
var arg2pW5 = new(W5)
var arg3pW5 = new(W5)

// type W6 [6]uint64
var rv6 = W6{}
var arg1vW6 = W6{}
var arg2vW6 = W6{}
var arg3vW6 = W6{}
var rp6 = new(W6)
var arg1pW6 = new(W6)
var arg2pW6 = new(W6)
var arg3pW6 = new(W6)

// type W7 [7]uint64
var rv7 = W7{}
var arg1vW7 = W7{}
var arg2vW7 = W7{}
var arg3vW7 = W7{}
var rp7 = new(W7)
var arg1pW7 = new(W7)
var arg2pW7 = new(W7)
var arg3pW7 = new(W7)

// type W8 [8]uint64
var rv8 = W8{}
var arg1vW8 = W8{}
var arg2vW8 = W8{}
var arg3vW8 = W8{}
var rp8 = new(W8)
var arg1pW8 = new(W8)
var arg2pW8 = new(W8)
var arg3pW8 = new(W8)

// type W9 [9]uint64
var rv9 = W9{}
var arg1vW9 = W9{}
var arg2vW9 = W9{}
var arg3vW9 = W9{}
var rp9 = new(W9)
var arg1pW9 = new(W9)
var arg2pW9 = new(W9)
var arg3pW9 = new(W9)

// type W10 [10]uint64
var rv10 = W10{}
var arg1vW10 = W10{}
var arg2vW10 = W10{}
var arg3vW10 = W10{}
var rp10 = new(W10)
var arg1pW10 = new(W10)
var arg2pW10 = new(W10)
var arg3pW10 = new(W10)

//go:noinline
func FuncNoinlineGlobal1Arg0ReturnValue1Word(arg W1) {
	_ = arg
}

//go:noinline
func FuncNoinlineGlobal2Arg0ReturnValue1Word(arg1, arg2 W1) {
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg0ReturnValue1Word(arg1, arg2, arg3 W1) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinlineGlobal1Arg1ReturnValue1Word(arg W1) W1 {
	_ = arg
	return arg
}

//go:noinline
func FuncNoinlineGlobal2Arg1ReturnValue1Word(arg1, arg2 W1) W1 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinlineGlobal3Arg1ReturnValue1Word(arg1, arg2, arg3 W1) W1 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinlineGlobal1Arg2ReturnValue1Word(arg W1) (W1, W1) {
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinlineGlobal2Arg2ReturnValue1Word(arg1, arg2 W1) (W1, W1) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg2ReturnValue1Word(arg1, arg2, arg3 W1) (W1, W1) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal1Arg0ReturnValue2Word(arg W2) {
	_ = arg
}

//go:noinline
func FuncNoinlineGlobal2Arg0ReturnValue2Word(arg1, arg2 W2) {
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg0ReturnValue2Word(arg1, arg2, arg3 W2) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinlineGlobal1Arg1ReturnValue2Word(arg W2) W2 {
	_ = arg
	return arg
}

//go:noinline
func FuncNoinlineGlobal2Arg1ReturnValue2Word(arg1, arg2 W2) W2 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinlineGlobal3Arg1ReturnValue2Word(arg1, arg2, arg3 W2) W2 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinlineGlobal1Arg2ReturnValue2Word(arg W2) (W2, W2) {
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinlineGlobal2Arg2ReturnValue2Word(arg1, arg2 W2) (W2, W2) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg2ReturnValue2Word(arg1, arg2, arg3 W2) (W2, W2) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal1Arg0ReturnValue3Word(arg W3) {
	_ = arg
}

//go:noinline
func FuncNoinlineGlobal2Arg0ReturnValue3Word(arg1, arg2 W3) {
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg0ReturnValue3Word(arg1, arg2, arg3 W3) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinlineGlobal1Arg1ReturnValue3Word(arg W3) W3 {
	_ = arg
	return arg
}

//go:noinline
func FuncNoinlineGlobal2Arg1ReturnValue3Word(arg1, arg2 W3) W3 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinlineGlobal3Arg1ReturnValue3Word(arg1, arg2, arg3 W3) W3 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinlineGlobal1Arg2ReturnValue3Word(arg W3) (W3, W3) {
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinlineGlobal2Arg2ReturnValue3Word(arg1, arg2 W3) (W3, W3) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg2ReturnValue3Word(arg1, arg2, arg3 W3) (W3, W3) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal1Arg0ReturnValue4Word(arg W4) {
	_ = arg
}

//go:noinline
func FuncNoinlineGlobal2Arg0ReturnValue4Word(arg1, arg2 W4) {
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg0ReturnValue4Word(arg1, arg2, arg3 W4) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinlineGlobal1Arg1ReturnValue4Word(arg W4) W4 {
	_ = arg
	return arg
}

//go:noinline
func FuncNoinlineGlobal2Arg1ReturnValue4Word(arg1, arg2 W4) W4 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinlineGlobal3Arg1ReturnValue4Word(arg1, arg2, arg3 W4) W4 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinlineGlobal1Arg2ReturnValue4Word(arg W4) (W4, W4) {
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinlineGlobal2Arg2ReturnValue4Word(arg1, arg2 W4) (W4, W4) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg2ReturnValue4Word(arg1, arg2, arg3 W4) (W4, W4) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal1Arg0ReturnValue5Word(arg W5) {
	_ = arg
}

//go:noinline
func FuncNoinlineGlobal2Arg0ReturnValue5Word(arg1, arg2 W5) {
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg0ReturnValue5Word(arg1, arg2, arg3 W5) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinlineGlobal1Arg1ReturnValue5Word(arg W5) W5 {
	_ = arg
	return arg
}

//go:noinline
func FuncNoinlineGlobal2Arg1ReturnValue5Word(arg1, arg2 W5) W5 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinlineGlobal3Arg1ReturnValue5Word(arg1, arg2, arg3 W5) W5 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinlineGlobal1Arg2ReturnValue5Word(arg W5) (W5, W5) {
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinlineGlobal2Arg2ReturnValue5Word(arg1, arg2 W5) (W5, W5) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg2ReturnValue5Word(arg1, arg2, arg3 W5) (W5, W5) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal1Arg0ReturnValue6Word(arg W6) {
	_ = arg
}

//go:noinline
func FuncNoinlineGlobal2Arg0ReturnValue6Word(arg1, arg2 W6) {
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg0ReturnValue6Word(arg1, arg2, arg3 W6) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinlineGlobal1Arg1ReturnValue6Word(arg W6) W6 {
	_ = arg
	return arg
}

//go:noinline
func FuncNoinlineGlobal2Arg1ReturnValue6Word(arg1, arg2 W6) W6 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinlineGlobal3Arg1ReturnValue6Word(arg1, arg2, arg3 W6) W6 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinlineGlobal1Arg2ReturnValue6Word(arg W6) (W6, W6) {
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinlineGlobal2Arg2ReturnValue6Word(arg1, arg2 W6) (W6, W6) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg2ReturnValue6Word(arg1, arg2, arg3 W6) (W6, W6) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal1Arg0ReturnValue7Word(arg W7) {
	_ = arg
}

//go:noinline
func FuncNoinlineGlobal2Arg0ReturnValue7Word(arg1, arg2 W7) {
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg0ReturnValue7Word(arg1, arg2, arg3 W7) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinlineGlobal1Arg1ReturnValue7Word(arg W7) W7 {
	_ = arg
	return arg
}

//go:noinline
func FuncNoinlineGlobal2Arg1ReturnValue7Word(arg1, arg2 W7) W7 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinlineGlobal3Arg1ReturnValue7Word(arg1, arg2, arg3 W7) W7 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinlineGlobal1Arg2ReturnValue7Word(arg W7) (W7, W7) {
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinlineGlobal2Arg2ReturnValue7Word(arg1, arg2 W7) (W7, W7) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg2ReturnValue7Word(arg1, arg2, arg3 W7) (W7, W7) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal1Arg0ReturnValue8Word(arg W8) {
	_ = arg
}

//go:noinline
func FuncNoinlineGlobal2Arg0ReturnValue8Word(arg1, arg2 W8) {
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg0ReturnValue8Word(arg1, arg2, arg3 W8) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinlineGlobal1Arg1ReturnValue8Word(arg W8) W8 {
	_ = arg
	return arg
}

//go:noinline
func FuncNoinlineGlobal2Arg1ReturnValue8Word(arg1, arg2 W8) W8 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinlineGlobal3Arg1ReturnValue8Word(arg1, arg2, arg3 W8) W8 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinlineGlobal1Arg2ReturnValue8Word(arg W8) (W8, W8) {
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinlineGlobal2Arg2ReturnValue8Word(arg1, arg2 W8) (W8, W8) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg2ReturnValue8Word(arg1, arg2, arg3 W8) (W8, W8) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal1Arg0ReturnValue9Word(arg W9) {
	_ = arg
}

//go:noinline
func FuncNoinlineGlobal2Arg0ReturnValue9Word(arg1, arg2 W9) {
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg0ReturnValue9Word(arg1, arg2, arg3 W9) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinlineGlobal1Arg1ReturnValue9Word(arg W9) W9 {
	_ = arg
	return arg
}

//go:noinline
func FuncNoinlineGlobal2Arg1ReturnValue9Word(arg1, arg2 W9) W9 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinlineGlobal3Arg1ReturnValue9Word(arg1, arg2, arg3 W9) W9 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinlineGlobal1Arg2ReturnValue9Word(arg W9) (W9, W9) {
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinlineGlobal2Arg2ReturnValue9Word(arg1, arg2 W9) (W9, W9) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg2ReturnValue9Word(arg1, arg2, arg3 W9) (W9, W9) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal1Arg0ReturnValue10Word(arg W10) {
	_ = arg
}

//go:noinline
func FuncNoinlineGlobal2Arg0ReturnValue10Word(arg1, arg2 W10) {
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg0ReturnValue10Word(arg1, arg2, arg3 W10) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinlineGlobal1Arg1ReturnValue10Word(arg W10) W10 {
	_ = arg
	return arg
}

//go:noinline
func FuncNoinlineGlobal2Arg1ReturnValue10Word(arg1, arg2 W10) W10 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinlineGlobal3Arg1ReturnValue10Word(arg1, arg2, arg3 W10) W10 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinlineGlobal1Arg2ReturnValue10Word(arg W10) (W10, W10) {
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinlineGlobal2Arg2ReturnValue10Word(arg1, arg2 W10) (W10, W10) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg2ReturnValue10Word(arg1, arg2, arg3 W10) (W10, W10) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal1Arg0ReturnPointer1Word(arg *W1) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func FuncNoinlineGlobal2Arg0ReturnPointer1Word(arg1, arg2 *W1) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg0ReturnPointer1Word(arg1, arg2, arg3 *W1) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinlineGlobal1Arg1ReturnPointer1Word(arg *W1) *W1 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func FuncNoinlineGlobal2Arg1ReturnPointer1Word(arg1, arg2 *W1) *W1 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinlineGlobal3Arg1ReturnPointer1Word(arg1, arg2, arg3 *W1) *W1 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinlineGlobal1Arg2ReturnPointer1Word(arg *W1) (*W1, *W1) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinlineGlobal2Arg2ReturnPointer1Word(arg1, arg2 *W1) (*W1, *W1) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg2ReturnPointer1Word(arg1, arg2, arg3 *W1) (*W1, *W1) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal1Arg0ReturnPointer2Word(arg *W2) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func FuncNoinlineGlobal2Arg0ReturnPointer2Word(arg1, arg2 *W2) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg0ReturnPointer2Word(arg1, arg2, arg3 *W2) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinlineGlobal1Arg1ReturnPointer2Word(arg *W2) *W2 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func FuncNoinlineGlobal2Arg1ReturnPointer2Word(arg1, arg2 *W2) *W2 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinlineGlobal3Arg1ReturnPointer2Word(arg1, arg2, arg3 *W2) *W2 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinlineGlobal1Arg2ReturnPointer2Word(arg *W2) (*W2, *W2) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinlineGlobal2Arg2ReturnPointer2Word(arg1, arg2 *W2) (*W2, *W2) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg2ReturnPointer2Word(arg1, arg2, arg3 *W2) (*W2, *W2) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal1Arg0ReturnPointer3Word(arg *W3) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func FuncNoinlineGlobal2Arg0ReturnPointer3Word(arg1, arg2 *W3) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg0ReturnPointer3Word(arg1, arg2, arg3 *W3) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinlineGlobal1Arg1ReturnPointer3Word(arg *W3) *W3 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func FuncNoinlineGlobal2Arg1ReturnPointer3Word(arg1, arg2 *W3) *W3 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinlineGlobal3Arg1ReturnPointer3Word(arg1, arg2, arg3 *W3) *W3 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinlineGlobal1Arg2ReturnPointer3Word(arg *W3) (*W3, *W3) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinlineGlobal2Arg2ReturnPointer3Word(arg1, arg2 *W3) (*W3, *W3) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg2ReturnPointer3Word(arg1, arg2, arg3 *W3) (*W3, *W3) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal1Arg0ReturnPointer4Word(arg *W4) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func FuncNoinlineGlobal2Arg0ReturnPointer4Word(arg1, arg2 *W4) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg0ReturnPointer4Word(arg1, arg2, arg3 *W4) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinlineGlobal1Arg1ReturnPointer4Word(arg *W4) *W4 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func FuncNoinlineGlobal2Arg1ReturnPointer4Word(arg1, arg2 *W4) *W4 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinlineGlobal3Arg1ReturnPointer4Word(arg1, arg2, arg3 *W4) *W4 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinlineGlobal1Arg2ReturnPointer4Word(arg *W4) (*W4, *W4) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinlineGlobal2Arg2ReturnPointer4Word(arg1, arg2 *W4) (*W4, *W4) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg2ReturnPointer4Word(arg1, arg2, arg3 *W4) (*W4, *W4) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal1Arg0ReturnPointer5Word(arg *W5) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func FuncNoinlineGlobal2Arg0ReturnPointer5Word(arg1, arg2 *W5) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg0ReturnPointer5Word(arg1, arg2, arg3 *W5) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinlineGlobal1Arg1ReturnPointer5Word(arg *W5) *W5 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func FuncNoinlineGlobal2Arg1ReturnPointer5Word(arg1, arg2 *W5) *W5 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinlineGlobal3Arg1ReturnPointer5Word(arg1, arg2, arg3 *W5) *W5 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinlineGlobal1Arg2ReturnPointer5Word(arg *W5) (*W5, *W5) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinlineGlobal2Arg2ReturnPointer5Word(arg1, arg2 *W5) (*W5, *W5) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg2ReturnPointer5Word(arg1, arg2, arg3 *W5) (*W5, *W5) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal1Arg0ReturnPointer6Word(arg *W6) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func FuncNoinlineGlobal2Arg0ReturnPointer6Word(arg1, arg2 *W6) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg0ReturnPointer6Word(arg1, arg2, arg3 *W6) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinlineGlobal1Arg1ReturnPointer6Word(arg *W6) *W6 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func FuncNoinlineGlobal2Arg1ReturnPointer6Word(arg1, arg2 *W6) *W6 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinlineGlobal3Arg1ReturnPointer6Word(arg1, arg2, arg3 *W6) *W6 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinlineGlobal1Arg2ReturnPointer6Word(arg *W6) (*W6, *W6) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinlineGlobal2Arg2ReturnPointer6Word(arg1, arg2 *W6) (*W6, *W6) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg2ReturnPointer6Word(arg1, arg2, arg3 *W6) (*W6, *W6) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal1Arg0ReturnPointer7Word(arg *W7) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func FuncNoinlineGlobal2Arg0ReturnPointer7Word(arg1, arg2 *W7) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg0ReturnPointer7Word(arg1, arg2, arg3 *W7) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinlineGlobal1Arg1ReturnPointer7Word(arg *W7) *W7 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func FuncNoinlineGlobal2Arg1ReturnPointer7Word(arg1, arg2 *W7) *W7 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinlineGlobal3Arg1ReturnPointer7Word(arg1, arg2, arg3 *W7) *W7 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinlineGlobal1Arg2ReturnPointer7Word(arg *W7) (*W7, *W7) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinlineGlobal2Arg2ReturnPointer7Word(arg1, arg2 *W7) (*W7, *W7) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg2ReturnPointer7Word(arg1, arg2, arg3 *W7) (*W7, *W7) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal1Arg0ReturnPointer8Word(arg *W8) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func FuncNoinlineGlobal2Arg0ReturnPointer8Word(arg1, arg2 *W8) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg0ReturnPointer8Word(arg1, arg2, arg3 *W8) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinlineGlobal1Arg1ReturnPointer8Word(arg *W8) *W8 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func FuncNoinlineGlobal2Arg1ReturnPointer8Word(arg1, arg2 *W8) *W8 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinlineGlobal3Arg1ReturnPointer8Word(arg1, arg2, arg3 *W8) *W8 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinlineGlobal1Arg2ReturnPointer8Word(arg *W8) (*W8, *W8) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinlineGlobal2Arg2ReturnPointer8Word(arg1, arg2 *W8) (*W8, *W8) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg2ReturnPointer8Word(arg1, arg2, arg3 *W8) (*W8, *W8) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal1Arg0ReturnPointer9Word(arg *W9) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func FuncNoinlineGlobal2Arg0ReturnPointer9Word(arg1, arg2 *W9) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg0ReturnPointer9Word(arg1, arg2, arg3 *W9) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinlineGlobal1Arg1ReturnPointer9Word(arg *W9) *W9 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func FuncNoinlineGlobal2Arg1ReturnPointer9Word(arg1, arg2 *W9) *W9 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinlineGlobal3Arg1ReturnPointer9Word(arg1, arg2, arg3 *W9) *W9 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinlineGlobal1Arg2ReturnPointer9Word(arg *W9) (*W9, *W9) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinlineGlobal2Arg2ReturnPointer9Word(arg1, arg2 *W9) (*W9, *W9) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg2ReturnPointer9Word(arg1, arg2, arg3 *W9) (*W9, *W9) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal1Arg0ReturnPointer10Word(arg *W10) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func FuncNoinlineGlobal2Arg0ReturnPointer10Word(arg1, arg2 *W10) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg0ReturnPointer10Word(arg1, arg2, arg3 *W10) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinlineGlobal1Arg1ReturnPointer10Word(arg *W10) *W10 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func FuncNoinlineGlobal2Arg1ReturnPointer10Word(arg1, arg2 *W10) *W10 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinlineGlobal3Arg1ReturnPointer10Word(arg1, arg2, arg3 *W10) *W10 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinlineGlobal1Arg2ReturnPointer10Word(arg *W10) (*W10, *W10) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinlineGlobal2Arg2ReturnPointer10Word(arg1, arg2 *W10) (*W10, *W10) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinlineGlobal3Arg2ReturnPointer10Word(arg1, arg2, arg3 *W10) (*W10, *W10) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W1) ValueReciverNoinlineGlobal1Arg0ReturnValue1Word(arg W1) {
	_ = arg
}

//go:noinline
func (r W1) ValueReciverNoinlineGlobal2Arg0ReturnValue1Word(arg1, arg2 W1) {
	_, _ = arg1, arg2
}

//go:noinline
func (r W1) ValueReciverNoinlineGlobal3Arg0ReturnValue1Word(arg1, arg2, arg3 W1) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W1) ValueReciverNoinlineGlobal1Arg1ReturnValue1Word(arg W1) W1 {
	_ = arg
	return arg
}

//go:noinline
func (r W1) ValueReciverNoinlineGlobal2Arg1ReturnValue1Word(arg1, arg2 W1) W1 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W1) ValueReciverNoinlineGlobal3Arg1ReturnValue1Word(arg1, arg2, arg3 W1) W1 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W1) ValueReciverNoinlineGlobal1Arg2ReturnValue1Word(arg W1) (W1, W1) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r W1) ValueReciverNoinlineGlobal2Arg2ReturnValue1Word(arg1, arg2 W1) (W1, W1) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W1) ValueReciverNoinlineGlobal3Arg2ReturnValue1Word(arg1, arg2, arg3 W1) (W1, W1) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W2) ValueReciverNoinlineGlobal1Arg0ReturnValue2Word(arg W2) {
	_ = arg
}

//go:noinline
func (r W2) ValueReciverNoinlineGlobal2Arg0ReturnValue2Word(arg1, arg2 W2) {
	_, _ = arg1, arg2
}

//go:noinline
func (r W2) ValueReciverNoinlineGlobal3Arg0ReturnValue2Word(arg1, arg2, arg3 W2) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W2) ValueReciverNoinlineGlobal1Arg1ReturnValue2Word(arg W2) W2 {
	_ = arg
	return arg
}

//go:noinline
func (r W2) ValueReciverNoinlineGlobal2Arg1ReturnValue2Word(arg1, arg2 W2) W2 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W2) ValueReciverNoinlineGlobal3Arg1ReturnValue2Word(arg1, arg2, arg3 W2) W2 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W2) ValueReciverNoinlineGlobal1Arg2ReturnValue2Word(arg W2) (W2, W2) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r W2) ValueReciverNoinlineGlobal2Arg2ReturnValue2Word(arg1, arg2 W2) (W2, W2) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W2) ValueReciverNoinlineGlobal3Arg2ReturnValue2Word(arg1, arg2, arg3 W2) (W2, W2) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W3) ValueReciverNoinlineGlobal1Arg0ReturnValue3Word(arg W3) {
	_ = arg
}

//go:noinline
func (r W3) ValueReciverNoinlineGlobal2Arg0ReturnValue3Word(arg1, arg2 W3) {
	_, _ = arg1, arg2
}

//go:noinline
func (r W3) ValueReciverNoinlineGlobal3Arg0ReturnValue3Word(arg1, arg2, arg3 W3) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W3) ValueReciverNoinlineGlobal1Arg1ReturnValue3Word(arg W3) W3 {
	_ = arg
	return arg
}

//go:noinline
func (r W3) ValueReciverNoinlineGlobal2Arg1ReturnValue3Word(arg1, arg2 W3) W3 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W3) ValueReciverNoinlineGlobal3Arg1ReturnValue3Word(arg1, arg2, arg3 W3) W3 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W3) ValueReciverNoinlineGlobal1Arg2ReturnValue3Word(arg W3) (W3, W3) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r W3) ValueReciverNoinlineGlobal2Arg2ReturnValue3Word(arg1, arg2 W3) (W3, W3) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W3) ValueReciverNoinlineGlobal3Arg2ReturnValue3Word(arg1, arg2, arg3 W3) (W3, W3) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W4) ValueReciverNoinlineGlobal1Arg0ReturnValue4Word(arg W4) {
	_ = arg
}

//go:noinline
func (r W4) ValueReciverNoinlineGlobal2Arg0ReturnValue4Word(arg1, arg2 W4) {
	_, _ = arg1, arg2
}

//go:noinline
func (r W4) ValueReciverNoinlineGlobal3Arg0ReturnValue4Word(arg1, arg2, arg3 W4) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W4) ValueReciverNoinlineGlobal1Arg1ReturnValue4Word(arg W4) W4 {
	_ = arg
	return arg
}

//go:noinline
func (r W4) ValueReciverNoinlineGlobal2Arg1ReturnValue4Word(arg1, arg2 W4) W4 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W4) ValueReciverNoinlineGlobal3Arg1ReturnValue4Word(arg1, arg2, arg3 W4) W4 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W4) ValueReciverNoinlineGlobal1Arg2ReturnValue4Word(arg W4) (W4, W4) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r W4) ValueReciverNoinlineGlobal2Arg2ReturnValue4Word(arg1, arg2 W4) (W4, W4) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W4) ValueReciverNoinlineGlobal3Arg2ReturnValue4Word(arg1, arg2, arg3 W4) (W4, W4) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W5) ValueReciverNoinlineGlobal1Arg0ReturnValue5Word(arg W5) {
	_ = arg
}

//go:noinline
func (r W5) ValueReciverNoinlineGlobal2Arg0ReturnValue5Word(arg1, arg2 W5) {
	_, _ = arg1, arg2
}

//go:noinline
func (r W5) ValueReciverNoinlineGlobal3Arg0ReturnValue5Word(arg1, arg2, arg3 W5) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W5) ValueReciverNoinlineGlobal1Arg1ReturnValue5Word(arg W5) W5 {
	_ = arg
	return arg
}

//go:noinline
func (r W5) ValueReciverNoinlineGlobal2Arg1ReturnValue5Word(arg1, arg2 W5) W5 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W5) ValueReciverNoinlineGlobal3Arg1ReturnValue5Word(arg1, arg2, arg3 W5) W5 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W5) ValueReciverNoinlineGlobal1Arg2ReturnValue5Word(arg W5) (W5, W5) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r W5) ValueReciverNoinlineGlobal2Arg2ReturnValue5Word(arg1, arg2 W5) (W5, W5) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W5) ValueReciverNoinlineGlobal3Arg2ReturnValue5Word(arg1, arg2, arg3 W5) (W5, W5) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W6) ValueReciverNoinlineGlobal1Arg0ReturnValue6Word(arg W6) {
	_ = arg
}

//go:noinline
func (r W6) ValueReciverNoinlineGlobal2Arg0ReturnValue6Word(arg1, arg2 W6) {
	_, _ = arg1, arg2
}

//go:noinline
func (r W6) ValueReciverNoinlineGlobal3Arg0ReturnValue6Word(arg1, arg2, arg3 W6) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W6) ValueReciverNoinlineGlobal1Arg1ReturnValue6Word(arg W6) W6 {
	_ = arg
	return arg
}

//go:noinline
func (r W6) ValueReciverNoinlineGlobal2Arg1ReturnValue6Word(arg1, arg2 W6) W6 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W6) ValueReciverNoinlineGlobal3Arg1ReturnValue6Word(arg1, arg2, arg3 W6) W6 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W6) ValueReciverNoinlineGlobal1Arg2ReturnValue6Word(arg W6) (W6, W6) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r W6) ValueReciverNoinlineGlobal2Arg2ReturnValue6Word(arg1, arg2 W6) (W6, W6) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W6) ValueReciverNoinlineGlobal3Arg2ReturnValue6Word(arg1, arg2, arg3 W6) (W6, W6) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W7) ValueReciverNoinlineGlobal1Arg0ReturnValue7Word(arg W7) {
	_ = arg
}

//go:noinline
func (r W7) ValueReciverNoinlineGlobal2Arg0ReturnValue7Word(arg1, arg2 W7) {
	_, _ = arg1, arg2
}

//go:noinline
func (r W7) ValueReciverNoinlineGlobal3Arg0ReturnValue7Word(arg1, arg2, arg3 W7) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W7) ValueReciverNoinlineGlobal1Arg1ReturnValue7Word(arg W7) W7 {
	_ = arg
	return arg
}

//go:noinline
func (r W7) ValueReciverNoinlineGlobal2Arg1ReturnValue7Word(arg1, arg2 W7) W7 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W7) ValueReciverNoinlineGlobal3Arg1ReturnValue7Word(arg1, arg2, arg3 W7) W7 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W7) ValueReciverNoinlineGlobal1Arg2ReturnValue7Word(arg W7) (W7, W7) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r W7) ValueReciverNoinlineGlobal2Arg2ReturnValue7Word(arg1, arg2 W7) (W7, W7) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W7) ValueReciverNoinlineGlobal3Arg2ReturnValue7Word(arg1, arg2, arg3 W7) (W7, W7) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W8) ValueReciverNoinlineGlobal1Arg0ReturnValue8Word(arg W8) {
	_ = arg
}

//go:noinline
func (r W8) ValueReciverNoinlineGlobal2Arg0ReturnValue8Word(arg1, arg2 W8) {
	_, _ = arg1, arg2
}

//go:noinline
func (r W8) ValueReciverNoinlineGlobal3Arg0ReturnValue8Word(arg1, arg2, arg3 W8) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W8) ValueReciverNoinlineGlobal1Arg1ReturnValue8Word(arg W8) W8 {
	_ = arg
	return arg
}

//go:noinline
func (r W8) ValueReciverNoinlineGlobal2Arg1ReturnValue8Word(arg1, arg2 W8) W8 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W8) ValueReciverNoinlineGlobal3Arg1ReturnValue8Word(arg1, arg2, arg3 W8) W8 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W8) ValueReciverNoinlineGlobal1Arg2ReturnValue8Word(arg W8) (W8, W8) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r W8) ValueReciverNoinlineGlobal2Arg2ReturnValue8Word(arg1, arg2 W8) (W8, W8) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W8) ValueReciverNoinlineGlobal3Arg2ReturnValue8Word(arg1, arg2, arg3 W8) (W8, W8) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W9) ValueReciverNoinlineGlobal1Arg0ReturnValue9Word(arg W9) {
	_ = arg
}

//go:noinline
func (r W9) ValueReciverNoinlineGlobal2Arg0ReturnValue9Word(arg1, arg2 W9) {
	_, _ = arg1, arg2
}

//go:noinline
func (r W9) ValueReciverNoinlineGlobal3Arg0ReturnValue9Word(arg1, arg2, arg3 W9) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W9) ValueReciverNoinlineGlobal1Arg1ReturnValue9Word(arg W9) W9 {
	_ = arg
	return arg
}

//go:noinline
func (r W9) ValueReciverNoinlineGlobal2Arg1ReturnValue9Word(arg1, arg2 W9) W9 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W9) ValueReciverNoinlineGlobal3Arg1ReturnValue9Word(arg1, arg2, arg3 W9) W9 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W9) ValueReciverNoinlineGlobal1Arg2ReturnValue9Word(arg W9) (W9, W9) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r W9) ValueReciverNoinlineGlobal2Arg2ReturnValue9Word(arg1, arg2 W9) (W9, W9) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W9) ValueReciverNoinlineGlobal3Arg2ReturnValue9Word(arg1, arg2, arg3 W9) (W9, W9) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W10) ValueReciverNoinlineGlobal1Arg0ReturnValue10Word(arg W10) {
	_ = arg
}

//go:noinline
func (r W10) ValueReciverNoinlineGlobal2Arg0ReturnValue10Word(arg1, arg2 W10) {
	_, _ = arg1, arg2
}

//go:noinline
func (r W10) ValueReciverNoinlineGlobal3Arg0ReturnValue10Word(arg1, arg2, arg3 W10) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W10) ValueReciverNoinlineGlobal1Arg1ReturnValue10Word(arg W10) W10 {
	_ = arg
	return arg
}

//go:noinline
func (r W10) ValueReciverNoinlineGlobal2Arg1ReturnValue10Word(arg1, arg2 W10) W10 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W10) ValueReciverNoinlineGlobal3Arg1ReturnValue10Word(arg1, arg2, arg3 W10) W10 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W10) ValueReciverNoinlineGlobal1Arg2ReturnValue10Word(arg W10) (W10, W10) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r W10) ValueReciverNoinlineGlobal2Arg2ReturnValue10Word(arg1, arg2 W10) (W10, W10) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W10) ValueReciverNoinlineGlobal3Arg2ReturnValue10Word(arg1, arg2, arg3 W10) (W10, W10) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W1) ValueReciverNoinlineGlobal1Arg0ReturnPointer1Word(arg *W1) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r W1) ValueReciverNoinlineGlobal2Arg0ReturnPointer1Word(arg1, arg2 *W1) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r W1) ValueReciverNoinlineGlobal3Arg0ReturnPointer1Word(arg1, arg2, arg3 *W1) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W1) ValueReciverNoinlineGlobal1Arg1ReturnPointer1Word(arg *W1) *W1 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r W1) ValueReciverNoinlineGlobal2Arg1ReturnPointer1Word(arg1, arg2 *W1) *W1 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W1) ValueReciverNoinlineGlobal3Arg1ReturnPointer1Word(arg1, arg2, arg3 *W1) *W1 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W1) ValueReciverNoinlineGlobal1Arg2ReturnPointer1Word(arg *W1) (*W1, *W1) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r W1) ValueReciverNoinlineGlobal2Arg2ReturnPointer1Word(arg1, arg2 *W1) (*W1, *W1) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W1) ValueReciverNoinlineGlobal3Arg2ReturnPointer1Word(arg1, arg2, arg3 *W1) (*W1, *W1) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W2) ValueReciverNoinlineGlobal1Arg0ReturnPointer2Word(arg *W2) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r W2) ValueReciverNoinlineGlobal2Arg0ReturnPointer2Word(arg1, arg2 *W2) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r W2) ValueReciverNoinlineGlobal3Arg0ReturnPointer2Word(arg1, arg2, arg3 *W2) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W2) ValueReciverNoinlineGlobal1Arg1ReturnPointer2Word(arg *W2) *W2 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r W2) ValueReciverNoinlineGlobal2Arg1ReturnPointer2Word(arg1, arg2 *W2) *W2 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W2) ValueReciverNoinlineGlobal3Arg1ReturnPointer2Word(arg1, arg2, arg3 *W2) *W2 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W2) ValueReciverNoinlineGlobal1Arg2ReturnPointer2Word(arg *W2) (*W2, *W2) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r W2) ValueReciverNoinlineGlobal2Arg2ReturnPointer2Word(arg1, arg2 *W2) (*W2, *W2) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W2) ValueReciverNoinlineGlobal3Arg2ReturnPointer2Word(arg1, arg2, arg3 *W2) (*W2, *W2) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W3) ValueReciverNoinlineGlobal1Arg0ReturnPointer3Word(arg *W3) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r W3) ValueReciverNoinlineGlobal2Arg0ReturnPointer3Word(arg1, arg2 *W3) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r W3) ValueReciverNoinlineGlobal3Arg0ReturnPointer3Word(arg1, arg2, arg3 *W3) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W3) ValueReciverNoinlineGlobal1Arg1ReturnPointer3Word(arg *W3) *W3 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r W3) ValueReciverNoinlineGlobal2Arg1ReturnPointer3Word(arg1, arg2 *W3) *W3 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W3) ValueReciverNoinlineGlobal3Arg1ReturnPointer3Word(arg1, arg2, arg3 *W3) *W3 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W3) ValueReciverNoinlineGlobal1Arg2ReturnPointer3Word(arg *W3) (*W3, *W3) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r W3) ValueReciverNoinlineGlobal2Arg2ReturnPointer3Word(arg1, arg2 *W3) (*W3, *W3) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W3) ValueReciverNoinlineGlobal3Arg2ReturnPointer3Word(arg1, arg2, arg3 *W3) (*W3, *W3) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W4) ValueReciverNoinlineGlobal1Arg0ReturnPointer4Word(arg *W4) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r W4) ValueReciverNoinlineGlobal2Arg0ReturnPointer4Word(arg1, arg2 *W4) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r W4) ValueReciverNoinlineGlobal3Arg0ReturnPointer4Word(arg1, arg2, arg3 *W4) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W4) ValueReciverNoinlineGlobal1Arg1ReturnPointer4Word(arg *W4) *W4 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r W4) ValueReciverNoinlineGlobal2Arg1ReturnPointer4Word(arg1, arg2 *W4) *W4 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W4) ValueReciverNoinlineGlobal3Arg1ReturnPointer4Word(arg1, arg2, arg3 *W4) *W4 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W4) ValueReciverNoinlineGlobal1Arg2ReturnPointer4Word(arg *W4) (*W4, *W4) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r W4) ValueReciverNoinlineGlobal2Arg2ReturnPointer4Word(arg1, arg2 *W4) (*W4, *W4) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W4) ValueReciverNoinlineGlobal3Arg2ReturnPointer4Word(arg1, arg2, arg3 *W4) (*W4, *W4) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W5) ValueReciverNoinlineGlobal1Arg0ReturnPointer5Word(arg *W5) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r W5) ValueReciverNoinlineGlobal2Arg0ReturnPointer5Word(arg1, arg2 *W5) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r W5) ValueReciverNoinlineGlobal3Arg0ReturnPointer5Word(arg1, arg2, arg3 *W5) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W5) ValueReciverNoinlineGlobal1Arg1ReturnPointer5Word(arg *W5) *W5 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r W5) ValueReciverNoinlineGlobal2Arg1ReturnPointer5Word(arg1, arg2 *W5) *W5 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W5) ValueReciverNoinlineGlobal3Arg1ReturnPointer5Word(arg1, arg2, arg3 *W5) *W5 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W5) ValueReciverNoinlineGlobal1Arg2ReturnPointer5Word(arg *W5) (*W5, *W5) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r W5) ValueReciverNoinlineGlobal2Arg2ReturnPointer5Word(arg1, arg2 *W5) (*W5, *W5) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W5) ValueReciverNoinlineGlobal3Arg2ReturnPointer5Word(arg1, arg2, arg3 *W5) (*W5, *W5) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W6) ValueReciverNoinlineGlobal1Arg0ReturnPointer6Word(arg *W6) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r W6) ValueReciverNoinlineGlobal2Arg0ReturnPointer6Word(arg1, arg2 *W6) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r W6) ValueReciverNoinlineGlobal3Arg0ReturnPointer6Word(arg1, arg2, arg3 *W6) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W6) ValueReciverNoinlineGlobal1Arg1ReturnPointer6Word(arg *W6) *W6 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r W6) ValueReciverNoinlineGlobal2Arg1ReturnPointer6Word(arg1, arg2 *W6) *W6 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W6) ValueReciverNoinlineGlobal3Arg1ReturnPointer6Word(arg1, arg2, arg3 *W6) *W6 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W6) ValueReciverNoinlineGlobal1Arg2ReturnPointer6Word(arg *W6) (*W6, *W6) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r W6) ValueReciverNoinlineGlobal2Arg2ReturnPointer6Word(arg1, arg2 *W6) (*W6, *W6) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W6) ValueReciverNoinlineGlobal3Arg2ReturnPointer6Word(arg1, arg2, arg3 *W6) (*W6, *W6) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W7) ValueReciverNoinlineGlobal1Arg0ReturnPointer7Word(arg *W7) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r W7) ValueReciverNoinlineGlobal2Arg0ReturnPointer7Word(arg1, arg2 *W7) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r W7) ValueReciverNoinlineGlobal3Arg0ReturnPointer7Word(arg1, arg2, arg3 *W7) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W7) ValueReciverNoinlineGlobal1Arg1ReturnPointer7Word(arg *W7) *W7 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r W7) ValueReciverNoinlineGlobal2Arg1ReturnPointer7Word(arg1, arg2 *W7) *W7 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W7) ValueReciverNoinlineGlobal3Arg1ReturnPointer7Word(arg1, arg2, arg3 *W7) *W7 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W7) ValueReciverNoinlineGlobal1Arg2ReturnPointer7Word(arg *W7) (*W7, *W7) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r W7) ValueReciverNoinlineGlobal2Arg2ReturnPointer7Word(arg1, arg2 *W7) (*W7, *W7) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W7) ValueReciverNoinlineGlobal3Arg2ReturnPointer7Word(arg1, arg2, arg3 *W7) (*W7, *W7) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W8) ValueReciverNoinlineGlobal1Arg0ReturnPointer8Word(arg *W8) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r W8) ValueReciverNoinlineGlobal2Arg0ReturnPointer8Word(arg1, arg2 *W8) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r W8) ValueReciverNoinlineGlobal3Arg0ReturnPointer8Word(arg1, arg2, arg3 *W8) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W8) ValueReciverNoinlineGlobal1Arg1ReturnPointer8Word(arg *W8) *W8 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r W8) ValueReciverNoinlineGlobal2Arg1ReturnPointer8Word(arg1, arg2 *W8) *W8 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W8) ValueReciverNoinlineGlobal3Arg1ReturnPointer8Word(arg1, arg2, arg3 *W8) *W8 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W8) ValueReciverNoinlineGlobal1Arg2ReturnPointer8Word(arg *W8) (*W8, *W8) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r W8) ValueReciverNoinlineGlobal2Arg2ReturnPointer8Word(arg1, arg2 *W8) (*W8, *W8) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W8) ValueReciverNoinlineGlobal3Arg2ReturnPointer8Word(arg1, arg2, arg3 *W8) (*W8, *W8) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W9) ValueReciverNoinlineGlobal1Arg0ReturnPointer9Word(arg *W9) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r W9) ValueReciverNoinlineGlobal2Arg0ReturnPointer9Word(arg1, arg2 *W9) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r W9) ValueReciverNoinlineGlobal3Arg0ReturnPointer9Word(arg1, arg2, arg3 *W9) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W9) ValueReciverNoinlineGlobal1Arg1ReturnPointer9Word(arg *W9) *W9 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r W9) ValueReciverNoinlineGlobal2Arg1ReturnPointer9Word(arg1, arg2 *W9) *W9 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W9) ValueReciverNoinlineGlobal3Arg1ReturnPointer9Word(arg1, arg2, arg3 *W9) *W9 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W9) ValueReciverNoinlineGlobal1Arg2ReturnPointer9Word(arg *W9) (*W9, *W9) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r W9) ValueReciverNoinlineGlobal2Arg2ReturnPointer9Word(arg1, arg2 *W9) (*W9, *W9) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W9) ValueReciverNoinlineGlobal3Arg2ReturnPointer9Word(arg1, arg2, arg3 *W9) (*W9, *W9) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W10) ValueReciverNoinlineGlobal1Arg0ReturnPointer10Word(arg *W10) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r W10) ValueReciverNoinlineGlobal2Arg0ReturnPointer10Word(arg1, arg2 *W10) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r W10) ValueReciverNoinlineGlobal3Arg0ReturnPointer10Word(arg1, arg2, arg3 *W10) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W10) ValueReciverNoinlineGlobal1Arg1ReturnPointer10Word(arg *W10) *W10 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r W10) ValueReciverNoinlineGlobal2Arg1ReturnPointer10Word(arg1, arg2 *W10) *W10 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W10) ValueReciverNoinlineGlobal3Arg1ReturnPointer10Word(arg1, arg2, arg3 *W10) *W10 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W10) ValueReciverNoinlineGlobal1Arg2ReturnPointer10Word(arg *W10) (*W10, *W10) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r W10) ValueReciverNoinlineGlobal2Arg2ReturnPointer10Word(arg1, arg2 *W10) (*W10, *W10) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W10) ValueReciverNoinlineGlobal3Arg2ReturnPointer10Word(arg1, arg2, arg3 *W10) (*W10, *W10) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W1) PointerReciverNoinlineGlobal1Arg0ReturnValue1Word(arg W1) {
	_ = arg
}

//go:noinline
func (r *W1) PointerReciverNoinlineGlobal2Arg0ReturnValue1Word(arg1, arg2 W1) {
	_, _ = arg1, arg2
}

//go:noinline
func (r *W1) PointerReciverNoinlineGlobal3Arg0ReturnValue1Word(arg1, arg2, arg3 W1) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W1) PointerReciverNoinlineGlobal1Arg1ReturnValue1Word(arg W1) W1 {
	_ = arg
	return arg
}

//go:noinline
func (r *W1) PointerReciverNoinlineGlobal2Arg1ReturnValue1Word(arg1, arg2 W1) W1 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W1) PointerReciverNoinlineGlobal3Arg1ReturnValue1Word(arg1, arg2, arg3 W1) W1 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W1) PointerReciverNoinlineGlobal1Arg2ReturnValue1Word(arg W1) (W1, W1) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W1) PointerReciverNoinlineGlobal2Arg2ReturnValue1Word(arg1, arg2 W1) (W1, W1) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W1) PointerReciverNoinlineGlobal3Arg2ReturnValue1Word(arg1, arg2, arg3 W1) (W1, W1) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W2) PointerReciverNoinlineGlobal1Arg0ReturnValue2Word(arg W2) {
	_ = arg
}

//go:noinline
func (r *W2) PointerReciverNoinlineGlobal2Arg0ReturnValue2Word(arg1, arg2 W2) {
	_, _ = arg1, arg2
}

//go:noinline
func (r *W2) PointerReciverNoinlineGlobal3Arg0ReturnValue2Word(arg1, arg2, arg3 W2) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W2) PointerReciverNoinlineGlobal1Arg1ReturnValue2Word(arg W2) W2 {
	_ = arg
	return arg
}

//go:noinline
func (r *W2) PointerReciverNoinlineGlobal2Arg1ReturnValue2Word(arg1, arg2 W2) W2 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W2) PointerReciverNoinlineGlobal3Arg1ReturnValue2Word(arg1, arg2, arg3 W2) W2 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W2) PointerReciverNoinlineGlobal1Arg2ReturnValue2Word(arg W2) (W2, W2) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W2) PointerReciverNoinlineGlobal2Arg2ReturnValue2Word(arg1, arg2 W2) (W2, W2) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W2) PointerReciverNoinlineGlobal3Arg2ReturnValue2Word(arg1, arg2, arg3 W2) (W2, W2) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W3) PointerReciverNoinlineGlobal1Arg0ReturnValue3Word(arg W3) {
	_ = arg
}

//go:noinline
func (r *W3) PointerReciverNoinlineGlobal2Arg0ReturnValue3Word(arg1, arg2 W3) {
	_, _ = arg1, arg2
}

//go:noinline
func (r *W3) PointerReciverNoinlineGlobal3Arg0ReturnValue3Word(arg1, arg2, arg3 W3) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W3) PointerReciverNoinlineGlobal1Arg1ReturnValue3Word(arg W3) W3 {
	_ = arg
	return arg
}

//go:noinline
func (r *W3) PointerReciverNoinlineGlobal2Arg1ReturnValue3Word(arg1, arg2 W3) W3 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W3) PointerReciverNoinlineGlobal3Arg1ReturnValue3Word(arg1, arg2, arg3 W3) W3 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W3) PointerReciverNoinlineGlobal1Arg2ReturnValue3Word(arg W3) (W3, W3) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W3) PointerReciverNoinlineGlobal2Arg2ReturnValue3Word(arg1, arg2 W3) (W3, W3) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W3) PointerReciverNoinlineGlobal3Arg2ReturnValue3Word(arg1, arg2, arg3 W3) (W3, W3) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W4) PointerReciverNoinlineGlobal1Arg0ReturnValue4Word(arg W4) {
	_ = arg
}

//go:noinline
func (r *W4) PointerReciverNoinlineGlobal2Arg0ReturnValue4Word(arg1, arg2 W4) {
	_, _ = arg1, arg2
}

//go:noinline
func (r *W4) PointerReciverNoinlineGlobal3Arg0ReturnValue4Word(arg1, arg2, arg3 W4) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W4) PointerReciverNoinlineGlobal1Arg1ReturnValue4Word(arg W4) W4 {
	_ = arg
	return arg
}

//go:noinline
func (r *W4) PointerReciverNoinlineGlobal2Arg1ReturnValue4Word(arg1, arg2 W4) W4 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W4) PointerReciverNoinlineGlobal3Arg1ReturnValue4Word(arg1, arg2, arg3 W4) W4 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W4) PointerReciverNoinlineGlobal1Arg2ReturnValue4Word(arg W4) (W4, W4) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W4) PointerReciverNoinlineGlobal2Arg2ReturnValue4Word(arg1, arg2 W4) (W4, W4) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W4) PointerReciverNoinlineGlobal3Arg2ReturnValue4Word(arg1, arg2, arg3 W4) (W4, W4) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W5) PointerReciverNoinlineGlobal1Arg0ReturnValue5Word(arg W5) {
	_ = arg
}

//go:noinline
func (r *W5) PointerReciverNoinlineGlobal2Arg0ReturnValue5Word(arg1, arg2 W5) {
	_, _ = arg1, arg2
}

//go:noinline
func (r *W5) PointerReciverNoinlineGlobal3Arg0ReturnValue5Word(arg1, arg2, arg3 W5) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W5) PointerReciverNoinlineGlobal1Arg1ReturnValue5Word(arg W5) W5 {
	_ = arg
	return arg
}

//go:noinline
func (r *W5) PointerReciverNoinlineGlobal2Arg1ReturnValue5Word(arg1, arg2 W5) W5 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W5) PointerReciverNoinlineGlobal3Arg1ReturnValue5Word(arg1, arg2, arg3 W5) W5 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W5) PointerReciverNoinlineGlobal1Arg2ReturnValue5Word(arg W5) (W5, W5) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W5) PointerReciverNoinlineGlobal2Arg2ReturnValue5Word(arg1, arg2 W5) (W5, W5) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W5) PointerReciverNoinlineGlobal3Arg2ReturnValue5Word(arg1, arg2, arg3 W5) (W5, W5) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W6) PointerReciverNoinlineGlobal1Arg0ReturnValue6Word(arg W6) {
	_ = arg
}

//go:noinline
func (r *W6) PointerReciverNoinlineGlobal2Arg0ReturnValue6Word(arg1, arg2 W6) {
	_, _ = arg1, arg2
}

//go:noinline
func (r *W6) PointerReciverNoinlineGlobal3Arg0ReturnValue6Word(arg1, arg2, arg3 W6) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W6) PointerReciverNoinlineGlobal1Arg1ReturnValue6Word(arg W6) W6 {
	_ = arg
	return arg
}

//go:noinline
func (r *W6) PointerReciverNoinlineGlobal2Arg1ReturnValue6Word(arg1, arg2 W6) W6 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W6) PointerReciverNoinlineGlobal3Arg1ReturnValue6Word(arg1, arg2, arg3 W6) W6 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W6) PointerReciverNoinlineGlobal1Arg2ReturnValue6Word(arg W6) (W6, W6) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W6) PointerReciverNoinlineGlobal2Arg2ReturnValue6Word(arg1, arg2 W6) (W6, W6) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W6) PointerReciverNoinlineGlobal3Arg2ReturnValue6Word(arg1, arg2, arg3 W6) (W6, W6) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W7) PointerReciverNoinlineGlobal1Arg0ReturnValue7Word(arg W7) {
	_ = arg
}

//go:noinline
func (r *W7) PointerReciverNoinlineGlobal2Arg0ReturnValue7Word(arg1, arg2 W7) {
	_, _ = arg1, arg2
}

//go:noinline
func (r *W7) PointerReciverNoinlineGlobal3Arg0ReturnValue7Word(arg1, arg2, arg3 W7) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W7) PointerReciverNoinlineGlobal1Arg1ReturnValue7Word(arg W7) W7 {
	_ = arg
	return arg
}

//go:noinline
func (r *W7) PointerReciverNoinlineGlobal2Arg1ReturnValue7Word(arg1, arg2 W7) W7 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W7) PointerReciverNoinlineGlobal3Arg1ReturnValue7Word(arg1, arg2, arg3 W7) W7 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W7) PointerReciverNoinlineGlobal1Arg2ReturnValue7Word(arg W7) (W7, W7) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W7) PointerReciverNoinlineGlobal2Arg2ReturnValue7Word(arg1, arg2 W7) (W7, W7) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W7) PointerReciverNoinlineGlobal3Arg2ReturnValue7Word(arg1, arg2, arg3 W7) (W7, W7) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W8) PointerReciverNoinlineGlobal1Arg0ReturnValue8Word(arg W8) {
	_ = arg
}

//go:noinline
func (r *W8) PointerReciverNoinlineGlobal2Arg0ReturnValue8Word(arg1, arg2 W8) {
	_, _ = arg1, arg2
}

//go:noinline
func (r *W8) PointerReciverNoinlineGlobal3Arg0ReturnValue8Word(arg1, arg2, arg3 W8) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W8) PointerReciverNoinlineGlobal1Arg1ReturnValue8Word(arg W8) W8 {
	_ = arg
	return arg
}

//go:noinline
func (r *W8) PointerReciverNoinlineGlobal2Arg1ReturnValue8Word(arg1, arg2 W8) W8 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W8) PointerReciverNoinlineGlobal3Arg1ReturnValue8Word(arg1, arg2, arg3 W8) W8 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W8) PointerReciverNoinlineGlobal1Arg2ReturnValue8Word(arg W8) (W8, W8) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W8) PointerReciverNoinlineGlobal2Arg2ReturnValue8Word(arg1, arg2 W8) (W8, W8) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W8) PointerReciverNoinlineGlobal3Arg2ReturnValue8Word(arg1, arg2, arg3 W8) (W8, W8) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W9) PointerReciverNoinlineGlobal1Arg0ReturnValue9Word(arg W9) {
	_ = arg
}

//go:noinline
func (r *W9) PointerReciverNoinlineGlobal2Arg0ReturnValue9Word(arg1, arg2 W9) {
	_, _ = arg1, arg2
}

//go:noinline
func (r *W9) PointerReciverNoinlineGlobal3Arg0ReturnValue9Word(arg1, arg2, arg3 W9) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W9) PointerReciverNoinlineGlobal1Arg1ReturnValue9Word(arg W9) W9 {
	_ = arg
	return arg
}

//go:noinline
func (r *W9) PointerReciverNoinlineGlobal2Arg1ReturnValue9Word(arg1, arg2 W9) W9 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W9) PointerReciverNoinlineGlobal3Arg1ReturnValue9Word(arg1, arg2, arg3 W9) W9 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W9) PointerReciverNoinlineGlobal1Arg2ReturnValue9Word(arg W9) (W9, W9) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W9) PointerReciverNoinlineGlobal2Arg2ReturnValue9Word(arg1, arg2 W9) (W9, W9) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W9) PointerReciverNoinlineGlobal3Arg2ReturnValue9Word(arg1, arg2, arg3 W9) (W9, W9) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W10) PointerReciverNoinlineGlobal1Arg0ReturnValue10Word(arg W10) {
	_ = arg
}

//go:noinline
func (r *W10) PointerReciverNoinlineGlobal2Arg0ReturnValue10Word(arg1, arg2 W10) {
	_, _ = arg1, arg2
}

//go:noinline
func (r *W10) PointerReciverNoinlineGlobal3Arg0ReturnValue10Word(arg1, arg2, arg3 W10) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W10) PointerReciverNoinlineGlobal1Arg1ReturnValue10Word(arg W10) W10 {
	_ = arg
	return arg
}

//go:noinline
func (r *W10) PointerReciverNoinlineGlobal2Arg1ReturnValue10Word(arg1, arg2 W10) W10 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W10) PointerReciverNoinlineGlobal3Arg1ReturnValue10Word(arg1, arg2, arg3 W10) W10 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W10) PointerReciverNoinlineGlobal1Arg2ReturnValue10Word(arg W10) (W10, W10) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W10) PointerReciverNoinlineGlobal2Arg2ReturnValue10Word(arg1, arg2 W10) (W10, W10) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W10) PointerReciverNoinlineGlobal3Arg2ReturnValue10Word(arg1, arg2, arg3 W10) (W10, W10) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W1) PointerReciverNoinlineGlobal1Arg0ReturnPointer1Word(arg *W1) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r *W1) PointerReciverNoinlineGlobal2Arg0ReturnPointer1Word(arg1, arg2 *W1) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r *W1) PointerReciverNoinlineGlobal3Arg0ReturnPointer1Word(arg1, arg2, arg3 *W1) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W1) PointerReciverNoinlineGlobal1Arg1ReturnPointer1Word(arg *W1) *W1 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r *W1) PointerReciverNoinlineGlobal2Arg1ReturnPointer1Word(arg1, arg2 *W1) *W1 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W1) PointerReciverNoinlineGlobal3Arg1ReturnPointer1Word(arg1, arg2, arg3 *W1) *W1 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W1) PointerReciverNoinlineGlobal1Arg2ReturnPointer1Word(arg *W1) (*W1, *W1) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W1) PointerReciverNoinlineGlobal2Arg2ReturnPointer1Word(arg1, arg2 *W1) (*W1, *W1) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W1) PointerReciverNoinlineGlobal3Arg2ReturnPointer1Word(arg1, arg2, arg3 *W1) (*W1, *W1) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W2) PointerReciverNoinlineGlobal1Arg0ReturnPointer2Word(arg *W2) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r *W2) PointerReciverNoinlineGlobal2Arg0ReturnPointer2Word(arg1, arg2 *W2) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r *W2) PointerReciverNoinlineGlobal3Arg0ReturnPointer2Word(arg1, arg2, arg3 *W2) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W2) PointerReciverNoinlineGlobal1Arg1ReturnPointer2Word(arg *W2) *W2 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r *W2) PointerReciverNoinlineGlobal2Arg1ReturnPointer2Word(arg1, arg2 *W2) *W2 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W2) PointerReciverNoinlineGlobal3Arg1ReturnPointer2Word(arg1, arg2, arg3 *W2) *W2 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W2) PointerReciverNoinlineGlobal1Arg2ReturnPointer2Word(arg *W2) (*W2, *W2) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W2) PointerReciverNoinlineGlobal2Arg2ReturnPointer2Word(arg1, arg2 *W2) (*W2, *W2) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W2) PointerReciverNoinlineGlobal3Arg2ReturnPointer2Word(arg1, arg2, arg3 *W2) (*W2, *W2) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W3) PointerReciverNoinlineGlobal1Arg0ReturnPointer3Word(arg *W3) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r *W3) PointerReciverNoinlineGlobal2Arg0ReturnPointer3Word(arg1, arg2 *W3) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r *W3) PointerReciverNoinlineGlobal3Arg0ReturnPointer3Word(arg1, arg2, arg3 *W3) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W3) PointerReciverNoinlineGlobal1Arg1ReturnPointer3Word(arg *W3) *W3 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r *W3) PointerReciverNoinlineGlobal2Arg1ReturnPointer3Word(arg1, arg2 *W3) *W3 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W3) PointerReciverNoinlineGlobal3Arg1ReturnPointer3Word(arg1, arg2, arg3 *W3) *W3 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W3) PointerReciverNoinlineGlobal1Arg2ReturnPointer3Word(arg *W3) (*W3, *W3) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W3) PointerReciverNoinlineGlobal2Arg2ReturnPointer3Word(arg1, arg2 *W3) (*W3, *W3) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W3) PointerReciverNoinlineGlobal3Arg2ReturnPointer3Word(arg1, arg2, arg3 *W3) (*W3, *W3) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W4) PointerReciverNoinlineGlobal1Arg0ReturnPointer4Word(arg *W4) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r *W4) PointerReciverNoinlineGlobal2Arg0ReturnPointer4Word(arg1, arg2 *W4) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r *W4) PointerReciverNoinlineGlobal3Arg0ReturnPointer4Word(arg1, arg2, arg3 *W4) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W4) PointerReciverNoinlineGlobal1Arg1ReturnPointer4Word(arg *W4) *W4 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r *W4) PointerReciverNoinlineGlobal2Arg1ReturnPointer4Word(arg1, arg2 *W4) *W4 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W4) PointerReciverNoinlineGlobal3Arg1ReturnPointer4Word(arg1, arg2, arg3 *W4) *W4 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W4) PointerReciverNoinlineGlobal1Arg2ReturnPointer4Word(arg *W4) (*W4, *W4) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W4) PointerReciverNoinlineGlobal2Arg2ReturnPointer4Word(arg1, arg2 *W4) (*W4, *W4) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W4) PointerReciverNoinlineGlobal3Arg2ReturnPointer4Word(arg1, arg2, arg3 *W4) (*W4, *W4) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W5) PointerReciverNoinlineGlobal1Arg0ReturnPointer5Word(arg *W5) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r *W5) PointerReciverNoinlineGlobal2Arg0ReturnPointer5Word(arg1, arg2 *W5) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r *W5) PointerReciverNoinlineGlobal3Arg0ReturnPointer5Word(arg1, arg2, arg3 *W5) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W5) PointerReciverNoinlineGlobal1Arg1ReturnPointer5Word(arg *W5) *W5 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r *W5) PointerReciverNoinlineGlobal2Arg1ReturnPointer5Word(arg1, arg2 *W5) *W5 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W5) PointerReciverNoinlineGlobal3Arg1ReturnPointer5Word(arg1, arg2, arg3 *W5) *W5 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W5) PointerReciverNoinlineGlobal1Arg2ReturnPointer5Word(arg *W5) (*W5, *W5) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W5) PointerReciverNoinlineGlobal2Arg2ReturnPointer5Word(arg1, arg2 *W5) (*W5, *W5) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W5) PointerReciverNoinlineGlobal3Arg2ReturnPointer5Word(arg1, arg2, arg3 *W5) (*W5, *W5) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W6) PointerReciverNoinlineGlobal1Arg0ReturnPointer6Word(arg *W6) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r *W6) PointerReciverNoinlineGlobal2Arg0ReturnPointer6Word(arg1, arg2 *W6) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r *W6) PointerReciverNoinlineGlobal3Arg0ReturnPointer6Word(arg1, arg2, arg3 *W6) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W6) PointerReciverNoinlineGlobal1Arg1ReturnPointer6Word(arg *W6) *W6 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r *W6) PointerReciverNoinlineGlobal2Arg1ReturnPointer6Word(arg1, arg2 *W6) *W6 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W6) PointerReciverNoinlineGlobal3Arg1ReturnPointer6Word(arg1, arg2, arg3 *W6) *W6 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W6) PointerReciverNoinlineGlobal1Arg2ReturnPointer6Word(arg *W6) (*W6, *W6) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W6) PointerReciverNoinlineGlobal2Arg2ReturnPointer6Word(arg1, arg2 *W6) (*W6, *W6) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W6) PointerReciverNoinlineGlobal3Arg2ReturnPointer6Word(arg1, arg2, arg3 *W6) (*W6, *W6) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W7) PointerReciverNoinlineGlobal1Arg0ReturnPointer7Word(arg *W7) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r *W7) PointerReciverNoinlineGlobal2Arg0ReturnPointer7Word(arg1, arg2 *W7) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r *W7) PointerReciverNoinlineGlobal3Arg0ReturnPointer7Word(arg1, arg2, arg3 *W7) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W7) PointerReciverNoinlineGlobal1Arg1ReturnPointer7Word(arg *W7) *W7 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r *W7) PointerReciverNoinlineGlobal2Arg1ReturnPointer7Word(arg1, arg2 *W7) *W7 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W7) PointerReciverNoinlineGlobal3Arg1ReturnPointer7Word(arg1, arg2, arg3 *W7) *W7 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W7) PointerReciverNoinlineGlobal1Arg2ReturnPointer7Word(arg *W7) (*W7, *W7) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W7) PointerReciverNoinlineGlobal2Arg2ReturnPointer7Word(arg1, arg2 *W7) (*W7, *W7) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W7) PointerReciverNoinlineGlobal3Arg2ReturnPointer7Word(arg1, arg2, arg3 *W7) (*W7, *W7) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W8) PointerReciverNoinlineGlobal1Arg0ReturnPointer8Word(arg *W8) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r *W8) PointerReciverNoinlineGlobal2Arg0ReturnPointer8Word(arg1, arg2 *W8) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r *W8) PointerReciverNoinlineGlobal3Arg0ReturnPointer8Word(arg1, arg2, arg3 *W8) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W8) PointerReciverNoinlineGlobal1Arg1ReturnPointer8Word(arg *W8) *W8 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r *W8) PointerReciverNoinlineGlobal2Arg1ReturnPointer8Word(arg1, arg2 *W8) *W8 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W8) PointerReciverNoinlineGlobal3Arg1ReturnPointer8Word(arg1, arg2, arg3 *W8) *W8 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W8) PointerReciverNoinlineGlobal1Arg2ReturnPointer8Word(arg *W8) (*W8, *W8) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W8) PointerReciverNoinlineGlobal2Arg2ReturnPointer8Word(arg1, arg2 *W8) (*W8, *W8) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W8) PointerReciverNoinlineGlobal3Arg2ReturnPointer8Word(arg1, arg2, arg3 *W8) (*W8, *W8) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W9) PointerReciverNoinlineGlobal1Arg0ReturnPointer9Word(arg *W9) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r *W9) PointerReciverNoinlineGlobal2Arg0ReturnPointer9Word(arg1, arg2 *W9) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r *W9) PointerReciverNoinlineGlobal3Arg0ReturnPointer9Word(arg1, arg2, arg3 *W9) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W9) PointerReciverNoinlineGlobal1Arg1ReturnPointer9Word(arg *W9) *W9 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r *W9) PointerReciverNoinlineGlobal2Arg1ReturnPointer9Word(arg1, arg2 *W9) *W9 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W9) PointerReciverNoinlineGlobal3Arg1ReturnPointer9Word(arg1, arg2, arg3 *W9) *W9 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W9) PointerReciverNoinlineGlobal1Arg2ReturnPointer9Word(arg *W9) (*W9, *W9) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W9) PointerReciverNoinlineGlobal2Arg2ReturnPointer9Word(arg1, arg2 *W9) (*W9, *W9) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W9) PointerReciverNoinlineGlobal3Arg2ReturnPointer9Word(arg1, arg2, arg3 *W9) (*W9, *W9) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W10) PointerReciverNoinlineGlobal1Arg0ReturnPointer10Word(arg *W10) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r *W10) PointerReciverNoinlineGlobal2Arg0ReturnPointer10Word(arg1, arg2 *W10) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r *W10) PointerReciverNoinlineGlobal3Arg0ReturnPointer10Word(arg1, arg2, arg3 *W10) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	if arg3 == nil {
		return
	}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W10) PointerReciverNoinlineGlobal1Arg1ReturnPointer10Word(arg *W10) *W10 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r *W10) PointerReciverNoinlineGlobal2Arg1ReturnPointer10Word(arg1, arg2 *W10) *W10 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W10) PointerReciverNoinlineGlobal3Arg1ReturnPointer10Word(arg1, arg2, arg3 *W10) *W10 {
	if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W10) PointerReciverNoinlineGlobal1Arg2ReturnPointer10Word(arg *W10) (*W10, *W10) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W10) PointerReciverNoinlineGlobal2Arg2ReturnPointer10Word(arg1, arg2 *W10) (*W10, *W10) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W10) PointerReciverNoinlineGlobal3Arg2ReturnPointer10Word(arg1, arg2, arg3 *W10) (*W10, *W10) {
	if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

func BenchmarkFuncNoinlineGlobal1Arg0ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal1Arg0ReturnValue1Word(arg1vW1)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg0ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal2Arg0ReturnValue1Word(arg1vW1, arg2vW1)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg0ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal3Arg0ReturnValue1Word(arg1vW1, arg2vW1, arg3vW1)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg1ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal1Arg1ReturnValue1Word(arg1vW1)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg1ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal2Arg1ReturnValue1Word(arg1vW1, arg2vW1)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg1ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal3Arg1ReturnValue1Word(arg1vW1, arg2vW1, arg3vW1)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg2ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal1Arg2ReturnValue1Word(arg1vW1)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg2ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal2Arg2ReturnValue1Word(arg1vW1, arg2vW1)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg2ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal3Arg2ReturnValue1Word(arg1vW1, arg2vW1, arg3vW1)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinlineGlobal1Arg0ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal1Arg0ReturnValue2Word(arg1vW2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg0ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal2Arg0ReturnValue2Word(arg1vW2, arg2vW2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg0ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal3Arg0ReturnValue2Word(arg1vW2, arg2vW2, arg3vW2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg1ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal1Arg1ReturnValue2Word(arg1vW2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg1ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal2Arg1ReturnValue2Word(arg1vW2, arg2vW2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg1ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal3Arg1ReturnValue2Word(arg1vW2, arg2vW2, arg3vW2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg2ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal1Arg2ReturnValue2Word(arg1vW2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg2ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal2Arg2ReturnValue2Word(arg1vW2, arg2vW2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg2ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal3Arg2ReturnValue2Word(arg1vW2, arg2vW2, arg3vW2)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinlineGlobal1Arg0ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal1Arg0ReturnValue3Word(arg1vW3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg0ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal2Arg0ReturnValue3Word(arg1vW3, arg2vW3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg0ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal3Arg0ReturnValue3Word(arg1vW3, arg2vW3, arg3vW3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg1ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal1Arg1ReturnValue3Word(arg1vW3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg1ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal2Arg1ReturnValue3Word(arg1vW3, arg2vW3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg1ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal3Arg1ReturnValue3Word(arg1vW3, arg2vW3, arg3vW3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg2ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal1Arg2ReturnValue3Word(arg1vW3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg2ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal2Arg2ReturnValue3Word(arg1vW3, arg2vW3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg2ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal3Arg2ReturnValue3Word(arg1vW3, arg2vW3, arg3vW3)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinlineGlobal1Arg0ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal1Arg0ReturnValue4Word(arg1vW4)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg0ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal2Arg0ReturnValue4Word(arg1vW4, arg2vW4)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg0ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal3Arg0ReturnValue4Word(arg1vW4, arg2vW4, arg3vW4)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg1ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal1Arg1ReturnValue4Word(arg1vW4)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg1ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal2Arg1ReturnValue4Word(arg1vW4, arg2vW4)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg1ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal3Arg1ReturnValue4Word(arg1vW4, arg2vW4, arg3vW4)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg2ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal1Arg2ReturnValue4Word(arg1vW4)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg2ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal2Arg2ReturnValue4Word(arg1vW4, arg2vW4)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg2ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal3Arg2ReturnValue4Word(arg1vW4, arg2vW4, arg3vW4)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinlineGlobal1Arg0ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal1Arg0ReturnValue5Word(arg1vW5)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg0ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal2Arg0ReturnValue5Word(arg1vW5, arg2vW5)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg0ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal3Arg0ReturnValue5Word(arg1vW5, arg2vW5, arg3vW5)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg1ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal1Arg1ReturnValue5Word(arg1vW5)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg1ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal2Arg1ReturnValue5Word(arg1vW5, arg2vW5)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg1ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal3Arg1ReturnValue5Word(arg1vW5, arg2vW5, arg3vW5)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg2ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal1Arg2ReturnValue5Word(arg1vW5)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg2ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal2Arg2ReturnValue5Word(arg1vW5, arg2vW5)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg2ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal3Arg2ReturnValue5Word(arg1vW5, arg2vW5, arg3vW5)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinlineGlobal1Arg0ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal1Arg0ReturnValue6Word(arg1vW6)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg0ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal2Arg0ReturnValue6Word(arg1vW6, arg2vW6)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg0ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal3Arg0ReturnValue6Word(arg1vW6, arg2vW6, arg3vW6)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg1ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal1Arg1ReturnValue6Word(arg1vW6)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg1ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal2Arg1ReturnValue6Word(arg1vW6, arg2vW6)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg1ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal3Arg1ReturnValue6Word(arg1vW6, arg2vW6, arg3vW6)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg2ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal1Arg2ReturnValue6Word(arg1vW6)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg2ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal2Arg2ReturnValue6Word(arg1vW6, arg2vW6)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg2ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal3Arg2ReturnValue6Word(arg1vW6, arg2vW6, arg3vW6)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinlineGlobal1Arg0ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal1Arg0ReturnValue7Word(arg1vW7)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg0ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal2Arg0ReturnValue7Word(arg1vW7, arg2vW7)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg0ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal3Arg0ReturnValue7Word(arg1vW7, arg2vW7, arg3vW7)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg1ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal1Arg1ReturnValue7Word(arg1vW7)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg1ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal2Arg1ReturnValue7Word(arg1vW7, arg2vW7)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg1ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal3Arg1ReturnValue7Word(arg1vW7, arg2vW7, arg3vW7)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg2ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal1Arg2ReturnValue7Word(arg1vW7)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg2ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal2Arg2ReturnValue7Word(arg1vW7, arg2vW7)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg2ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal3Arg2ReturnValue7Word(arg1vW7, arg2vW7, arg3vW7)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinlineGlobal1Arg0ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal1Arg0ReturnValue8Word(arg1vW8)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg0ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal2Arg0ReturnValue8Word(arg1vW8, arg2vW8)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg0ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal3Arg0ReturnValue8Word(arg1vW8, arg2vW8, arg3vW8)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg1ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal1Arg1ReturnValue8Word(arg1vW8)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg1ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal2Arg1ReturnValue8Word(arg1vW8, arg2vW8)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg1ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal3Arg1ReturnValue8Word(arg1vW8, arg2vW8, arg3vW8)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg2ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal1Arg2ReturnValue8Word(arg1vW8)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg2ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal2Arg2ReturnValue8Word(arg1vW8, arg2vW8)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg2ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal3Arg2ReturnValue8Word(arg1vW8, arg2vW8, arg3vW8)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinlineGlobal1Arg0ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal1Arg0ReturnValue9Word(arg1vW9)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg0ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal2Arg0ReturnValue9Word(arg1vW9, arg2vW9)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg0ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal3Arg0ReturnValue9Word(arg1vW9, arg2vW9, arg3vW9)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg1ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal1Arg1ReturnValue9Word(arg1vW9)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg1ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal2Arg1ReturnValue9Word(arg1vW9, arg2vW9)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg1ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal3Arg1ReturnValue9Word(arg1vW9, arg2vW9, arg3vW9)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg2ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal1Arg2ReturnValue9Word(arg1vW9)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg2ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal2Arg2ReturnValue9Word(arg1vW9, arg2vW9)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg2ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal3Arg2ReturnValue9Word(arg1vW9, arg2vW9, arg3vW9)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinlineGlobal1Arg0ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal1Arg0ReturnValue10Word(arg1vW10)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg0ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal2Arg0ReturnValue10Word(arg1vW10, arg2vW10)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg0ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal3Arg0ReturnValue10Word(arg1vW10, arg2vW10, arg3vW10)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg1ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal1Arg1ReturnValue10Word(arg1vW10)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg1ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal2Arg1ReturnValue10Word(arg1vW10, arg2vW10)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg1ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal3Arg1ReturnValue10Word(arg1vW10, arg2vW10, arg3vW10)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg2ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal1Arg2ReturnValue10Word(arg1vW10)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg2ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal2Arg2ReturnValue10Word(arg1vW10, arg2vW10)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg2ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal3Arg2ReturnValue10Word(arg1vW10, arg2vW10, arg3vW10)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg0ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal1Arg0ReturnPointer1Word(arg1pW1)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg0ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal2Arg0ReturnPointer1Word(arg1pW1, arg2pW1)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg0ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal3Arg0ReturnPointer1Word(arg1pW1, arg2pW1, arg3pW1)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg1ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal1Arg1ReturnPointer1Word(arg1pW1)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg1ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal2Arg1ReturnPointer1Word(arg1pW1, arg2pW1)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg1ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal3Arg1ReturnPointer1Word(arg1pW1, arg2pW1, arg3pW1)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg2ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal1Arg2ReturnPointer1Word(arg1pW1)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg2ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal2Arg2ReturnPointer1Word(arg1pW1, arg2pW1)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg2ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal3Arg2ReturnPointer1Word(arg1pW1, arg2pW1, arg3pW1)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinlineGlobal1Arg0ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal1Arg0ReturnPointer2Word(arg1pW2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg0ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal2Arg0ReturnPointer2Word(arg1pW2, arg2pW2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg0ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal3Arg0ReturnPointer2Word(arg1pW2, arg2pW2, arg3pW2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg1ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal1Arg1ReturnPointer2Word(arg1pW2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg1ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal2Arg1ReturnPointer2Word(arg1pW2, arg2pW2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg1ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal3Arg1ReturnPointer2Word(arg1pW2, arg2pW2, arg3pW2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg2ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal1Arg2ReturnPointer2Word(arg1pW2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg2ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal2Arg2ReturnPointer2Word(arg1pW2, arg2pW2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg2ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal3Arg2ReturnPointer2Word(arg1pW2, arg2pW2, arg3pW2)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinlineGlobal1Arg0ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal1Arg0ReturnPointer3Word(arg1pW3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg0ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal2Arg0ReturnPointer3Word(arg1pW3, arg2pW3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg0ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal3Arg0ReturnPointer3Word(arg1pW3, arg2pW3, arg3pW3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg1ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal1Arg1ReturnPointer3Word(arg1pW3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg1ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal2Arg1ReturnPointer3Word(arg1pW3, arg2pW3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg1ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal3Arg1ReturnPointer3Word(arg1pW3, arg2pW3, arg3pW3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg2ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal1Arg2ReturnPointer3Word(arg1pW3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg2ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal2Arg2ReturnPointer3Word(arg1pW3, arg2pW3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg2ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal3Arg2ReturnPointer3Word(arg1pW3, arg2pW3, arg3pW3)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinlineGlobal1Arg0ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal1Arg0ReturnPointer4Word(arg1pW4)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg0ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal2Arg0ReturnPointer4Word(arg1pW4, arg2pW4)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg0ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal3Arg0ReturnPointer4Word(arg1pW4, arg2pW4, arg3pW4)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg1ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal1Arg1ReturnPointer4Word(arg1pW4)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg1ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal2Arg1ReturnPointer4Word(arg1pW4, arg2pW4)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg1ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal3Arg1ReturnPointer4Word(arg1pW4, arg2pW4, arg3pW4)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg2ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal1Arg2ReturnPointer4Word(arg1pW4)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg2ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal2Arg2ReturnPointer4Word(arg1pW4, arg2pW4)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg2ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal3Arg2ReturnPointer4Word(arg1pW4, arg2pW4, arg3pW4)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinlineGlobal1Arg0ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal1Arg0ReturnPointer5Word(arg1pW5)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg0ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal2Arg0ReturnPointer5Word(arg1pW5, arg2pW5)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg0ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal3Arg0ReturnPointer5Word(arg1pW5, arg2pW5, arg3pW5)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg1ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal1Arg1ReturnPointer5Word(arg1pW5)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg1ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal2Arg1ReturnPointer5Word(arg1pW5, arg2pW5)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg1ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal3Arg1ReturnPointer5Word(arg1pW5, arg2pW5, arg3pW5)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg2ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal1Arg2ReturnPointer5Word(arg1pW5)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg2ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal2Arg2ReturnPointer5Word(arg1pW5, arg2pW5)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg2ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal3Arg2ReturnPointer5Word(arg1pW5, arg2pW5, arg3pW5)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinlineGlobal1Arg0ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal1Arg0ReturnPointer6Word(arg1pW6)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg0ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal2Arg0ReturnPointer6Word(arg1pW6, arg2pW6)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg0ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal3Arg0ReturnPointer6Word(arg1pW6, arg2pW6, arg3pW6)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg1ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal1Arg1ReturnPointer6Word(arg1pW6)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg1ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal2Arg1ReturnPointer6Word(arg1pW6, arg2pW6)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg1ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal3Arg1ReturnPointer6Word(arg1pW6, arg2pW6, arg3pW6)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg2ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal1Arg2ReturnPointer6Word(arg1pW6)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg2ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal2Arg2ReturnPointer6Word(arg1pW6, arg2pW6)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg2ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal3Arg2ReturnPointer6Word(arg1pW6, arg2pW6, arg3pW6)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinlineGlobal1Arg0ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal1Arg0ReturnPointer7Word(arg1pW7)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg0ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal2Arg0ReturnPointer7Word(arg1pW7, arg2pW7)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg0ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal3Arg0ReturnPointer7Word(arg1pW7, arg2pW7, arg3pW7)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg1ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal1Arg1ReturnPointer7Word(arg1pW7)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg1ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal2Arg1ReturnPointer7Word(arg1pW7, arg2pW7)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg1ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal3Arg1ReturnPointer7Word(arg1pW7, arg2pW7, arg3pW7)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg2ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal1Arg2ReturnPointer7Word(arg1pW7)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg2ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal2Arg2ReturnPointer7Word(arg1pW7, arg2pW7)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg2ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal3Arg2ReturnPointer7Word(arg1pW7, arg2pW7, arg3pW7)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinlineGlobal1Arg0ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal1Arg0ReturnPointer8Word(arg1pW8)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg0ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal2Arg0ReturnPointer8Word(arg1pW8, arg2pW8)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg0ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal3Arg0ReturnPointer8Word(arg1pW8, arg2pW8, arg3pW8)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg1ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal1Arg1ReturnPointer8Word(arg1pW8)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg1ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal2Arg1ReturnPointer8Word(arg1pW8, arg2pW8)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg1ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal3Arg1ReturnPointer8Word(arg1pW8, arg2pW8, arg3pW8)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg2ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal1Arg2ReturnPointer8Word(arg1pW8)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg2ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal2Arg2ReturnPointer8Word(arg1pW8, arg2pW8)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg2ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal3Arg2ReturnPointer8Word(arg1pW8, arg2pW8, arg3pW8)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinlineGlobal1Arg0ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal1Arg0ReturnPointer9Word(arg1pW9)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg0ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal2Arg0ReturnPointer9Word(arg1pW9, arg2pW9)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg0ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal3Arg0ReturnPointer9Word(arg1pW9, arg2pW9, arg3pW9)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg1ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal1Arg1ReturnPointer9Word(arg1pW9)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg1ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal2Arg1ReturnPointer9Word(arg1pW9, arg2pW9)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg1ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal3Arg1ReturnPointer9Word(arg1pW9, arg2pW9, arg3pW9)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg2ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal1Arg2ReturnPointer9Word(arg1pW9)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg2ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal2Arg2ReturnPointer9Word(arg1pW9, arg2pW9)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg2ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal3Arg2ReturnPointer9Word(arg1pW9, arg2pW9, arg3pW9)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinlineGlobal1Arg0ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal1Arg0ReturnPointer10Word(arg1pW10)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg0ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal2Arg0ReturnPointer10Word(arg1pW10, arg2pW10)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg0ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinlineGlobal3Arg0ReturnPointer10Word(arg1pW10, arg2pW10, arg3pW10)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg1ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal1Arg1ReturnPointer10Word(arg1pW10)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg1ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal2Arg1ReturnPointer10Word(arg1pW10, arg2pW10)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg1ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinlineGlobal3Arg1ReturnPointer10Word(arg1pW10, arg2pW10, arg3pW10)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal1Arg2ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal1Arg2ReturnPointer10Word(arg1pW10)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal2Arg2ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal2Arg2ReturnPointer10Word(arg1pW10, arg2pW10)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinlineGlobal3Arg2ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinlineGlobal3Arg2ReturnPointer10Word(arg1pW10, arg2pW10, arg3pW10)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg0ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv1.ValueReciverNoinlineGlobal1Arg0ReturnValue1Word(arg1vW1)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg0ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv1.ValueReciverNoinlineGlobal2Arg0ReturnValue1Word(arg1vW1, arg2vW1)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg0ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv1.ValueReciverNoinlineGlobal3Arg0ReturnValue1Word(arg1vW1, arg2vW1, arg3vW1)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg1ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv1.ValueReciverNoinlineGlobal1Arg1ReturnValue1Word(arg1vW1)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg1ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv1.ValueReciverNoinlineGlobal2Arg1ReturnValue1Word(arg1vW1, arg2vW1)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg1ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv1.ValueReciverNoinlineGlobal3Arg1ReturnValue1Word(arg1vW1, arg2vW1, arg3vW1)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg2ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv1.ValueReciverNoinlineGlobal1Arg2ReturnValue1Word(arg1vW1)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg2ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv1.ValueReciverNoinlineGlobal2Arg2ReturnValue1Word(arg1vW1, arg2vW1)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg2ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv1.ValueReciverNoinlineGlobal3Arg2ReturnValue1Word(arg1vW1, arg2vW1, arg3vW1)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinlineGlobal1Arg0ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv2.ValueReciverNoinlineGlobal1Arg0ReturnValue2Word(arg1vW2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg0ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv2.ValueReciverNoinlineGlobal2Arg0ReturnValue2Word(arg1vW2, arg2vW2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg0ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv2.ValueReciverNoinlineGlobal3Arg0ReturnValue2Word(arg1vW2, arg2vW2, arg3vW2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg1ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv2.ValueReciverNoinlineGlobal1Arg1ReturnValue2Word(arg1vW2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg1ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv2.ValueReciverNoinlineGlobal2Arg1ReturnValue2Word(arg1vW2, arg2vW2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg1ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv2.ValueReciverNoinlineGlobal3Arg1ReturnValue2Word(arg1vW2, arg2vW2, arg3vW2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg2ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv2.ValueReciverNoinlineGlobal1Arg2ReturnValue2Word(arg1vW2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg2ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv2.ValueReciverNoinlineGlobal2Arg2ReturnValue2Word(arg1vW2, arg2vW2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg2ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv2.ValueReciverNoinlineGlobal3Arg2ReturnValue2Word(arg1vW2, arg2vW2, arg3vW2)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinlineGlobal1Arg0ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv3.ValueReciverNoinlineGlobal1Arg0ReturnValue3Word(arg1vW3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg0ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv3.ValueReciverNoinlineGlobal2Arg0ReturnValue3Word(arg1vW3, arg2vW3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg0ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv3.ValueReciverNoinlineGlobal3Arg0ReturnValue3Word(arg1vW3, arg2vW3, arg3vW3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg1ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv3.ValueReciverNoinlineGlobal1Arg1ReturnValue3Word(arg1vW3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg1ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv3.ValueReciverNoinlineGlobal2Arg1ReturnValue3Word(arg1vW3, arg2vW3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg1ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv3.ValueReciverNoinlineGlobal3Arg1ReturnValue3Word(arg1vW3, arg2vW3, arg3vW3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg2ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv3.ValueReciverNoinlineGlobal1Arg2ReturnValue3Word(arg1vW3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg2ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv3.ValueReciverNoinlineGlobal2Arg2ReturnValue3Word(arg1vW3, arg2vW3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg2ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv3.ValueReciverNoinlineGlobal3Arg2ReturnValue3Word(arg1vW3, arg2vW3, arg3vW3)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinlineGlobal1Arg0ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv4.ValueReciverNoinlineGlobal1Arg0ReturnValue4Word(arg1vW4)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg0ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv4.ValueReciverNoinlineGlobal2Arg0ReturnValue4Word(arg1vW4, arg2vW4)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg0ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv4.ValueReciverNoinlineGlobal3Arg0ReturnValue4Word(arg1vW4, arg2vW4, arg3vW4)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg1ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv4.ValueReciverNoinlineGlobal1Arg1ReturnValue4Word(arg1vW4)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg1ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv4.ValueReciverNoinlineGlobal2Arg1ReturnValue4Word(arg1vW4, arg2vW4)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg1ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv4.ValueReciverNoinlineGlobal3Arg1ReturnValue4Word(arg1vW4, arg2vW4, arg3vW4)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg2ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv4.ValueReciverNoinlineGlobal1Arg2ReturnValue4Word(arg1vW4)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg2ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv4.ValueReciverNoinlineGlobal2Arg2ReturnValue4Word(arg1vW4, arg2vW4)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg2ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv4.ValueReciverNoinlineGlobal3Arg2ReturnValue4Word(arg1vW4, arg2vW4, arg3vW4)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinlineGlobal1Arg0ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv5.ValueReciverNoinlineGlobal1Arg0ReturnValue5Word(arg1vW5)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg0ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv5.ValueReciverNoinlineGlobal2Arg0ReturnValue5Word(arg1vW5, arg2vW5)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg0ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv5.ValueReciverNoinlineGlobal3Arg0ReturnValue5Word(arg1vW5, arg2vW5, arg3vW5)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg1ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv5.ValueReciverNoinlineGlobal1Arg1ReturnValue5Word(arg1vW5)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg1ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv5.ValueReciverNoinlineGlobal2Arg1ReturnValue5Word(arg1vW5, arg2vW5)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg1ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv5.ValueReciverNoinlineGlobal3Arg1ReturnValue5Word(arg1vW5, arg2vW5, arg3vW5)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg2ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv5.ValueReciverNoinlineGlobal1Arg2ReturnValue5Word(arg1vW5)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg2ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv5.ValueReciverNoinlineGlobal2Arg2ReturnValue5Word(arg1vW5, arg2vW5)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg2ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv5.ValueReciverNoinlineGlobal3Arg2ReturnValue5Word(arg1vW5, arg2vW5, arg3vW5)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinlineGlobal1Arg0ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv6.ValueReciverNoinlineGlobal1Arg0ReturnValue6Word(arg1vW6)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg0ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv6.ValueReciverNoinlineGlobal2Arg0ReturnValue6Word(arg1vW6, arg2vW6)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg0ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv6.ValueReciverNoinlineGlobal3Arg0ReturnValue6Word(arg1vW6, arg2vW6, arg3vW6)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg1ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv6.ValueReciverNoinlineGlobal1Arg1ReturnValue6Word(arg1vW6)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg1ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv6.ValueReciverNoinlineGlobal2Arg1ReturnValue6Word(arg1vW6, arg2vW6)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg1ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv6.ValueReciverNoinlineGlobal3Arg1ReturnValue6Word(arg1vW6, arg2vW6, arg3vW6)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg2ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv6.ValueReciverNoinlineGlobal1Arg2ReturnValue6Word(arg1vW6)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg2ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv6.ValueReciverNoinlineGlobal2Arg2ReturnValue6Word(arg1vW6, arg2vW6)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg2ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv6.ValueReciverNoinlineGlobal3Arg2ReturnValue6Word(arg1vW6, arg2vW6, arg3vW6)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinlineGlobal1Arg0ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv7.ValueReciverNoinlineGlobal1Arg0ReturnValue7Word(arg1vW7)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg0ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv7.ValueReciverNoinlineGlobal2Arg0ReturnValue7Word(arg1vW7, arg2vW7)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg0ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv7.ValueReciverNoinlineGlobal3Arg0ReturnValue7Word(arg1vW7, arg2vW7, arg3vW7)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg1ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv7.ValueReciverNoinlineGlobal1Arg1ReturnValue7Word(arg1vW7)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg1ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv7.ValueReciverNoinlineGlobal2Arg1ReturnValue7Word(arg1vW7, arg2vW7)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg1ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv7.ValueReciverNoinlineGlobal3Arg1ReturnValue7Word(arg1vW7, arg2vW7, arg3vW7)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg2ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv7.ValueReciverNoinlineGlobal1Arg2ReturnValue7Word(arg1vW7)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg2ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv7.ValueReciverNoinlineGlobal2Arg2ReturnValue7Word(arg1vW7, arg2vW7)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg2ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv7.ValueReciverNoinlineGlobal3Arg2ReturnValue7Word(arg1vW7, arg2vW7, arg3vW7)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinlineGlobal1Arg0ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv8.ValueReciverNoinlineGlobal1Arg0ReturnValue8Word(arg1vW8)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg0ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv8.ValueReciverNoinlineGlobal2Arg0ReturnValue8Word(arg1vW8, arg2vW8)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg0ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv8.ValueReciverNoinlineGlobal3Arg0ReturnValue8Word(arg1vW8, arg2vW8, arg3vW8)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg1ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv8.ValueReciverNoinlineGlobal1Arg1ReturnValue8Word(arg1vW8)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg1ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv8.ValueReciverNoinlineGlobal2Arg1ReturnValue8Word(arg1vW8, arg2vW8)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg1ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv8.ValueReciverNoinlineGlobal3Arg1ReturnValue8Word(arg1vW8, arg2vW8, arg3vW8)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg2ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv8.ValueReciverNoinlineGlobal1Arg2ReturnValue8Word(arg1vW8)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg2ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv8.ValueReciverNoinlineGlobal2Arg2ReturnValue8Word(arg1vW8, arg2vW8)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg2ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv8.ValueReciverNoinlineGlobal3Arg2ReturnValue8Word(arg1vW8, arg2vW8, arg3vW8)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinlineGlobal1Arg0ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv9.ValueReciverNoinlineGlobal1Arg0ReturnValue9Word(arg1vW9)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg0ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv9.ValueReciverNoinlineGlobal2Arg0ReturnValue9Word(arg1vW9, arg2vW9)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg0ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv9.ValueReciverNoinlineGlobal3Arg0ReturnValue9Word(arg1vW9, arg2vW9, arg3vW9)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg1ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv9.ValueReciverNoinlineGlobal1Arg1ReturnValue9Word(arg1vW9)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg1ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv9.ValueReciverNoinlineGlobal2Arg1ReturnValue9Word(arg1vW9, arg2vW9)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg1ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv9.ValueReciverNoinlineGlobal3Arg1ReturnValue9Word(arg1vW9, arg2vW9, arg3vW9)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg2ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv9.ValueReciverNoinlineGlobal1Arg2ReturnValue9Word(arg1vW9)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg2ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv9.ValueReciverNoinlineGlobal2Arg2ReturnValue9Word(arg1vW9, arg2vW9)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg2ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv9.ValueReciverNoinlineGlobal3Arg2ReturnValue9Word(arg1vW9, arg2vW9, arg3vW9)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinlineGlobal1Arg0ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv10.ValueReciverNoinlineGlobal1Arg0ReturnValue10Word(arg1vW10)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg0ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv10.ValueReciverNoinlineGlobal2Arg0ReturnValue10Word(arg1vW10, arg2vW10)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg0ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv10.ValueReciverNoinlineGlobal3Arg0ReturnValue10Word(arg1vW10, arg2vW10, arg3vW10)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg1ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv10.ValueReciverNoinlineGlobal1Arg1ReturnValue10Word(arg1vW10)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg1ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv10.ValueReciverNoinlineGlobal2Arg1ReturnValue10Word(arg1vW10, arg2vW10)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg1ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv10.ValueReciverNoinlineGlobal3Arg1ReturnValue10Word(arg1vW10, arg2vW10, arg3vW10)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg2ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv10.ValueReciverNoinlineGlobal1Arg2ReturnValue10Word(arg1vW10)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg2ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv10.ValueReciverNoinlineGlobal2Arg2ReturnValue10Word(arg1vW10, arg2vW10)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg2ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv10.ValueReciverNoinlineGlobal3Arg2ReturnValue10Word(arg1vW10, arg2vW10, arg3vW10)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg0ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv1.ValueReciverNoinlineGlobal1Arg0ReturnPointer1Word(arg1pW1)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg0ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv1.ValueReciverNoinlineGlobal2Arg0ReturnPointer1Word(arg1pW1, arg2pW1)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg0ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv1.ValueReciverNoinlineGlobal3Arg0ReturnPointer1Word(arg1pW1, arg2pW1, arg3pW1)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg1ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv1.ValueReciverNoinlineGlobal1Arg1ReturnPointer1Word(arg1pW1)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg1ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv1.ValueReciverNoinlineGlobal2Arg1ReturnPointer1Word(arg1pW1, arg2pW1)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg1ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv1.ValueReciverNoinlineGlobal3Arg1ReturnPointer1Word(arg1pW1, arg2pW1, arg3pW1)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg2ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv1.ValueReciverNoinlineGlobal1Arg2ReturnPointer1Word(arg1pW1)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg2ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv1.ValueReciverNoinlineGlobal2Arg2ReturnPointer1Word(arg1pW1, arg2pW1)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg2ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv1.ValueReciverNoinlineGlobal3Arg2ReturnPointer1Word(arg1pW1, arg2pW1, arg3pW1)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinlineGlobal1Arg0ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv2.ValueReciverNoinlineGlobal1Arg0ReturnPointer2Word(arg1pW2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg0ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv2.ValueReciverNoinlineGlobal2Arg0ReturnPointer2Word(arg1pW2, arg2pW2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg0ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv2.ValueReciverNoinlineGlobal3Arg0ReturnPointer2Word(arg1pW2, arg2pW2, arg3pW2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg1ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv2.ValueReciverNoinlineGlobal1Arg1ReturnPointer2Word(arg1pW2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg1ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv2.ValueReciverNoinlineGlobal2Arg1ReturnPointer2Word(arg1pW2, arg2pW2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg1ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv2.ValueReciverNoinlineGlobal3Arg1ReturnPointer2Word(arg1pW2, arg2pW2, arg3pW2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg2ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv2.ValueReciverNoinlineGlobal1Arg2ReturnPointer2Word(arg1pW2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg2ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv2.ValueReciverNoinlineGlobal2Arg2ReturnPointer2Word(arg1pW2, arg2pW2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg2ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv2.ValueReciverNoinlineGlobal3Arg2ReturnPointer2Word(arg1pW2, arg2pW2, arg3pW2)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinlineGlobal1Arg0ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv3.ValueReciverNoinlineGlobal1Arg0ReturnPointer3Word(arg1pW3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg0ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv3.ValueReciverNoinlineGlobal2Arg0ReturnPointer3Word(arg1pW3, arg2pW3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg0ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv3.ValueReciverNoinlineGlobal3Arg0ReturnPointer3Word(arg1pW3, arg2pW3, arg3pW3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg1ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv3.ValueReciverNoinlineGlobal1Arg1ReturnPointer3Word(arg1pW3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg1ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv3.ValueReciverNoinlineGlobal2Arg1ReturnPointer3Word(arg1pW3, arg2pW3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg1ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv3.ValueReciverNoinlineGlobal3Arg1ReturnPointer3Word(arg1pW3, arg2pW3, arg3pW3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg2ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv3.ValueReciverNoinlineGlobal1Arg2ReturnPointer3Word(arg1pW3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg2ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv3.ValueReciverNoinlineGlobal2Arg2ReturnPointer3Word(arg1pW3, arg2pW3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg2ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv3.ValueReciverNoinlineGlobal3Arg2ReturnPointer3Word(arg1pW3, arg2pW3, arg3pW3)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinlineGlobal1Arg0ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv4.ValueReciverNoinlineGlobal1Arg0ReturnPointer4Word(arg1pW4)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg0ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv4.ValueReciverNoinlineGlobal2Arg0ReturnPointer4Word(arg1pW4, arg2pW4)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg0ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv4.ValueReciverNoinlineGlobal3Arg0ReturnPointer4Word(arg1pW4, arg2pW4, arg3pW4)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg1ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv4.ValueReciverNoinlineGlobal1Arg1ReturnPointer4Word(arg1pW4)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg1ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv4.ValueReciverNoinlineGlobal2Arg1ReturnPointer4Word(arg1pW4, arg2pW4)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg1ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv4.ValueReciverNoinlineGlobal3Arg1ReturnPointer4Word(arg1pW4, arg2pW4, arg3pW4)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg2ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv4.ValueReciverNoinlineGlobal1Arg2ReturnPointer4Word(arg1pW4)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg2ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv4.ValueReciverNoinlineGlobal2Arg2ReturnPointer4Word(arg1pW4, arg2pW4)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg2ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv4.ValueReciverNoinlineGlobal3Arg2ReturnPointer4Word(arg1pW4, arg2pW4, arg3pW4)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinlineGlobal1Arg0ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv5.ValueReciverNoinlineGlobal1Arg0ReturnPointer5Word(arg1pW5)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg0ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv5.ValueReciverNoinlineGlobal2Arg0ReturnPointer5Word(arg1pW5, arg2pW5)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg0ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv5.ValueReciverNoinlineGlobal3Arg0ReturnPointer5Word(arg1pW5, arg2pW5, arg3pW5)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg1ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv5.ValueReciverNoinlineGlobal1Arg1ReturnPointer5Word(arg1pW5)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg1ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv5.ValueReciverNoinlineGlobal2Arg1ReturnPointer5Word(arg1pW5, arg2pW5)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg1ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv5.ValueReciverNoinlineGlobal3Arg1ReturnPointer5Word(arg1pW5, arg2pW5, arg3pW5)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg2ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv5.ValueReciverNoinlineGlobal1Arg2ReturnPointer5Word(arg1pW5)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg2ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv5.ValueReciverNoinlineGlobal2Arg2ReturnPointer5Word(arg1pW5, arg2pW5)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg2ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv5.ValueReciverNoinlineGlobal3Arg2ReturnPointer5Word(arg1pW5, arg2pW5, arg3pW5)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinlineGlobal1Arg0ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv6.ValueReciverNoinlineGlobal1Arg0ReturnPointer6Word(arg1pW6)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg0ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv6.ValueReciverNoinlineGlobal2Arg0ReturnPointer6Word(arg1pW6, arg2pW6)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg0ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv6.ValueReciverNoinlineGlobal3Arg0ReturnPointer6Word(arg1pW6, arg2pW6, arg3pW6)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg1ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv6.ValueReciverNoinlineGlobal1Arg1ReturnPointer6Word(arg1pW6)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg1ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv6.ValueReciverNoinlineGlobal2Arg1ReturnPointer6Word(arg1pW6, arg2pW6)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg1ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv6.ValueReciverNoinlineGlobal3Arg1ReturnPointer6Word(arg1pW6, arg2pW6, arg3pW6)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg2ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv6.ValueReciverNoinlineGlobal1Arg2ReturnPointer6Word(arg1pW6)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg2ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv6.ValueReciverNoinlineGlobal2Arg2ReturnPointer6Word(arg1pW6, arg2pW6)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg2ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv6.ValueReciverNoinlineGlobal3Arg2ReturnPointer6Word(arg1pW6, arg2pW6, arg3pW6)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinlineGlobal1Arg0ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv7.ValueReciverNoinlineGlobal1Arg0ReturnPointer7Word(arg1pW7)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg0ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv7.ValueReciverNoinlineGlobal2Arg0ReturnPointer7Word(arg1pW7, arg2pW7)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg0ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv7.ValueReciverNoinlineGlobal3Arg0ReturnPointer7Word(arg1pW7, arg2pW7, arg3pW7)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg1ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv7.ValueReciverNoinlineGlobal1Arg1ReturnPointer7Word(arg1pW7)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg1ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv7.ValueReciverNoinlineGlobal2Arg1ReturnPointer7Word(arg1pW7, arg2pW7)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg1ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv7.ValueReciverNoinlineGlobal3Arg1ReturnPointer7Word(arg1pW7, arg2pW7, arg3pW7)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg2ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv7.ValueReciverNoinlineGlobal1Arg2ReturnPointer7Word(arg1pW7)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg2ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv7.ValueReciverNoinlineGlobal2Arg2ReturnPointer7Word(arg1pW7, arg2pW7)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg2ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv7.ValueReciverNoinlineGlobal3Arg2ReturnPointer7Word(arg1pW7, arg2pW7, arg3pW7)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinlineGlobal1Arg0ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv8.ValueReciverNoinlineGlobal1Arg0ReturnPointer8Word(arg1pW8)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg0ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv8.ValueReciverNoinlineGlobal2Arg0ReturnPointer8Word(arg1pW8, arg2pW8)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg0ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv8.ValueReciverNoinlineGlobal3Arg0ReturnPointer8Word(arg1pW8, arg2pW8, arg3pW8)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg1ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv8.ValueReciverNoinlineGlobal1Arg1ReturnPointer8Word(arg1pW8)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg1ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv8.ValueReciverNoinlineGlobal2Arg1ReturnPointer8Word(arg1pW8, arg2pW8)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg1ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv8.ValueReciverNoinlineGlobal3Arg1ReturnPointer8Word(arg1pW8, arg2pW8, arg3pW8)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg2ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv8.ValueReciverNoinlineGlobal1Arg2ReturnPointer8Word(arg1pW8)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg2ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv8.ValueReciverNoinlineGlobal2Arg2ReturnPointer8Word(arg1pW8, arg2pW8)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg2ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv8.ValueReciverNoinlineGlobal3Arg2ReturnPointer8Word(arg1pW8, arg2pW8, arg3pW8)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinlineGlobal1Arg0ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv9.ValueReciverNoinlineGlobal1Arg0ReturnPointer9Word(arg1pW9)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg0ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv9.ValueReciverNoinlineGlobal2Arg0ReturnPointer9Word(arg1pW9, arg2pW9)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg0ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv9.ValueReciverNoinlineGlobal3Arg0ReturnPointer9Word(arg1pW9, arg2pW9, arg3pW9)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg1ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv9.ValueReciverNoinlineGlobal1Arg1ReturnPointer9Word(arg1pW9)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg1ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv9.ValueReciverNoinlineGlobal2Arg1ReturnPointer9Word(arg1pW9, arg2pW9)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg1ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv9.ValueReciverNoinlineGlobal3Arg1ReturnPointer9Word(arg1pW9, arg2pW9, arg3pW9)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg2ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv9.ValueReciverNoinlineGlobal1Arg2ReturnPointer9Word(arg1pW9)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg2ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv9.ValueReciverNoinlineGlobal2Arg2ReturnPointer9Word(arg1pW9, arg2pW9)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg2ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv9.ValueReciverNoinlineGlobal3Arg2ReturnPointer9Word(arg1pW9, arg2pW9, arg3pW9)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinlineGlobal1Arg0ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv10.ValueReciverNoinlineGlobal1Arg0ReturnPointer10Word(arg1pW10)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg0ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv10.ValueReciverNoinlineGlobal2Arg0ReturnPointer10Word(arg1pW10, arg2pW10)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg0ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rv10.ValueReciverNoinlineGlobal3Arg0ReturnPointer10Word(arg1pW10, arg2pW10, arg3pW10)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg1ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv10.ValueReciverNoinlineGlobal1Arg1ReturnPointer10Word(arg1pW10)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg1ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv10.ValueReciverNoinlineGlobal2Arg1ReturnPointer10Word(arg1pW10, arg2pW10)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg1ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rv10.ValueReciverNoinlineGlobal3Arg1ReturnPointer10Word(arg1pW10, arg2pW10, arg3pW10)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal1Arg2ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv10.ValueReciverNoinlineGlobal1Arg2ReturnPointer10Word(arg1pW10)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal2Arg2ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv10.ValueReciverNoinlineGlobal2Arg2ReturnPointer10Word(arg1pW10, arg2pW10)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinlineGlobal3Arg2ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rv10.ValueReciverNoinlineGlobal3Arg2ReturnPointer10Word(arg1pW10, arg2pW10, arg3pW10)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg0ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp1.PointerReciverNoinlineGlobal1Arg0ReturnValue1Word(arg1vW1)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg0ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp1.PointerReciverNoinlineGlobal2Arg0ReturnValue1Word(arg1vW1, arg2vW1)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg0ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp1.PointerReciverNoinlineGlobal3Arg0ReturnValue1Word(arg1vW1, arg2vW1, arg3vW1)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg1ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp1.PointerReciverNoinlineGlobal1Arg1ReturnValue1Word(arg1vW1)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg1ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp1.PointerReciverNoinlineGlobal2Arg1ReturnValue1Word(arg1vW1, arg2vW1)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg1ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp1.PointerReciverNoinlineGlobal3Arg1ReturnValue1Word(arg1vW1, arg2vW1, arg3vW1)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg2ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp1.PointerReciverNoinlineGlobal1Arg2ReturnValue1Word(arg1vW1)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg2ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp1.PointerReciverNoinlineGlobal2Arg2ReturnValue1Word(arg1vW1, arg2vW1)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg2ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp1.PointerReciverNoinlineGlobal3Arg2ReturnValue1Word(arg1vW1, arg2vW1, arg3vW1)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinlineGlobal1Arg0ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp2.PointerReciverNoinlineGlobal1Arg0ReturnValue2Word(arg1vW2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg0ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp2.PointerReciverNoinlineGlobal2Arg0ReturnValue2Word(arg1vW2, arg2vW2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg0ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp2.PointerReciverNoinlineGlobal3Arg0ReturnValue2Word(arg1vW2, arg2vW2, arg3vW2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg1ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp2.PointerReciverNoinlineGlobal1Arg1ReturnValue2Word(arg1vW2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg1ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp2.PointerReciverNoinlineGlobal2Arg1ReturnValue2Word(arg1vW2, arg2vW2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg1ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp2.PointerReciverNoinlineGlobal3Arg1ReturnValue2Word(arg1vW2, arg2vW2, arg3vW2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg2ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp2.PointerReciverNoinlineGlobal1Arg2ReturnValue2Word(arg1vW2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg2ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp2.PointerReciverNoinlineGlobal2Arg2ReturnValue2Word(arg1vW2, arg2vW2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg2ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp2.PointerReciverNoinlineGlobal3Arg2ReturnValue2Word(arg1vW2, arg2vW2, arg3vW2)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinlineGlobal1Arg0ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp3.PointerReciverNoinlineGlobal1Arg0ReturnValue3Word(arg1vW3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg0ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp3.PointerReciverNoinlineGlobal2Arg0ReturnValue3Word(arg1vW3, arg2vW3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg0ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp3.PointerReciverNoinlineGlobal3Arg0ReturnValue3Word(arg1vW3, arg2vW3, arg3vW3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg1ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp3.PointerReciverNoinlineGlobal1Arg1ReturnValue3Word(arg1vW3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg1ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp3.PointerReciverNoinlineGlobal2Arg1ReturnValue3Word(arg1vW3, arg2vW3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg1ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp3.PointerReciverNoinlineGlobal3Arg1ReturnValue3Word(arg1vW3, arg2vW3, arg3vW3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg2ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp3.PointerReciverNoinlineGlobal1Arg2ReturnValue3Word(arg1vW3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg2ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp3.PointerReciverNoinlineGlobal2Arg2ReturnValue3Word(arg1vW3, arg2vW3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg2ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp3.PointerReciverNoinlineGlobal3Arg2ReturnValue3Word(arg1vW3, arg2vW3, arg3vW3)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinlineGlobal1Arg0ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp4.PointerReciverNoinlineGlobal1Arg0ReturnValue4Word(arg1vW4)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg0ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp4.PointerReciverNoinlineGlobal2Arg0ReturnValue4Word(arg1vW4, arg2vW4)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg0ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp4.PointerReciverNoinlineGlobal3Arg0ReturnValue4Word(arg1vW4, arg2vW4, arg3vW4)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg1ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp4.PointerReciverNoinlineGlobal1Arg1ReturnValue4Word(arg1vW4)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg1ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp4.PointerReciverNoinlineGlobal2Arg1ReturnValue4Word(arg1vW4, arg2vW4)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg1ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp4.PointerReciverNoinlineGlobal3Arg1ReturnValue4Word(arg1vW4, arg2vW4, arg3vW4)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg2ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp4.PointerReciverNoinlineGlobal1Arg2ReturnValue4Word(arg1vW4)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg2ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp4.PointerReciverNoinlineGlobal2Arg2ReturnValue4Word(arg1vW4, arg2vW4)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg2ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp4.PointerReciverNoinlineGlobal3Arg2ReturnValue4Word(arg1vW4, arg2vW4, arg3vW4)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinlineGlobal1Arg0ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp5.PointerReciverNoinlineGlobal1Arg0ReturnValue5Word(arg1vW5)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg0ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp5.PointerReciverNoinlineGlobal2Arg0ReturnValue5Word(arg1vW5, arg2vW5)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg0ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp5.PointerReciverNoinlineGlobal3Arg0ReturnValue5Word(arg1vW5, arg2vW5, arg3vW5)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg1ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp5.PointerReciverNoinlineGlobal1Arg1ReturnValue5Word(arg1vW5)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg1ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp5.PointerReciverNoinlineGlobal2Arg1ReturnValue5Word(arg1vW5, arg2vW5)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg1ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp5.PointerReciverNoinlineGlobal3Arg1ReturnValue5Word(arg1vW5, arg2vW5, arg3vW5)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg2ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp5.PointerReciverNoinlineGlobal1Arg2ReturnValue5Word(arg1vW5)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg2ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp5.PointerReciverNoinlineGlobal2Arg2ReturnValue5Word(arg1vW5, arg2vW5)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg2ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp5.PointerReciverNoinlineGlobal3Arg2ReturnValue5Word(arg1vW5, arg2vW5, arg3vW5)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinlineGlobal1Arg0ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp6.PointerReciverNoinlineGlobal1Arg0ReturnValue6Word(arg1vW6)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg0ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp6.PointerReciverNoinlineGlobal2Arg0ReturnValue6Word(arg1vW6, arg2vW6)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg0ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp6.PointerReciverNoinlineGlobal3Arg0ReturnValue6Word(arg1vW6, arg2vW6, arg3vW6)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg1ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp6.PointerReciverNoinlineGlobal1Arg1ReturnValue6Word(arg1vW6)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg1ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp6.PointerReciverNoinlineGlobal2Arg1ReturnValue6Word(arg1vW6, arg2vW6)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg1ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp6.PointerReciverNoinlineGlobal3Arg1ReturnValue6Word(arg1vW6, arg2vW6, arg3vW6)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg2ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp6.PointerReciverNoinlineGlobal1Arg2ReturnValue6Word(arg1vW6)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg2ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp6.PointerReciverNoinlineGlobal2Arg2ReturnValue6Word(arg1vW6, arg2vW6)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg2ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp6.PointerReciverNoinlineGlobal3Arg2ReturnValue6Word(arg1vW6, arg2vW6, arg3vW6)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinlineGlobal1Arg0ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp7.PointerReciverNoinlineGlobal1Arg0ReturnValue7Word(arg1vW7)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg0ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp7.PointerReciverNoinlineGlobal2Arg0ReturnValue7Word(arg1vW7, arg2vW7)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg0ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp7.PointerReciverNoinlineGlobal3Arg0ReturnValue7Word(arg1vW7, arg2vW7, arg3vW7)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg1ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp7.PointerReciverNoinlineGlobal1Arg1ReturnValue7Word(arg1vW7)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg1ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp7.PointerReciverNoinlineGlobal2Arg1ReturnValue7Word(arg1vW7, arg2vW7)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg1ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp7.PointerReciverNoinlineGlobal3Arg1ReturnValue7Word(arg1vW7, arg2vW7, arg3vW7)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg2ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp7.PointerReciverNoinlineGlobal1Arg2ReturnValue7Word(arg1vW7)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg2ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp7.PointerReciverNoinlineGlobal2Arg2ReturnValue7Word(arg1vW7, arg2vW7)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg2ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp7.PointerReciverNoinlineGlobal3Arg2ReturnValue7Word(arg1vW7, arg2vW7, arg3vW7)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinlineGlobal1Arg0ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp8.PointerReciverNoinlineGlobal1Arg0ReturnValue8Word(arg1vW8)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg0ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp8.PointerReciverNoinlineGlobal2Arg0ReturnValue8Word(arg1vW8, arg2vW8)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg0ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp8.PointerReciverNoinlineGlobal3Arg0ReturnValue8Word(arg1vW8, arg2vW8, arg3vW8)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg1ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp8.PointerReciverNoinlineGlobal1Arg1ReturnValue8Word(arg1vW8)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg1ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp8.PointerReciverNoinlineGlobal2Arg1ReturnValue8Word(arg1vW8, arg2vW8)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg1ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp8.PointerReciverNoinlineGlobal3Arg1ReturnValue8Word(arg1vW8, arg2vW8, arg3vW8)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg2ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp8.PointerReciverNoinlineGlobal1Arg2ReturnValue8Word(arg1vW8)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg2ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp8.PointerReciverNoinlineGlobal2Arg2ReturnValue8Word(arg1vW8, arg2vW8)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg2ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp8.PointerReciverNoinlineGlobal3Arg2ReturnValue8Word(arg1vW8, arg2vW8, arg3vW8)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinlineGlobal1Arg0ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp9.PointerReciverNoinlineGlobal1Arg0ReturnValue9Word(arg1vW9)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg0ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp9.PointerReciverNoinlineGlobal2Arg0ReturnValue9Word(arg1vW9, arg2vW9)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg0ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp9.PointerReciverNoinlineGlobal3Arg0ReturnValue9Word(arg1vW9, arg2vW9, arg3vW9)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg1ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp9.PointerReciverNoinlineGlobal1Arg1ReturnValue9Word(arg1vW9)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg1ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp9.PointerReciverNoinlineGlobal2Arg1ReturnValue9Word(arg1vW9, arg2vW9)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg1ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp9.PointerReciverNoinlineGlobal3Arg1ReturnValue9Word(arg1vW9, arg2vW9, arg3vW9)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg2ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp9.PointerReciverNoinlineGlobal1Arg2ReturnValue9Word(arg1vW9)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg2ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp9.PointerReciverNoinlineGlobal2Arg2ReturnValue9Word(arg1vW9, arg2vW9)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg2ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp9.PointerReciverNoinlineGlobal3Arg2ReturnValue9Word(arg1vW9, arg2vW9, arg3vW9)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinlineGlobal1Arg0ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp10.PointerReciverNoinlineGlobal1Arg0ReturnValue10Word(arg1vW10)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg0ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp10.PointerReciverNoinlineGlobal2Arg0ReturnValue10Word(arg1vW10, arg2vW10)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg0ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp10.PointerReciverNoinlineGlobal3Arg0ReturnValue10Word(arg1vW10, arg2vW10, arg3vW10)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg1ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp10.PointerReciverNoinlineGlobal1Arg1ReturnValue10Word(arg1vW10)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg1ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp10.PointerReciverNoinlineGlobal2Arg1ReturnValue10Word(arg1vW10, arg2vW10)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg1ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp10.PointerReciverNoinlineGlobal3Arg1ReturnValue10Word(arg1vW10, arg2vW10, arg3vW10)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg2ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp10.PointerReciverNoinlineGlobal1Arg2ReturnValue10Word(arg1vW10)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg2ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp10.PointerReciverNoinlineGlobal2Arg2ReturnValue10Word(arg1vW10, arg2vW10)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg2ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp10.PointerReciverNoinlineGlobal3Arg2ReturnValue10Word(arg1vW10, arg2vW10, arg3vW10)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg0ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp1.PointerReciverNoinlineGlobal1Arg0ReturnPointer1Word(arg1pW1)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg0ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp1.PointerReciverNoinlineGlobal2Arg0ReturnPointer1Word(arg1pW1, arg2pW1)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg0ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp1.PointerReciverNoinlineGlobal3Arg0ReturnPointer1Word(arg1pW1, arg2pW1, arg3pW1)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg1ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp1.PointerReciverNoinlineGlobal1Arg1ReturnPointer1Word(arg1pW1)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg1ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp1.PointerReciverNoinlineGlobal2Arg1ReturnPointer1Word(arg1pW1, arg2pW1)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg1ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp1.PointerReciverNoinlineGlobal3Arg1ReturnPointer1Word(arg1pW1, arg2pW1, arg3pW1)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg2ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp1.PointerReciverNoinlineGlobal1Arg2ReturnPointer1Word(arg1pW1)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg2ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp1.PointerReciverNoinlineGlobal2Arg2ReturnPointer1Word(arg1pW1, arg2pW1)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg2ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp1.PointerReciverNoinlineGlobal3Arg2ReturnPointer1Word(arg1pW1, arg2pW1, arg3pW1)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinlineGlobal1Arg0ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp2.PointerReciverNoinlineGlobal1Arg0ReturnPointer2Word(arg1pW2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg0ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp2.PointerReciverNoinlineGlobal2Arg0ReturnPointer2Word(arg1pW2, arg2pW2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg0ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp2.PointerReciverNoinlineGlobal3Arg0ReturnPointer2Word(arg1pW2, arg2pW2, arg3pW2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg1ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp2.PointerReciverNoinlineGlobal1Arg1ReturnPointer2Word(arg1pW2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg1ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp2.PointerReciverNoinlineGlobal2Arg1ReturnPointer2Word(arg1pW2, arg2pW2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg1ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp2.PointerReciverNoinlineGlobal3Arg1ReturnPointer2Word(arg1pW2, arg2pW2, arg3pW2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg2ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp2.PointerReciverNoinlineGlobal1Arg2ReturnPointer2Word(arg1pW2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg2ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp2.PointerReciverNoinlineGlobal2Arg2ReturnPointer2Word(arg1pW2, arg2pW2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg2ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp2.PointerReciverNoinlineGlobal3Arg2ReturnPointer2Word(arg1pW2, arg2pW2, arg3pW2)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinlineGlobal1Arg0ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp3.PointerReciverNoinlineGlobal1Arg0ReturnPointer3Word(arg1pW3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg0ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp3.PointerReciverNoinlineGlobal2Arg0ReturnPointer3Word(arg1pW3, arg2pW3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg0ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp3.PointerReciverNoinlineGlobal3Arg0ReturnPointer3Word(arg1pW3, arg2pW3, arg3pW3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg1ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp3.PointerReciverNoinlineGlobal1Arg1ReturnPointer3Word(arg1pW3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg1ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp3.PointerReciverNoinlineGlobal2Arg1ReturnPointer3Word(arg1pW3, arg2pW3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg1ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp3.PointerReciverNoinlineGlobal3Arg1ReturnPointer3Word(arg1pW3, arg2pW3, arg3pW3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg2ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp3.PointerReciverNoinlineGlobal1Arg2ReturnPointer3Word(arg1pW3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg2ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp3.PointerReciverNoinlineGlobal2Arg2ReturnPointer3Word(arg1pW3, arg2pW3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg2ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp3.PointerReciverNoinlineGlobal3Arg2ReturnPointer3Word(arg1pW3, arg2pW3, arg3pW3)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinlineGlobal1Arg0ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp4.PointerReciverNoinlineGlobal1Arg0ReturnPointer4Word(arg1pW4)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg0ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp4.PointerReciverNoinlineGlobal2Arg0ReturnPointer4Word(arg1pW4, arg2pW4)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg0ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp4.PointerReciverNoinlineGlobal3Arg0ReturnPointer4Word(arg1pW4, arg2pW4, arg3pW4)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg1ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp4.PointerReciverNoinlineGlobal1Arg1ReturnPointer4Word(arg1pW4)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg1ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp4.PointerReciverNoinlineGlobal2Arg1ReturnPointer4Word(arg1pW4, arg2pW4)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg1ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp4.PointerReciverNoinlineGlobal3Arg1ReturnPointer4Word(arg1pW4, arg2pW4, arg3pW4)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg2ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp4.PointerReciverNoinlineGlobal1Arg2ReturnPointer4Word(arg1pW4)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg2ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp4.PointerReciverNoinlineGlobal2Arg2ReturnPointer4Word(arg1pW4, arg2pW4)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg2ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp4.PointerReciverNoinlineGlobal3Arg2ReturnPointer4Word(arg1pW4, arg2pW4, arg3pW4)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinlineGlobal1Arg0ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp5.PointerReciverNoinlineGlobal1Arg0ReturnPointer5Word(arg1pW5)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg0ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp5.PointerReciverNoinlineGlobal2Arg0ReturnPointer5Word(arg1pW5, arg2pW5)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg0ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp5.PointerReciverNoinlineGlobal3Arg0ReturnPointer5Word(arg1pW5, arg2pW5, arg3pW5)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg1ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp5.PointerReciverNoinlineGlobal1Arg1ReturnPointer5Word(arg1pW5)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg1ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp5.PointerReciverNoinlineGlobal2Arg1ReturnPointer5Word(arg1pW5, arg2pW5)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg1ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp5.PointerReciverNoinlineGlobal3Arg1ReturnPointer5Word(arg1pW5, arg2pW5, arg3pW5)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg2ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp5.PointerReciverNoinlineGlobal1Arg2ReturnPointer5Word(arg1pW5)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg2ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp5.PointerReciverNoinlineGlobal2Arg2ReturnPointer5Word(arg1pW5, arg2pW5)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg2ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp5.PointerReciverNoinlineGlobal3Arg2ReturnPointer5Word(arg1pW5, arg2pW5, arg3pW5)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinlineGlobal1Arg0ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp6.PointerReciverNoinlineGlobal1Arg0ReturnPointer6Word(arg1pW6)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg0ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp6.PointerReciverNoinlineGlobal2Arg0ReturnPointer6Word(arg1pW6, arg2pW6)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg0ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp6.PointerReciverNoinlineGlobal3Arg0ReturnPointer6Word(arg1pW6, arg2pW6, arg3pW6)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg1ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp6.PointerReciverNoinlineGlobal1Arg1ReturnPointer6Word(arg1pW6)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg1ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp6.PointerReciverNoinlineGlobal2Arg1ReturnPointer6Word(arg1pW6, arg2pW6)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg1ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp6.PointerReciverNoinlineGlobal3Arg1ReturnPointer6Word(arg1pW6, arg2pW6, arg3pW6)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg2ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp6.PointerReciverNoinlineGlobal1Arg2ReturnPointer6Word(arg1pW6)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg2ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp6.PointerReciverNoinlineGlobal2Arg2ReturnPointer6Word(arg1pW6, arg2pW6)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg2ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp6.PointerReciverNoinlineGlobal3Arg2ReturnPointer6Word(arg1pW6, arg2pW6, arg3pW6)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinlineGlobal1Arg0ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp7.PointerReciverNoinlineGlobal1Arg0ReturnPointer7Word(arg1pW7)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg0ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp7.PointerReciverNoinlineGlobal2Arg0ReturnPointer7Word(arg1pW7, arg2pW7)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg0ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp7.PointerReciverNoinlineGlobal3Arg0ReturnPointer7Word(arg1pW7, arg2pW7, arg3pW7)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg1ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp7.PointerReciverNoinlineGlobal1Arg1ReturnPointer7Word(arg1pW7)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg1ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp7.PointerReciverNoinlineGlobal2Arg1ReturnPointer7Word(arg1pW7, arg2pW7)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg1ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp7.PointerReciverNoinlineGlobal3Arg1ReturnPointer7Word(arg1pW7, arg2pW7, arg3pW7)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg2ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp7.PointerReciverNoinlineGlobal1Arg2ReturnPointer7Word(arg1pW7)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg2ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp7.PointerReciverNoinlineGlobal2Arg2ReturnPointer7Word(arg1pW7, arg2pW7)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg2ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp7.PointerReciverNoinlineGlobal3Arg2ReturnPointer7Word(arg1pW7, arg2pW7, arg3pW7)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinlineGlobal1Arg0ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp8.PointerReciverNoinlineGlobal1Arg0ReturnPointer8Word(arg1pW8)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg0ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp8.PointerReciverNoinlineGlobal2Arg0ReturnPointer8Word(arg1pW8, arg2pW8)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg0ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp8.PointerReciverNoinlineGlobal3Arg0ReturnPointer8Word(arg1pW8, arg2pW8, arg3pW8)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg1ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp8.PointerReciverNoinlineGlobal1Arg1ReturnPointer8Word(arg1pW8)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg1ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp8.PointerReciverNoinlineGlobal2Arg1ReturnPointer8Word(arg1pW8, arg2pW8)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg1ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp8.PointerReciverNoinlineGlobal3Arg1ReturnPointer8Word(arg1pW8, arg2pW8, arg3pW8)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg2ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp8.PointerReciverNoinlineGlobal1Arg2ReturnPointer8Word(arg1pW8)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg2ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp8.PointerReciverNoinlineGlobal2Arg2ReturnPointer8Word(arg1pW8, arg2pW8)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg2ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp8.PointerReciverNoinlineGlobal3Arg2ReturnPointer8Word(arg1pW8, arg2pW8, arg3pW8)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinlineGlobal1Arg0ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp9.PointerReciverNoinlineGlobal1Arg0ReturnPointer9Word(arg1pW9)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg0ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp9.PointerReciverNoinlineGlobal2Arg0ReturnPointer9Word(arg1pW9, arg2pW9)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg0ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp9.PointerReciverNoinlineGlobal3Arg0ReturnPointer9Word(arg1pW9, arg2pW9, arg3pW9)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg1ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp9.PointerReciverNoinlineGlobal1Arg1ReturnPointer9Word(arg1pW9)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg1ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp9.PointerReciverNoinlineGlobal2Arg1ReturnPointer9Word(arg1pW9, arg2pW9)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg1ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp9.PointerReciverNoinlineGlobal3Arg1ReturnPointer9Word(arg1pW9, arg2pW9, arg3pW9)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg2ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp9.PointerReciverNoinlineGlobal1Arg2ReturnPointer9Word(arg1pW9)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg2ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp9.PointerReciverNoinlineGlobal2Arg2ReturnPointer9Word(arg1pW9, arg2pW9)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg2ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp9.PointerReciverNoinlineGlobal3Arg2ReturnPointer9Word(arg1pW9, arg2pW9, arg3pW9)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinlineGlobal1Arg0ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp10.PointerReciverNoinlineGlobal1Arg0ReturnPointer10Word(arg1pW10)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg0ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp10.PointerReciverNoinlineGlobal2Arg0ReturnPointer10Word(arg1pW10, arg2pW10)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg0ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rp10.PointerReciverNoinlineGlobal3Arg0ReturnPointer10Word(arg1pW10, arg2pW10, arg3pW10)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg1ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp10.PointerReciverNoinlineGlobal1Arg1ReturnPointer10Word(arg1pW10)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg1ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp10.PointerReciverNoinlineGlobal2Arg1ReturnPointer10Word(arg1pW10, arg2pW10)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg1ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = rp10.PointerReciverNoinlineGlobal3Arg1ReturnPointer10Word(arg1pW10, arg2pW10, arg3pW10)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal1Arg2ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp10.PointerReciverNoinlineGlobal1Arg2ReturnPointer10Word(arg1pW10)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal2Arg2ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp10.PointerReciverNoinlineGlobal2Arg2ReturnPointer10Word(arg1pW10, arg2pW10)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinlineGlobal3Arg2ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rp10.PointerReciverNoinlineGlobal3Arg2ReturnPointer10Word(arg1pW10, arg2pW10, arg3pW10)
	}
	b.StopTimer()
}
