# Gofield

Computes electric fields when point charge configuration is given. Uses the command line interface (CLI).

## Usage

```sh
# config.txt
#
# Point charge: Q <x> <y> <charge [C]>
# Measure point: P <x> <y>

Q 0.1 0 0.01
Q -0.1 0 0.01
P 0 0.5
P 0 1
P 0 1.5
```

```sh
$ ./gofield "config.txt"

### INPUTS ###
r(Q1) = 0.1i + 0j
r(Q2) = -0.1i + 0j
### OUTPUTS ###
E(0i + 0.5j) = 0i + -6.8e+07j
|E| = 67.79245850063175 [MV/m]
E(0i + 1j) = 0i + -1.8e+07j
|E| = 17.708808471279617 [MV/m]
E(0i + 1.5j) = 0i + -7.9e+06j
|E| = 7.935969716399953 [MV/m]
```

## Build

```sh
go build
```