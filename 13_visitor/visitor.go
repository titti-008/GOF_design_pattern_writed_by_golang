package main

type visitor interface {
	visitFile(f file)
	visitDir(d *dir)
}
