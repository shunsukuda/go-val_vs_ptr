#!/bin/bash
result_file="bench_result.txt"
rm -rf $result_file
cmd="go test -bench FuncNoinline1Arg0Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench ValueReciverNoinline1Arg0Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench PointerReciverNoinline1Arg0Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench FuncNoinline2Arg0Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench ValueReciverNoinline2Arg0Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench PointerReciverNoinline2Arg0Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench FuncNoinline3Arg0Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench ValueReciverNoinline3Arg0Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench PointerReciverNoinline3Arg0Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench FuncNoinline1Arg1Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench ValueReciverNoinline1Arg1Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench PointerReciverNoinline1Arg1Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench FuncNoinline2Arg1Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench ValueReciverNoinline2Arg1Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench PointerReciverNoinline2Arg1Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench FuncNoinline3Arg1Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench ValueReciverNoinline3Arg1Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench PointerReciverNoinline3Arg1Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench FuncNoinline1Arg2Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench ValueReciverNoinline1Arg2Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench PointerReciverNoinline1Arg2Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench FuncNoinline2Arg2Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench ValueReciverNoinline2Arg2Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench PointerReciverNoinline2Arg2Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench FuncNoinline3Arg2Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench ValueReciverNoinline3Arg2Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench PointerReciverNoinline3Arg2Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 