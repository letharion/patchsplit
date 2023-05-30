# PatchSplit

PatchSplit is a Go program that reads a given file containing multiple patches and splits it into separate patch files.
[Splitpatch](https://github.com/jaalto/splitpatch) didn't handle git patch headers correctly, so I wrote this to solve that and practice some go.

## Installation

```bash
git clone https://github.com/letharion/patchsplit.git
cd patchsplit
go build
```

## Usage

Once the program is built, you can use it as follows:

```bash

./patchsplit <file_path>
```
