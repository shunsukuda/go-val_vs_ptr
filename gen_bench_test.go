package main

import "testing"

type W1 [1]uint64
type W2 [2]uint64
type W3 [3]uint64
type W4 [4]uint64
type W5 [5]uint64
type W6 [6]uint64
type W7 [7]uint64
type W8 [8]uint64
type W9 [9]uint64
type W10 [10]uint64

//go:noinline
func FuncNoinline1Arg0ReturnValue1Word(arg W1) {
	_ = arg
}

//go:noinline
func FuncNoinline2Arg0ReturnValue1Word(arg1, arg2 W1) {
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinline3Arg0ReturnValue1Word(arg1, arg2, arg3 W1) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinline1Arg1ReturnValue1Word(arg W1) W1 {
	_ = arg
	return arg
}

//go:noinline
func FuncNoinline2Arg1ReturnValue1Word(arg1, arg2 W1) W1 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinline3Arg1ReturnValue1Word(arg1, arg2, arg3 W1) W1 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinline1Arg2ReturnValue1Word(arg W1) (W1, W1) {
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinline2Arg2ReturnValue1Word(arg1, arg2 W1) (W1, W1) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinline3Arg2ReturnValue1Word(arg1, arg2, arg3 W1) (W1, W1) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinline1Arg0ReturnValue2Word(arg W2) {
	_ = arg
}

//go:noinline
func FuncNoinline2Arg0ReturnValue2Word(arg1, arg2 W2) {
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinline3Arg0ReturnValue2Word(arg1, arg2, arg3 W2) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinline1Arg1ReturnValue2Word(arg W2) W2 {
	_ = arg
	return arg
}

//go:noinline
func FuncNoinline2Arg1ReturnValue2Word(arg1, arg2 W2) W2 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinline3Arg1ReturnValue2Word(arg1, arg2, arg3 W2) W2 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinline1Arg2ReturnValue2Word(arg W2) (W2, W2) {
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinline2Arg2ReturnValue2Word(arg1, arg2 W2) (W2, W2) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinline3Arg2ReturnValue2Word(arg1, arg2, arg3 W2) (W2, W2) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinline1Arg0ReturnValue3Word(arg W3) {
	_ = arg
}

//go:noinline
func FuncNoinline2Arg0ReturnValue3Word(arg1, arg2 W3) {
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinline3Arg0ReturnValue3Word(arg1, arg2, arg3 W3) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinline1Arg1ReturnValue3Word(arg W3) W3 {
	_ = arg
	return arg
}

//go:noinline
func FuncNoinline2Arg1ReturnValue3Word(arg1, arg2 W3) W3 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinline3Arg1ReturnValue3Word(arg1, arg2, arg3 W3) W3 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinline1Arg2ReturnValue3Word(arg W3) (W3, W3) {
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinline2Arg2ReturnValue3Word(arg1, arg2 W3) (W3, W3) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinline3Arg2ReturnValue3Word(arg1, arg2, arg3 W3) (W3, W3) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinline1Arg0ReturnValue4Word(arg W4) {
	_ = arg
}

//go:noinline
func FuncNoinline2Arg0ReturnValue4Word(arg1, arg2 W4) {
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinline3Arg0ReturnValue4Word(arg1, arg2, arg3 W4) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinline1Arg1ReturnValue4Word(arg W4) W4 {
	_ = arg
	return arg
}

//go:noinline
func FuncNoinline2Arg1ReturnValue4Word(arg1, arg2 W4) W4 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinline3Arg1ReturnValue4Word(arg1, arg2, arg3 W4) W4 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinline1Arg2ReturnValue4Word(arg W4) (W4, W4) {
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinline2Arg2ReturnValue4Word(arg1, arg2 W4) (W4, W4) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinline3Arg2ReturnValue4Word(arg1, arg2, arg3 W4) (W4, W4) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinline1Arg0ReturnValue5Word(arg W5) {
	_ = arg
}

//go:noinline
func FuncNoinline2Arg0ReturnValue5Word(arg1, arg2 W5) {
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinline3Arg0ReturnValue5Word(arg1, arg2, arg3 W5) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinline1Arg1ReturnValue5Word(arg W5) W5 {
	_ = arg
	return arg
}

//go:noinline
func FuncNoinline2Arg1ReturnValue5Word(arg1, arg2 W5) W5 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinline3Arg1ReturnValue5Word(arg1, arg2, arg3 W5) W5 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinline1Arg2ReturnValue5Word(arg W5) (W5, W5) {
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinline2Arg2ReturnValue5Word(arg1, arg2 W5) (W5, W5) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinline3Arg2ReturnValue5Word(arg1, arg2, arg3 W5) (W5, W5) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinline1Arg0ReturnValue6Word(arg W6) {
	_ = arg
}

//go:noinline
func FuncNoinline2Arg0ReturnValue6Word(arg1, arg2 W6) {
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinline3Arg0ReturnValue6Word(arg1, arg2, arg3 W6) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinline1Arg1ReturnValue6Word(arg W6) W6 {
	_ = arg
	return arg
}

//go:noinline
func FuncNoinline2Arg1ReturnValue6Word(arg1, arg2 W6) W6 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinline3Arg1ReturnValue6Word(arg1, arg2, arg3 W6) W6 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinline1Arg2ReturnValue6Word(arg W6) (W6, W6) {
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinline2Arg2ReturnValue6Word(arg1, arg2 W6) (W6, W6) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinline3Arg2ReturnValue6Word(arg1, arg2, arg3 W6) (W6, W6) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinline1Arg0ReturnValue7Word(arg W7) {
	_ = arg
}

//go:noinline
func FuncNoinline2Arg0ReturnValue7Word(arg1, arg2 W7) {
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinline3Arg0ReturnValue7Word(arg1, arg2, arg3 W7) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinline1Arg1ReturnValue7Word(arg W7) W7 {
	_ = arg
	return arg
}

//go:noinline
func FuncNoinline2Arg1ReturnValue7Word(arg1, arg2 W7) W7 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinline3Arg1ReturnValue7Word(arg1, arg2, arg3 W7) W7 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinline1Arg2ReturnValue7Word(arg W7) (W7, W7) {
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinline2Arg2ReturnValue7Word(arg1, arg2 W7) (W7, W7) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinline3Arg2ReturnValue7Word(arg1, arg2, arg3 W7) (W7, W7) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinline1Arg0ReturnValue8Word(arg W8) {
	_ = arg
}

//go:noinline
func FuncNoinline2Arg0ReturnValue8Word(arg1, arg2 W8) {
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinline3Arg0ReturnValue8Word(arg1, arg2, arg3 W8) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinline1Arg1ReturnValue8Word(arg W8) W8 {
	_ = arg
	return arg
}

//go:noinline
func FuncNoinline2Arg1ReturnValue8Word(arg1, arg2 W8) W8 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinline3Arg1ReturnValue8Word(arg1, arg2, arg3 W8) W8 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinline1Arg2ReturnValue8Word(arg W8) (W8, W8) {
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinline2Arg2ReturnValue8Word(arg1, arg2 W8) (W8, W8) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinline3Arg2ReturnValue8Word(arg1, arg2, arg3 W8) (W8, W8) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinline1Arg0ReturnValue9Word(arg W9) {
	_ = arg
}

//go:noinline
func FuncNoinline2Arg0ReturnValue9Word(arg1, arg2 W9) {
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinline3Arg0ReturnValue9Word(arg1, arg2, arg3 W9) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinline1Arg1ReturnValue9Word(arg W9) W9 {
	_ = arg
	return arg
}

//go:noinline
func FuncNoinline2Arg1ReturnValue9Word(arg1, arg2 W9) W9 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinline3Arg1ReturnValue9Word(arg1, arg2, arg3 W9) W9 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinline1Arg2ReturnValue9Word(arg W9) (W9, W9) {
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinline2Arg2ReturnValue9Word(arg1, arg2 W9) (W9, W9) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinline3Arg2ReturnValue9Word(arg1, arg2, arg3 W9) (W9, W9) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinline1Arg0ReturnValue10Word(arg W10) {
	_ = arg
}

//go:noinline
func FuncNoinline2Arg0ReturnValue10Word(arg1, arg2 W10) {
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinline3Arg0ReturnValue10Word(arg1, arg2, arg3 W10) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func FuncNoinline1Arg1ReturnValue10Word(arg W10) W10 {
	_ = arg
	return arg
}

//go:noinline
func FuncNoinline2Arg1ReturnValue10Word(arg1, arg2 W10) W10 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func FuncNoinline3Arg1ReturnValue10Word(arg1, arg2, arg3 W10) W10 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func FuncNoinline1Arg2ReturnValue10Word(arg W10) (W10, W10) {
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinline2Arg2ReturnValue10Word(arg1, arg2 W10) (W10, W10) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func FuncNoinline3Arg2ReturnValue10Word(arg1, arg2, arg3 W10) (W10, W10) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func FuncNoinline1Arg0ReturnPointer1Word(arg *W1) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func FuncNoinline2Arg0ReturnPointer1Word(arg1, arg2 *W1) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinline3Arg0ReturnPointer1Word(arg1, arg2, arg3 *W1) {
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
func FuncNoinline1Arg1ReturnPointer1Word(arg *W1) *W1 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func FuncNoinline2Arg1ReturnPointer1Word(arg1, arg2 *W1) *W1 {
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
func FuncNoinline3Arg1ReturnPointer1Word(arg1, arg2, arg3 *W1) *W1 {
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
func FuncNoinline1Arg2ReturnPointer1Word(arg *W1) (*W1, *W1) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinline2Arg2ReturnPointer1Word(arg1, arg2 *W1) (*W1, *W1) {
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
func FuncNoinline3Arg2ReturnPointer1Word(arg1, arg2, arg3 *W1) (*W1, *W1) {
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
func FuncNoinline1Arg0ReturnPointer2Word(arg *W2) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func FuncNoinline2Arg0ReturnPointer2Word(arg1, arg2 *W2) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinline3Arg0ReturnPointer2Word(arg1, arg2, arg3 *W2) {
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
func FuncNoinline1Arg1ReturnPointer2Word(arg *W2) *W2 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func FuncNoinline2Arg1ReturnPointer2Word(arg1, arg2 *W2) *W2 {
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
func FuncNoinline3Arg1ReturnPointer2Word(arg1, arg2, arg3 *W2) *W2 {
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
func FuncNoinline1Arg2ReturnPointer2Word(arg *W2) (*W2, *W2) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinline2Arg2ReturnPointer2Word(arg1, arg2 *W2) (*W2, *W2) {
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
func FuncNoinline3Arg2ReturnPointer2Word(arg1, arg2, arg3 *W2) (*W2, *W2) {
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
func FuncNoinline1Arg0ReturnPointer3Word(arg *W3) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func FuncNoinline2Arg0ReturnPointer3Word(arg1, arg2 *W3) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinline3Arg0ReturnPointer3Word(arg1, arg2, arg3 *W3) {
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
func FuncNoinline1Arg1ReturnPointer3Word(arg *W3) *W3 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func FuncNoinline2Arg1ReturnPointer3Word(arg1, arg2 *W3) *W3 {
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
func FuncNoinline3Arg1ReturnPointer3Word(arg1, arg2, arg3 *W3) *W3 {
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
func FuncNoinline1Arg2ReturnPointer3Word(arg *W3) (*W3, *W3) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinline2Arg2ReturnPointer3Word(arg1, arg2 *W3) (*W3, *W3) {
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
func FuncNoinline3Arg2ReturnPointer3Word(arg1, arg2, arg3 *W3) (*W3, *W3) {
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
func FuncNoinline1Arg0ReturnPointer4Word(arg *W4) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func FuncNoinline2Arg0ReturnPointer4Word(arg1, arg2 *W4) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinline3Arg0ReturnPointer4Word(arg1, arg2, arg3 *W4) {
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
func FuncNoinline1Arg1ReturnPointer4Word(arg *W4) *W4 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func FuncNoinline2Arg1ReturnPointer4Word(arg1, arg2 *W4) *W4 {
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
func FuncNoinline3Arg1ReturnPointer4Word(arg1, arg2, arg3 *W4) *W4 {
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
func FuncNoinline1Arg2ReturnPointer4Word(arg *W4) (*W4, *W4) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinline2Arg2ReturnPointer4Word(arg1, arg2 *W4) (*W4, *W4) {
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
func FuncNoinline3Arg2ReturnPointer4Word(arg1, arg2, arg3 *W4) (*W4, *W4) {
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
func FuncNoinline1Arg0ReturnPointer5Word(arg *W5) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func FuncNoinline2Arg0ReturnPointer5Word(arg1, arg2 *W5) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinline3Arg0ReturnPointer5Word(arg1, arg2, arg3 *W5) {
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
func FuncNoinline1Arg1ReturnPointer5Word(arg *W5) *W5 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func FuncNoinline2Arg1ReturnPointer5Word(arg1, arg2 *W5) *W5 {
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
func FuncNoinline3Arg1ReturnPointer5Word(arg1, arg2, arg3 *W5) *W5 {
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
func FuncNoinline1Arg2ReturnPointer5Word(arg *W5) (*W5, *W5) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinline2Arg2ReturnPointer5Word(arg1, arg2 *W5) (*W5, *W5) {
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
func FuncNoinline3Arg2ReturnPointer5Word(arg1, arg2, arg3 *W5) (*W5, *W5) {
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
func FuncNoinline1Arg0ReturnPointer6Word(arg *W6) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func FuncNoinline2Arg0ReturnPointer6Word(arg1, arg2 *W6) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinline3Arg0ReturnPointer6Word(arg1, arg2, arg3 *W6) {
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
func FuncNoinline1Arg1ReturnPointer6Word(arg *W6) *W6 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func FuncNoinline2Arg1ReturnPointer6Word(arg1, arg2 *W6) *W6 {
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
func FuncNoinline3Arg1ReturnPointer6Word(arg1, arg2, arg3 *W6) *W6 {
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
func FuncNoinline1Arg2ReturnPointer6Word(arg *W6) (*W6, *W6) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinline2Arg2ReturnPointer6Word(arg1, arg2 *W6) (*W6, *W6) {
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
func FuncNoinline3Arg2ReturnPointer6Word(arg1, arg2, arg3 *W6) (*W6, *W6) {
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
func FuncNoinline1Arg0ReturnPointer7Word(arg *W7) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func FuncNoinline2Arg0ReturnPointer7Word(arg1, arg2 *W7) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinline3Arg0ReturnPointer7Word(arg1, arg2, arg3 *W7) {
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
func FuncNoinline1Arg1ReturnPointer7Word(arg *W7) *W7 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func FuncNoinline2Arg1ReturnPointer7Word(arg1, arg2 *W7) *W7 {
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
func FuncNoinline3Arg1ReturnPointer7Word(arg1, arg2, arg3 *W7) *W7 {
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
func FuncNoinline1Arg2ReturnPointer7Word(arg *W7) (*W7, *W7) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinline2Arg2ReturnPointer7Word(arg1, arg2 *W7) (*W7, *W7) {
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
func FuncNoinline3Arg2ReturnPointer7Word(arg1, arg2, arg3 *W7) (*W7, *W7) {
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
func FuncNoinline1Arg0ReturnPointer8Word(arg *W8) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func FuncNoinline2Arg0ReturnPointer8Word(arg1, arg2 *W8) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinline3Arg0ReturnPointer8Word(arg1, arg2, arg3 *W8) {
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
func FuncNoinline1Arg1ReturnPointer8Word(arg *W8) *W8 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func FuncNoinline2Arg1ReturnPointer8Word(arg1, arg2 *W8) *W8 {
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
func FuncNoinline3Arg1ReturnPointer8Word(arg1, arg2, arg3 *W8) *W8 {
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
func FuncNoinline1Arg2ReturnPointer8Word(arg *W8) (*W8, *W8) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinline2Arg2ReturnPointer8Word(arg1, arg2 *W8) (*W8, *W8) {
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
func FuncNoinline3Arg2ReturnPointer8Word(arg1, arg2, arg3 *W8) (*W8, *W8) {
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
func FuncNoinline1Arg0ReturnPointer9Word(arg *W9) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func FuncNoinline2Arg0ReturnPointer9Word(arg1, arg2 *W9) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinline3Arg0ReturnPointer9Word(arg1, arg2, arg3 *W9) {
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
func FuncNoinline1Arg1ReturnPointer9Word(arg *W9) *W9 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func FuncNoinline2Arg1ReturnPointer9Word(arg1, arg2 *W9) *W9 {
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
func FuncNoinline3Arg1ReturnPointer9Word(arg1, arg2, arg3 *W9) *W9 {
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
func FuncNoinline1Arg2ReturnPointer9Word(arg *W9) (*W9, *W9) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinline2Arg2ReturnPointer9Word(arg1, arg2 *W9) (*W9, *W9) {
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
func FuncNoinline3Arg2ReturnPointer9Word(arg1, arg2, arg3 *W9) (*W9, *W9) {
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
func FuncNoinline1Arg0ReturnPointer10Word(arg *W10) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func FuncNoinline2Arg0ReturnPointer10Word(arg1, arg2 *W10) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func FuncNoinline3Arg0ReturnPointer10Word(arg1, arg2, arg3 *W10) {
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
func FuncNoinline1Arg1ReturnPointer10Word(arg *W10) *W10 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func FuncNoinline2Arg1ReturnPointer10Word(arg1, arg2 *W10) *W10 {
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
func FuncNoinline3Arg1ReturnPointer10Word(arg1, arg2, arg3 *W10) *W10 {
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
func FuncNoinline1Arg2ReturnPointer10Word(arg *W10) (*W10, *W10) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func FuncNoinline2Arg2ReturnPointer10Word(arg1, arg2 *W10) (*W10, *W10) {
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
func FuncNoinline3Arg2ReturnPointer10Word(arg1, arg2, arg3 *W10) (*W10, *W10) {
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
func (r W1) ValueReciverNoinline1Arg0ReturnValue1Word(arg W1) {
	_ = arg
}

//go:noinline
func (r W1) ValueReciverNoinline2Arg0ReturnValue1Word(arg1, arg2 W1) {
	_, _ = arg1, arg2
}

//go:noinline
func (r W1) ValueReciverNoinline3Arg0ReturnValue1Word(arg1, arg2, arg3 W1) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W1) ValueReciverNoinline1Arg1ReturnValue1Word(arg W1) W1 {
	_ = arg
	return arg
}

//go:noinline
func (r W1) ValueReciverNoinline2Arg1ReturnValue1Word(arg1, arg2 W1) W1 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W1) ValueReciverNoinline3Arg1ReturnValue1Word(arg1, arg2, arg3 W1) W1 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W1) ValueReciverNoinline1Arg2ReturnValue1Word(arg W1) (W1, W1) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r W1) ValueReciverNoinline2Arg2ReturnValue1Word(arg1, arg2 W1) (W1, W1) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W1) ValueReciverNoinline3Arg2ReturnValue1Word(arg1, arg2, arg3 W1) (W1, W1) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W2) ValueReciverNoinline1Arg0ReturnValue2Word(arg W2) {
	_ = arg
}

//go:noinline
func (r W2) ValueReciverNoinline2Arg0ReturnValue2Word(arg1, arg2 W2) {
	_, _ = arg1, arg2
}

//go:noinline
func (r W2) ValueReciverNoinline3Arg0ReturnValue2Word(arg1, arg2, arg3 W2) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W2) ValueReciverNoinline1Arg1ReturnValue2Word(arg W2) W2 {
	_ = arg
	return arg
}

//go:noinline
func (r W2) ValueReciverNoinline2Arg1ReturnValue2Word(arg1, arg2 W2) W2 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W2) ValueReciverNoinline3Arg1ReturnValue2Word(arg1, arg2, arg3 W2) W2 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W2) ValueReciverNoinline1Arg2ReturnValue2Word(arg W2) (W2, W2) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r W2) ValueReciverNoinline2Arg2ReturnValue2Word(arg1, arg2 W2) (W2, W2) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W2) ValueReciverNoinline3Arg2ReturnValue2Word(arg1, arg2, arg3 W2) (W2, W2) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W3) ValueReciverNoinline1Arg0ReturnValue3Word(arg W3) {
	_ = arg
}

//go:noinline
func (r W3) ValueReciverNoinline2Arg0ReturnValue3Word(arg1, arg2 W3) {
	_, _ = arg1, arg2
}

//go:noinline
func (r W3) ValueReciverNoinline3Arg0ReturnValue3Word(arg1, arg2, arg3 W3) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W3) ValueReciverNoinline1Arg1ReturnValue3Word(arg W3) W3 {
	_ = arg
	return arg
}

//go:noinline
func (r W3) ValueReciverNoinline2Arg1ReturnValue3Word(arg1, arg2 W3) W3 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W3) ValueReciverNoinline3Arg1ReturnValue3Word(arg1, arg2, arg3 W3) W3 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W3) ValueReciverNoinline1Arg2ReturnValue3Word(arg W3) (W3, W3) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r W3) ValueReciverNoinline2Arg2ReturnValue3Word(arg1, arg2 W3) (W3, W3) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W3) ValueReciverNoinline3Arg2ReturnValue3Word(arg1, arg2, arg3 W3) (W3, W3) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W4) ValueReciverNoinline1Arg0ReturnValue4Word(arg W4) {
	_ = arg
}

//go:noinline
func (r W4) ValueReciverNoinline2Arg0ReturnValue4Word(arg1, arg2 W4) {
	_, _ = arg1, arg2
}

//go:noinline
func (r W4) ValueReciverNoinline3Arg0ReturnValue4Word(arg1, arg2, arg3 W4) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W4) ValueReciverNoinline1Arg1ReturnValue4Word(arg W4) W4 {
	_ = arg
	return arg
}

//go:noinline
func (r W4) ValueReciverNoinline2Arg1ReturnValue4Word(arg1, arg2 W4) W4 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W4) ValueReciverNoinline3Arg1ReturnValue4Word(arg1, arg2, arg3 W4) W4 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W4) ValueReciverNoinline1Arg2ReturnValue4Word(arg W4) (W4, W4) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r W4) ValueReciverNoinline2Arg2ReturnValue4Word(arg1, arg2 W4) (W4, W4) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W4) ValueReciverNoinline3Arg2ReturnValue4Word(arg1, arg2, arg3 W4) (W4, W4) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W5) ValueReciverNoinline1Arg0ReturnValue5Word(arg W5) {
	_ = arg
}

//go:noinline
func (r W5) ValueReciverNoinline2Arg0ReturnValue5Word(arg1, arg2 W5) {
	_, _ = arg1, arg2
}

//go:noinline
func (r W5) ValueReciverNoinline3Arg0ReturnValue5Word(arg1, arg2, arg3 W5) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W5) ValueReciverNoinline1Arg1ReturnValue5Word(arg W5) W5 {
	_ = arg
	return arg
}

//go:noinline
func (r W5) ValueReciverNoinline2Arg1ReturnValue5Word(arg1, arg2 W5) W5 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W5) ValueReciverNoinline3Arg1ReturnValue5Word(arg1, arg2, arg3 W5) W5 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W5) ValueReciverNoinline1Arg2ReturnValue5Word(arg W5) (W5, W5) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r W5) ValueReciverNoinline2Arg2ReturnValue5Word(arg1, arg2 W5) (W5, W5) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W5) ValueReciverNoinline3Arg2ReturnValue5Word(arg1, arg2, arg3 W5) (W5, W5) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W6) ValueReciverNoinline1Arg0ReturnValue6Word(arg W6) {
	_ = arg
}

//go:noinline
func (r W6) ValueReciverNoinline2Arg0ReturnValue6Word(arg1, arg2 W6) {
	_, _ = arg1, arg2
}

//go:noinline
func (r W6) ValueReciverNoinline3Arg0ReturnValue6Word(arg1, arg2, arg3 W6) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W6) ValueReciverNoinline1Arg1ReturnValue6Word(arg W6) W6 {
	_ = arg
	return arg
}

//go:noinline
func (r W6) ValueReciverNoinline2Arg1ReturnValue6Word(arg1, arg2 W6) W6 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W6) ValueReciverNoinline3Arg1ReturnValue6Word(arg1, arg2, arg3 W6) W6 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W6) ValueReciverNoinline1Arg2ReturnValue6Word(arg W6) (W6, W6) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r W6) ValueReciverNoinline2Arg2ReturnValue6Word(arg1, arg2 W6) (W6, W6) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W6) ValueReciverNoinline3Arg2ReturnValue6Word(arg1, arg2, arg3 W6) (W6, W6) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W7) ValueReciverNoinline1Arg0ReturnValue7Word(arg W7) {
	_ = arg
}

//go:noinline
func (r W7) ValueReciverNoinline2Arg0ReturnValue7Word(arg1, arg2 W7) {
	_, _ = arg1, arg2
}

//go:noinline
func (r W7) ValueReciverNoinline3Arg0ReturnValue7Word(arg1, arg2, arg3 W7) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W7) ValueReciverNoinline1Arg1ReturnValue7Word(arg W7) W7 {
	_ = arg
	return arg
}

//go:noinline
func (r W7) ValueReciverNoinline2Arg1ReturnValue7Word(arg1, arg2 W7) W7 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W7) ValueReciverNoinline3Arg1ReturnValue7Word(arg1, arg2, arg3 W7) W7 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W7) ValueReciverNoinline1Arg2ReturnValue7Word(arg W7) (W7, W7) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r W7) ValueReciverNoinline2Arg2ReturnValue7Word(arg1, arg2 W7) (W7, W7) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W7) ValueReciverNoinline3Arg2ReturnValue7Word(arg1, arg2, arg3 W7) (W7, W7) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W8) ValueReciverNoinline1Arg0ReturnValue8Word(arg W8) {
	_ = arg
}

//go:noinline
func (r W8) ValueReciverNoinline2Arg0ReturnValue8Word(arg1, arg2 W8) {
	_, _ = arg1, arg2
}

//go:noinline
func (r W8) ValueReciverNoinline3Arg0ReturnValue8Word(arg1, arg2, arg3 W8) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W8) ValueReciverNoinline1Arg1ReturnValue8Word(arg W8) W8 {
	_ = arg
	return arg
}

//go:noinline
func (r W8) ValueReciverNoinline2Arg1ReturnValue8Word(arg1, arg2 W8) W8 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W8) ValueReciverNoinline3Arg1ReturnValue8Word(arg1, arg2, arg3 W8) W8 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W8) ValueReciverNoinline1Arg2ReturnValue8Word(arg W8) (W8, W8) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r W8) ValueReciverNoinline2Arg2ReturnValue8Word(arg1, arg2 W8) (W8, W8) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W8) ValueReciverNoinline3Arg2ReturnValue8Word(arg1, arg2, arg3 W8) (W8, W8) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W9) ValueReciverNoinline1Arg0ReturnValue9Word(arg W9) {
	_ = arg
}

//go:noinline
func (r W9) ValueReciverNoinline2Arg0ReturnValue9Word(arg1, arg2 W9) {
	_, _ = arg1, arg2
}

//go:noinline
func (r W9) ValueReciverNoinline3Arg0ReturnValue9Word(arg1, arg2, arg3 W9) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W9) ValueReciverNoinline1Arg1ReturnValue9Word(arg W9) W9 {
	_ = arg
	return arg
}

//go:noinline
func (r W9) ValueReciverNoinline2Arg1ReturnValue9Word(arg1, arg2 W9) W9 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W9) ValueReciverNoinline3Arg1ReturnValue9Word(arg1, arg2, arg3 W9) W9 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W9) ValueReciverNoinline1Arg2ReturnValue9Word(arg W9) (W9, W9) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r W9) ValueReciverNoinline2Arg2ReturnValue9Word(arg1, arg2 W9) (W9, W9) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W9) ValueReciverNoinline3Arg2ReturnValue9Word(arg1, arg2, arg3 W9) (W9, W9) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W10) ValueReciverNoinline1Arg0ReturnValue10Word(arg W10) {
	_ = arg
}

//go:noinline
func (r W10) ValueReciverNoinline2Arg0ReturnValue10Word(arg1, arg2 W10) {
	_, _ = arg1, arg2
}

//go:noinline
func (r W10) ValueReciverNoinline3Arg0ReturnValue10Word(arg1, arg2, arg3 W10) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r W10) ValueReciverNoinline1Arg1ReturnValue10Word(arg W10) W10 {
	_ = arg
	return arg
}

//go:noinline
func (r W10) ValueReciverNoinline2Arg1ReturnValue10Word(arg1, arg2 W10) W10 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r W10) ValueReciverNoinline3Arg1ReturnValue10Word(arg1, arg2, arg3 W10) W10 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r W10) ValueReciverNoinline1Arg2ReturnValue10Word(arg W10) (W10, W10) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r W10) ValueReciverNoinline2Arg2ReturnValue10Word(arg1, arg2 W10) (W10, W10) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r W10) ValueReciverNoinline3Arg2ReturnValue10Word(arg1, arg2, arg3 W10) (W10, W10) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r W1) ValueReciverNoinline1Arg0ReturnPointer1Word(arg *W1) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r W1) ValueReciverNoinline2Arg0ReturnPointer1Word(arg1, arg2 *W1) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r W1) ValueReciverNoinline3Arg0ReturnPointer1Word(arg1, arg2, arg3 *W1) {
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
func (r W1) ValueReciverNoinline1Arg1ReturnPointer1Word(arg *W1) *W1 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r W1) ValueReciverNoinline2Arg1ReturnPointer1Word(arg1, arg2 *W1) *W1 {
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
func (r W1) ValueReciverNoinline3Arg1ReturnPointer1Word(arg1, arg2, arg3 *W1) *W1 {
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
func (r W1) ValueReciverNoinline1Arg2ReturnPointer1Word(arg *W1) (*W1, *W1) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r W1) ValueReciverNoinline2Arg2ReturnPointer1Word(arg1, arg2 *W1) (*W1, *W1) {
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
func (r W1) ValueReciverNoinline3Arg2ReturnPointer1Word(arg1, arg2, arg3 *W1) (*W1, *W1) {
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
func (r W2) ValueReciverNoinline1Arg0ReturnPointer2Word(arg *W2) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r W2) ValueReciverNoinline2Arg0ReturnPointer2Word(arg1, arg2 *W2) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r W2) ValueReciverNoinline3Arg0ReturnPointer2Word(arg1, arg2, arg3 *W2) {
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
func (r W2) ValueReciverNoinline1Arg1ReturnPointer2Word(arg *W2) *W2 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r W2) ValueReciverNoinline2Arg1ReturnPointer2Word(arg1, arg2 *W2) *W2 {
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
func (r W2) ValueReciverNoinline3Arg1ReturnPointer2Word(arg1, arg2, arg3 *W2) *W2 {
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
func (r W2) ValueReciverNoinline1Arg2ReturnPointer2Word(arg *W2) (*W2, *W2) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r W2) ValueReciverNoinline2Arg2ReturnPointer2Word(arg1, arg2 *W2) (*W2, *W2) {
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
func (r W2) ValueReciverNoinline3Arg2ReturnPointer2Word(arg1, arg2, arg3 *W2) (*W2, *W2) {
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
func (r W3) ValueReciverNoinline1Arg0ReturnPointer3Word(arg *W3) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r W3) ValueReciverNoinline2Arg0ReturnPointer3Word(arg1, arg2 *W3) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r W3) ValueReciverNoinline3Arg0ReturnPointer3Word(arg1, arg2, arg3 *W3) {
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
func (r W3) ValueReciverNoinline1Arg1ReturnPointer3Word(arg *W3) *W3 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r W3) ValueReciverNoinline2Arg1ReturnPointer3Word(arg1, arg2 *W3) *W3 {
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
func (r W3) ValueReciverNoinline3Arg1ReturnPointer3Word(arg1, arg2, arg3 *W3) *W3 {
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
func (r W3) ValueReciverNoinline1Arg2ReturnPointer3Word(arg *W3) (*W3, *W3) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r W3) ValueReciverNoinline2Arg2ReturnPointer3Word(arg1, arg2 *W3) (*W3, *W3) {
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
func (r W3) ValueReciverNoinline3Arg2ReturnPointer3Word(arg1, arg2, arg3 *W3) (*W3, *W3) {
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
func (r W4) ValueReciverNoinline1Arg0ReturnPointer4Word(arg *W4) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r W4) ValueReciverNoinline2Arg0ReturnPointer4Word(arg1, arg2 *W4) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r W4) ValueReciverNoinline3Arg0ReturnPointer4Word(arg1, arg2, arg3 *W4) {
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
func (r W4) ValueReciverNoinline1Arg1ReturnPointer4Word(arg *W4) *W4 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r W4) ValueReciverNoinline2Arg1ReturnPointer4Word(arg1, arg2 *W4) *W4 {
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
func (r W4) ValueReciverNoinline3Arg1ReturnPointer4Word(arg1, arg2, arg3 *W4) *W4 {
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
func (r W4) ValueReciverNoinline1Arg2ReturnPointer4Word(arg *W4) (*W4, *W4) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r W4) ValueReciverNoinline2Arg2ReturnPointer4Word(arg1, arg2 *W4) (*W4, *W4) {
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
func (r W4) ValueReciverNoinline3Arg2ReturnPointer4Word(arg1, arg2, arg3 *W4) (*W4, *W4) {
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
func (r W5) ValueReciverNoinline1Arg0ReturnPointer5Word(arg *W5) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r W5) ValueReciverNoinline2Arg0ReturnPointer5Word(arg1, arg2 *W5) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r W5) ValueReciverNoinline3Arg0ReturnPointer5Word(arg1, arg2, arg3 *W5) {
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
func (r W5) ValueReciverNoinline1Arg1ReturnPointer5Word(arg *W5) *W5 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r W5) ValueReciverNoinline2Arg1ReturnPointer5Word(arg1, arg2 *W5) *W5 {
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
func (r W5) ValueReciverNoinline3Arg1ReturnPointer5Word(arg1, arg2, arg3 *W5) *W5 {
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
func (r W5) ValueReciverNoinline1Arg2ReturnPointer5Word(arg *W5) (*W5, *W5) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r W5) ValueReciverNoinline2Arg2ReturnPointer5Word(arg1, arg2 *W5) (*W5, *W5) {
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
func (r W5) ValueReciverNoinline3Arg2ReturnPointer5Word(arg1, arg2, arg3 *W5) (*W5, *W5) {
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
func (r W6) ValueReciverNoinline1Arg0ReturnPointer6Word(arg *W6) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r W6) ValueReciverNoinline2Arg0ReturnPointer6Word(arg1, arg2 *W6) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r W6) ValueReciverNoinline3Arg0ReturnPointer6Word(arg1, arg2, arg3 *W6) {
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
func (r W6) ValueReciverNoinline1Arg1ReturnPointer6Word(arg *W6) *W6 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r W6) ValueReciverNoinline2Arg1ReturnPointer6Word(arg1, arg2 *W6) *W6 {
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
func (r W6) ValueReciverNoinline3Arg1ReturnPointer6Word(arg1, arg2, arg3 *W6) *W6 {
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
func (r W6) ValueReciverNoinline1Arg2ReturnPointer6Word(arg *W6) (*W6, *W6) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r W6) ValueReciverNoinline2Arg2ReturnPointer6Word(arg1, arg2 *W6) (*W6, *W6) {
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
func (r W6) ValueReciverNoinline3Arg2ReturnPointer6Word(arg1, arg2, arg3 *W6) (*W6, *W6) {
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
func (r W7) ValueReciverNoinline1Arg0ReturnPointer7Word(arg *W7) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r W7) ValueReciverNoinline2Arg0ReturnPointer7Word(arg1, arg2 *W7) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r W7) ValueReciverNoinline3Arg0ReturnPointer7Word(arg1, arg2, arg3 *W7) {
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
func (r W7) ValueReciverNoinline1Arg1ReturnPointer7Word(arg *W7) *W7 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r W7) ValueReciverNoinline2Arg1ReturnPointer7Word(arg1, arg2 *W7) *W7 {
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
func (r W7) ValueReciverNoinline3Arg1ReturnPointer7Word(arg1, arg2, arg3 *W7) *W7 {
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
func (r W7) ValueReciverNoinline1Arg2ReturnPointer7Word(arg *W7) (*W7, *W7) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r W7) ValueReciverNoinline2Arg2ReturnPointer7Word(arg1, arg2 *W7) (*W7, *W7) {
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
func (r W7) ValueReciverNoinline3Arg2ReturnPointer7Word(arg1, arg2, arg3 *W7) (*W7, *W7) {
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
func (r W8) ValueReciverNoinline1Arg0ReturnPointer8Word(arg *W8) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r W8) ValueReciverNoinline2Arg0ReturnPointer8Word(arg1, arg2 *W8) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r W8) ValueReciverNoinline3Arg0ReturnPointer8Word(arg1, arg2, arg3 *W8) {
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
func (r W8) ValueReciverNoinline1Arg1ReturnPointer8Word(arg *W8) *W8 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r W8) ValueReciverNoinline2Arg1ReturnPointer8Word(arg1, arg2 *W8) *W8 {
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
func (r W8) ValueReciverNoinline3Arg1ReturnPointer8Word(arg1, arg2, arg3 *W8) *W8 {
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
func (r W8) ValueReciverNoinline1Arg2ReturnPointer8Word(arg *W8) (*W8, *W8) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r W8) ValueReciverNoinline2Arg2ReturnPointer8Word(arg1, arg2 *W8) (*W8, *W8) {
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
func (r W8) ValueReciverNoinline3Arg2ReturnPointer8Word(arg1, arg2, arg3 *W8) (*W8, *W8) {
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
func (r W9) ValueReciverNoinline1Arg0ReturnPointer9Word(arg *W9) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r W9) ValueReciverNoinline2Arg0ReturnPointer9Word(arg1, arg2 *W9) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r W9) ValueReciverNoinline3Arg0ReturnPointer9Word(arg1, arg2, arg3 *W9) {
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
func (r W9) ValueReciverNoinline1Arg1ReturnPointer9Word(arg *W9) *W9 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r W9) ValueReciverNoinline2Arg1ReturnPointer9Word(arg1, arg2 *W9) *W9 {
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
func (r W9) ValueReciverNoinline3Arg1ReturnPointer9Word(arg1, arg2, arg3 *W9) *W9 {
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
func (r W9) ValueReciverNoinline1Arg2ReturnPointer9Word(arg *W9) (*W9, *W9) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r W9) ValueReciverNoinline2Arg2ReturnPointer9Word(arg1, arg2 *W9) (*W9, *W9) {
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
func (r W9) ValueReciverNoinline3Arg2ReturnPointer9Word(arg1, arg2, arg3 *W9) (*W9, *W9) {
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
func (r W10) ValueReciverNoinline1Arg0ReturnPointer10Word(arg *W10) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r W10) ValueReciverNoinline2Arg0ReturnPointer10Word(arg1, arg2 *W10) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r W10) ValueReciverNoinline3Arg0ReturnPointer10Word(arg1, arg2, arg3 *W10) {
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
func (r W10) ValueReciverNoinline1Arg1ReturnPointer10Word(arg *W10) *W10 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r W10) ValueReciverNoinline2Arg1ReturnPointer10Word(arg1, arg2 *W10) *W10 {
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
func (r W10) ValueReciverNoinline3Arg1ReturnPointer10Word(arg1, arg2, arg3 *W10) *W10 {
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
func (r W10) ValueReciverNoinline1Arg2ReturnPointer10Word(arg *W10) (*W10, *W10) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r W10) ValueReciverNoinline2Arg2ReturnPointer10Word(arg1, arg2 *W10) (*W10, *W10) {
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
func (r W10) ValueReciverNoinline3Arg2ReturnPointer10Word(arg1, arg2, arg3 *W10) (*W10, *W10) {
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
func (r *W1) PointerReciverNoinline1Arg0ReturnValue1Word(arg W1) {
	_ = arg
}

//go:noinline
func (r *W1) PointerReciverNoinline2Arg0ReturnValue1Word(arg1, arg2 W1) {
	_, _ = arg1, arg2
}

//go:noinline
func (r *W1) PointerReciverNoinline3Arg0ReturnValue1Word(arg1, arg2, arg3 W1) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W1) PointerReciverNoinline1Arg1ReturnValue1Word(arg W1) W1 {
	_ = arg
	return arg
}

//go:noinline
func (r *W1) PointerReciverNoinline2Arg1ReturnValue1Word(arg1, arg2 W1) W1 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W1) PointerReciverNoinline3Arg1ReturnValue1Word(arg1, arg2, arg3 W1) W1 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W1) PointerReciverNoinline1Arg2ReturnValue1Word(arg W1) (W1, W1) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W1) PointerReciverNoinline2Arg2ReturnValue1Word(arg1, arg2 W1) (W1, W1) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W1) PointerReciverNoinline3Arg2ReturnValue1Word(arg1, arg2, arg3 W1) (W1, W1) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W2) PointerReciverNoinline1Arg0ReturnValue2Word(arg W2) {
	_ = arg
}

//go:noinline
func (r *W2) PointerReciverNoinline2Arg0ReturnValue2Word(arg1, arg2 W2) {
	_, _ = arg1, arg2
}

//go:noinline
func (r *W2) PointerReciverNoinline3Arg0ReturnValue2Word(arg1, arg2, arg3 W2) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W2) PointerReciverNoinline1Arg1ReturnValue2Word(arg W2) W2 {
	_ = arg
	return arg
}

//go:noinline
func (r *W2) PointerReciverNoinline2Arg1ReturnValue2Word(arg1, arg2 W2) W2 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W2) PointerReciverNoinline3Arg1ReturnValue2Word(arg1, arg2, arg3 W2) W2 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W2) PointerReciverNoinline1Arg2ReturnValue2Word(arg W2) (W2, W2) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W2) PointerReciverNoinline2Arg2ReturnValue2Word(arg1, arg2 W2) (W2, W2) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W2) PointerReciverNoinline3Arg2ReturnValue2Word(arg1, arg2, arg3 W2) (W2, W2) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W3) PointerReciverNoinline1Arg0ReturnValue3Word(arg W3) {
	_ = arg
}

//go:noinline
func (r *W3) PointerReciverNoinline2Arg0ReturnValue3Word(arg1, arg2 W3) {
	_, _ = arg1, arg2
}

//go:noinline
func (r *W3) PointerReciverNoinline3Arg0ReturnValue3Word(arg1, arg2, arg3 W3) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W3) PointerReciverNoinline1Arg1ReturnValue3Word(arg W3) W3 {
	_ = arg
	return arg
}

//go:noinline
func (r *W3) PointerReciverNoinline2Arg1ReturnValue3Word(arg1, arg2 W3) W3 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W3) PointerReciverNoinline3Arg1ReturnValue3Word(arg1, arg2, arg3 W3) W3 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W3) PointerReciverNoinline1Arg2ReturnValue3Word(arg W3) (W3, W3) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W3) PointerReciverNoinline2Arg2ReturnValue3Word(arg1, arg2 W3) (W3, W3) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W3) PointerReciverNoinline3Arg2ReturnValue3Word(arg1, arg2, arg3 W3) (W3, W3) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W4) PointerReciverNoinline1Arg0ReturnValue4Word(arg W4) {
	_ = arg
}

//go:noinline
func (r *W4) PointerReciverNoinline2Arg0ReturnValue4Word(arg1, arg2 W4) {
	_, _ = arg1, arg2
}

//go:noinline
func (r *W4) PointerReciverNoinline3Arg0ReturnValue4Word(arg1, arg2, arg3 W4) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W4) PointerReciverNoinline1Arg1ReturnValue4Word(arg W4) W4 {
	_ = arg
	return arg
}

//go:noinline
func (r *W4) PointerReciverNoinline2Arg1ReturnValue4Word(arg1, arg2 W4) W4 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W4) PointerReciverNoinline3Arg1ReturnValue4Word(arg1, arg2, arg3 W4) W4 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W4) PointerReciverNoinline1Arg2ReturnValue4Word(arg W4) (W4, W4) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W4) PointerReciverNoinline2Arg2ReturnValue4Word(arg1, arg2 W4) (W4, W4) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W4) PointerReciverNoinline3Arg2ReturnValue4Word(arg1, arg2, arg3 W4) (W4, W4) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W5) PointerReciverNoinline1Arg0ReturnValue5Word(arg W5) {
	_ = arg
}

//go:noinline
func (r *W5) PointerReciverNoinline2Arg0ReturnValue5Word(arg1, arg2 W5) {
	_, _ = arg1, arg2
}

//go:noinline
func (r *W5) PointerReciverNoinline3Arg0ReturnValue5Word(arg1, arg2, arg3 W5) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W5) PointerReciverNoinline1Arg1ReturnValue5Word(arg W5) W5 {
	_ = arg
	return arg
}

//go:noinline
func (r *W5) PointerReciverNoinline2Arg1ReturnValue5Word(arg1, arg2 W5) W5 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W5) PointerReciverNoinline3Arg1ReturnValue5Word(arg1, arg2, arg3 W5) W5 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W5) PointerReciverNoinline1Arg2ReturnValue5Word(arg W5) (W5, W5) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W5) PointerReciverNoinline2Arg2ReturnValue5Word(arg1, arg2 W5) (W5, W5) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W5) PointerReciverNoinline3Arg2ReturnValue5Word(arg1, arg2, arg3 W5) (W5, W5) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W6) PointerReciverNoinline1Arg0ReturnValue6Word(arg W6) {
	_ = arg
}

//go:noinline
func (r *W6) PointerReciverNoinline2Arg0ReturnValue6Word(arg1, arg2 W6) {
	_, _ = arg1, arg2
}

//go:noinline
func (r *W6) PointerReciverNoinline3Arg0ReturnValue6Word(arg1, arg2, arg3 W6) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W6) PointerReciverNoinline1Arg1ReturnValue6Word(arg W6) W6 {
	_ = arg
	return arg
}

//go:noinline
func (r *W6) PointerReciverNoinline2Arg1ReturnValue6Word(arg1, arg2 W6) W6 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W6) PointerReciverNoinline3Arg1ReturnValue6Word(arg1, arg2, arg3 W6) W6 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W6) PointerReciverNoinline1Arg2ReturnValue6Word(arg W6) (W6, W6) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W6) PointerReciverNoinline2Arg2ReturnValue6Word(arg1, arg2 W6) (W6, W6) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W6) PointerReciverNoinline3Arg2ReturnValue6Word(arg1, arg2, arg3 W6) (W6, W6) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W7) PointerReciverNoinline1Arg0ReturnValue7Word(arg W7) {
	_ = arg
}

//go:noinline
func (r *W7) PointerReciverNoinline2Arg0ReturnValue7Word(arg1, arg2 W7) {
	_, _ = arg1, arg2
}

//go:noinline
func (r *W7) PointerReciverNoinline3Arg0ReturnValue7Word(arg1, arg2, arg3 W7) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W7) PointerReciverNoinline1Arg1ReturnValue7Word(arg W7) W7 {
	_ = arg
	return arg
}

//go:noinline
func (r *W7) PointerReciverNoinline2Arg1ReturnValue7Word(arg1, arg2 W7) W7 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W7) PointerReciverNoinline3Arg1ReturnValue7Word(arg1, arg2, arg3 W7) W7 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W7) PointerReciverNoinline1Arg2ReturnValue7Word(arg W7) (W7, W7) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W7) PointerReciverNoinline2Arg2ReturnValue7Word(arg1, arg2 W7) (W7, W7) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W7) PointerReciverNoinline3Arg2ReturnValue7Word(arg1, arg2, arg3 W7) (W7, W7) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W8) PointerReciverNoinline1Arg0ReturnValue8Word(arg W8) {
	_ = arg
}

//go:noinline
func (r *W8) PointerReciverNoinline2Arg0ReturnValue8Word(arg1, arg2 W8) {
	_, _ = arg1, arg2
}

//go:noinline
func (r *W8) PointerReciverNoinline3Arg0ReturnValue8Word(arg1, arg2, arg3 W8) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W8) PointerReciverNoinline1Arg1ReturnValue8Word(arg W8) W8 {
	_ = arg
	return arg
}

//go:noinline
func (r *W8) PointerReciverNoinline2Arg1ReturnValue8Word(arg1, arg2 W8) W8 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W8) PointerReciverNoinline3Arg1ReturnValue8Word(arg1, arg2, arg3 W8) W8 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W8) PointerReciverNoinline1Arg2ReturnValue8Word(arg W8) (W8, W8) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W8) PointerReciverNoinline2Arg2ReturnValue8Word(arg1, arg2 W8) (W8, W8) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W8) PointerReciverNoinline3Arg2ReturnValue8Word(arg1, arg2, arg3 W8) (W8, W8) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W9) PointerReciverNoinline1Arg0ReturnValue9Word(arg W9) {
	_ = arg
}

//go:noinline
func (r *W9) PointerReciverNoinline2Arg0ReturnValue9Word(arg1, arg2 W9) {
	_, _ = arg1, arg2
}

//go:noinline
func (r *W9) PointerReciverNoinline3Arg0ReturnValue9Word(arg1, arg2, arg3 W9) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W9) PointerReciverNoinline1Arg1ReturnValue9Word(arg W9) W9 {
	_ = arg
	return arg
}

//go:noinline
func (r *W9) PointerReciverNoinline2Arg1ReturnValue9Word(arg1, arg2 W9) W9 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W9) PointerReciverNoinline3Arg1ReturnValue9Word(arg1, arg2, arg3 W9) W9 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W9) PointerReciverNoinline1Arg2ReturnValue9Word(arg W9) (W9, W9) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W9) PointerReciverNoinline2Arg2ReturnValue9Word(arg1, arg2 W9) (W9, W9) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W9) PointerReciverNoinline3Arg2ReturnValue9Word(arg1, arg2, arg3 W9) (W9, W9) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W10) PointerReciverNoinline1Arg0ReturnValue10Word(arg W10) {
	_ = arg
}

//go:noinline
func (r *W10) PointerReciverNoinline2Arg0ReturnValue10Word(arg1, arg2 W10) {
	_, _ = arg1, arg2
}

//go:noinline
func (r *W10) PointerReciverNoinline3Arg0ReturnValue10Word(arg1, arg2, arg3 W10) {
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func (r *W10) PointerReciverNoinline1Arg1ReturnValue10Word(arg W10) W10 {
	_ = arg
	return arg
}

//go:noinline
func (r *W10) PointerReciverNoinline2Arg1ReturnValue10Word(arg1, arg2 W10) W10 {
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func (r *W10) PointerReciverNoinline3Arg1ReturnValue10Word(arg1, arg2, arg3 W10) W10 {
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func (r *W10) PointerReciverNoinline1Arg2ReturnValue10Word(arg W10) (W10, W10) {
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W10) PointerReciverNoinline2Arg2ReturnValue10Word(arg1, arg2 W10) (W10, W10) {
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func (r *W10) PointerReciverNoinline3Arg2ReturnValue10Word(arg1, arg2, arg3 W10) (W10, W10) {
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

//go:noinline
func (r *W1) PointerReciverNoinline1Arg0ReturnPointer1Word(arg *W1) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r *W1) PointerReciverNoinline2Arg0ReturnPointer1Word(arg1, arg2 *W1) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r *W1) PointerReciverNoinline3Arg0ReturnPointer1Word(arg1, arg2, arg3 *W1) {
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
func (r *W1) PointerReciverNoinline1Arg1ReturnPointer1Word(arg *W1) *W1 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r *W1) PointerReciverNoinline2Arg1ReturnPointer1Word(arg1, arg2 *W1) *W1 {
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
func (r *W1) PointerReciverNoinline3Arg1ReturnPointer1Word(arg1, arg2, arg3 *W1) *W1 {
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
func (r *W1) PointerReciverNoinline1Arg2ReturnPointer1Word(arg *W1) (*W1, *W1) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W1) PointerReciverNoinline2Arg2ReturnPointer1Word(arg1, arg2 *W1) (*W1, *W1) {
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
func (r *W1) PointerReciverNoinline3Arg2ReturnPointer1Word(arg1, arg2, arg3 *W1) (*W1, *W1) {
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
func (r *W2) PointerReciverNoinline1Arg0ReturnPointer2Word(arg *W2) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r *W2) PointerReciverNoinline2Arg0ReturnPointer2Word(arg1, arg2 *W2) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r *W2) PointerReciverNoinline3Arg0ReturnPointer2Word(arg1, arg2, arg3 *W2) {
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
func (r *W2) PointerReciverNoinline1Arg1ReturnPointer2Word(arg *W2) *W2 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r *W2) PointerReciverNoinline2Arg1ReturnPointer2Word(arg1, arg2 *W2) *W2 {
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
func (r *W2) PointerReciverNoinline3Arg1ReturnPointer2Word(arg1, arg2, arg3 *W2) *W2 {
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
func (r *W2) PointerReciverNoinline1Arg2ReturnPointer2Word(arg *W2) (*W2, *W2) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W2) PointerReciverNoinline2Arg2ReturnPointer2Word(arg1, arg2 *W2) (*W2, *W2) {
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
func (r *W2) PointerReciverNoinline3Arg2ReturnPointer2Word(arg1, arg2, arg3 *W2) (*W2, *W2) {
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
func (r *W3) PointerReciverNoinline1Arg0ReturnPointer3Word(arg *W3) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r *W3) PointerReciverNoinline2Arg0ReturnPointer3Word(arg1, arg2 *W3) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r *W3) PointerReciverNoinline3Arg0ReturnPointer3Word(arg1, arg2, arg3 *W3) {
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
func (r *W3) PointerReciverNoinline1Arg1ReturnPointer3Word(arg *W3) *W3 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r *W3) PointerReciverNoinline2Arg1ReturnPointer3Word(arg1, arg2 *W3) *W3 {
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
func (r *W3) PointerReciverNoinline3Arg1ReturnPointer3Word(arg1, arg2, arg3 *W3) *W3 {
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
func (r *W3) PointerReciverNoinline1Arg2ReturnPointer3Word(arg *W3) (*W3, *W3) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W3) PointerReciverNoinline2Arg2ReturnPointer3Word(arg1, arg2 *W3) (*W3, *W3) {
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
func (r *W3) PointerReciverNoinline3Arg2ReturnPointer3Word(arg1, arg2, arg3 *W3) (*W3, *W3) {
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
func (r *W4) PointerReciverNoinline1Arg0ReturnPointer4Word(arg *W4) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r *W4) PointerReciverNoinline2Arg0ReturnPointer4Word(arg1, arg2 *W4) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r *W4) PointerReciverNoinline3Arg0ReturnPointer4Word(arg1, arg2, arg3 *W4) {
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
func (r *W4) PointerReciverNoinline1Arg1ReturnPointer4Word(arg *W4) *W4 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r *W4) PointerReciverNoinline2Arg1ReturnPointer4Word(arg1, arg2 *W4) *W4 {
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
func (r *W4) PointerReciverNoinline3Arg1ReturnPointer4Word(arg1, arg2, arg3 *W4) *W4 {
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
func (r *W4) PointerReciverNoinline1Arg2ReturnPointer4Word(arg *W4) (*W4, *W4) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W4) PointerReciverNoinline2Arg2ReturnPointer4Word(arg1, arg2 *W4) (*W4, *W4) {
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
func (r *W4) PointerReciverNoinline3Arg2ReturnPointer4Word(arg1, arg2, arg3 *W4) (*W4, *W4) {
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
func (r *W5) PointerReciverNoinline1Arg0ReturnPointer5Word(arg *W5) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r *W5) PointerReciverNoinline2Arg0ReturnPointer5Word(arg1, arg2 *W5) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r *W5) PointerReciverNoinline3Arg0ReturnPointer5Word(arg1, arg2, arg3 *W5) {
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
func (r *W5) PointerReciverNoinline1Arg1ReturnPointer5Word(arg *W5) *W5 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r *W5) PointerReciverNoinline2Arg1ReturnPointer5Word(arg1, arg2 *W5) *W5 {
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
func (r *W5) PointerReciverNoinline3Arg1ReturnPointer5Word(arg1, arg2, arg3 *W5) *W5 {
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
func (r *W5) PointerReciverNoinline1Arg2ReturnPointer5Word(arg *W5) (*W5, *W5) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W5) PointerReciverNoinline2Arg2ReturnPointer5Word(arg1, arg2 *W5) (*W5, *W5) {
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
func (r *W5) PointerReciverNoinline3Arg2ReturnPointer5Word(arg1, arg2, arg3 *W5) (*W5, *W5) {
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
func (r *W6) PointerReciverNoinline1Arg0ReturnPointer6Word(arg *W6) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r *W6) PointerReciverNoinline2Arg0ReturnPointer6Word(arg1, arg2 *W6) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r *W6) PointerReciverNoinline3Arg0ReturnPointer6Word(arg1, arg2, arg3 *W6) {
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
func (r *W6) PointerReciverNoinline1Arg1ReturnPointer6Word(arg *W6) *W6 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r *W6) PointerReciverNoinline2Arg1ReturnPointer6Word(arg1, arg2 *W6) *W6 {
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
func (r *W6) PointerReciverNoinline3Arg1ReturnPointer6Word(arg1, arg2, arg3 *W6) *W6 {
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
func (r *W6) PointerReciverNoinline1Arg2ReturnPointer6Word(arg *W6) (*W6, *W6) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W6) PointerReciverNoinline2Arg2ReturnPointer6Word(arg1, arg2 *W6) (*W6, *W6) {
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
func (r *W6) PointerReciverNoinline3Arg2ReturnPointer6Word(arg1, arg2, arg3 *W6) (*W6, *W6) {
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
func (r *W7) PointerReciverNoinline1Arg0ReturnPointer7Word(arg *W7) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r *W7) PointerReciverNoinline2Arg0ReturnPointer7Word(arg1, arg2 *W7) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r *W7) PointerReciverNoinline3Arg0ReturnPointer7Word(arg1, arg2, arg3 *W7) {
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
func (r *W7) PointerReciverNoinline1Arg1ReturnPointer7Word(arg *W7) *W7 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r *W7) PointerReciverNoinline2Arg1ReturnPointer7Word(arg1, arg2 *W7) *W7 {
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
func (r *W7) PointerReciverNoinline3Arg1ReturnPointer7Word(arg1, arg2, arg3 *W7) *W7 {
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
func (r *W7) PointerReciverNoinline1Arg2ReturnPointer7Word(arg *W7) (*W7, *W7) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W7) PointerReciverNoinline2Arg2ReturnPointer7Word(arg1, arg2 *W7) (*W7, *W7) {
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
func (r *W7) PointerReciverNoinline3Arg2ReturnPointer7Word(arg1, arg2, arg3 *W7) (*W7, *W7) {
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
func (r *W8) PointerReciverNoinline1Arg0ReturnPointer8Word(arg *W8) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r *W8) PointerReciverNoinline2Arg0ReturnPointer8Word(arg1, arg2 *W8) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r *W8) PointerReciverNoinline3Arg0ReturnPointer8Word(arg1, arg2, arg3 *W8) {
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
func (r *W8) PointerReciverNoinline1Arg1ReturnPointer8Word(arg *W8) *W8 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r *W8) PointerReciverNoinline2Arg1ReturnPointer8Word(arg1, arg2 *W8) *W8 {
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
func (r *W8) PointerReciverNoinline3Arg1ReturnPointer8Word(arg1, arg2, arg3 *W8) *W8 {
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
func (r *W8) PointerReciverNoinline1Arg2ReturnPointer8Word(arg *W8) (*W8, *W8) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W8) PointerReciverNoinline2Arg2ReturnPointer8Word(arg1, arg2 *W8) (*W8, *W8) {
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
func (r *W8) PointerReciverNoinline3Arg2ReturnPointer8Word(arg1, arg2, arg3 *W8) (*W8, *W8) {
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
func (r *W9) PointerReciverNoinline1Arg0ReturnPointer9Word(arg *W9) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r *W9) PointerReciverNoinline2Arg0ReturnPointer9Word(arg1, arg2 *W9) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r *W9) PointerReciverNoinline3Arg0ReturnPointer9Word(arg1, arg2, arg3 *W9) {
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
func (r *W9) PointerReciverNoinline1Arg1ReturnPointer9Word(arg *W9) *W9 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r *W9) PointerReciverNoinline2Arg1ReturnPointer9Word(arg1, arg2 *W9) *W9 {
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
func (r *W9) PointerReciverNoinline3Arg1ReturnPointer9Word(arg1, arg2, arg3 *W9) *W9 {
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
func (r *W9) PointerReciverNoinline1Arg2ReturnPointer9Word(arg *W9) (*W9, *W9) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W9) PointerReciverNoinline2Arg2ReturnPointer9Word(arg1, arg2 *W9) (*W9, *W9) {
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
func (r *W9) PointerReciverNoinline3Arg2ReturnPointer9Word(arg1, arg2, arg3 *W9) (*W9, *W9) {
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
func (r *W10) PointerReciverNoinline1Arg0ReturnPointer10Word(arg *W10) {
	if arg == nil {
		return
	}
	_ = arg
}

//go:noinline
func (r *W10) PointerReciverNoinline2Arg0ReturnPointer10Word(arg1, arg2 *W10) {
	if arg1 == nil {
		return
	}
	if arg2 == nil {
		return
	}
	_, _ = arg1, arg2
}

//go:noinline
func (r *W10) PointerReciverNoinline3Arg0ReturnPointer10Word(arg1, arg2, arg3 *W10) {
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
func (r *W10) PointerReciverNoinline1Arg1ReturnPointer10Word(arg *W10) *W10 {
	if arg == nil {
		return nil
	}
	_ = arg
	return arg
}

//go:noinline
func (r *W10) PointerReciverNoinline2Arg1ReturnPointer10Word(arg1, arg2 *W10) *W10 {
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
func (r *W10) PointerReciverNoinline3Arg1ReturnPointer10Word(arg1, arg2, arg3 *W10) *W10 {
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
func (r *W10) PointerReciverNoinline1Arg2ReturnPointer10Word(arg *W10) (*W10, *W10) {
	if arg == nil {
		return nil, nil
	}
	_ = arg
	return arg, arg
}

//go:noinline
func (r *W10) PointerReciverNoinline2Arg2ReturnPointer10Word(arg1, arg2 *W10) (*W10, *W10) {
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
func (r *W10) PointerReciverNoinline3Arg2ReturnPointer10Word(arg1, arg2, arg3 *W10) (*W10, *W10) {
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

func BenchmarkFuncNoinline1Arg0ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline1Arg0ReturnValue1Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg0ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W1 = W1{}
	var arg2 W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline2Arg0ReturnValue1Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg0ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W1 = W1{}
	var arg2 W1 = W1{}
	var arg3 W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline3Arg0ReturnValue1Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg1ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline1Arg1ReturnValue1Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg1ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W1 = W1{}
	var arg2 W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline2Arg1ReturnValue1Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg1ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W1 = W1{}
	var arg2 W1 = W1{}
	var arg3 W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline3Arg1ReturnValue1Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg2ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline1Arg2ReturnValue1Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg2ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W1 = W1{}
	var arg2 W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline2Arg2ReturnValue1Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg2ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W1 = W1{}
	var arg2 W1 = W1{}
	var arg3 W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline3Arg2ReturnValue1Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinline1Arg0ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline1Arg0ReturnValue2Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg0ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W2 = W2{}
	var arg2 W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline2Arg0ReturnValue2Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg0ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W2 = W2{}
	var arg2 W2 = W2{}
	var arg3 W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline3Arg0ReturnValue2Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg1ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline1Arg1ReturnValue2Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg1ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W2 = W2{}
	var arg2 W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline2Arg1ReturnValue2Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg1ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W2 = W2{}
	var arg2 W2 = W2{}
	var arg3 W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline3Arg1ReturnValue2Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg2ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline1Arg2ReturnValue2Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg2ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W2 = W2{}
	var arg2 W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline2Arg2ReturnValue2Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg2ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W2 = W2{}
	var arg2 W2 = W2{}
	var arg3 W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline3Arg2ReturnValue2Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinline1Arg0ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline1Arg0ReturnValue3Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg0ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W3 = W3{}
	var arg2 W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline2Arg0ReturnValue3Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg0ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W3 = W3{}
	var arg2 W3 = W3{}
	var arg3 W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline3Arg0ReturnValue3Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg1ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline1Arg1ReturnValue3Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg1ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W3 = W3{}
	var arg2 W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline2Arg1ReturnValue3Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg1ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W3 = W3{}
	var arg2 W3 = W3{}
	var arg3 W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline3Arg1ReturnValue3Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg2ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline1Arg2ReturnValue3Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg2ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W3 = W3{}
	var arg2 W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline2Arg2ReturnValue3Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg2ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W3 = W3{}
	var arg2 W3 = W3{}
	var arg3 W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline3Arg2ReturnValue3Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinline1Arg0ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline1Arg0ReturnValue4Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg0ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W4 = W4{}
	var arg2 W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline2Arg0ReturnValue4Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg0ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W4 = W4{}
	var arg2 W4 = W4{}
	var arg3 W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline3Arg0ReturnValue4Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg1ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline1Arg1ReturnValue4Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg1ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W4 = W4{}
	var arg2 W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline2Arg1ReturnValue4Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg1ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W4 = W4{}
	var arg2 W4 = W4{}
	var arg3 W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline3Arg1ReturnValue4Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg2ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline1Arg2ReturnValue4Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg2ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W4 = W4{}
	var arg2 W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline2Arg2ReturnValue4Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg2ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W4 = W4{}
	var arg2 W4 = W4{}
	var arg3 W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline3Arg2ReturnValue4Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinline1Arg0ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline1Arg0ReturnValue5Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg0ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W5 = W5{}
	var arg2 W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline2Arg0ReturnValue5Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg0ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W5 = W5{}
	var arg2 W5 = W5{}
	var arg3 W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline3Arg0ReturnValue5Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg1ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline1Arg1ReturnValue5Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg1ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W5 = W5{}
	var arg2 W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline2Arg1ReturnValue5Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg1ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W5 = W5{}
	var arg2 W5 = W5{}
	var arg3 W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline3Arg1ReturnValue5Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg2ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline1Arg2ReturnValue5Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg2ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W5 = W5{}
	var arg2 W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline2Arg2ReturnValue5Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg2ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W5 = W5{}
	var arg2 W5 = W5{}
	var arg3 W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline3Arg2ReturnValue5Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinline1Arg0ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline1Arg0ReturnValue6Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg0ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W6 = W6{}
	var arg2 W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline2Arg0ReturnValue6Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg0ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W6 = W6{}
	var arg2 W6 = W6{}
	var arg3 W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline3Arg0ReturnValue6Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg1ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline1Arg1ReturnValue6Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg1ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W6 = W6{}
	var arg2 W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline2Arg1ReturnValue6Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg1ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W6 = W6{}
	var arg2 W6 = W6{}
	var arg3 W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline3Arg1ReturnValue6Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg2ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline1Arg2ReturnValue6Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg2ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W6 = W6{}
	var arg2 W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline2Arg2ReturnValue6Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg2ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W6 = W6{}
	var arg2 W6 = W6{}
	var arg3 W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline3Arg2ReturnValue6Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinline1Arg0ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline1Arg0ReturnValue7Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg0ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W7 = W7{}
	var arg2 W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline2Arg0ReturnValue7Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg0ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W7 = W7{}
	var arg2 W7 = W7{}
	var arg3 W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline3Arg0ReturnValue7Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg1ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline1Arg1ReturnValue7Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg1ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W7 = W7{}
	var arg2 W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline2Arg1ReturnValue7Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg1ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W7 = W7{}
	var arg2 W7 = W7{}
	var arg3 W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline3Arg1ReturnValue7Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg2ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline1Arg2ReturnValue7Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg2ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W7 = W7{}
	var arg2 W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline2Arg2ReturnValue7Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg2ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W7 = W7{}
	var arg2 W7 = W7{}
	var arg3 W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline3Arg2ReturnValue7Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinline1Arg0ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline1Arg0ReturnValue8Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg0ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W8 = W8{}
	var arg2 W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline2Arg0ReturnValue8Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg0ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W8 = W8{}
	var arg2 W8 = W8{}
	var arg3 W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline3Arg0ReturnValue8Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg1ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline1Arg1ReturnValue8Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg1ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W8 = W8{}
	var arg2 W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline2Arg1ReturnValue8Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg1ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W8 = W8{}
	var arg2 W8 = W8{}
	var arg3 W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline3Arg1ReturnValue8Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg2ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline1Arg2ReturnValue8Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg2ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W8 = W8{}
	var arg2 W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline2Arg2ReturnValue8Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg2ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W8 = W8{}
	var arg2 W8 = W8{}
	var arg3 W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline3Arg2ReturnValue8Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinline1Arg0ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline1Arg0ReturnValue9Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg0ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W9 = W9{}
	var arg2 W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline2Arg0ReturnValue9Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg0ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W9 = W9{}
	var arg2 W9 = W9{}
	var arg3 W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline3Arg0ReturnValue9Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg1ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline1Arg1ReturnValue9Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg1ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W9 = W9{}
	var arg2 W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline2Arg1ReturnValue9Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg1ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W9 = W9{}
	var arg2 W9 = W9{}
	var arg3 W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline3Arg1ReturnValue9Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg2ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline1Arg2ReturnValue9Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg2ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W9 = W9{}
	var arg2 W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline2Arg2ReturnValue9Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg2ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W9 = W9{}
	var arg2 W9 = W9{}
	var arg3 W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline3Arg2ReturnValue9Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinline1Arg0ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline1Arg0ReturnValue10Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg0ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W10 = W10{}
	var arg2 W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline2Arg0ReturnValue10Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg0ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W10 = W10{}
	var arg2 W10 = W10{}
	var arg3 W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline3Arg0ReturnValue10Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg1ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline1Arg1ReturnValue10Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg1ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W10 = W10{}
	var arg2 W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline2Arg1ReturnValue10Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg1ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W10 = W10{}
	var arg2 W10 = W10{}
	var arg3 W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline3Arg1ReturnValue10Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg2ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline1Arg2ReturnValue10Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg2ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W10 = W10{}
	var arg2 W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline2Arg2ReturnValue10Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg2ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W10 = W10{}
	var arg2 W10 = W10{}
	var arg3 W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline3Arg2ReturnValue10Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg0ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline1Arg0ReturnPointer1Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg0ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W1 = new(W1)
	var arg2 *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline2Arg0ReturnPointer1Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg0ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W1 = new(W1)
	var arg2 *W1 = new(W1)
	var arg3 *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline3Arg0ReturnPointer1Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg1ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline1Arg1ReturnPointer1Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg1ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W1 = new(W1)
	var arg2 *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline2Arg1ReturnPointer1Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg1ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W1 = new(W1)
	var arg2 *W1 = new(W1)
	var arg3 *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline3Arg1ReturnPointer1Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg2ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline1Arg2ReturnPointer1Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg2ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W1 = new(W1)
	var arg2 *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline2Arg2ReturnPointer1Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg2ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W1 = new(W1)
	var arg2 *W1 = new(W1)
	var arg3 *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline3Arg2ReturnPointer1Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinline1Arg0ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline1Arg0ReturnPointer2Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg0ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W2 = new(W2)
	var arg2 *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline2Arg0ReturnPointer2Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg0ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W2 = new(W2)
	var arg2 *W2 = new(W2)
	var arg3 *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline3Arg0ReturnPointer2Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg1ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline1Arg1ReturnPointer2Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg1ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W2 = new(W2)
	var arg2 *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline2Arg1ReturnPointer2Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg1ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W2 = new(W2)
	var arg2 *W2 = new(W2)
	var arg3 *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline3Arg1ReturnPointer2Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg2ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline1Arg2ReturnPointer2Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg2ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W2 = new(W2)
	var arg2 *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline2Arg2ReturnPointer2Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg2ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W2 = new(W2)
	var arg2 *W2 = new(W2)
	var arg3 *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline3Arg2ReturnPointer2Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinline1Arg0ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline1Arg0ReturnPointer3Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg0ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W3 = new(W3)
	var arg2 *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline2Arg0ReturnPointer3Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg0ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W3 = new(W3)
	var arg2 *W3 = new(W3)
	var arg3 *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline3Arg0ReturnPointer3Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg1ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline1Arg1ReturnPointer3Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg1ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W3 = new(W3)
	var arg2 *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline2Arg1ReturnPointer3Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg1ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W3 = new(W3)
	var arg2 *W3 = new(W3)
	var arg3 *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline3Arg1ReturnPointer3Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg2ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline1Arg2ReturnPointer3Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg2ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W3 = new(W3)
	var arg2 *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline2Arg2ReturnPointer3Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg2ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W3 = new(W3)
	var arg2 *W3 = new(W3)
	var arg3 *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline3Arg2ReturnPointer3Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinline1Arg0ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline1Arg0ReturnPointer4Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg0ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W4 = new(W4)
	var arg2 *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline2Arg0ReturnPointer4Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg0ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W4 = new(W4)
	var arg2 *W4 = new(W4)
	var arg3 *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline3Arg0ReturnPointer4Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg1ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline1Arg1ReturnPointer4Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg1ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W4 = new(W4)
	var arg2 *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline2Arg1ReturnPointer4Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg1ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W4 = new(W4)
	var arg2 *W4 = new(W4)
	var arg3 *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline3Arg1ReturnPointer4Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg2ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline1Arg2ReturnPointer4Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg2ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W4 = new(W4)
	var arg2 *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline2Arg2ReturnPointer4Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg2ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W4 = new(W4)
	var arg2 *W4 = new(W4)
	var arg3 *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline3Arg2ReturnPointer4Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinline1Arg0ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline1Arg0ReturnPointer5Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg0ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W5 = new(W5)
	var arg2 *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline2Arg0ReturnPointer5Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg0ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W5 = new(W5)
	var arg2 *W5 = new(W5)
	var arg3 *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline3Arg0ReturnPointer5Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg1ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline1Arg1ReturnPointer5Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg1ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W5 = new(W5)
	var arg2 *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline2Arg1ReturnPointer5Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg1ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W5 = new(W5)
	var arg2 *W5 = new(W5)
	var arg3 *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline3Arg1ReturnPointer5Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg2ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline1Arg2ReturnPointer5Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg2ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W5 = new(W5)
	var arg2 *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline2Arg2ReturnPointer5Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg2ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W5 = new(W5)
	var arg2 *W5 = new(W5)
	var arg3 *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline3Arg2ReturnPointer5Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinline1Arg0ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline1Arg0ReturnPointer6Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg0ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W6 = new(W6)
	var arg2 *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline2Arg0ReturnPointer6Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg0ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W6 = new(W6)
	var arg2 *W6 = new(W6)
	var arg3 *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline3Arg0ReturnPointer6Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg1ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline1Arg1ReturnPointer6Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg1ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W6 = new(W6)
	var arg2 *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline2Arg1ReturnPointer6Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg1ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W6 = new(W6)
	var arg2 *W6 = new(W6)
	var arg3 *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline3Arg1ReturnPointer6Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg2ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline1Arg2ReturnPointer6Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg2ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W6 = new(W6)
	var arg2 *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline2Arg2ReturnPointer6Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg2ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W6 = new(W6)
	var arg2 *W6 = new(W6)
	var arg3 *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline3Arg2ReturnPointer6Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinline1Arg0ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline1Arg0ReturnPointer7Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg0ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W7 = new(W7)
	var arg2 *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline2Arg0ReturnPointer7Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg0ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W7 = new(W7)
	var arg2 *W7 = new(W7)
	var arg3 *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline3Arg0ReturnPointer7Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg1ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline1Arg1ReturnPointer7Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg1ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W7 = new(W7)
	var arg2 *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline2Arg1ReturnPointer7Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg1ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W7 = new(W7)
	var arg2 *W7 = new(W7)
	var arg3 *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline3Arg1ReturnPointer7Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg2ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline1Arg2ReturnPointer7Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg2ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W7 = new(W7)
	var arg2 *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline2Arg2ReturnPointer7Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg2ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W7 = new(W7)
	var arg2 *W7 = new(W7)
	var arg3 *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline3Arg2ReturnPointer7Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinline1Arg0ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline1Arg0ReturnPointer8Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg0ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W8 = new(W8)
	var arg2 *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline2Arg0ReturnPointer8Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg0ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W8 = new(W8)
	var arg2 *W8 = new(W8)
	var arg3 *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline3Arg0ReturnPointer8Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg1ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline1Arg1ReturnPointer8Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg1ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W8 = new(W8)
	var arg2 *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline2Arg1ReturnPointer8Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg1ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W8 = new(W8)
	var arg2 *W8 = new(W8)
	var arg3 *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline3Arg1ReturnPointer8Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg2ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline1Arg2ReturnPointer8Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg2ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W8 = new(W8)
	var arg2 *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline2Arg2ReturnPointer8Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg2ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W8 = new(W8)
	var arg2 *W8 = new(W8)
	var arg3 *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline3Arg2ReturnPointer8Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinline1Arg0ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline1Arg0ReturnPointer9Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg0ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W9 = new(W9)
	var arg2 *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline2Arg0ReturnPointer9Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg0ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W9 = new(W9)
	var arg2 *W9 = new(W9)
	var arg3 *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline3Arg0ReturnPointer9Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg1ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline1Arg1ReturnPointer9Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg1ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W9 = new(W9)
	var arg2 *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline2Arg1ReturnPointer9Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg1ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W9 = new(W9)
	var arg2 *W9 = new(W9)
	var arg3 *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline3Arg1ReturnPointer9Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg2ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline1Arg2ReturnPointer9Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg2ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W9 = new(W9)
	var arg2 *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline2Arg2ReturnPointer9Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg2ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W9 = new(W9)
	var arg2 *W9 = new(W9)
	var arg3 *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline3Arg2ReturnPointer9Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkFuncNoinline1Arg0ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline1Arg0ReturnPointer10Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg0ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W10 = new(W10)
	var arg2 *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline2Arg0ReturnPointer10Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg0ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W10 = new(W10)
	var arg2 *W10 = new(W10)
	var arg3 *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FuncNoinline3Arg0ReturnPointer10Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg1ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline1Arg1ReturnPointer10Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg1ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W10 = new(W10)
	var arg2 *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline2Arg1ReturnPointer10Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg1ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W10 = new(W10)
	var arg2 *W10 = new(W10)
	var arg3 *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = FuncNoinline3Arg1ReturnPointer10Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline1Arg2ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline1Arg2ReturnPointer10Word(arg)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline2Arg2ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W10 = new(W10)
	var arg2 *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline2Arg2ReturnPointer10Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkFuncNoinline3Arg2ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W10 = new(W10)
	var arg2 *W10 = new(W10)
	var arg3 *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FuncNoinline3Arg2ReturnPointer10Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg0ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg W1 = W1{}
	var r W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline1Arg0ReturnValue1Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg0ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W1 = W1{}
	var arg2 W1 = W1{}
	var r W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline2Arg0ReturnValue1Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg0ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W1 = W1{}
	var arg2 W1 = W1{}
	var arg3 W1 = W1{}
	var r W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline3Arg0ReturnValue1Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg1ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg W1 = W1{}
	var r W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline1Arg1ReturnValue1Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg1ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W1 = W1{}
	var arg2 W1 = W1{}
	var r W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline2Arg1ReturnValue1Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg1ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W1 = W1{}
	var arg2 W1 = W1{}
	var arg3 W1 = W1{}
	var r W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline3Arg1ReturnValue1Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg2ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg W1 = W1{}
	var r W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline1Arg2ReturnValue1Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg2ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W1 = W1{}
	var arg2 W1 = W1{}
	var r W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline2Arg2ReturnValue1Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg2ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W1 = W1{}
	var arg2 W1 = W1{}
	var arg3 W1 = W1{}
	var r W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline3Arg2ReturnValue1Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinline1Arg0ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg W2 = W2{}
	var r W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline1Arg0ReturnValue2Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg0ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W2 = W2{}
	var arg2 W2 = W2{}
	var r W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline2Arg0ReturnValue2Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg0ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W2 = W2{}
	var arg2 W2 = W2{}
	var arg3 W2 = W2{}
	var r W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline3Arg0ReturnValue2Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg1ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg W2 = W2{}
	var r W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline1Arg1ReturnValue2Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg1ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W2 = W2{}
	var arg2 W2 = W2{}
	var r W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline2Arg1ReturnValue2Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg1ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W2 = W2{}
	var arg2 W2 = W2{}
	var arg3 W2 = W2{}
	var r W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline3Arg1ReturnValue2Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg2ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg W2 = W2{}
	var r W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline1Arg2ReturnValue2Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg2ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W2 = W2{}
	var arg2 W2 = W2{}
	var r W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline2Arg2ReturnValue2Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg2ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W2 = W2{}
	var arg2 W2 = W2{}
	var arg3 W2 = W2{}
	var r W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline3Arg2ReturnValue2Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinline1Arg0ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg W3 = W3{}
	var r W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline1Arg0ReturnValue3Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg0ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W3 = W3{}
	var arg2 W3 = W3{}
	var r W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline2Arg0ReturnValue3Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg0ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W3 = W3{}
	var arg2 W3 = W3{}
	var arg3 W3 = W3{}
	var r W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline3Arg0ReturnValue3Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg1ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg W3 = W3{}
	var r W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline1Arg1ReturnValue3Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg1ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W3 = W3{}
	var arg2 W3 = W3{}
	var r W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline2Arg1ReturnValue3Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg1ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W3 = W3{}
	var arg2 W3 = W3{}
	var arg3 W3 = W3{}
	var r W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline3Arg1ReturnValue3Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg2ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg W3 = W3{}
	var r W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline1Arg2ReturnValue3Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg2ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W3 = W3{}
	var arg2 W3 = W3{}
	var r W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline2Arg2ReturnValue3Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg2ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W3 = W3{}
	var arg2 W3 = W3{}
	var arg3 W3 = W3{}
	var r W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline3Arg2ReturnValue3Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinline1Arg0ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg W4 = W4{}
	var r W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline1Arg0ReturnValue4Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg0ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W4 = W4{}
	var arg2 W4 = W4{}
	var r W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline2Arg0ReturnValue4Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg0ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W4 = W4{}
	var arg2 W4 = W4{}
	var arg3 W4 = W4{}
	var r W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline3Arg0ReturnValue4Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg1ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg W4 = W4{}
	var r W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline1Arg1ReturnValue4Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg1ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W4 = W4{}
	var arg2 W4 = W4{}
	var r W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline2Arg1ReturnValue4Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg1ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W4 = W4{}
	var arg2 W4 = W4{}
	var arg3 W4 = W4{}
	var r W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline3Arg1ReturnValue4Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg2ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg W4 = W4{}
	var r W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline1Arg2ReturnValue4Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg2ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W4 = W4{}
	var arg2 W4 = W4{}
	var r W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline2Arg2ReturnValue4Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg2ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W4 = W4{}
	var arg2 W4 = W4{}
	var arg3 W4 = W4{}
	var r W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline3Arg2ReturnValue4Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinline1Arg0ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg W5 = W5{}
	var r W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline1Arg0ReturnValue5Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg0ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W5 = W5{}
	var arg2 W5 = W5{}
	var r W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline2Arg0ReturnValue5Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg0ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W5 = W5{}
	var arg2 W5 = W5{}
	var arg3 W5 = W5{}
	var r W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline3Arg0ReturnValue5Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg1ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg W5 = W5{}
	var r W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline1Arg1ReturnValue5Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg1ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W5 = W5{}
	var arg2 W5 = W5{}
	var r W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline2Arg1ReturnValue5Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg1ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W5 = W5{}
	var arg2 W5 = W5{}
	var arg3 W5 = W5{}
	var r W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline3Arg1ReturnValue5Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg2ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg W5 = W5{}
	var r W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline1Arg2ReturnValue5Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg2ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W5 = W5{}
	var arg2 W5 = W5{}
	var r W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline2Arg2ReturnValue5Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg2ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W5 = W5{}
	var arg2 W5 = W5{}
	var arg3 W5 = W5{}
	var r W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline3Arg2ReturnValue5Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinline1Arg0ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg W6 = W6{}
	var r W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline1Arg0ReturnValue6Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg0ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W6 = W6{}
	var arg2 W6 = W6{}
	var r W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline2Arg0ReturnValue6Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg0ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W6 = W6{}
	var arg2 W6 = W6{}
	var arg3 W6 = W6{}
	var r W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline3Arg0ReturnValue6Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg1ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg W6 = W6{}
	var r W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline1Arg1ReturnValue6Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg1ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W6 = W6{}
	var arg2 W6 = W6{}
	var r W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline2Arg1ReturnValue6Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg1ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W6 = W6{}
	var arg2 W6 = W6{}
	var arg3 W6 = W6{}
	var r W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline3Arg1ReturnValue6Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg2ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg W6 = W6{}
	var r W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline1Arg2ReturnValue6Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg2ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W6 = W6{}
	var arg2 W6 = W6{}
	var r W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline2Arg2ReturnValue6Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg2ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W6 = W6{}
	var arg2 W6 = W6{}
	var arg3 W6 = W6{}
	var r W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline3Arg2ReturnValue6Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinline1Arg0ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg W7 = W7{}
	var r W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline1Arg0ReturnValue7Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg0ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W7 = W7{}
	var arg2 W7 = W7{}
	var r W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline2Arg0ReturnValue7Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg0ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W7 = W7{}
	var arg2 W7 = W7{}
	var arg3 W7 = W7{}
	var r W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline3Arg0ReturnValue7Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg1ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg W7 = W7{}
	var r W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline1Arg1ReturnValue7Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg1ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W7 = W7{}
	var arg2 W7 = W7{}
	var r W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline2Arg1ReturnValue7Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg1ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W7 = W7{}
	var arg2 W7 = W7{}
	var arg3 W7 = W7{}
	var r W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline3Arg1ReturnValue7Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg2ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg W7 = W7{}
	var r W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline1Arg2ReturnValue7Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg2ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W7 = W7{}
	var arg2 W7 = W7{}
	var r W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline2Arg2ReturnValue7Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg2ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W7 = W7{}
	var arg2 W7 = W7{}
	var arg3 W7 = W7{}
	var r W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline3Arg2ReturnValue7Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinline1Arg0ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg W8 = W8{}
	var r W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline1Arg0ReturnValue8Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg0ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W8 = W8{}
	var arg2 W8 = W8{}
	var r W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline2Arg0ReturnValue8Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg0ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W8 = W8{}
	var arg2 W8 = W8{}
	var arg3 W8 = W8{}
	var r W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline3Arg0ReturnValue8Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg1ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg W8 = W8{}
	var r W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline1Arg1ReturnValue8Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg1ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W8 = W8{}
	var arg2 W8 = W8{}
	var r W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline2Arg1ReturnValue8Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg1ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W8 = W8{}
	var arg2 W8 = W8{}
	var arg3 W8 = W8{}
	var r W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline3Arg1ReturnValue8Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg2ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg W8 = W8{}
	var r W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline1Arg2ReturnValue8Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg2ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W8 = W8{}
	var arg2 W8 = W8{}
	var r W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline2Arg2ReturnValue8Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg2ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W8 = W8{}
	var arg2 W8 = W8{}
	var arg3 W8 = W8{}
	var r W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline3Arg2ReturnValue8Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinline1Arg0ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg W9 = W9{}
	var r W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline1Arg0ReturnValue9Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg0ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W9 = W9{}
	var arg2 W9 = W9{}
	var r W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline2Arg0ReturnValue9Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg0ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W9 = W9{}
	var arg2 W9 = W9{}
	var arg3 W9 = W9{}
	var r W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline3Arg0ReturnValue9Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg1ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg W9 = W9{}
	var r W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline1Arg1ReturnValue9Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg1ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W9 = W9{}
	var arg2 W9 = W9{}
	var r W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline2Arg1ReturnValue9Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg1ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W9 = W9{}
	var arg2 W9 = W9{}
	var arg3 W9 = W9{}
	var r W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline3Arg1ReturnValue9Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg2ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg W9 = W9{}
	var r W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline1Arg2ReturnValue9Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg2ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W9 = W9{}
	var arg2 W9 = W9{}
	var r W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline2Arg2ReturnValue9Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg2ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W9 = W9{}
	var arg2 W9 = W9{}
	var arg3 W9 = W9{}
	var r W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline3Arg2ReturnValue9Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinline1Arg0ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg W10 = W10{}
	var r W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline1Arg0ReturnValue10Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg0ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W10 = W10{}
	var arg2 W10 = W10{}
	var r W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline2Arg0ReturnValue10Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg0ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W10 = W10{}
	var arg2 W10 = W10{}
	var arg3 W10 = W10{}
	var r W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline3Arg0ReturnValue10Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg1ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg W10 = W10{}
	var r W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline1Arg1ReturnValue10Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg1ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W10 = W10{}
	var arg2 W10 = W10{}
	var r W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline2Arg1ReturnValue10Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg1ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W10 = W10{}
	var arg2 W10 = W10{}
	var arg3 W10 = W10{}
	var r W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline3Arg1ReturnValue10Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg2ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg W10 = W10{}
	var r W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline1Arg2ReturnValue10Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg2ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W10 = W10{}
	var arg2 W10 = W10{}
	var r W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline2Arg2ReturnValue10Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg2ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W10 = W10{}
	var arg2 W10 = W10{}
	var arg3 W10 = W10{}
	var r W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline3Arg2ReturnValue10Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg0ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W1 = new(W1)
	var r W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline1Arg0ReturnPointer1Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg0ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W1 = new(W1)
	var arg2 *W1 = new(W1)
	var r W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline2Arg0ReturnPointer1Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg0ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W1 = new(W1)
	var arg2 *W1 = new(W1)
	var arg3 *W1 = new(W1)
	var r W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline3Arg0ReturnPointer1Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg1ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W1 = new(W1)
	var r W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline1Arg1ReturnPointer1Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg1ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W1 = new(W1)
	var arg2 *W1 = new(W1)
	var r W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline2Arg1ReturnPointer1Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg1ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W1 = new(W1)
	var arg2 *W1 = new(W1)
	var arg3 *W1 = new(W1)
	var r W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline3Arg1ReturnPointer1Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg2ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W1 = new(W1)
	var r W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline1Arg2ReturnPointer1Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg2ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W1 = new(W1)
	var arg2 *W1 = new(W1)
	var r W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline2Arg2ReturnPointer1Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg2ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W1 = new(W1)
	var arg2 *W1 = new(W1)
	var arg3 *W1 = new(W1)
	var r W1 = W1{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline3Arg2ReturnPointer1Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinline1Arg0ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W2 = new(W2)
	var r W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline1Arg0ReturnPointer2Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg0ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W2 = new(W2)
	var arg2 *W2 = new(W2)
	var r W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline2Arg0ReturnPointer2Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg0ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W2 = new(W2)
	var arg2 *W2 = new(W2)
	var arg3 *W2 = new(W2)
	var r W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline3Arg0ReturnPointer2Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg1ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W2 = new(W2)
	var r W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline1Arg1ReturnPointer2Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg1ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W2 = new(W2)
	var arg2 *W2 = new(W2)
	var r W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline2Arg1ReturnPointer2Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg1ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W2 = new(W2)
	var arg2 *W2 = new(W2)
	var arg3 *W2 = new(W2)
	var r W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline3Arg1ReturnPointer2Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg2ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W2 = new(W2)
	var r W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline1Arg2ReturnPointer2Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg2ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W2 = new(W2)
	var arg2 *W2 = new(W2)
	var r W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline2Arg2ReturnPointer2Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg2ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W2 = new(W2)
	var arg2 *W2 = new(W2)
	var arg3 *W2 = new(W2)
	var r W2 = W2{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline3Arg2ReturnPointer2Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinline1Arg0ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W3 = new(W3)
	var r W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline1Arg0ReturnPointer3Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg0ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W3 = new(W3)
	var arg2 *W3 = new(W3)
	var r W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline2Arg0ReturnPointer3Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg0ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W3 = new(W3)
	var arg2 *W3 = new(W3)
	var arg3 *W3 = new(W3)
	var r W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline3Arg0ReturnPointer3Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg1ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W3 = new(W3)
	var r W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline1Arg1ReturnPointer3Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg1ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W3 = new(W3)
	var arg2 *W3 = new(W3)
	var r W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline2Arg1ReturnPointer3Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg1ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W3 = new(W3)
	var arg2 *W3 = new(W3)
	var arg3 *W3 = new(W3)
	var r W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline3Arg1ReturnPointer3Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg2ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W3 = new(W3)
	var r W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline1Arg2ReturnPointer3Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg2ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W3 = new(W3)
	var arg2 *W3 = new(W3)
	var r W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline2Arg2ReturnPointer3Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg2ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W3 = new(W3)
	var arg2 *W3 = new(W3)
	var arg3 *W3 = new(W3)
	var r W3 = W3{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline3Arg2ReturnPointer3Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinline1Arg0ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W4 = new(W4)
	var r W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline1Arg0ReturnPointer4Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg0ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W4 = new(W4)
	var arg2 *W4 = new(W4)
	var r W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline2Arg0ReturnPointer4Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg0ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W4 = new(W4)
	var arg2 *W4 = new(W4)
	var arg3 *W4 = new(W4)
	var r W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline3Arg0ReturnPointer4Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg1ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W4 = new(W4)
	var r W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline1Arg1ReturnPointer4Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg1ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W4 = new(W4)
	var arg2 *W4 = new(W4)
	var r W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline2Arg1ReturnPointer4Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg1ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W4 = new(W4)
	var arg2 *W4 = new(W4)
	var arg3 *W4 = new(W4)
	var r W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline3Arg1ReturnPointer4Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg2ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W4 = new(W4)
	var r W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline1Arg2ReturnPointer4Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg2ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W4 = new(W4)
	var arg2 *W4 = new(W4)
	var r W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline2Arg2ReturnPointer4Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg2ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W4 = new(W4)
	var arg2 *W4 = new(W4)
	var arg3 *W4 = new(W4)
	var r W4 = W4{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline3Arg2ReturnPointer4Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinline1Arg0ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W5 = new(W5)
	var r W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline1Arg0ReturnPointer5Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg0ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W5 = new(W5)
	var arg2 *W5 = new(W5)
	var r W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline2Arg0ReturnPointer5Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg0ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W5 = new(W5)
	var arg2 *W5 = new(W5)
	var arg3 *W5 = new(W5)
	var r W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline3Arg0ReturnPointer5Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg1ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W5 = new(W5)
	var r W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline1Arg1ReturnPointer5Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg1ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W5 = new(W5)
	var arg2 *W5 = new(W5)
	var r W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline2Arg1ReturnPointer5Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg1ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W5 = new(W5)
	var arg2 *W5 = new(W5)
	var arg3 *W5 = new(W5)
	var r W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline3Arg1ReturnPointer5Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg2ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W5 = new(W5)
	var r W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline1Arg2ReturnPointer5Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg2ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W5 = new(W5)
	var arg2 *W5 = new(W5)
	var r W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline2Arg2ReturnPointer5Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg2ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W5 = new(W5)
	var arg2 *W5 = new(W5)
	var arg3 *W5 = new(W5)
	var r W5 = W5{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline3Arg2ReturnPointer5Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinline1Arg0ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W6 = new(W6)
	var r W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline1Arg0ReturnPointer6Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg0ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W6 = new(W6)
	var arg2 *W6 = new(W6)
	var r W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline2Arg0ReturnPointer6Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg0ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W6 = new(W6)
	var arg2 *W6 = new(W6)
	var arg3 *W6 = new(W6)
	var r W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline3Arg0ReturnPointer6Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg1ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W6 = new(W6)
	var r W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline1Arg1ReturnPointer6Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg1ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W6 = new(W6)
	var arg2 *W6 = new(W6)
	var r W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline2Arg1ReturnPointer6Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg1ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W6 = new(W6)
	var arg2 *W6 = new(W6)
	var arg3 *W6 = new(W6)
	var r W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline3Arg1ReturnPointer6Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg2ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W6 = new(W6)
	var r W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline1Arg2ReturnPointer6Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg2ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W6 = new(W6)
	var arg2 *W6 = new(W6)
	var r W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline2Arg2ReturnPointer6Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg2ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W6 = new(W6)
	var arg2 *W6 = new(W6)
	var arg3 *W6 = new(W6)
	var r W6 = W6{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline3Arg2ReturnPointer6Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinline1Arg0ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W7 = new(W7)
	var r W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline1Arg0ReturnPointer7Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg0ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W7 = new(W7)
	var arg2 *W7 = new(W7)
	var r W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline2Arg0ReturnPointer7Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg0ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W7 = new(W7)
	var arg2 *W7 = new(W7)
	var arg3 *W7 = new(W7)
	var r W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline3Arg0ReturnPointer7Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg1ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W7 = new(W7)
	var r W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline1Arg1ReturnPointer7Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg1ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W7 = new(W7)
	var arg2 *W7 = new(W7)
	var r W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline2Arg1ReturnPointer7Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg1ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W7 = new(W7)
	var arg2 *W7 = new(W7)
	var arg3 *W7 = new(W7)
	var r W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline3Arg1ReturnPointer7Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg2ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W7 = new(W7)
	var r W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline1Arg2ReturnPointer7Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg2ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W7 = new(W7)
	var arg2 *W7 = new(W7)
	var r W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline2Arg2ReturnPointer7Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg2ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W7 = new(W7)
	var arg2 *W7 = new(W7)
	var arg3 *W7 = new(W7)
	var r W7 = W7{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline3Arg2ReturnPointer7Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinline1Arg0ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W8 = new(W8)
	var r W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline1Arg0ReturnPointer8Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg0ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W8 = new(W8)
	var arg2 *W8 = new(W8)
	var r W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline2Arg0ReturnPointer8Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg0ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W8 = new(W8)
	var arg2 *W8 = new(W8)
	var arg3 *W8 = new(W8)
	var r W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline3Arg0ReturnPointer8Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg1ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W8 = new(W8)
	var r W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline1Arg1ReturnPointer8Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg1ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W8 = new(W8)
	var arg2 *W8 = new(W8)
	var r W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline2Arg1ReturnPointer8Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg1ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W8 = new(W8)
	var arg2 *W8 = new(W8)
	var arg3 *W8 = new(W8)
	var r W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline3Arg1ReturnPointer8Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg2ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W8 = new(W8)
	var r W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline1Arg2ReturnPointer8Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg2ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W8 = new(W8)
	var arg2 *W8 = new(W8)
	var r W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline2Arg2ReturnPointer8Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg2ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W8 = new(W8)
	var arg2 *W8 = new(W8)
	var arg3 *W8 = new(W8)
	var r W8 = W8{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline3Arg2ReturnPointer8Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinline1Arg0ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W9 = new(W9)
	var r W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline1Arg0ReturnPointer9Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg0ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W9 = new(W9)
	var arg2 *W9 = new(W9)
	var r W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline2Arg0ReturnPointer9Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg0ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W9 = new(W9)
	var arg2 *W9 = new(W9)
	var arg3 *W9 = new(W9)
	var r W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline3Arg0ReturnPointer9Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg1ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W9 = new(W9)
	var r W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline1Arg1ReturnPointer9Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg1ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W9 = new(W9)
	var arg2 *W9 = new(W9)
	var r W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline2Arg1ReturnPointer9Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg1ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W9 = new(W9)
	var arg2 *W9 = new(W9)
	var arg3 *W9 = new(W9)
	var r W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline3Arg1ReturnPointer9Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg2ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W9 = new(W9)
	var r W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline1Arg2ReturnPointer9Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg2ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W9 = new(W9)
	var arg2 *W9 = new(W9)
	var r W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline2Arg2ReturnPointer9Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg2ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W9 = new(W9)
	var arg2 *W9 = new(W9)
	var arg3 *W9 = new(W9)
	var r W9 = W9{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline3Arg2ReturnPointer9Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkValueReciverNoinline1Arg0ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W10 = new(W10)
	var r W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline1Arg0ReturnPointer10Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg0ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W10 = new(W10)
	var arg2 *W10 = new(W10)
	var r W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline2Arg0ReturnPointer10Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg0ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W10 = new(W10)
	var arg2 *W10 = new(W10)
	var arg3 *W10 = new(W10)
	var r W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.ValueReciverNoinline3Arg0ReturnPointer10Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg1ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W10 = new(W10)
	var r W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline1Arg1ReturnPointer10Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg1ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W10 = new(W10)
	var arg2 *W10 = new(W10)
	var r W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline2Arg1ReturnPointer10Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg1ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W10 = new(W10)
	var arg2 *W10 = new(W10)
	var arg3 *W10 = new(W10)
	var r W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.ValueReciverNoinline3Arg1ReturnPointer10Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline1Arg2ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W10 = new(W10)
	var r W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline1Arg2ReturnPointer10Word(arg)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline2Arg2ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W10 = new(W10)
	var arg2 *W10 = new(W10)
	var r W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline2Arg2ReturnPointer10Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkValueReciverNoinline3Arg2ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W10 = new(W10)
	var arg2 *W10 = new(W10)
	var arg3 *W10 = new(W10)
	var r W10 = W10{}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.ValueReciverNoinline3Arg2ReturnPointer10Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg0ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg W1 = W1{}
	var r *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline1Arg0ReturnValue1Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg0ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W1 = W1{}
	var arg2 W1 = W1{}
	var r *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline2Arg0ReturnValue1Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg0ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W1 = W1{}
	var arg2 W1 = W1{}
	var arg3 W1 = W1{}
	var r *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline3Arg0ReturnValue1Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg1ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg W1 = W1{}
	var r *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline1Arg1ReturnValue1Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg1ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W1 = W1{}
	var arg2 W1 = W1{}
	var r *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline2Arg1ReturnValue1Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg1ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W1 = W1{}
	var arg2 W1 = W1{}
	var arg3 W1 = W1{}
	var r *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline3Arg1ReturnValue1Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg2ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg W1 = W1{}
	var r *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline1Arg2ReturnValue1Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg2ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W1 = W1{}
	var arg2 W1 = W1{}
	var r *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline2Arg2ReturnValue1Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg2ReturnValue1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W1 = W1{}
	var arg2 W1 = W1{}
	var arg3 W1 = W1{}
	var r *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline3Arg2ReturnValue1Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinline1Arg0ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg W2 = W2{}
	var r *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline1Arg0ReturnValue2Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg0ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W2 = W2{}
	var arg2 W2 = W2{}
	var r *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline2Arg0ReturnValue2Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg0ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W2 = W2{}
	var arg2 W2 = W2{}
	var arg3 W2 = W2{}
	var r *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline3Arg0ReturnValue2Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg1ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg W2 = W2{}
	var r *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline1Arg1ReturnValue2Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg1ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W2 = W2{}
	var arg2 W2 = W2{}
	var r *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline2Arg1ReturnValue2Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg1ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W2 = W2{}
	var arg2 W2 = W2{}
	var arg3 W2 = W2{}
	var r *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline3Arg1ReturnValue2Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg2ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg W2 = W2{}
	var r *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline1Arg2ReturnValue2Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg2ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W2 = W2{}
	var arg2 W2 = W2{}
	var r *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline2Arg2ReturnValue2Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg2ReturnValue2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W2 = W2{}
	var arg2 W2 = W2{}
	var arg3 W2 = W2{}
	var r *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline3Arg2ReturnValue2Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinline1Arg0ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg W3 = W3{}
	var r *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline1Arg0ReturnValue3Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg0ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W3 = W3{}
	var arg2 W3 = W3{}
	var r *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline2Arg0ReturnValue3Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg0ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W3 = W3{}
	var arg2 W3 = W3{}
	var arg3 W3 = W3{}
	var r *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline3Arg0ReturnValue3Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg1ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg W3 = W3{}
	var r *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline1Arg1ReturnValue3Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg1ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W3 = W3{}
	var arg2 W3 = W3{}
	var r *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline2Arg1ReturnValue3Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg1ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W3 = W3{}
	var arg2 W3 = W3{}
	var arg3 W3 = W3{}
	var r *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline3Arg1ReturnValue3Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg2ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg W3 = W3{}
	var r *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline1Arg2ReturnValue3Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg2ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W3 = W3{}
	var arg2 W3 = W3{}
	var r *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline2Arg2ReturnValue3Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg2ReturnValue3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W3 = W3{}
	var arg2 W3 = W3{}
	var arg3 W3 = W3{}
	var r *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline3Arg2ReturnValue3Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinline1Arg0ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg W4 = W4{}
	var r *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline1Arg0ReturnValue4Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg0ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W4 = W4{}
	var arg2 W4 = W4{}
	var r *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline2Arg0ReturnValue4Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg0ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W4 = W4{}
	var arg2 W4 = W4{}
	var arg3 W4 = W4{}
	var r *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline3Arg0ReturnValue4Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg1ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg W4 = W4{}
	var r *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline1Arg1ReturnValue4Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg1ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W4 = W4{}
	var arg2 W4 = W4{}
	var r *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline2Arg1ReturnValue4Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg1ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W4 = W4{}
	var arg2 W4 = W4{}
	var arg3 W4 = W4{}
	var r *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline3Arg1ReturnValue4Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg2ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg W4 = W4{}
	var r *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline1Arg2ReturnValue4Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg2ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W4 = W4{}
	var arg2 W4 = W4{}
	var r *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline2Arg2ReturnValue4Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg2ReturnValue4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W4 = W4{}
	var arg2 W4 = W4{}
	var arg3 W4 = W4{}
	var r *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline3Arg2ReturnValue4Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinline1Arg0ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg W5 = W5{}
	var r *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline1Arg0ReturnValue5Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg0ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W5 = W5{}
	var arg2 W5 = W5{}
	var r *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline2Arg0ReturnValue5Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg0ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W5 = W5{}
	var arg2 W5 = W5{}
	var arg3 W5 = W5{}
	var r *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline3Arg0ReturnValue5Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg1ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg W5 = W5{}
	var r *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline1Arg1ReturnValue5Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg1ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W5 = W5{}
	var arg2 W5 = W5{}
	var r *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline2Arg1ReturnValue5Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg1ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W5 = W5{}
	var arg2 W5 = W5{}
	var arg3 W5 = W5{}
	var r *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline3Arg1ReturnValue5Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg2ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg W5 = W5{}
	var r *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline1Arg2ReturnValue5Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg2ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W5 = W5{}
	var arg2 W5 = W5{}
	var r *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline2Arg2ReturnValue5Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg2ReturnValue5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W5 = W5{}
	var arg2 W5 = W5{}
	var arg3 W5 = W5{}
	var r *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline3Arg2ReturnValue5Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinline1Arg0ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg W6 = W6{}
	var r *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline1Arg0ReturnValue6Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg0ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W6 = W6{}
	var arg2 W6 = W6{}
	var r *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline2Arg0ReturnValue6Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg0ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W6 = W6{}
	var arg2 W6 = W6{}
	var arg3 W6 = W6{}
	var r *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline3Arg0ReturnValue6Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg1ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg W6 = W6{}
	var r *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline1Arg1ReturnValue6Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg1ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W6 = W6{}
	var arg2 W6 = W6{}
	var r *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline2Arg1ReturnValue6Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg1ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W6 = W6{}
	var arg2 W6 = W6{}
	var arg3 W6 = W6{}
	var r *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline3Arg1ReturnValue6Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg2ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg W6 = W6{}
	var r *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline1Arg2ReturnValue6Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg2ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W6 = W6{}
	var arg2 W6 = W6{}
	var r *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline2Arg2ReturnValue6Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg2ReturnValue6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W6 = W6{}
	var arg2 W6 = W6{}
	var arg3 W6 = W6{}
	var r *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline3Arg2ReturnValue6Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinline1Arg0ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg W7 = W7{}
	var r *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline1Arg0ReturnValue7Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg0ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W7 = W7{}
	var arg2 W7 = W7{}
	var r *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline2Arg0ReturnValue7Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg0ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W7 = W7{}
	var arg2 W7 = W7{}
	var arg3 W7 = W7{}
	var r *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline3Arg0ReturnValue7Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg1ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg W7 = W7{}
	var r *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline1Arg1ReturnValue7Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg1ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W7 = W7{}
	var arg2 W7 = W7{}
	var r *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline2Arg1ReturnValue7Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg1ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W7 = W7{}
	var arg2 W7 = W7{}
	var arg3 W7 = W7{}
	var r *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline3Arg1ReturnValue7Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg2ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg W7 = W7{}
	var r *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline1Arg2ReturnValue7Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg2ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W7 = W7{}
	var arg2 W7 = W7{}
	var r *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline2Arg2ReturnValue7Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg2ReturnValue7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W7 = W7{}
	var arg2 W7 = W7{}
	var arg3 W7 = W7{}
	var r *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline3Arg2ReturnValue7Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinline1Arg0ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg W8 = W8{}
	var r *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline1Arg0ReturnValue8Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg0ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W8 = W8{}
	var arg2 W8 = W8{}
	var r *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline2Arg0ReturnValue8Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg0ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W8 = W8{}
	var arg2 W8 = W8{}
	var arg3 W8 = W8{}
	var r *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline3Arg0ReturnValue8Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg1ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg W8 = W8{}
	var r *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline1Arg1ReturnValue8Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg1ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W8 = W8{}
	var arg2 W8 = W8{}
	var r *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline2Arg1ReturnValue8Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg1ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W8 = W8{}
	var arg2 W8 = W8{}
	var arg3 W8 = W8{}
	var r *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline3Arg1ReturnValue8Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg2ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg W8 = W8{}
	var r *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline1Arg2ReturnValue8Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg2ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W8 = W8{}
	var arg2 W8 = W8{}
	var r *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline2Arg2ReturnValue8Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg2ReturnValue8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W8 = W8{}
	var arg2 W8 = W8{}
	var arg3 W8 = W8{}
	var r *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline3Arg2ReturnValue8Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinline1Arg0ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg W9 = W9{}
	var r *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline1Arg0ReturnValue9Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg0ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W9 = W9{}
	var arg2 W9 = W9{}
	var r *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline2Arg0ReturnValue9Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg0ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W9 = W9{}
	var arg2 W9 = W9{}
	var arg3 W9 = W9{}
	var r *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline3Arg0ReturnValue9Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg1ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg W9 = W9{}
	var r *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline1Arg1ReturnValue9Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg1ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W9 = W9{}
	var arg2 W9 = W9{}
	var r *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline2Arg1ReturnValue9Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg1ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W9 = W9{}
	var arg2 W9 = W9{}
	var arg3 W9 = W9{}
	var r *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline3Arg1ReturnValue9Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg2ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg W9 = W9{}
	var r *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline1Arg2ReturnValue9Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg2ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W9 = W9{}
	var arg2 W9 = W9{}
	var r *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline2Arg2ReturnValue9Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg2ReturnValue9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W9 = W9{}
	var arg2 W9 = W9{}
	var arg3 W9 = W9{}
	var r *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline3Arg2ReturnValue9Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinline1Arg0ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg W10 = W10{}
	var r *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline1Arg0ReturnValue10Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg0ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W10 = W10{}
	var arg2 W10 = W10{}
	var r *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline2Arg0ReturnValue10Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg0ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W10 = W10{}
	var arg2 W10 = W10{}
	var arg3 W10 = W10{}
	var r *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline3Arg0ReturnValue10Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg1ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg W10 = W10{}
	var r *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline1Arg1ReturnValue10Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg1ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W10 = W10{}
	var arg2 W10 = W10{}
	var r *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline2Arg1ReturnValue10Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg1ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W10 = W10{}
	var arg2 W10 = W10{}
	var arg3 W10 = W10{}
	var r *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline3Arg1ReturnValue10Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg2ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg W10 = W10{}
	var r *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline1Arg2ReturnValue10Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg2ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W10 = W10{}
	var arg2 W10 = W10{}
	var r *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline2Arg2ReturnValue10Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg2ReturnValue10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 W10 = W10{}
	var arg2 W10 = W10{}
	var arg3 W10 = W10{}
	var r *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline3Arg2ReturnValue10Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg0ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W1 = new(W1)
	var r *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline1Arg0ReturnPointer1Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg0ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W1 = new(W1)
	var arg2 *W1 = new(W1)
	var r *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline2Arg0ReturnPointer1Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg0ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W1 = new(W1)
	var arg2 *W1 = new(W1)
	var arg3 *W1 = new(W1)
	var r *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline3Arg0ReturnPointer1Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg1ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W1 = new(W1)
	var r *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline1Arg1ReturnPointer1Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg1ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W1 = new(W1)
	var arg2 *W1 = new(W1)
	var r *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline2Arg1ReturnPointer1Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg1ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W1 = new(W1)
	var arg2 *W1 = new(W1)
	var arg3 *W1 = new(W1)
	var r *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline3Arg1ReturnPointer1Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg2ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W1 = new(W1)
	var r *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline1Arg2ReturnPointer1Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg2ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W1 = new(W1)
	var arg2 *W1 = new(W1)
	var r *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline2Arg2ReturnPointer1Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg2ReturnPointer1Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W1 = new(W1)
	var arg2 *W1 = new(W1)
	var arg3 *W1 = new(W1)
	var r *W1 = new(W1)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline3Arg2ReturnPointer1Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinline1Arg0ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W2 = new(W2)
	var r *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline1Arg0ReturnPointer2Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg0ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W2 = new(W2)
	var arg2 *W2 = new(W2)
	var r *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline2Arg0ReturnPointer2Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg0ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W2 = new(W2)
	var arg2 *W2 = new(W2)
	var arg3 *W2 = new(W2)
	var r *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline3Arg0ReturnPointer2Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg1ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W2 = new(W2)
	var r *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline1Arg1ReturnPointer2Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg1ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W2 = new(W2)
	var arg2 *W2 = new(W2)
	var r *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline2Arg1ReturnPointer2Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg1ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W2 = new(W2)
	var arg2 *W2 = new(W2)
	var arg3 *W2 = new(W2)
	var r *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline3Arg1ReturnPointer2Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg2ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W2 = new(W2)
	var r *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline1Arg2ReturnPointer2Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg2ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W2 = new(W2)
	var arg2 *W2 = new(W2)
	var r *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline2Arg2ReturnPointer2Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg2ReturnPointer2Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W2 = new(W2)
	var arg2 *W2 = new(W2)
	var arg3 *W2 = new(W2)
	var r *W2 = new(W2)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline3Arg2ReturnPointer2Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinline1Arg0ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W3 = new(W3)
	var r *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline1Arg0ReturnPointer3Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg0ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W3 = new(W3)
	var arg2 *W3 = new(W3)
	var r *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline2Arg0ReturnPointer3Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg0ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W3 = new(W3)
	var arg2 *W3 = new(W3)
	var arg3 *W3 = new(W3)
	var r *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline3Arg0ReturnPointer3Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg1ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W3 = new(W3)
	var r *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline1Arg1ReturnPointer3Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg1ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W3 = new(W3)
	var arg2 *W3 = new(W3)
	var r *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline2Arg1ReturnPointer3Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg1ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W3 = new(W3)
	var arg2 *W3 = new(W3)
	var arg3 *W3 = new(W3)
	var r *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline3Arg1ReturnPointer3Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg2ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W3 = new(W3)
	var r *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline1Arg2ReturnPointer3Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg2ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W3 = new(W3)
	var arg2 *W3 = new(W3)
	var r *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline2Arg2ReturnPointer3Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg2ReturnPointer3Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W3 = new(W3)
	var arg2 *W3 = new(W3)
	var arg3 *W3 = new(W3)
	var r *W3 = new(W3)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline3Arg2ReturnPointer3Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinline1Arg0ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W4 = new(W4)
	var r *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline1Arg0ReturnPointer4Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg0ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W4 = new(W4)
	var arg2 *W4 = new(W4)
	var r *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline2Arg0ReturnPointer4Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg0ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W4 = new(W4)
	var arg2 *W4 = new(W4)
	var arg3 *W4 = new(W4)
	var r *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline3Arg0ReturnPointer4Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg1ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W4 = new(W4)
	var r *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline1Arg1ReturnPointer4Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg1ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W4 = new(W4)
	var arg2 *W4 = new(W4)
	var r *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline2Arg1ReturnPointer4Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg1ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W4 = new(W4)
	var arg2 *W4 = new(W4)
	var arg3 *W4 = new(W4)
	var r *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline3Arg1ReturnPointer4Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg2ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W4 = new(W4)
	var r *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline1Arg2ReturnPointer4Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg2ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W4 = new(W4)
	var arg2 *W4 = new(W4)
	var r *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline2Arg2ReturnPointer4Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg2ReturnPointer4Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W4 = new(W4)
	var arg2 *W4 = new(W4)
	var arg3 *W4 = new(W4)
	var r *W4 = new(W4)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline3Arg2ReturnPointer4Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinline1Arg0ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W5 = new(W5)
	var r *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline1Arg0ReturnPointer5Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg0ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W5 = new(W5)
	var arg2 *W5 = new(W5)
	var r *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline2Arg0ReturnPointer5Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg0ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W5 = new(W5)
	var arg2 *W5 = new(W5)
	var arg3 *W5 = new(W5)
	var r *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline3Arg0ReturnPointer5Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg1ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W5 = new(W5)
	var r *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline1Arg1ReturnPointer5Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg1ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W5 = new(W5)
	var arg2 *W5 = new(W5)
	var r *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline2Arg1ReturnPointer5Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg1ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W5 = new(W5)
	var arg2 *W5 = new(W5)
	var arg3 *W5 = new(W5)
	var r *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline3Arg1ReturnPointer5Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg2ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W5 = new(W5)
	var r *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline1Arg2ReturnPointer5Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg2ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W5 = new(W5)
	var arg2 *W5 = new(W5)
	var r *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline2Arg2ReturnPointer5Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg2ReturnPointer5Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W5 = new(W5)
	var arg2 *W5 = new(W5)
	var arg3 *W5 = new(W5)
	var r *W5 = new(W5)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline3Arg2ReturnPointer5Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinline1Arg0ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W6 = new(W6)
	var r *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline1Arg0ReturnPointer6Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg0ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W6 = new(W6)
	var arg2 *W6 = new(W6)
	var r *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline2Arg0ReturnPointer6Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg0ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W6 = new(W6)
	var arg2 *W6 = new(W6)
	var arg3 *W6 = new(W6)
	var r *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline3Arg0ReturnPointer6Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg1ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W6 = new(W6)
	var r *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline1Arg1ReturnPointer6Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg1ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W6 = new(W6)
	var arg2 *W6 = new(W6)
	var r *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline2Arg1ReturnPointer6Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg1ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W6 = new(W6)
	var arg2 *W6 = new(W6)
	var arg3 *W6 = new(W6)
	var r *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline3Arg1ReturnPointer6Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg2ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W6 = new(W6)
	var r *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline1Arg2ReturnPointer6Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg2ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W6 = new(W6)
	var arg2 *W6 = new(W6)
	var r *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline2Arg2ReturnPointer6Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg2ReturnPointer6Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W6 = new(W6)
	var arg2 *W6 = new(W6)
	var arg3 *W6 = new(W6)
	var r *W6 = new(W6)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline3Arg2ReturnPointer6Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinline1Arg0ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W7 = new(W7)
	var r *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline1Arg0ReturnPointer7Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg0ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W7 = new(W7)
	var arg2 *W7 = new(W7)
	var r *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline2Arg0ReturnPointer7Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg0ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W7 = new(W7)
	var arg2 *W7 = new(W7)
	var arg3 *W7 = new(W7)
	var r *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline3Arg0ReturnPointer7Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg1ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W7 = new(W7)
	var r *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline1Arg1ReturnPointer7Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg1ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W7 = new(W7)
	var arg2 *W7 = new(W7)
	var r *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline2Arg1ReturnPointer7Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg1ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W7 = new(W7)
	var arg2 *W7 = new(W7)
	var arg3 *W7 = new(W7)
	var r *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline3Arg1ReturnPointer7Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg2ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W7 = new(W7)
	var r *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline1Arg2ReturnPointer7Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg2ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W7 = new(W7)
	var arg2 *W7 = new(W7)
	var r *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline2Arg2ReturnPointer7Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg2ReturnPointer7Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W7 = new(W7)
	var arg2 *W7 = new(W7)
	var arg3 *W7 = new(W7)
	var r *W7 = new(W7)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline3Arg2ReturnPointer7Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinline1Arg0ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W8 = new(W8)
	var r *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline1Arg0ReturnPointer8Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg0ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W8 = new(W8)
	var arg2 *W8 = new(W8)
	var r *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline2Arg0ReturnPointer8Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg0ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W8 = new(W8)
	var arg2 *W8 = new(W8)
	var arg3 *W8 = new(W8)
	var r *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline3Arg0ReturnPointer8Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg1ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W8 = new(W8)
	var r *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline1Arg1ReturnPointer8Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg1ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W8 = new(W8)
	var arg2 *W8 = new(W8)
	var r *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline2Arg1ReturnPointer8Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg1ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W8 = new(W8)
	var arg2 *W8 = new(W8)
	var arg3 *W8 = new(W8)
	var r *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline3Arg1ReturnPointer8Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg2ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W8 = new(W8)
	var r *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline1Arg2ReturnPointer8Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg2ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W8 = new(W8)
	var arg2 *W8 = new(W8)
	var r *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline2Arg2ReturnPointer8Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg2ReturnPointer8Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W8 = new(W8)
	var arg2 *W8 = new(W8)
	var arg3 *W8 = new(W8)
	var r *W8 = new(W8)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline3Arg2ReturnPointer8Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinline1Arg0ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W9 = new(W9)
	var r *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline1Arg0ReturnPointer9Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg0ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W9 = new(W9)
	var arg2 *W9 = new(W9)
	var r *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline2Arg0ReturnPointer9Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg0ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W9 = new(W9)
	var arg2 *W9 = new(W9)
	var arg3 *W9 = new(W9)
	var r *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline3Arg0ReturnPointer9Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg1ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W9 = new(W9)
	var r *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline1Arg1ReturnPointer9Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg1ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W9 = new(W9)
	var arg2 *W9 = new(W9)
	var r *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline2Arg1ReturnPointer9Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg1ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W9 = new(W9)
	var arg2 *W9 = new(W9)
	var arg3 *W9 = new(W9)
	var r *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline3Arg1ReturnPointer9Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg2ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W9 = new(W9)
	var r *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline1Arg2ReturnPointer9Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg2ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W9 = new(W9)
	var arg2 *W9 = new(W9)
	var r *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline2Arg2ReturnPointer9Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg2ReturnPointer9Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W9 = new(W9)
	var arg2 *W9 = new(W9)
	var arg3 *W9 = new(W9)
	var r *W9 = new(W9)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline3Arg2ReturnPointer9Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
func BenchmarkPointerReciverNoinline1Arg0ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W10 = new(W10)
	var r *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline1Arg0ReturnPointer10Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg0ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W10 = new(W10)
	var arg2 *W10 = new(W10)
	var r *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline2Arg0ReturnPointer10Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg0ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W10 = new(W10)
	var arg2 *W10 = new(W10)
	var arg3 *W10 = new(W10)
	var r *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.PointerReciverNoinline3Arg0ReturnPointer10Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg1ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W10 = new(W10)
	var r *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline1Arg1ReturnPointer10Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg1ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W10 = new(W10)
	var arg2 *W10 = new(W10)
	var r *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline2Arg1ReturnPointer10Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg1ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W10 = new(W10)
	var arg2 *W10 = new(W10)
	var arg3 *W10 = new(W10)
	var r *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = r.PointerReciverNoinline3Arg1ReturnPointer10Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline1Arg2ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg *W10 = new(W10)
	var r *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline1Arg2ReturnPointer10Word(arg)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline2Arg2ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W10 = new(W10)
	var arg2 *W10 = new(W10)
	var r *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline2Arg2ReturnPointer10Word(arg1, arg2)
	}
	b.StopTimer()
}

func BenchmarkPointerReciverNoinline3Arg2ReturnPointer10Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 *W10 = new(W10)
	var arg2 *W10 = new(W10)
	var arg3 *W10 = new(W10)
	var r *W10 = new(W10)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = r.PointerReciverNoinline3Arg2ReturnPointer10Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
