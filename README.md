<img src="insulin3png.png" alt="drawing" width="350"/>

# Insulin :syringe:

Flexible, universal and fast Smart Contract testing framework written in Go

It uses [Goblin](https://github.com/franela/goblin) as a testing library for a mocha-like experience.

Insulin also integrates with [MythX](https://mythx.io/) out of the box for easy smart contract security analysis

Insulin is a great remedy for those you had too much sweet :candy:

## Installation
```bash
git clone git@github.com:jeffprestes/insulin.git
```
And then
```bash
make all
```
You need to have the [Go](https://golang.org) installed

## Usage

#### Installing
```bash
insulin init
```
This will create a contracts folder where you should leave your `.sol` files

#### Compiling
```bash
insulin compile
```
This will create the contracts artifactis in the `artifacts` folder

#### Testing
```bash
insulin test make
```
This will create auxiliary test files.
Write your own test in the file `test/your_contract_test.go`


To run the tests
```bash
insulin test run
```
