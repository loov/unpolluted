package syscalls

import "golang.org/x/sys/unix"

var _ = unix.Exit

var Name = map[uint64]string{
	unix.SYS_RESTART_SYSCALL:        "restart_syscall",
	unix.SYS_EXIT:                   "exit",
	unix.SYS_FORK:                   "fork",
	unix.SYS_READ:                   "read",
	unix.SYS_WRITE:                  "write",
	unix.SYS_OPEN:                   "open",
	unix.SYS_CLOSE:                  "close",
	unix.SYS_WAITPID:                "waitpid",
	unix.SYS_CREAT:                  "creat",
	unix.SYS_LINK:                   "link",
	unix.SYS_UNLINK:                 "unlink",
	unix.SYS_EXECVE:                 "execve",
	unix.SYS_CHDIR:                  "chdir",
	unix.SYS_TIME:                   "time",
	unix.SYS_MKNOD:                  "mknod",
	unix.SYS_CHMOD:                  "chmod",
	unix.SYS_LCHOWN:                 "lchown",
	unix.SYS_BREAK:                  "break",
	unix.SYS_OLDSTAT:                "oldstat",
	unix.SYS_LSEEK:                  "lseek",
	unix.SYS_GETPID:                 "getpid",
	unix.SYS_MOUNT:                  "mount",
	unix.SYS_UMOUNT:                 "umount",
	unix.SYS_SETUID:                 "setuid",
	unix.SYS_GETUID:                 "getuid",
	unix.SYS_STIME:                  "stime",
	unix.SYS_PTRACE:                 "ptrace",
	unix.SYS_ALARM:                  "alarm",
	unix.SYS_OLDFSTAT:               "oldfstat",
	unix.SYS_PAUSE:                  "pause",
	unix.SYS_UTIME:                  "utime",
	unix.SYS_STTY:                   "stty",
	unix.SYS_GTTY:                   "gtty",
	unix.SYS_ACCESS:                 "access",
	unix.SYS_NICE:                   "nice",
	unix.SYS_FTIME:                  "ftime",
	unix.SYS_SYNC:                   "sync",
	unix.SYS_KILL:                   "kill",
	unix.SYS_RENAME:                 "rename",
	unix.SYS_MKDIR:                  "mkdir",
	unix.SYS_RMDIR:                  "rmdir",
	unix.SYS_DUP:                    "dup",
	unix.SYS_PIPE:                   "pipe",
	unix.SYS_TIMES:                  "times",
	unix.SYS_PROF:                   "prof",
	unix.SYS_BRK:                    "brk",
	unix.SYS_SETGID:                 "setgid",
	unix.SYS_GETGID:                 "getgid",
	unix.SYS_SIGNAL:                 "signal",
	unix.SYS_GETEUID:                "geteuid",
	unix.SYS_GETEGID:                "getegid",
	unix.SYS_ACCT:                   "acct",
	unix.SYS_UMOUNT2:                "umount2",
	unix.SYS_LOCK:                   "lock",
	unix.SYS_IOCTL:                  "ioctl",
	unix.SYS_FCNTL:                  "fcntl",
	unix.SYS_MPX:                    "mpx",
	unix.SYS_SETPGID:                "setpgid",
	unix.SYS_ULIMIT:                 "ulimit",
	unix.SYS_OLDOLDUNAME:            "oldolduname",
	unix.SYS_UMASK:                  "umask",
	unix.SYS_CHROOT:                 "chroot",
	unix.SYS_USTAT:                  "ustat",
	unix.SYS_DUP2:                   "dup2",
	unix.SYS_GETPPID:                "getppid",
	unix.SYS_GETPGRP:                "getpgrp",
	unix.SYS_SETSID:                 "setsid",
	unix.SYS_SIGACTION:              "sigaction",
	unix.SYS_SGETMASK:               "sgetmask",
	unix.SYS_SSETMASK:               "ssetmask",
	unix.SYS_SETREUID:               "setreuid",
	unix.SYS_SETREGID:               "setregid",
	unix.SYS_SIGSUSPEND:             "sigsuspend",
	unix.SYS_SIGPENDING:             "sigpending",
	unix.SYS_SETHOSTNAME:            "sethostname",
	unix.SYS_SETRLIMIT:              "setrlimit",
	unix.SYS_GETRLIMIT:              "getrlimit",
	unix.SYS_GETRUSAGE:              "getrusage",
	unix.SYS_GETTIMEOFDAY:           "gettimeofday",
	unix.SYS_SETTIMEOFDAY:           "settimeofday",
	unix.SYS_GETGROUPS:              "getgroups",
	unix.SYS_SETGROUPS:              "setgroups",
	unix.SYS_SELECT:                 "select",
	unix.SYS_SYMLINK:                "symlink",
	unix.SYS_OLDLSTAT:               "oldlstat",
	unix.SYS_READLINK:               "readlink",
	unix.SYS_USELIB:                 "uselib",
	unix.SYS_SWAPON:                 "swapon",
	unix.SYS_REBOOT:                 "reboot",
	unix.SYS_READDIR:                "readdir",
	unix.SYS_MMAP:                   "mmap",
	unix.SYS_MUNMAP:                 "munmap",
	unix.SYS_TRUNCATE:               "truncate",
	unix.SYS_FTRUNCATE:              "ftruncate",
	unix.SYS_FCHMOD:                 "fchmod",
	unix.SYS_FCHOWN:                 "fchown",
	unix.SYS_GETPRIORITY:            "getpriority",
	unix.SYS_SETPRIORITY:            "setpriority",
	unix.SYS_PROFIL:                 "profil",
	unix.SYS_STATFS:                 "statfs",
	unix.SYS_FSTATFS:                "fstatfs",
	unix.SYS_IOPERM:                 "ioperm",
	unix.SYS_SOCKETCALL:             "socketcall",
	unix.SYS_SYSLOG:                 "syslog",
	unix.SYS_SETITIMER:              "setitimer",
	unix.SYS_GETITIMER:              "getitimer",
	unix.SYS_STAT:                   "stat",
	unix.SYS_LSTAT:                  "lstat",
	unix.SYS_FSTAT:                  "fstat",
	unix.SYS_OLDUNAME:               "olduname",
	unix.SYS_IOPL:                   "iopl",
	unix.SYS_VHANGUP:                "vhangup",
	unix.SYS_IDLE:                   "idle",
	unix.SYS_VM86:                   "vm86",
	unix.SYS_WAIT4:                  "wait4",
	unix.SYS_SWAPOFF:                "swapoff",
	unix.SYS_SYSINFO:                "sysinfo",
	unix.SYS_IPC:                    "ipc",
	unix.SYS_FSYNC:                  "fsync",
	unix.SYS_SIGRETURN:              "sigreturn",
	unix.SYS_CLONE:                  "clone",
	unix.SYS_SETDOMAINNAME:          "setdomainname",
	unix.SYS_UNAME:                  "uname",
	unix.SYS_MODIFY_LDT:             "modify_ldt",
	unix.SYS_ADJTIMEX:               "adjtimex",
	unix.SYS_MPROTECT:               "mprotect",
	unix.SYS_SIGPROCMASK:            "sigprocmask",
	unix.SYS_CREATE_MODULE:          "create_module",
	unix.SYS_INIT_MODULE:            "init_module",
	unix.SYS_DELETE_MODULE:          "delete_module",
	unix.SYS_GET_KERNEL_SYMS:        "get_kernel_syms",
	unix.SYS_QUOTACTL:               "quotactl",
	unix.SYS_GETPGID:                "getpgid",
	unix.SYS_FCHDIR:                 "fchdir",
	unix.SYS_BDFLUSH:                "bdflush",
	unix.SYS_SYSFS:                  "sysfs",
	unix.SYS_PERSONALITY:            "personality",
	unix.SYS_AFS_SYSCALL:            "afs_syscall",
	unix.SYS_SETFSUID:               "setfsuid",
	unix.SYS_SETFSGID:               "setfsgid",
	unix.SYS__LLSEEK:                "_llseek",
	unix.SYS_GETDENTS:               "getdents",
	unix.SYS__NEWSELECT:             "_newselect",
	unix.SYS_FLOCK:                  "flock",
	unix.SYS_MSYNC:                  "msync",
	unix.SYS_READV:                  "readv",
	unix.SYS_WRITEV:                 "writev",
	unix.SYS_GETSID:                 "getsid",
	unix.SYS_FDATASYNC:              "fdatasync",
	unix.SYS__SYSCTL:                "_sysctl",
	unix.SYS_MLOCK:                  "mlock",
	unix.SYS_MUNLOCK:                "munlock",
	unix.SYS_MLOCKALL:               "mlockall",
	unix.SYS_MUNLOCKALL:             "munlockall",
	unix.SYS_SCHED_SETPARAM:         "sched_setparam",
	unix.SYS_SCHED_GETPARAM:         "sched_getparam",
	unix.SYS_SCHED_SETSCHEDULER:     "sched_setscheduler",
	unix.SYS_SCHED_GETSCHEDULER:     "sched_getscheduler",
	unix.SYS_SCHED_YIELD:            "sched_yield",
	unix.SYS_SCHED_GET_PRIORITY_MAX: "sched_get_priority_max",
	unix.SYS_SCHED_GET_PRIORITY_MIN: "sched_get_priority_min",
	unix.SYS_SCHED_RR_GET_INTERVAL:  "sched_rr_get_interval",
	unix.SYS_NANOSLEEP:              "nanosleep",
	unix.SYS_MREMAP:                 "mremap",
	unix.SYS_SETRESUID:              "setresuid",
	unix.SYS_GETRESUID:              "getresuid",
	unix.SYS_QUERY_MODULE:           "query_module",
	unix.SYS_POLL:                   "poll",
	unix.SYS_NFSSERVCTL:             "nfsservctl",
	unix.SYS_SETRESGID:              "setresgid",
	unix.SYS_GETRESGID:              "getresgid",
	unix.SYS_PRCTL:                  "prctl",
	unix.SYS_RT_SIGRETURN:           "rt_sigreturn",
	unix.SYS_RT_SIGACTION:           "rt_sigaction",
	unix.SYS_RT_SIGPROCMASK:         "rt_sigprocmask",
	unix.SYS_RT_SIGPENDING:          "rt_sigpending",
	unix.SYS_RT_SIGTIMEDWAIT:        "rt_sigtimedwait",
	unix.SYS_RT_SIGQUEUEINFO:        "rt_sigqueueinfo",
	unix.SYS_RT_SIGSUSPEND:          "rt_sigsuspend",
	unix.SYS_PREAD64:                "pread64",
	unix.SYS_PWRITE64:               "pwrite64",
	unix.SYS_CHOWN:                  "chown",
	unix.SYS_GETCWD:                 "getcwd",
	unix.SYS_CAPGET:                 "capget",
	unix.SYS_CAPSET:                 "capset",
	unix.SYS_SIGALTSTACK:            "sigaltstack",
	unix.SYS_SENDFILE:               "sendfile",
	unix.SYS_GETPMSG:                "getpmsg",
	unix.SYS_PUTPMSG:                "putpmsg",
	unix.SYS_VFORK:                  "vfork",
	unix.SYS_UGETRLIMIT:             "ugetrlimit",
	unix.SYS_READAHEAD:              "readahead",
	unix.SYS_PCICONFIG_READ:         "pciconfig_read",
	unix.SYS_PCICONFIG_WRITE:        "pciconfig_write",
	unix.SYS_PCICONFIG_IOBASE:       "pciconfig_iobase",
	unix.SYS_MULTIPLEXER:            "multiplexer",
	unix.SYS_GETDENTS64:             "getdents64",
	unix.SYS_PIVOT_ROOT:             "pivot_root",
	unix.SYS_MADVISE:                "madvise",
	unix.SYS_MINCORE:                "mincore",
	unix.SYS_GETTID:                 "gettid",
	unix.SYS_TKILL:                  "tkill",
	unix.SYS_SETXATTR:               "setxattr",
	unix.SYS_LSETXATTR:              "lsetxattr",
	unix.SYS_FSETXATTR:              "fsetxattr",
	unix.SYS_GETXATTR:               "getxattr",
	unix.SYS_LGETXATTR:              "lgetxattr",
	unix.SYS_FGETXATTR:              "fgetxattr",
	unix.SYS_LISTXATTR:              "listxattr",
	unix.SYS_LLISTXATTR:             "llistxattr",
	unix.SYS_FLISTXATTR:             "flistxattr",
	unix.SYS_REMOVEXATTR:            "removexattr",
	unix.SYS_LREMOVEXATTR:           "lremovexattr",
	unix.SYS_FREMOVEXATTR:           "fremovexattr",
	unix.SYS_FUTEX:                  "futex",
	unix.SYS_SCHED_SETAFFINITY:      "sched_setaffinity",
	unix.SYS_SCHED_GETAFFINITY:      "sched_getaffinity",
	unix.SYS_TUXCALL:                "tuxcall",
	unix.SYS_IO_SETUP:               "io_setup",
	unix.SYS_IO_DESTROY:             "io_destroy",
	unix.SYS_IO_GETEVENTS:           "io_getevents",
	unix.SYS_IO_SUBMIT:              "io_submit",
	unix.SYS_IO_CANCEL:              "io_cancel",
	unix.SYS_SET_TID_ADDRESS:        "set_tid_address",
	unix.SYS_FADVISE64:              "fadvise64",
	unix.SYS_EXIT_GROUP:             "exit_group",
	unix.SYS_LOOKUP_DCOOKIE:         "lookup_dcookie",
	unix.SYS_EPOLL_CREATE:           "epoll_create",
	unix.SYS_EPOLL_CTL:              "epoll_ctl",
	unix.SYS_EPOLL_WAIT:             "epoll_wait",
	unix.SYS_REMAP_FILE_PAGES:       "remap_file_pages",
	unix.SYS_TIMER_CREATE:           "timer_create",
	unix.SYS_TIMER_SETTIME:          "timer_settime",
	unix.SYS_TIMER_GETTIME:          "timer_gettime",
	unix.SYS_TIMER_GETOVERRUN:       "timer_getoverrun",
	unix.SYS_TIMER_DELETE:           "timer_delete",
	unix.SYS_CLOCK_SETTIME:          "clock_settime",
	unix.SYS_CLOCK_GETTIME:          "clock_gettime",
	unix.SYS_CLOCK_GETRES:           "clock_getres",
	unix.SYS_CLOCK_NANOSLEEP:        "clock_nanosleep",
	unix.SYS_SWAPCONTEXT:            "swapcontext",
	unix.SYS_TGKILL:                 "tgkill",
	unix.SYS_UTIMES:                 "utimes",
	unix.SYS_STATFS64:               "statfs64",
	unix.SYS_FSTATFS64:              "fstatfs64",
	unix.SYS_RTAS:                   "rtas",
	unix.SYS_SYS_DEBUG_SETCONTEXT:   "sys_debug_setcontext",
	unix.SYS_MIGRATE_PAGES:          "migrate_pages",
	unix.SYS_MBIND:                  "mbind",
	unix.SYS_GET_MEMPOLICY:          "get_mempolicy",
	unix.SYS_SET_MEMPOLICY:          "set_mempolicy",
	unix.SYS_MQ_OPEN:                "mq_open",
	unix.SYS_MQ_UNLINK:              "mq_unlink",
	unix.SYS_MQ_TIMEDSEND:           "mq_timedsend",
	unix.SYS_MQ_TIMEDRECEIVE:        "mq_timedreceive",
	unix.SYS_MQ_NOTIFY:              "mq_notify",
	unix.SYS_MQ_GETSETATTR:          "mq_getsetattr",
	unix.SYS_KEXEC_LOAD:             "kexec_load",
	unix.SYS_ADD_KEY:                "add_key",
	unix.SYS_REQUEST_KEY:            "request_key",
	unix.SYS_KEYCTL:                 "keyctl",
	unix.SYS_WAITID:                 "waitid",
	unix.SYS_IOPRIO_SET:             "ioprio_set",
	unix.SYS_IOPRIO_GET:             "ioprio_get",
	unix.SYS_INOTIFY_INIT:           "inotify_init",
	unix.SYS_INOTIFY_ADD_WATCH:      "inotify_add_watch",
	unix.SYS_INOTIFY_RM_WATCH:       "inotify_rm_watch",
	unix.SYS_SPU_RUN:                "spu_run",
	unix.SYS_SPU_CREATE:             "spu_create",
	unix.SYS_PSELECT6:               "pselect6",
	unix.SYS_PPOLL:                  "ppoll",
	unix.SYS_UNSHARE:                "unshare",
	unix.SYS_SPLICE:                 "splice",
	unix.SYS_TEE:                    "tee",
	unix.SYS_VMSPLICE:               "vmsplice",
	unix.SYS_OPENAT:                 "openat",
	unix.SYS_MKDIRAT:                "mkdirat",
	unix.SYS_MKNODAT:                "mknodat",
	unix.SYS_FCHOWNAT:               "fchownat",
	unix.SYS_FUTIMESAT:              "futimesat",
	unix.SYS_NEWFSTATAT:             "newfstatat",
	unix.SYS_UNLINKAT:               "unlinkat",
	unix.SYS_RENAMEAT:               "renameat",
	unix.SYS_LINKAT:                 "linkat",
	unix.SYS_SYMLINKAT:              "symlinkat",
	unix.SYS_READLINKAT:             "readlinkat",
	unix.SYS_FCHMODAT:               "fchmodat",
	unix.SYS_FACCESSAT:              "faccessat",
	unix.SYS_GET_ROBUST_LIST:        "get_robust_list",
	unix.SYS_SET_ROBUST_LIST:        "set_robust_list",
	unix.SYS_MOVE_PAGES:             "move_pages",
	unix.SYS_GETCPU:                 "getcpu",
	unix.SYS_EPOLL_PWAIT:            "epoll_pwait",
	unix.SYS_UTIMENSAT:              "utimensat",
	unix.SYS_SIGNALFD:               "signalfd",
	unix.SYS_TIMERFD_CREATE:         "timerfd_create",
	unix.SYS_EVENTFD:                "eventfd",
	unix.SYS_SYNC_FILE_RANGE2:       "sync_file_range2",
	unix.SYS_FALLOCATE:              "fallocate",
	unix.SYS_SUBPAGE_PROT:           "subpage_prot",
	unix.SYS_TIMERFD_SETTIME:        "timerfd_settime",
	unix.SYS_TIMERFD_GETTIME:        "timerfd_gettime",
	unix.SYS_SIGNALFD4:              "signalfd4",
	unix.SYS_EVENTFD2:               "eventfd2",
	unix.SYS_EPOLL_CREATE1:          "epoll_create1",
	unix.SYS_DUP3:                   "dup3",
	unix.SYS_PIPE2:                  "pipe2",
	unix.SYS_INOTIFY_INIT1:          "inotify_init1",
	unix.SYS_PERF_EVENT_OPEN:        "perf_event_open",
	unix.SYS_PREADV:                 "preadv",
	unix.SYS_PWRITEV:                "pwritev",
	unix.SYS_RT_TGSIGQUEUEINFO:      "rt_tgsigqueueinfo",
	unix.SYS_FANOTIFY_INIT:          "fanotify_init",
	unix.SYS_FANOTIFY_MARK:          "fanotify_mark",
	unix.SYS_PRLIMIT64:              "prlimit64",
	unix.SYS_SOCKET:                 "socket",
	unix.SYS_BIND:                   "bind",
	unix.SYS_CONNECT:                "connect",
	unix.SYS_LISTEN:                 "listen",
	unix.SYS_ACCEPT:                 "accept",
	unix.SYS_GETSOCKNAME:            "getsockname",
	unix.SYS_GETPEERNAME:            "getpeername",
	unix.SYS_SOCKETPAIR:             "socketpair",
	unix.SYS_SEND:                   "send",
	unix.SYS_SENDTO:                 "sendto",
	unix.SYS_RECV:                   "recv",
	unix.SYS_RECVFROM:               "recvfrom",
	unix.SYS_SHUTDOWN:               "shutdown",
	unix.SYS_SETSOCKOPT:             "setsockopt",
	unix.SYS_GETSOCKOPT:             "getsockopt",
	unix.SYS_SENDMSG:                "sendmsg",
	unix.SYS_RECVMSG:                "recvmsg",
	unix.SYS_RECVMMSG:               "recvmmsg",
	unix.SYS_ACCEPT4:                "accept4",
	unix.SYS_NAME_TO_HANDLE_AT:      "name_to_handle_at",
	unix.SYS_OPEN_BY_HANDLE_AT:      "open_by_handle_at",
	unix.SYS_CLOCK_ADJTIME:          "clock_adjtime",
	unix.SYS_SYNCFS:                 "syncfs",
	unix.SYS_SENDMMSG:               "sendmmsg",
	unix.SYS_SETNS:                  "setns",
	unix.SYS_PROCESS_VM_READV:       "process_vm_readv",
	unix.SYS_PROCESS_VM_WRITEV:      "process_vm_writev",
	unix.SYS_FINIT_MODULE:           "finit_module",
	unix.SYS_KCMP:                   "kcmp",
	unix.SYS_SCHED_SETATTR:          "sched_setattr",
	unix.SYS_SCHED_GETATTR:          "sched_getattr",
	unix.SYS_RENAMEAT2:              "renameat2",
	unix.SYS_SECCOMP:                "seccomp",
	unix.SYS_GETRANDOM:              "getrandom",
	unix.SYS_MEMFD_CREATE:           "memfd_create",
	unix.SYS_BPF:                    "bpf",
	unix.SYS_EXECVEAT:               "execveat",
	unix.SYS_SWITCH_ENDIAN:          "switch_endian",
	unix.SYS_USERFAULTFD:            "userfaultfd",
	unix.SYS_MEMBARRIER:             "membarrier",
	unix.SYS_MLOCK2:                 "mlock2",
	unix.SYS_COPY_FILE_RANGE:        "copy_file_range",
	unix.SYS_PREADV2:                "preadv2",
	unix.SYS_PWRITEV2:               "pwritev2",
	unix.SYS_KEXEC_FILE_LOAD:        "kexec_file_load",
	unix.SYS_STATX:                  "statx",
	unix.SYS_PKEY_ALLOC:             "pkey_alloc",
	unix.SYS_PKEY_FREE:              "pkey_free",
	unix.SYS_PKEY_MPROTECT:          "pkey_mprotect",
}
