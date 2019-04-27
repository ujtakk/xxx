XXX (tentative)
============================================================
General purpose binary builder

Overview
--------------------------------------------------

XXX (tentative) is a yet another general purpose binary builder.
XXX simply concatenate values listed in source texts and pack them to
the binary sequence.

Requirements
--------------------------------------------------

* Go (>= 1.10)

XXX is written by Go only with standard packages.
So, you can build XXX by:
```
$ go build
```

Usage
--------------------------------------------------

XXX supports values in formats of binary (0b101, ...), octal (0644, ...),
decimal (42, ...), and hexadecimal (0x1F, ...).
Values in a single line are separated by spaces (or tabs).
XXX restricts the bitwidth of each line to be a multiple of 8.
As for decimal format, you have to care the bitwidth of each value.

XXX provides some useful (but minimum) supports for building sequences.
`=` command enables to name the specific value.
`.` command enables to include other sources.
You can also put comments by `#`.
In each line, letters from `#` to the end of line are regarded as comments.

Here shows a example source (examples/level-4.xxx):
```
$ cat examples/level-4.xxx
############################################################
# header
############################################################

. examples/lib.xxx

############################################################
# alias
############################################################

= fuga 0x44
= hoge 0x12 # hosdfasdf
= wasd 33 # hosdfasdf

############################################################
# body
############################################################

nop                             # Initialize by NOP
0x12    32      0b1100000 02    # HOGE 32 15 2
hoge    wasd    0b1100000 02    # HOGE 33 15 2
0x12    34      0b1100000 02    # HOGE 34 15 2
fuga    35      0b1100000 02    # FUGA 35 15 2
0x12    36      0b1100000 02    # HOGE 36 15 2
jmp     0x0000                  # Jump to the origin
0x12    38      0b1100000 02    # HOGE 38 15 2
```

This source is packed by the command below:
(Sequence is dumped for stdout when `-o` option is omitted.)
```
$ xxx -o test.bin examples/level-4.xxx
```

Then, the result is:
```
$ xxd -b test.bin
00000000: 00000000 00010010 10000011 00000010 00010010 10000111  ......
00000006: 00000010 00010010 10001011 00000010 01000100 10001111  ....D.
0000000c: 00000010 00010010 10010011 00000010 00100010 00000000  ....".
00000012: 00000000 00010010 10011011 00000010                    ....
```

License
--------------------------------------------------

MIT License (see `LICENSE` file).
