module main

go 1.22.1

replace github.com/goruby/goruby => ../..

require github.com/goruby/goruby v0.0.0

require (
	github.com/chzyer/test v1.0.0 // indirect
	github.com/goruby/readline v0.0.0-20171103131923-a4d5111b6178
	golang.org/x/sys v0.0.0-20220310020820-b874c991c1a5 // indirect
)

require github.com/pkg/errors v0.9.1 // indirect
