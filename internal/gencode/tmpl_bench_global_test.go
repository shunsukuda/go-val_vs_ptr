package main

import "testing"

{{range $Words := $.Words}}
//type W{{$Words}} [{{$Words}}]uint64 
var rv{{$Words}} = W{{$Words}}{}
var arg1vW{{$Words}} = W{{$Words}}{}
var arg2vW{{$Words}} = W{{$Words}}{}
var arg3vW{{$Words}} = W{{$Words}}{}
var rp{{$Words}} = new(W{{$Words}})
var arg1pW{{$Words}} = new(W{{$Words}})
var arg2pW{{$Words}} = new(W{{$Words}})
var arg3pW{{$Words}} = new(W{{$Words}})
{{- end}} {{/* range $Words := $.Words */}}
{{- range $FuncType := $.FuncType}}
{{- range $Type := $.Type}}
{{- range $Words := $.Words}}
//go:noinline
func {{if ne $FuncType "Func" -}}
	(r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}}) {{end -}}
	{{$FuncType}}NoinlineGlobal1Arg0Return{{$Type}}{{$Words}}Word(arg {{if eq $Type "Pointer"}}*{{end}}W{{$Words}}) {
	{{- if eq $Type "Pointer"}}if arg == nil {
		return
	}{{- end -}} {{/* if */}}
	_ = arg
}

//go:noinline
func {{if ne $FuncType "Func" -}}
	(r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}}) {{end -}}
	{{$FuncType}}NoinlineGlobal2Arg0Return{{$Type}}{{$Words}}Word(arg1, arg2 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}}) {
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
	{{$FuncType}}NoinlineGlobal3Arg0Return{{$Type}}{{$Words}}Word(arg1, arg2, arg3 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}}) {
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
	{{$FuncType}}NoinlineGlobal1Arg1Return{{$Type}}{{$Words}}Word(arg {{if eq $Type "Pointer"}}*{{end}}W{{$Words}}) {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} {
	{{- if eq $Type "Pointer"}}if arg == nil {
		return nil
	}{{- end -}} {{/* if */}}
	_ = arg
	return arg
}

//go:noinline
func {{if ne $FuncType "Func" -}}
	(r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}}) {{end -}}
	{{$FuncType}}NoinlineGlobal2Arg1Return{{$Type}}{{$Words}}Word(arg1, arg2 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}}) {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} {
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
	{{$FuncType}}NoinlineGlobal3Arg1Return{{$Type}}{{$Words}}Word(arg1, arg2, arg3 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}}) {{if eq $Type "Pointer"}}*{{end}}W{{$Words}} {
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
	{{$FuncType}}NoinlineGlobal1Arg2Return{{$Type}}{{$Words}}Word(arg {{if eq $Type "Pointer"}}*{{end}}W{{$Words}}) ({{if eq $Type "Pointer"}}*{{end}}W{{$Words}}, {{if eq $Type "Pointer"}}*{{end}}W{{$Words}}) {
	{{- if eq $Type "Pointer"}}if arg == nil {
		return nil, nil
	}{{- end -}} {{/* if */}}
	_ = arg
	return arg, arg
}

//go:noinline
func {{if ne $FuncType "Func" -}}
	(r {{if eq $FuncType "PointerReciver"}}*{{end}}W{{$Words}}) {{end}}
	{{- $FuncType}}NoinlineGlobal2Arg2Return{{$Type}}{{$Words}}Word(arg1, arg2 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}})  ({{if eq $Type "Pointer"}}*{{end}}W{{$Words}}, {{if eq $Type "Pointer"}}*{{end}}W{{$Words}}) {
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
	{{- $FuncType}}NoinlineGlobal3Arg2Return{{$Type}}{{$Words}}Word(arg1, arg2, arg3 {{if eq $Type "Pointer"}}*{{end}}W{{$Words}})  ({{if eq $Type "Pointer"}}*{{end}}W{{$Words}}, {{if eq $Type "Pointer"}}*{{end}}W{{$Words}}) {
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
func Benchmark{{$FuncType}}NoinlineGlobal1Arg0Return{{$Type}}{{$Words}}Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		{{if ne $FuncType "Func"}}r{{if eq $FuncType "ValueReciver"}}v{{else}}p{{end}}{{$Words}}.{{end}}{{$FuncType}}NoinlineGlobal1Arg0Return{{$Type}}{{$Words}}Word(arg1
			{{- if eq $Type "Pointer"}}p{{else}}v{{end}}W{{$Words}})
	}
	b.StopTimer()
}

func Benchmark{{$FuncType}}NoinlineGlobal2Arg0Return{{$Type}}{{$Words}}Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		{{if ne $FuncType "Func"}}r{{if eq $FuncType "ValueReciver"}}v{{else}}p{{end}}{{$Words}}.{{end}}{{$FuncType}}NoinlineGlobal2Arg0Return{{$Type}}{{$Words}}Word(arg1
			{{- if eq $Type "Pointer"}}p{{else}}v{{end}}W{{$Words}}, arg2
			{{- if eq $Type "Pointer"}}p{{else}}v{{end}}W{{$Words}})
	}
	b.StopTimer()
}

func Benchmark{{$FuncType}}NoinlineGlobal3Arg0Return{{$Type}}{{$Words}}Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		{{if ne $FuncType "Func"}}r{{if eq $FuncType "ValueReciver"}}v{{else}}p{{end}}{{$Words}}.{{end}}{{$FuncType}}NoinlineGlobal3Arg0Return{{$Type}}{{$Words}}Word(arg1
			{{- if eq $Type "Pointer"}}p{{else}}v{{end}}W{{$Words}}, arg2
			{{- if eq $Type "Pointer"}}p{{else}}v{{end}}W{{$Words}}, arg3
			{{- if eq $Type "Pointer"}}p{{else}}v{{end}}W{{$Words}})
	}
	b.StopTimer()
}

func Benchmark{{$FuncType}}NoinlineGlobal1Arg1Return{{$Type}}{{$Words}}Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = {{if ne $FuncType "Func"}}r{{if eq $FuncType "ValueReciver"}}v{{else}}p{{end}}{{$Words}}.{{end}}{{$FuncType}}NoinlineGlobal1Arg1Return{{$Type}}{{$Words}}Word(arg1
			{{- if eq $Type "Pointer"}}p{{else}}v{{end}}W{{$Words}})
	}
	b.StopTimer()
}

func Benchmark{{$FuncType}}NoinlineGlobal2Arg1Return{{$Type}}{{$Words}}Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = {{if ne $FuncType "Func"}}r{{if eq $FuncType "ValueReciver"}}v{{else}}p{{end}}{{$Words}}.{{end}}{{$FuncType}}NoinlineGlobal2Arg1Return{{$Type}}{{$Words}}Word(arg1
			{{- if eq $Type "Pointer"}}p{{else}}v{{end}}W{{$Words}}, arg2
			{{- if eq $Type "Pointer"}}p{{else}}v{{end}}W{{$Words}})
	}
	b.StopTimer()
}

func Benchmark{{$FuncType}}NoinlineGlobal3Arg1Return{{$Type}}{{$Words}}Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = {{if ne $FuncType "Func"}}r{{if eq $FuncType "ValueReciver"}}v{{else}}p{{end}}{{$Words}}.{{end}}{{$FuncType}}NoinlineGlobal3Arg1Return{{$Type}}{{$Words}}Word(arg1
			{{- if eq $Type "Pointer"}}p{{else}}v{{end}}W{{$Words}}, arg2
			{{- if eq $Type "Pointer"}}p{{else}}v{{end}}W{{$Words}}, arg3
			{{- if eq $Type "Pointer"}}p{{else}}v{{end}}W{{$Words}})
	}
	b.StopTimer()
}


func Benchmark{{$FuncType}}NoinlineGlobal1Arg2Return{{$Type}}{{$Words}}Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = {{if ne $FuncType "Func"}}r{{if eq $FuncType "ValueReciver"}}v{{else}}p{{end}}{{$Words}}.{{end}}{{$FuncType}}NoinlineGlobal1Arg2Return{{$Type}}{{$Words}}Word(arg1
			{{- if eq $Type "Pointer"}}p{{else}}v{{end}}W{{$Words}})
	}
	b.StopTimer()
}

func Benchmark{{$FuncType}}NoinlineGlobal2Arg2Return{{$Type}}{{$Words}}Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = {{if ne $FuncType "Func"}}r{{if eq $FuncType "ValueReciver"}}v{{else}}p{{end}}{{$Words}}.{{end}}{{$FuncType}}NoinlineGlobal2Arg2Return{{$Type}}{{$Words}}Word(arg1
			{{- if eq $Type "Pointer"}}p{{else}}v{{end}}W{{$Words}}, arg2
			{{- if eq $Type "Pointer"}}p{{else}}v{{end}}W{{$Words}})
	}
	b.StopTimer()
}

func Benchmark{{$FuncType}}NoinlineGlobal3Arg2Return{{$Type}}{{$Words}}Word(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = {{if ne $FuncType "Func"}}r{{if eq $FuncType "ValueReciver"}}v{{else}}p{{end}}{{$Words}}.{{end}}{{$FuncType}}NoinlineGlobal3Arg2Return{{$Type}}{{$Words}}Word(arg1
			{{- if eq $Type "Pointer"}}p{{else}}v{{end}}W{{$Words}}, arg2
			{{- if eq $Type "Pointer"}}p{{else}}v{{end}}W{{$Words}}, arg3
			{{- if eq $Type "Pointer"}}p{{else}}v{{end}}W{{$Words}})
	}
	b.StopTimer()
}
{{- end}} {{/* range $Words := $.Words */}}
{{- end}} {{/* range $Type := $.Type */}}
{{- end}} {{/* range $FuncType := $.FuncType */}}