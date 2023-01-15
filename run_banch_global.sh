#!/bin/bash
result_file="bench_global_result.txt"
rm -rf $result_file
cmd="go test -bench FuncNoinlineGlobal1Arg0Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench ValueReciverNoinlineGlobal1Arg0Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench PointerReciverNoinlineGlobal1Arg0Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench FuncNoinlineGlobal2Arg0Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench ValueReciverNoinlineGlobal2Arg0Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench PointerReciverNoinlineGlobal2Arg0Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench FuncNoinlineGlobal3Arg0Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench ValueReciverNoinlineGlobal3Arg0Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench PointerReciverNoinlineGlobal3Arg0Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench FuncNoinlineGlobal1Arg1Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench ValueReciverNoinlineGlobal1Arg1Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench PointerReciverNoinlineGlobal1Arg1Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench FuncNoinlineGlobal2Arg1Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench ValueReciverNoinlineGlobal2Arg1Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench PointerReciverNoinlineGlobal2Arg1Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench FuncNoinlineGlobal3Arg1Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench ValueReciverNoinlineGlobal3Arg1Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench PointerReciverNoinlineGlobal3Arg1Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench FuncNoinlineGlobal1Arg2Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench ValueReciverNoinlineGlobal1Arg2Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench PointerReciverNoinlineGlobal1Arg2Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench FuncNoinlineGlobal2Arg2Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench ValueReciverNoinlineGlobal2Arg2Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench PointerReciverNoinlineGlobal2Arg2Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench FuncNoinlineGlobal3Arg2Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench ValueReciverNoinlineGlobal3Arg2Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 
cmd="go test -bench PointerReciverNoinlineGlobal3Arg2Return"
echo $cmd >> $result_file
$cmd >> $result_file
echo -e "" >> $result_file 