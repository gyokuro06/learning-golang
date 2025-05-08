module example.com/hello

go 1.24.1

require (
	example.com/greetings v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
)

replace example.com/greetings => ../greetings
