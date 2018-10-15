package ptrace

import (
	"syscall"

	"github.com/loov/leakcheck/api"
)

func registersToCall(pid int, registers syscall.PtraceRegs) api.Call {
	raw := api.Syscall{
		Number: registers.Orig_rax,
	}

	switch raw.Number {
	case syscall.SYS_OPEN:
		return api.Open{
			Syscall:  raw,
			Path:     stringArgument(pid, uintptr(registers.Rdi)),
			ResultFD: int64(registers.Rax),
			Failed:   registers.Rax < 0,
		}
	case syscall.SYS_OPENAT:
		return api.Open{
			Syscall:  raw,
			Path:     stringArgument(pid, uintptr(registers.Rsi)),
			ResultFD: int64(registers.Rax),
			Failed:   registers.Rax < 0,
		}

	case syscall.SYS_CLOSE:
		return api.Close{
			Syscall: raw,
			FD:      int64(registers.Rdi),
			Failed:  registers.Rax != 0,
		}

	case syscall.SYS_UNLINK:
		return api.Unlink{
			Syscall: raw,
			Path:    stringArgument(pid, uintptr(registers.Rdi)),
			Failed:  registers.Rax != 0,
		}
	case syscall.SYS_UNLINKAT:
		return api.Unlink{
			Syscall: raw,
			Path:    stringArgument(pid, uintptr(registers.Rsi)),
			Failed:  registers.Rax != 0,
		}

	case syscall.SYS_SOCKET:
		return api.Socket{
			Syscall:  raw,
			ResultFD: int64(registers.Rax),
			Failed:   registers.Rax < 0,
		}
	case syscall.SYS_BIND:
		return api.Bind{
			Syscall: raw,
			FD:      int64(registers.Orig_rax),
			Failed:  registers.Rax != 0,
		}
	}

	return raw
}