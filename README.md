# mz - print mz .exe header info

Prints the header of an MZ EXE (i.e. MS-DOS EXE) in a human-readable form.

For more information check out:

- https://www.fileformat.info/format/exe/corion-mz.htm
- https://wiki.osdev.org/MZ

## usage

Pass it the file as an argument:

```
mz myfile.exe
```

Or from stdin:

```
mz < myfile.exe
curl http://my-sketchy-website.example.com/file.exe | mz
```

## output example

```
3C501.EXE: MZ header OK!
Bytes in last page:                 0x0052
Number of pages (inc last):         0x003b
Number of relocation entries:       0x0003
Header size (paragraphs):           0x0020
Min. Memory allocated (paragraphs): 0x0009
Max. Memory allocated (paragraphs): 0xffff
Initial Stack Segment:              0x0726
Initial Stack Pointer:              0x0080
Checksum (0 for none):              0xf55b
Initial Instruction Pointer:        0x0002
Initial Code Segment:               0x0000
Offset of relocation table:         0x001e
Overlay number:                     0x0000
```

## copying

Licensed MIT and dedicated to [Moochikins](https://github.com/MoochMcGee)
