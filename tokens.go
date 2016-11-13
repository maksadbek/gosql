package main

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WS

	IDENT

	ASTERISK
	COMMA

	SELECT
	FROM
)
