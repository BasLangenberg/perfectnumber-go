# perfectnumber-go

An excercise in Go

## Reasoning

A couple of years ago, I wrote [this](https://github.com/BasLangenberg/perfectmath). Now I am learning Golang and I want to test if I can easily port this and use the concurrent features to search all perfect numbers < 10.000

## Timing

All executed on my XPS15 with 16GB of Ram and an i7 CPU. (August 2018)

### Original 10.000

- Original Python:

```shell
$ time python run.py
6
28
496
8128

real    0m1.405s
user    0m1.281s
sys     0m0.124s
```

- Go:

```shell
$ time go run main.go
6
28
496
8128

real    0m1.320s
user    0m0.000s
sys     0m0.031s
```

- Go /w Waitgroups:

```shell
$ time go run main.go
6
28
496
8128

real    0m1.243s
user    0m0.000s
sys     0m0.000s
```

### New: 100.000

- Original Python:

```shell
$ time python run.py
6
28
496
8128

real    2m38.102s
user    2m36.515s
sys     0m0.093s
```

- Go:

```shell
$ time go run main.go
6
28
496
8128

real    0m19.173s
user    0m0.000s
sys     0m0.015s
```

- Go /w Waitgroups (concurrency):

```shell
$ time go run main.go
6
28
496
8128

real    0m4.165s
user    0m0.000s
sys     0m0.000s
```
