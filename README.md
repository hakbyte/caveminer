# Caveminer

According to Wikipedia [code caves](https://en.wikipedia.org/wiki/Code_cave)
are *areas inside a program's file that are not used for code (instructions)*.
These spaces are usually filled with zeros by the compiler/linker as either
padding or placeholders for future code changes.

The interesting thing about code caves is that they can be used to hide
malicious code or data. To explore that I built *caveminer*, a tool that list
all code caves found in a Windows binary. This information can later be abused
in different ways (see this old [blog post](https://pentest.blog/art-of-anti-detection-2-pe-backdoor-manufacturing/)
for some ideas).

Example output:

```txt
C:\>.\caveminer.exe C:\Windows\System32\calc.exe 80
-----------------------------------
Section .text (3072 bytes)
No caves found.
-----------------------------------
Section .rdata (3584 bytes)
# Cave 1
	Size        : 86 bytes
	Offset Start: 1012h
	Offset End  : 1068h
# Cave 2
	Size        : 133 bytes
	Offset Start: 10a3h
	Offset End  : 1128h
# Cave 3
	Size        : 396 bytes
	Offset Start: 1c74h
	Offset End  : 1e00h
-----------------------------------
Section .data (512 bytes)
# Cave 1
	Size        : 431 bytes
	Offset Start: 1e51h
	Offset End  : 2000h
-----------------------------------
Section .pdata (512 bytes)
# Cave 1
	Size        : 274 bytes
	Offset Start: 20eeh
	Offset End  : 2200h
-----------------------------------
Section .rsrc (18432 bytes)
# Cave 1
	Size        : 128 bytes
	Offset Start: 42a8h
	Offset End  : 4328h
# Cave 2
	Size        : 386 bytes
	Offset Start: 6750h
	Offset End  : 68d2h
# Cave 3
	Size        : 243 bytes
	Offset Start: 690dh
	Offset End  : 6a00h
-----------------------------------
Section .reloc (512 bytes)
# Cave 1
	Size        : 470 bytes
	Offset Start: 6a2ah
	Offset End  : 6c00h
-----------------------------------
```
