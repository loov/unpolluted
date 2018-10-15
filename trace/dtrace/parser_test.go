package dtrace

import (
	"io"
	"strings"
	"testing"
)

const dtruss = `
SYSCALL(args) 		 = return
PASS
open("/dev/dtracehelper\0", 0x2, 0xFFFFFFFFEFBFEFD0)		 = 3 0
ioctl(0x3, 0x80086804, 0x7FFEEFBFEDE0)		 = 0 0
close(0x3)		 = 0 0
access("/AppleInternal/XBS/.isChrooted\0", 0x0, 0x0)		 = -1 Err#2
bsdthread_register(0x7FFF60BA9418, 0x7FFF60BA9408, 0x2000)		 = 1073742047 0
sysctlbyname(kern.bootargs, 0xD, 0x7FFEEFBFE230, 0x7FFEEFBFE228, 0x0)		 = 0 0
issetugid(0x0, 0x0, 0x0)		 = 0 0
ioctl(0x2, 0x4004667A, 0x7FFEEFBFDA44)		 = 0 0
mprotect(0x1277000, 0x1000, 0x0)		 = 0 0
mprotect(0x127C000, 0x1000, 0x0)		 = 0 0
mprotect(0x127D000, 0x1000, 0x0)		 = 0 0
mprotect(0x1282000, 0x1000, 0x0)		 = 0 0
mprotect(0x1275000, 0x90, 0x1)		 = 0 0
mprotect(0x1283000, 0x1000, 0x1)		 = 0 0
mprotect(0x1275000, 0x90, 0x3)		 = 0 0
mprotect(0x1275000, 0x90, 0x1)		 = 0 0
getpid(0x0, 0x0, 0x0)		 = 77991 0
stat64("/AppleInternal\0", 0x7FFEEFBFE6A0, 0x0)		 = -1 Err#2
csops(0x130A7, 0x7, 0x7FFEEFBFE1D0)		 = -1 Err#22
proc_info(0x2, 0x130A7, 0xD)		 = 64 0
csops(0x130A7, 0x7, 0x7FFEEFBFDA20)		 = -1 Err#22
dtrace: error on enabled probe ID 2198 (ID 557: syscall::sysctl:return): invalid kernel access in action #10 at DIF offset 28
dtrace: error on enabled probe ID 2198 (ID 557: syscall::sysctl:return): invalid kernel access in action #10 at DIF offset 28
mmap(0x0, 0x40000, 0x3, 0x1002, 0xFFFFFFFFFFFFFFFF, 0x0)		 = 0x1284000 0
mmap(0xC000000000, 0x4000000, 0x0, 0x1002, 0xFFFFFFFFFFFFFFFF, 0x0)		 = 0xC000000000 0
mmap(0xC000000000, 0x4000000, 0x3, 0x1012, 0xFFFFFFFFFFFFFFFF, 0x0)		 = 0xC000000000 0
mmap(0x0, 0x2000000, 0x3, 0x1002, 0xFFFFFFFFFFFFFFFF, 0x0)		 = 0x2000000 0
mmap(0x0, 0x210000, 0x3, 0x1002, 0xFFFFFFFFFFFFFFFF, 0x0)		 = 0x1400000 0
mmap(0x0, 0x10000, 0x3, 0x1002, 0xFFFFFFFFFFFFFFFF, 0x0)		 = 0x12C4000 0
mmap(0x0, 0x10000, 0x3, 0x1002, 0xFFFFFFFFFFFFFFFF, 0x0)		 = 0x12D4000 0
open("/dev/urandom\0", 0x0, 0x0)		 = 3 0
dtrace: error on enabled probe ID 2174 (ID 159: syscall::read:return): invalid kernel access in action #12 at DIF offset 68
close(0x3)		 = 0 0
__pthread_sigmask(0x3, 0x0, 0x1224D00)		 = 0 0
sigaltstack(0x0, 0x7FFEEFBFF980, 0x0)		 = 0 0
sigaltstack(0x7FFEEFBFF940, 0x0, 0x0)		 = 0 0
__pthread_sigmask(0x3, 0x7FFEEFBFF994, 0x0)		 = 0 0
sigaction(0x1, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0x1, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0x2, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0x2, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0x3, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0x3, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0x4, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0x4, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0x5, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0x5, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0x6, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0x6, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0x7, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0x7, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0x8, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0x8, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0xA, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0xA, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0xB, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0xB, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0xC, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0xC, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0xD, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0xD, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0xE, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0xE, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0xF, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0xF, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0x10, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0x10, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0x14, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0x14, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0x17, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0x17, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0x18, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0x18, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0x19, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0x19, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0x1A, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0x1A, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0x1B, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0x1B, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0x1C, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0x1C, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0x1D, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0x1D, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0x1E, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0x1E, 0x7FFEEFBFF808, 0x0)		 = 0 0
sigaction(0x1F, 0x0, 0x7FFEEFBFF940)		 = 0 0
sigaction(0x1F, 0x7FFEEFBFF808, 0x0)		 = 0 0
__pthread_sigmask(0x3, 0x120D12C, 0x7FFEEFBFF8D0)		 = 0 0
bsdthread_create(0x10589F0, 0xC000034000, 0x700009258000)		 = 153452544 0
thread_selfid(0x0, 0x0, 0x0)		 = 11007692 0
__pthread_sigmask(0x3, 0x7FFEEFBFF8D0, 0x0)		 = 0 0
__pthread_sigmask(0x3, 0x120D12C, 0x7FFEEFBFF7E0)		 = 0 0
sigaltstack(0x0, 0x700009257E70, 0x0)		 = 0 0
sigaltstack(0x700009257E30, 0x0, 0x0)		 = 0 0
__pthread_sigmask(0x3, 0x700009257E84, 0x0)		 = 0 0
bsdthread_create(0x10589F0, 0xC000034400, 0x70000926B000)		 = 153530368 0
__pthread_sigmask(0x3, 0x7FFEEFBFF7E0, 0x0)		 = 0 0
thread_selfid(0x0, 0x0, 0x0)		 = 11007693 0
sigaltstack(0x0, 0x70000926AE70, 0x0)		 = 0 0
sigaltstack(0x70000926AE30, 0x0, 0x0)		 = 0 0
__pthread_sigmask(0x3, 0x70000926AE84, 0x0)		 = 0 0
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__pthread_sigmask(0x3, 0x120D12C, 0x7FFEEFBFF638)		 = 0 0
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
bsdthread_create(0x10589F0, 0xC000034800, 0x70000927E000)		 = 153608192 0
__pthread_sigmask(0x3, 0x120D12C, 0x70000926AD08)		 = 0 0
__pthread_sigmask(0x3, 0x7FFEEFBFF638, 0x0)		 = 0 0
thread_selfid(0x0, 0x0, 0x0)		 = 11007694 0
bsdthread_create(0x10589F0, 0xC000052000, 0x700009291000)		 = 153686016 0
sigaltstack(0x0, 0x70000927DE70, 0x0)		 = 0 0
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__pthread_sigmask(0x3, 0x70000926AD08, 0x0)		 = 0 0
thread_selfid(0x0, 0x0, 0x0)		 = 11007695 0
sigaltstack(0x70000927DE30, 0x0, 0x0)		 = 0 0
psynch_cvsignal(0x1225000, 0x100, 0x0)		 = 257 0
sigaltstack(0x0, 0x700009290E70, 0x0)		 = 0 0
__pthread_sigmask(0x3, 0x70000927DE84, 0x0)		 = 0 0
psynch_cvwait(0x1225000, 0x100000100, 0x0)		 = 0 0
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
sigaltstack(0x700009290E30, 0x0, 0x0)		 = 0 0
__pthread_sigmask(0x3, 0x700009290E84, 0x0)		 = 0 0
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
stat64(".\0", 0xC000066108, 0x0)		 = 0 0
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
stat64("/Users/egon/go/src/github.com/loov/leakcheck\0", 0xC000066378, 0x0)		 = 0 0
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
fcntl(0x0, 0x3, 0x0)		 = 2 0
psynch_cvsignal(0xC000052380, 0x100, 0x0)		 = 257 0
psynch_cvwait(0xC000052380, 0x100000100, 0x0)		 = 0 0
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
mmap(0x0, 0x40000, 0x3, 0x1002, 0xFFFFFFFFFFFFFFFF, 0x0)		 = 0x1610000 0
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
fcntl(0x1, 0x3, 0x0)		 = 2 0
psynch_cvsignal(0xC000034B80, 0x100, 0x0)		 = 257 0
psynch_cvwait(0xC000034B80, 0x100000100, 0x0)		 = 0 0
fcntl(0x2, 0x3, 0x0)		 = 2 0
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
psynch_cvsignal(0xC000034B80, 0x10000000200, 0x100)		 = 257 0
psynch_cvwait(0xC000034B80, 0x10100000200, 0x100)		 = 0 0
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
getpid(0x0, 0x0, 0x0)		 = 77991 0
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
open("/tmp/leak-717100594.txt\0", 0x1000A02, 0x180)		 = 3 0
fstat64(0x3, 0xC00002C550, 0x0)		 = 0 0
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
kqueue(0x0, 0x0, 0x0)		 = 4 0
fcntl(0x4, 0x2, 0x1)		 = 0 0
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
kevent(0x4, 0xC00002C420, 0x2)		 = 0 0
fcntl(0x3, 0x3, 0x0)		 = 2 0
fcntl(0x3, 0x4, 0x6)		 = 0 0
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
psynch_cvsignal(0xC000034B80, 0x20000000300, 0x200)		 = 257 0
psynch_cvwait(0xC000034B80, 0x20100000300, 0x200)		 = 0 0
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
dtrace: error on enabled probe ID 2172 (ID 161: syscall::write:return): invalid kernel access in action #12 at DIF offset 68
__semwait_signal(0xA03, 0x0, 0x1)		 = -1 Err#60
psynch_cvwait(0xC000034780, 0x100000100, 0x0)		 = -1 Err#260
psynch_cvwait(0xC000034B80, 0x30100000400, 0x300)		 = -1 Err#260
psynch_cvwait(0xC000052380, 0x10100000200, 0x100)		 = -1 Err#260
`

func TestParser(t *testing.T) {
	parser := NewParser(strings.NewReader(dtruss))
	for {
		call, err := parser.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		t.Log(call)
	}
}
