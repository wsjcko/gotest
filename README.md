
### 编译
go tool compile -S -N -l main.go  //生成main.o中间文件
go tool link -v -o main main.o //生成可执行文件

### 基准测试
go test -bench=. -benchmem        

##### 内存性能
go test -bench=. -benchmem -memprofile=mem.out

##### CPU性能
go test -bench=. -benchmem -cpuprofile=cpu.out

##### 文件查看
生成的CPU、内存文件可以通过go tool pprof [file]进行查看，然后在pprof中通过list [file]方法查看CPU、内存的耗时情况
go tool pprof mem.out
list BenchmarkRemoveDuplicatesInPlace

File: gotest.test
Type: alloc_space
Time: May 2, 2022 at 12:34pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) list BenchmarkRemoveDuplicatesInPlace
Total: 358.31MB
ROUTINE ======================== gotest.BenchmarkRemoveDuplicatesInPlace in /home/wsjcko/test/remove_test.go
         0   356.30MB (flat, cum) 99.44% of Total
         .          .     12:   }
         .          .     13:   return s
         .          .     14:}
         .          .     15:
         .          .     16:func BenchmarkRemoveDuplicatesInPlace(b *testing.B) {
         .   356.30MB     17:   s := generateSlice()
         .          .     18:   b.ResetTimer()
         .          .     19:   for i := 0; i < b.N; i++ {
         .          .     20:           RemoveDuplicatesInPlace(s)
         .          .     21:   }
         .          .     22:}
(pprof) 



go tool pprof cpu.out
list BenchmarkRemoveDuplicatesInPlace

File: gotest.test
Type: cpu
Time: May 2, 2022 at 12:34pm (CST)
Duration: 7.63s, Total samples = 7.52s (98.54%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) list BenchmarkRemoveDuplicatesInPlace
Total: 7.52s
ROUTINE ======================== gotest.BenchmarkRemoveDuplicatesInPlace in /home/wsjcko/test/remove_test.go
         0      7.44s (flat, cum) 98.94% of Total
         .          .     12:   }
         .          .     13:   return s
         .          .     14:}
         .          .     15:
         .          .     16:func BenchmarkRemoveDuplicatesInPlace(b *testing.B) {
         .      260ms     17:   s := generateSlice()
         .          .     18:   b.ResetTimer()
         .          .     19:   for i := 0; i < b.N; i++ {
         .      7.18s     20:           RemoveDuplicatesInPlace(s)
         .          .     21:   }
         .          .     22:}
(pprof) 



### file 命令查看文件格式
file main.go
main.go: C source, UTF-8 Unicode text
file main.o 
current ar archive
file main   
main: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, not stripped


### go build -x 查看生成可执行文件过程
go build -x main.go 
不需要生成可执行文件,只是单纯查看过程: go build -n main.go 
不需要生成可执行文件,只是单纯查看过程,并执行程序 go run -x main.go

WORK=/tmp/go-build2682583105
mkdir -p $WORK/b001/
cat >$WORK/b001/importcfg << 'EOF' # internal
import config
packagefile fmt=/usr/local/go/pkg/linux_amd64/fmt.a
packagefile math=/usr/local/go/pkg/linux_amd64/math.a
packagefile math/rand=/usr/local/go/pkg/linux_amd64/math/rand.a
packagefile reflect=/usr/local/go/pkg/linux_amd64/reflect.a
packagefile sync=/usr/local/go/pkg/linux_amd64/sync.a
packagefile time=/usr/local/go/pkg/linux_amd64/time.a
packagefile runtime=/usr/local/go/pkg/linux_amd64/runtime.a
EOF
cd /home/wsjcko/test
/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b001/_pkg_.a -trimpath "$WORK/b001=>" -p main -complete -buildid UuhDfIDNdijc2Z6D33HK/UuhDfIDNdijc2Z6D33HK -goversion go1.18.1 -c=4 -nolocalimports -importcfg $WORK/b001/importcfg -pack ./main.go
/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b001/_pkg_.a # internal
cp $WORK/b001/_pkg_.a /home/wsjcko/.cache/go-build/cf/cf74e201cdc567e629449b0a966daab4a2588aad12c8e39d512f2dd3affda4e6-d # internal
cat >$WORK/b001/importcfg.link << 'EOF' # internal
packagefile command-line-arguments=$WORK/b001/_pkg_.a
packagefile fmt=/usr/local/go/pkg/linux_amd64/fmt.a
packagefile math=/usr/local/go/pkg/linux_amd64/math.a
packagefile math/rand=/usr/local/go/pkg/linux_amd64/math/rand.a
packagefile reflect=/usr/local/go/pkg/linux_amd64/reflect.a
packagefile sync=/usr/local/go/pkg/linux_amd64/sync.a
packagefile time=/usr/local/go/pkg/linux_amd64/time.a
packagefile runtime=/usr/local/go/pkg/linux_amd64/runtime.a
packagefile errors=/usr/local/go/pkg/linux_amd64/errors.a
packagefile internal/fmtsort=/usr/local/go/pkg/linux_amd64/internal/fmtsort.a
packagefile io=/usr/local/go/pkg/linux_amd64/io.a
packagefile os=/usr/local/go/pkg/linux_amd64/os.a
packagefile strconv=/usr/local/go/pkg/linux_amd64/strconv.a
packagefile unicode/utf8=/usr/local/go/pkg/linux_amd64/unicode/utf8.a
packagefile internal/cpu=/usr/local/go/pkg/linux_amd64/internal/cpu.a
packagefile math/bits=/usr/local/go/pkg/linux_amd64/math/bits.a
packagefile internal/abi=/usr/local/go/pkg/linux_amd64/internal/abi.a
packagefile internal/bytealg=/usr/local/go/pkg/linux_amd64/internal/bytealg.a
packagefile internal/goarch=/usr/local/go/pkg/linux_amd64/internal/goarch.a
packagefile internal/goexperiment=/usr/local/go/pkg/linux_amd64/internal/goexperiment.a
packagefile internal/itoa=/usr/local/go/pkg/linux_amd64/internal/itoa.a
packagefile internal/unsafeheader=/usr/local/go/pkg/linux_amd64/internal/unsafeheader.a
packagefile unicode=/usr/local/go/pkg/linux_amd64/unicode.a
packagefile internal/race=/usr/local/go/pkg/linux_amd64/internal/race.a
packagefile sync/atomic=/usr/local/go/pkg/linux_amd64/sync/atomic.a
packagefile syscall=/usr/local/go/pkg/linux_amd64/syscall.a
packagefile internal/goos=/usr/local/go/pkg/linux_amd64/internal/goos.a
packagefile runtime/internal/atomic=/usr/local/go/pkg/linux_amd64/runtime/internal/atomic.a
packagefile runtime/internal/math=/usr/local/go/pkg/linux_amd64/runtime/internal/math.a
packagefile runtime/internal/sys=/usr/local/go/pkg/linux_amd64/runtime/internal/sys.a
packagefile runtime/internal/syscall=/usr/local/go/pkg/linux_amd64/runtime/internal/syscall.a
packagefile internal/reflectlite=/usr/local/go/pkg/linux_amd64/internal/reflectlite.a
packagefile sort=/usr/local/go/pkg/linux_amd64/sort.a
packagefile internal/oserror=/usr/local/go/pkg/linux_amd64/internal/oserror.a
packagefile internal/poll=/usr/local/go/pkg/linux_amd64/internal/poll.a
packagefile internal/syscall/execenv=/usr/local/go/pkg/linux_amd64/internal/syscall/execenv.a
packagefile internal/syscall/unix=/usr/local/go/pkg/linux_amd64/internal/syscall/unix.a
packagefile internal/testlog=/usr/local/go/pkg/linux_amd64/internal/testlog.a
packagefile io/fs=/usr/local/go/pkg/linux_amd64/io/fs.a
packagefile path=/usr/local/go/pkg/linux_amd64/path.a
modinfo "0w\xaf\f\x92t\b\x02A\xe1\xc1\a\xe6\xd6\x18\xe6path\tcommand-line-arguments\nbuild\t-compiler=gc\nbuild\tCGO_ENABLED=1\nbuild\tCGO_CFLAGS=\nbuild\tCGO_CPPFLAGS=\nbuild\tCGO_CXXFLAGS=\nbuild\tCGO_LDFLAGS=\nbuild\tGOARCH=amd64\nbuild\tGOOS=linux\nbuild\tGOAMD64=v1\n\xf92C1\x86\x18 r\x00\x82B\x10A\x16\xd8\xf2"
EOF
mkdir -p $WORK/b001/exe/
cd .
/usr/local/go/pkg/tool/linux_amd64/link -o $WORK/b001/exe/a.out -importcfg $WORK/b001/importcfg.link -buildmode=exe -buildid=6-9c_0F3RiX2QXgxBCOc/UuhDfIDNdijc2Z6D33HK/K7OTO8wVdJ6HriHnKgGu/6-9c_0F3RiX2QXgxBCOc -extld=gcc $WORK/b001/_pkg_.a
/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b001/exe/a.out # internal
mv $WORK/b001/exe/a.out main
rm -r $WORK/b001/

### 编译工作目录，可以查看编译过程中创建的所有文件
go run -work main.go     
WORK=/tmp/go-build3129056699

### 显示使用到的汇编指令
go build -gcflags="-S" main.go

command-line-arguments
"".Do STEXT size=108 args=0x8 locals=0x18 funcid=0x0 align=0x0
        0x0000 00000 (/home/wsjcko/test/main.go:9)      TEXT    "".Do(SB), ABIInternal, $24-8
        0x0000 00000 (/home/wsjcko/test/main.go:9)      CMPQ    SP, 16(R14)
        0x0004 00004 (/home/wsjcko/test/main.go:9)      PCDATA  $0, $-2
        0x0004 00004 (/home/wsjcko/test/main.go:9)      JLS     91
        0x0006 00006 (/home/wsjcko/test/main.go:9)      PCDATA  $0, $-1
        0x0006 00006 (/home/wsjcko/test/main.go:9)      SUBQ    $24, SP
        0x000a 00010 (/home/wsjcko/test/main.go:9)      MOVQ    BP, 16(SP)
        0x000f 00015 (/home/wsjcko/test/main.go:9)      LEAQ    16(SP), BP
        0x0014 00020 (/home/wsjcko/test/main.go:9)      FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0014 00020 (/home/wsjcko/test/main.go:9)      FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0014 00020 (/home/wsjcko/test/main.go:9)      FUNCDATA        $5, "".Do.arginfo1(SB)
        0x0014 00020 (/home/wsjcko/test/main.go:9)      FUNCDATA        $6, "".Do.argliveinfo(SB)
        0x0014 00020 (/home/wsjcko/test/main.go:9)      PCDATA  $3, $1
        0x0014 00020 (/home/wsjcko/test/main.go:10)     CMPQ    AX, $13
        0x0018 00024 (/home/wsjcko/test/main.go:10)     JLE     46
        0x001a 00026 (/home/wsjcko/test/main.go:11)     MOVQ    $120000000000, AX
        0x0024 00036 (/home/wsjcko/test/main.go:11)     MOVQ    16(SP), BP
        0x0029 00041 (/home/wsjcko/test/main.go:11)     ADDQ    $24, SP
        0x002d 00045 (/home/wsjcko/test/main.go:11)     RET
        0x002e 00046 (<unknown line number>)    NOP
        0x002e 00046 (/home/wsjcko/test/main.go:13)     XORPS   X0, X0
        0x0031 00049 (/home/wsjcko/test/main.go:13)     CVTSQ2SD        AX, X0
        0x0036 00054 ($GOROOT/src/math/pow.go:42)       MOVSD   $f64.4005bf0a8b145769(SB), X1
        0x003e 00062 ($GOROOT/src/math/pow.go:42)       PCDATA  $1, $0
        0x003e 00062 ($GOROOT/src/math/pow.go:42)       NOP
        0x0040 00064 ($GOROOT/src/math/pow.go:42)       CALL    math.pow(SB)
        0x0045 00069 (/home/wsjcko/test/main.go:13)     CVTTSD2SQ       X0, AX
        0x004a 00074 (/home/wsjcko/test/main.go:13)     IMUL3Q  $100000000, AX, AX
        0x0051 00081 (/home/wsjcko/test/main.go:13)     MOVQ    16(SP), BP
        0x0056 00086 (/home/wsjcko/test/main.go:13)     ADDQ    $24, SP
        0x005a 00090 (/home/wsjcko/test/main.go:13)     RET
        0x005b 00091 (/home/wsjcko/test/main.go:13)     NOP
        0x005b 00091 (/home/wsjcko/test/main.go:9)      PCDATA  $1, $-1
        0x005b 00091 (/home/wsjcko/test/main.go:9)      PCDATA  $0, $-2
        0x005b 00091 (/home/wsjcko/test/main.go:9)      MOVQ    AX, 8(SP)
        0x0060 00096 (/home/wsjcko/test/main.go:9)      CALL    runtime.morestack_noctxt(SB)
        0x0065 00101 (/home/wsjcko/test/main.go:9)      MOVQ    8(SP), AX
        0x006a 00106 (/home/wsjcko/test/main.go:9)      PCDATA  $0, $-1
        0x006a 00106 (/home/wsjcko/test/main.go:9)      JMP     0
        0x0000 49 3b 66 10 76 55 48 83 ec 18 48 89 6c 24 10 48  I;f.vUH...H.l$.H
        0x0010 8d 6c 24 10 48 83 f8 0d 7e 14 48 b8 00 b0 8e f0  .l$.H...~.H.....
        0x0020 1b 00 00 00 48 8b 6c 24 10 48 83 c4 18 c3 0f 57  ....H.l$.H.....W
        0x0030 c0 f2 48 0f 2a c0 f2 0f 10 0d 00 00 00 00 66 90  ..H.*.........f.
        0x0040 e8 00 00 00 00 f2 48 0f 2c c0 48 69 c0 00 e1 f5  ......H.,.Hi....
        0x0050 05 48 8b 6c 24 10 48 83 c4 18 c3 48 89 44 24 08  .H.l$.H....H.D$.
        0x0060 e8 00 00 00 00 48 8b 44 24 08 eb 94              .....H.D$...
        rel 58+4 t=14 $f64.4005bf0a8b145769+0
        rel 65+4 t=7 math.pow+0
        rel 97+4 t=7 runtime.morestack_noctxt+0
"".main STEXT size=103 args=0x0 locals=0x40 funcid=0x0 align=0x0
        0x0000 00000 (/home/wsjcko/test/main.go:16)     TEXT    "".main(SB), ABIInternal, $64-0
        0x0000 00000 (/home/wsjcko/test/main.go:16)     CMPQ    SP, 16(R14)
        0x0004 00004 (/home/wsjcko/test/main.go:16)     PCDATA  $0, $-2
        0x0004 00004 (/home/wsjcko/test/main.go:16)     JLS     92
        0x0006 00006 (/home/wsjcko/test/main.go:16)     PCDATA  $0, $-1
        0x0006 00006 (/home/wsjcko/test/main.go:16)     SUBQ    $64, SP
        0x000a 00010 (/home/wsjcko/test/main.go:16)     MOVQ    BP, 56(SP)
        0x000f 00015 (/home/wsjcko/test/main.go:16)     LEAQ    56(SP), BP
        0x0014 00020 (/home/wsjcko/test/main.go:16)     FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0014 00020 (/home/wsjcko/test/main.go:16)     FUNCDATA        $1, gclocals·f207267fbf96a0178e8758c6e3e0ce28(SB)
        0x0014 00020 (/home/wsjcko/test/main.go:16)     FUNCDATA        $2, "".main.stkobj(SB)
        0x0014 00020 (/home/wsjcko/test/main.go:18)     MOVUPS  X15, ""..autotmp_8+40(SP)
        0x001a 00026 (/home/wsjcko/test/main.go:18)     LEAQ    type.int(SB), DX
        0x0021 00033 (/home/wsjcko/test/main.go:18)     MOVQ    DX, ""..autotmp_8+40(SP)
        0x0026 00038 (/home/wsjcko/test/main.go:18)     LEAQ    ""..stmp_0(SB), DX
        0x002d 00045 (/home/wsjcko/test/main.go:18)     MOVQ    DX, ""..autotmp_8+48(SP)
        0x0032 00050 (<unknown line number>)    NOP
        0x0032 00050 ($GOROOT/src/fmt/print.go:274)     MOVQ    os.Stdout(SB), BX
        0x0039 00057 ($GOROOT/src/fmt/print.go:274)     LEAQ    go.itab.*os.File,io.Writer(SB), AX
        0x0040 00064 ($GOROOT/src/fmt/print.go:274)     LEAQ    ""..autotmp_8+40(SP), CX
        0x0045 00069 ($GOROOT/src/fmt/print.go:274)     MOVL    $1, DI
        0x004a 00074 ($GOROOT/src/fmt/print.go:274)     MOVQ    DI, SI
        0x004d 00077 ($GOROOT/src/fmt/print.go:274)     PCDATA  $1, $0
        0x004d 00077 ($GOROOT/src/fmt/print.go:274)     CALL    fmt.Fprintln(SB)
        0x0052 00082 (/home/wsjcko/test/main.go:19)     MOVQ    56(SP), BP
        0x0057 00087 (/home/wsjcko/test/main.go:19)     ADDQ    $64, SP
        0x005b 00091 (/home/wsjcko/test/main.go:19)     RET
        0x005c 00092 (/home/wsjcko/test/main.go:19)     NOP
        0x005c 00092 (/home/wsjcko/test/main.go:16)     PCDATA  $1, $-1
        0x005c 00092 (/home/wsjcko/test/main.go:16)     PCDATA  $0, $-2
        0x005c 00092 (/home/wsjcko/test/main.go:16)     NOP
        0x0060 00096 (/home/wsjcko/test/main.go:16)     CALL    runtime.morestack_noctxt(SB)
        0x0065 00101 (/home/wsjcko/test/main.go:16)     PCDATA  $0, $-1
        0x0065 00101 (/home/wsjcko/test/main.go:16)     JMP     0
        0x0000 49 3b 66 10 76 56 48 83 ec 40 48 89 6c 24 38 48  I;f.vVH..@H.l$8H
        0x0010 8d 6c 24 38 44 0f 11 7c 24 28 48 8d 15 00 00 00  .l$8D..|$(H.....
        0x0020 00 48 89 54 24 28 48 8d 15 00 00 00 00 48 89 54  .H.T$(H......H.T
        0x0030 24 30 48 8b 1d 00 00 00 00 48 8d 05 00 00 00 00  $0H......H......
        0x0040 48 8d 4c 24 28 bf 01 00 00 00 48 89 fe e8 00 00  H.L$(.....H.....
        0x0050 00 00 48 8b 6c 24 38 48 83 c4 40 c3 0f 1f 40 00  ..H.l$8H..@...@.
        0x0060 e8 00 00 00 00 eb 99                             .......
        rel 2+0 t=23 type.int+0
        rel 2+0 t=23 type.*os.File+0
        rel 29+4 t=14 type.int+0
        rel 41+4 t=14 ""..stmp_0+0
        rel 53+4 t=14 os.Stdout+0
        rel 60+4 t=14 go.itab.*os.File,io.Writer+0
        rel 78+4 t=7 fmt.Fprintln+0
        rel 97+4 t=7 runtime.morestack_noctxt+0
go.cuinfo.producer.main SDWARFCUINFO dupok size=0
        0x0000 72 65 67 61 62 69                                regabi
go.cuinfo.packagename.main SDWARFCUINFO dupok size=0
        0x0000 6d 61 69 6e                                      main
go.info.math.Pow$abstract SDWARFABSFCN dupok size=29
        0x0000 05 6d 61 74 68 2e 50 6f 77 00 01 01 13 78 00 00  .math.Pow....x..
        0x0010 00 00 00 00 13 79 00 00 00 00 00 00 00           .....y.......
        rel 0+0 t=22 type.float64+0
        rel 16+4 t=31 go.info.float64+0
        rel 24+4 t=31 go.info.float64+0
go.info.fmt.Println$abstract SDWARFABSFCN dupok size=42
        0x0000 05 66 6d 74 2e 50 72 69 6e 74 6c 6e 00 01 01 13  .fmt.Println....
        0x0010 61 00 00 00 00 00 00 13 6e 00 01 00 00 00 00 13  a.......n.......
        0x0020 65 72 72 00 01 00 00 00 00 00                    err.......
        rel 0+0 t=22 type.[]interface {}+0
        rel 0+0 t=22 type.error+0
        rel 0+0 t=22 type.int+0
        rel 19+4 t=31 go.info.[]interface {}+0
        rel 27+4 t=31 go.info.int+0
        rel 37+4 t=31 go.info.error+0

### 查看main Do方法的汇编指令
go tool objdump -S -s "Do" ./main
go tool objdump -S main.o  

go build -gcflags '-l' -o  testmain .
 go tool objdump -s "main" main

### 减少生成的二进制文件的大小，你可以分离执行过程中不需要的信息
go build -ldflags="-s -w" main.go 