package main

import "testing"

{{range $Words := $.Words}}
type W{{$Words}} [{{$Words}}]uint64 
{{- end}} {{/* range $Words := $.Words */}}
{{- range $FuncType := $.FuncType}}
{{- range $Type := $.Type}}
{{- range $Words := $.Words}}
//go:noinline
func {{if ne $FuncType "Func" -}}
	(r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}}) {{end -}}
	{{$FuncType}}Noinline1Arg0Return{{$Type}}{{$Words}}Word(arg {{if eq $Type "Pointer"}}*{{end}}W{{$Words}}) {
	{{- if eq $Type "Pointer"}}if arg == nil {
		return
	}{{- end -}} {{/* if */}}
	_ = arg
}

//go:noinline
func {{if ne $FuncType "Func" -}}
	(r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}}) {{end -}}
	{{$FuncType}}Noinline2Arg0Return{{$Type}}{{$Words}}Word(arg1, arg2 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}}) {
	{{- if eq $Type "Pointer"}}if arg1 == nil {
		return
	}
	if arg2 == nil {
		return 
	}{{- end -}} {{/* if */}}
	_, _ = arg1, arg2
}

//go:noinline
func {{if ne $FuncType "Func" -}}
	(r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}}) {{end -}}
	{{$FuncType}}Noinline3Arg0Return{{$Type}}{{$Words}}Word(arg1, arg2, arg3 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}}) {
	{{- if eq $Type "Pointer"}}if arg1 == nil {
		return
	}
	if arg2 == nil {
		return 
	}
	if arg3 == nil {
		return
	}{{- end -}} {{/* if */}}
	_, _, _ = arg1, arg2, arg3
}

//go:noinline
func {{if ne $FuncType "Func" -}}
	(r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}}) {{end -}}
	{{$FuncType}}Noinline1Arg1Return{{$Type}}{{$Words}}Word(arg {{if eq $Type "Pointer"}}*{{end}}W{{$Words}}) {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} {
	{{- if eq $Type "Pointer"}}if arg == nil {
		return nil
	}{{- end -}} {{/* if */}}
	_ = arg
	return arg
}

//go:noinline
func {{if ne $FuncType "Func" -}}
	(r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}}) {{end -}}
	{{$FuncType}}Noinline2Arg1Return{{$Type}}{{$Words}}Word(arg1, arg2 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}}) {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} {
	{{- if eq $Type "Pointer"}}if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}{{- end -}} {{/* if */}}
	_, _ = arg1, arg2
	return arg1
}

//go:noinline
func {{if ne $FuncType "Func" -}}
	(r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}}) {{end -}}
	{{$FuncType}}Noinline3Arg1Return{{$Type}}{{$Words}}Word(arg1, arg2, arg3 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}}) {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} {
	{{- if eq $Type "Pointer"}}if arg1 == nil {
		return nil
	}
	if arg2 == nil {
		return nil
	}
	if arg3 == nil {
		return nil
	}{{- end -}} {{/* if */}}
	_, _, _ = arg1, arg2, arg3
	return arg1
}

//go:noinline
func {{if ne $FuncType "Func" -}}
	(r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}}) {{end -}}
	{{$FuncType}}Noinline1Arg2Return{{$Type}}{{$Words}}Word(arg {{if eq $Type "Pointer"}}*{{end}}W{{$Words}}) ({{if eq $Type "Pointer"}}*{{end}}W{{$Words}}, {{if eq $Type "Pointer"}}*{{end}}W{{$Words}}) {
	{{- if eq $Type "Pointer"}}if arg == nil {
		return nil, nil
	}{{- end -}} {{/* if */}}
	_ = arg
	return arg, arg
}

//go:noinline
func {{if ne $FuncType "Func" -}}
	(r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}}) {{end}}
	{{- $FuncType}}Noinline2Arg2Return{{$Type}}{{$Words}}Word(arg1, arg2 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}})  ({{if eq $Type "Pointer"}}*{{end}}W{{$Words}}, {{if eq $Type "Pointer"}}*{{end}}W{{$Words}}) {
	{{- if eq $Type "Pointer"}}if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}{{- end -}} {{/* if */}}
	_, _ = arg1, arg2
	return arg1, arg2
}

//go:noinline
func {{if ne $FuncType "Func" -}}
	(r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}}) {{end}}
	{{- $FuncType}}Noinline3Arg2Return{{$Type}}{{$Words}}Word(arg1, arg2, arg3 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}})  ({{if eq $Type "Pointer"}}*{{end}}W{{$Words}}, {{if eq $Type "Pointer"}}*{{end}}W{{$Words}}) {
	{{- if eq $Type "Pointer"}}if arg1 == nil {
		return nil, nil
	}
	if arg2 == nil {
		return nil, nil
	}
	if arg3 == nil {
		return nil, nil
	}{{- end -}} {{/* if */}}
	_, _, _ = arg1, arg2, arg3
	return arg1, arg2
}

{{- end}} {{/* range $Words := $.Words */}}
{{- end}} {{/* range $Type := $.Type */}}
{{- end}} {{/* range $FuncType := $.FuncType */}}

{{range $FuncType := $.FuncType}}
{{range $Type := $.Type}}
{{range $Words := $.Words}}
func Benchmark{{$FuncType}}Noinline1Arg0Return{{$Type}}{{$Words}}Word(b *testing.B) {
	b.ReportAllocs()
	var arg {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} = {{if eq $Type "Pointer"}}new({{end}}W{{$Words}}{{if eq $Type "Pointer"}}){{else}}{}{{end}}
	{{- if ne $FuncType "Func"}}
	var r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}} = {{if eq $FuncType "PointerReciver"}}new({{end}}W{{$Words}}{{if eq $FuncType "PointerReciver"}}){{else}}{}{{end}}
	{{- end}}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		{{if ne $FuncType "Func"}}r.{{end}}{{$FuncType}}Noinline1Arg0Return{{$Type}}{{$Words}}Word(arg)
	}
	b.StopTimer()
}

func Benchmark{{$FuncType}}Noinline2Arg0Return{{$Type}}{{$Words}}Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} = {{if eq $Type "Pointer"}}new({{end}}W{{$Words}}{{if eq $Type "Pointer"}}){{else}}{}{{end}}
	var arg2 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} = {{if eq $Type "Pointer"}}new({{end}}W{{$Words}}{{if eq $Type "Pointer"}}){{else}}{}{{end}}
	{{- if ne $FuncType "Func"}}
	var r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}} = {{if eq $FuncType "PointerReciver"}}new({{end}}W{{$Words}}{{if eq $FuncType "PointerReciver"}}){{else}}{}{{end}}
	{{- end}}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		{{if ne $FuncType "Func"}}r.{{end}}{{$FuncType}}Noinline2Arg0Return{{$Type}}{{$Words}}Word(arg1, arg2)
	}
	b.StopTimer()
}

func Benchmark{{$FuncType}}Noinline3Arg0Return{{$Type}}{{$Words}}Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} = {{if eq $Type "Pointer"}}new({{end}}W{{$Words}}{{if eq $Type "Pointer"}}){{else}}{}{{end}}
	var arg2 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} = {{if eq $Type "Pointer"}}new({{end}}W{{$Words}}{{if eq $Type "Pointer"}}){{else}}{}{{end}}
	var arg3 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} = {{if eq $Type "Pointer"}}new({{end}}W{{$Words}}{{if eq $Type "Pointer"}}){{else}}{}{{end}}
	{{- if ne $FuncType "Func"}}
	var r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}} = {{if eq $FuncType "PointerReciver"}}new({{end}}W{{$Words}}{{if eq $FuncType "PointerReciver"}}){{else}}{}{{end}}
	{{- end}}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		{{if ne $FuncType "Func"}}r.{{end}}{{$FuncType}}Noinline3Arg0Return{{$Type}}{{$Words}}Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}

func Benchmark{{$FuncType}}Noinline1Arg1Return{{$Type}}{{$Words}}Word(b *testing.B) {
	b.ReportAllocs()
	var arg {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} = {{if eq $Type "Pointer"}}new({{end}}W{{$Words}}{{if eq $Type "Pointer"}}){{else}}{}{{end}}
	{{- if ne $FuncType "Func"}}
	var r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}} = {{if eq $FuncType "PointerReciver"}}new({{end}}W{{$Words}}{{if eq $FuncType "PointerReciver"}}){{else}}{}{{end}}
	{{- end}}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = {{if ne $FuncType "Func"}}r.{{end}}{{$FuncType}}Noinline1Arg1Return{{$Type}}{{$Words}}Word(arg)
	}
	b.StopTimer()
}

func Benchmark{{$FuncType}}Noinline2Arg1Return{{$Type}}{{$Words}}Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} = {{if eq $Type "Pointer"}}new({{end}}W{{$Words}}{{if eq $Type "Pointer"}}){{else}}{}{{end}}
	var arg2 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} = {{if eq $Type "Pointer"}}new({{end}}W{{$Words}}{{if eq $Type "Pointer"}}){{else}}{}{{end}}
	{{- if ne $FuncType "Func"}}
	var r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}} = {{if eq $FuncType "PointerReciver"}}new({{end}}W{{$Words}}{{if eq $FuncType "PointerReciver"}}){{else}}{}{{end}}
	{{- end}}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = {{if ne $FuncType "Func"}}r.{{end}}{{$FuncType}}Noinline2Arg1Return{{$Type}}{{$Words}}Word(arg1, arg2)
	}
	b.StopTimer()
}

func Benchmark{{$FuncType}}Noinline3Arg1Return{{$Type}}{{$Words}}Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} = {{if eq $Type "Pointer"}}new({{end}}W{{$Words}}{{if eq $Type "Pointer"}}){{else}}{}{{end}}
	var arg2 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} = {{if eq $Type "Pointer"}}new({{end}}W{{$Words}}{{if eq $Type "Pointer"}}){{else}}{}{{end}}
	var arg3 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} = {{if eq $Type "Pointer"}}new({{end}}W{{$Words}}{{if eq $Type "Pointer"}}){{else}}{}{{end}}
	{{- if ne $FuncType "Func"}}
	var r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}} = {{if eq $FuncType "PointerReciver"}}new({{end}}W{{$Words}}{{if eq $FuncType "PointerReciver"}}){{else}}{}{{end}}
	{{- end}}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = {{if ne $FuncType "Func"}}r.{{end}}{{$FuncType}}Noinline3Arg1Return{{$Type}}{{$Words}}Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}


func Benchmark{{$FuncType}}Noinline1Arg2Return{{$Type}}{{$Words}}Word(b *testing.B) {
	b.ReportAllocs()
	var arg {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} = {{if eq $Type "Pointer"}}new({{end}}W{{$Words}}{{if eq $Type "Pointer"}}){{else}}{}{{end}}
	{{- if ne $FuncType "Func"}}
	var r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}} = {{if eq $FuncType "PointerReciver"}}new({{end}}W{{$Words}}{{if eq $FuncType "PointerReciver"}}){{else}}{}{{end}}
	{{- end}}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = {{if ne $FuncType "Func"}}r.{{end}}{{$FuncType}}Noinline1Arg2Return{{$Type}}{{$Words}}Word(arg)
	}
	b.StopTimer()
}

func Benchmark{{$FuncType}}Noinline2Arg2Return{{$Type}}{{$Words}}Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} = {{if eq $Type "Pointer"}}new({{end}}W{{$Words}}{{if eq $Type "Pointer"}}){{else}}{}{{end}}
	var arg2 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} = {{if eq $Type "Pointer"}}new({{end}}W{{$Words}}{{if eq $Type "Pointer"}}){{else}}{}{{end}}
	{{- if ne $FuncType "Func"}}
	var r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}} = {{if eq $FuncType "PointerReciver"}}new({{end}}W{{$Words}}{{if eq $FuncType "PointerReciver"}}){{else}}{}{{end}}
	{{- end}}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = {{if ne $FuncType "Func"}}r.{{end}}{{$FuncType}}Noinline2Arg2Return{{$Type}}{{$Words}}Word(arg1, arg2)
	}
	b.StopTimer()
}

func Benchmark{{$FuncType}}Noinline3Arg2Return{{$Type}}{{$Words}}Word(b *testing.B) {
	b.ReportAllocs()
	var arg1 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} = {{if eq $Type "Pointer"}}new({{end}}W{{$Words}}{{if eq $Type "Pointer"}}){{else}}{}{{end}}
	var arg2 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} = {{if eq $Type "Pointer"}}new({{end}}W{{$Words}}{{if eq $Type "Pointer"}}){{else}}{}{{end}}
	var arg3 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} = {{if eq $Type "Pointer"}}new({{end}}W{{$Words}}{{if eq $Type "Pointer"}}){{else}}{}{{end}}
	{{- if ne $FuncType "Func"}}
	var r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}} = {{if eq $FuncType "PointerReciver"}}new({{end}}W{{$Words}}{{if eq $FuncType "PointerReciver"}}){{else}}{}{{end}}
	{{- end}}
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = {{if ne $FuncType "Func"}}r.{{end}}{{$FuncType}}Noinline3Arg2Return{{$Type}}{{$Words}}Word(arg1, arg2, arg3)
	}
	b.StopTimer()
}
{{- end}} {{/* range $Words := $.Words */}}
{{- end}} {{/* range $Type := $.Type */}}
{{- end}} {{/* range $FuncType := $.FuncType */}}