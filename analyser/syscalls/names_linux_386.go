package syscalls

import "syscall"

var _ = syscall.Exit

var Name = map[int64]string{
	syscall.SYS_RESTART_SYSCALL:        "restart_syscall",
	syscall.SYS_EXIT:                   "exit",
	syscall.SYS_FORK:                   "fork",
	syscall.SYS_READ:                   "read",
	syscall.SYS_WRITE:                  "write",
	syscall.SYS_OPEN:                   "open",
	syscall.SYS_CLOSE:                  "close",
	syscall.SYS_WAITPID:                "waitpid",
	syscall.SYS_CREAT:                  "creat",
	syscall.SYS_LINK:                   "link",
	syscall.SYS_UNLINK:                 "unlink",
	syscall.SYS_EXECVE:                 "execve",
	syscall.SYS_CHDIR:                  "chdir",
	syscall.SYS_TIME:                   "time",
	syscall.SYS_MKNOD:                  "mknod",
	syscall.SYS_CHMOD:                  "chmod",
	syscall.SYS_LCHOWN:                 "lchown",
	syscall.SYS_BREAK:                  "break",
	syscall.SYS_OLDSTAT:                "oldstat",
	syscall.SYS_LSEEK:                  "lseek",
	syscall.SYS_GETPID:                 "getpid",
	syscall.SYS_MOUNT:                  "mount",
	syscall.SYS_UMOUNT:                 "umount",
	syscall.SYS_SETUID:                 "setuid",
	syscall.SYS_GETUID:                 "getuid",
	syscall.SYS_STIME:                  "stime",
	syscall.SYS_PTRACE:                 "ptrace",
	syscall.SYS_ALARM:                  "alarm",
	syscall.SYS_OLDFSTAT:               "oldfstat",
	syscall.SYS_PAUSE:                  "pause",
	syscall.SYS_UTIME:                  "utime",
	syscall.SYS_STTY:                   "stty",
	syscall.SYS_GTTY:                   "gtty",
	syscall.SYS_ACCESS:                 "access",
	syscall.SYS_NICE:                   "nice",
	syscall.SYS_FTIME:                  "ftime",
	syscall.SYS_SYNC:                   "sync",
	syscall.SYS_KILL:                   "kill",
	syscall.SYS_RENAME:                 "rename",
	syscall.SYS_MKDIR:                  "mkdir",
	syscall.SYS_RMDIR:                  "rmdir",
	syscall.SYS_DUP:                    "dup",
	syscall.SYS_PIPE:                   "pipe",
	syscall.SYS_TIMES:                  "times",
	syscall.SYS_PROF:                   "prof",
	syscall.SYS_BRK:                    "brk",
	syscall.SYS_SETGID:                 "setgid",
	syscall.SYS_GETGID:                 "getgid",
	syscall.SYS_SIGNAL:                 "signal",
	syscall.SYS_GETEUID:                "geteuid",
	syscall.SYS_GETEGID:                "getegid",
	syscall.SYS_ACCT:                   "acct",
	syscall.SYS_UMOUNT2:                "umount2",
	syscall.SYS_LOCK:                   "lock",
	syscall.SYS_IOCTL:                  "ioctl",
	syscall.SYS_FCNTL:                  "fcntl",
	syscall.SYS_MPX:                    "mpx",
	syscall.SYS_SETPGID:                "setpgid",
	syscall.SYS_ULIMIT:                 "ulimit",
	syscall.SYS_OLDOLDUNAME:            "oldolduname",
	syscall.SYS_UMASK:                  "umask",
	syscall.SYS_CHROOT:                 "chroot",
	syscall.SYS_USTAT:                  "ustat",
	syscall.SYS_DUP2:                   "dup2",
	syscall.SYS_GETPPID:                "getppid",
	syscall.SYS_GETPGRP:                "getpgrp",
	syscall.SYS_SETSID:                 "setsid",
	syscall.SYS_SIGACTION:              "sigaction",
	syscall.SYS_SGETMASK:               "sgetmask",
	syscall.SYS_SSETMASK:               "ssetmask",
	syscall.SYS_SETREUID:               "setreuid",
	syscall.SYS_SETREGID:               "setregid",
	syscall.SYS_SIGSUSPEND:             "sigsuspend",
	syscall.SYS_SIGPENDING:             "sigpending",
	syscall.SYS_SETHOSTNAME:            "sethostname",
	syscall.SYS_SETRLIMIT:              "setrlimit",
	syscall.SYS_GETRLIMIT:              "getrlimit",
	syscall.SYS_GETRUSAGE:              "getrusage",
	syscall.SYS_GETTIMEOFDAY:           "gettimeofday",
	syscall.SYS_SETTIMEOFDAY:           "settimeofday",
	syscall.SYS_GETGROUPS:              "getgroups",
	syscall.SYS_SETGROUPS:              "setgroups",
	syscall.SYS_SELECT:                 "select",
	syscall.SYS_SYMLINK:                "symlink",
	syscall.SYS_OLDLSTAT:               "oldlstat",
	syscall.SYS_READLINK:               "readlink",
	syscall.SYS_USELIB:                 "uselib",
	syscall.SYS_SWAPON:                 "swapon",
	syscall.SYS_REBOOT:                 "reboot",
	syscall.SYS_READDIR:                "readdir",
	syscall.SYS_MMAP:                   "mmap",
	syscall.SYS_MUNMAP:                 "munmap",
	syscall.SYS_TRUNCATE:               "truncate",
	syscall.SYS_FTRUNCATE:              "ftruncate",
	syscall.SYS_FCHMOD:                 "fchmod",
	syscall.SYS_FCHOWN:                 "fchown",
	syscall.SYS_GETPRIORITY:            "getpriority",
	syscall.SYS_SETPRIORITY:            "setpriority",
	syscall.SYS_PROFIL:                 "profil",
	syscall.SYS_STATFS:                 "statfs",
	syscall.SYS_FSTATFS:                "fstatfs",
	syscall.SYS_IOPERM:                 "ioperm",
	syscall.SYS_SOCKETCALL:             "socketcall",
	syscall.SYS_SYSLOG:                 "syslog",
	syscall.SYS_SETITIMER:              "setitimer",
	syscall.SYS_GETITIMER:              "getitimer",
	syscall.SYS_STAT:                   "stat",
	syscall.SYS_LSTAT:                  "lstat",
	syscall.SYS_FSTAT:                  "fstat",
	syscall.SYS_OLDUNAME:               "olduname",
	syscall.SYS_IOPL:                   "iopl",
	syscall.SYS_VHANGUP:                "vhangup",
	syscall.SYS_IDLE:                   "idle",
	syscall.SYS_VM86OLD:                "vm86old",
	syscall.SYS_WAIT4:                  "wait4",
	syscall.SYS_SWAPOFF:                "swapoff",
	syscall.SYS_SYSINFO:                "sysinfo",
	syscall.SYS_IPC:                    "ipc",
	syscall.SYS_FSYNC:                  "fsync",
	syscall.SYS_SIGRETURN:              "sigreturn",
	syscall.SYS_CLONE:                  "clone",
	syscall.SYS_SETDOMAINNAME:          "setdomainname",
	syscall.SYS_UNAME:                  "uname",
	syscall.SYS_MODIFY_LDT:             "modify_ldt",
	syscall.SYS_ADJTIMEX:               "adjtimex",
	syscall.SYS_MPROTECT:               "mprotect",
	syscall.SYS_SIGPROCMASK:            "sigprocmask",
	syscall.SYS_CREATE_MODULE:          "create_module",
	syscall.SYS_INIT_MODULE:            "init_module",
	syscall.SYS_DELETE_MODULE:          "delete_module",
	syscall.SYS_GET_KERNEL_SYMS:        "get_kernel_syms",
	syscall.SYS_QUOTACTL:               "quotactl",
	syscall.SYS_GETPGID:                "getpgid",
	syscall.SYS_FCHDIR:                 "fchdir",
	syscall.SYS_BDFLUSH:                "bdflush",
	syscall.SYS_SYSFS:                  "sysfs",
	syscall.SYS_PERSONALITY:            "personality",
	syscall.SYS_AFS_SYSCALL:            "afs_syscall",
	syscall.SYS_SETFSUID:               "setfsuid",
	syscall.SYS_SETFSGID:               "setfsgid",
	syscall.SYS__LLSEEK:                "_llseek",
	syscall.SYS_GETDENTS:               "getdents",
	syscall.SYS__NEWSELECT:             "_newselect",
	syscall.SYS_FLOCK:                  "flock",
	syscall.SYS_MSYNC:                  "msync",
	syscall.SYS_READV:                  "readv",
	syscall.SYS_WRITEV:                 "writev",
	syscall.SYS_GETSID:                 "getsid",
	syscall.SYS_FDATASYNC:              "fdatasync",
	syscall.SYS__SYSCTL:                "_sysctl",
	syscall.SYS_MLOCK:                  "mlock",
	syscall.SYS_MUNLOCK:                "munlock",
	syscall.SYS_MLOCKALL:               "mlockall",
	syscall.SYS_MUNLOCKALL:             "munlockall",
	syscall.SYS_SCHED_SETPARAM:         "sched_setparam",
	syscall.SYS_SCHED_GETPARAM:         "sched_getparam",
	syscall.SYS_SCHED_SETSCHEDULER:     "sched_setscheduler",
	syscall.SYS_SCHED_GETSCHEDULER:     "sched_getscheduler",
	syscall.SYS_SCHED_YIELD:            "sched_yield",
	syscall.SYS_SCHED_GET_PRIORITY_MAX: "sched_get_priority_max",
	syscall.SYS_SCHED_GET_PRIORITY_MIN: "sched_get_priority_min",
	syscall.SYS_SCHED_RR_GET_INTERVAL:  "sched_rr_get_interval",
	syscall.SYS_NANOSLEEP:              "nanosleep",
	syscall.SYS_MREMAP:                 "mremap",
	syscall.SYS_SETRESUID:              "setresuid",
	syscall.SYS_GETRESUID:              "getresuid",
	syscall.SYS_VM86:                   "vm86",
	syscall.SYS_QUERY_MODULE:           "query_module",
	syscall.SYS_POLL:                   "poll",
	syscall.SYS_NFSSERVCTL:             "nfsservctl",
	syscall.SYS_SETRESGID:              "setresgid",
	syscall.SYS_GETRESGID:              "getresgid",
	syscall.SYS_PRCTL:                  "prctl",
	syscall.SYS_RT_SIGRETURN:           "rt_sigreturn",
	syscall.SYS_RT_SIGACTION:           "rt_sigaction",
	syscall.SYS_RT_SIGPROCMASK:         "rt_sigprocmask",
	syscall.SYS_RT_SIGPENDING:          "rt_sigpending",
	syscall.SYS_RT_SIGTIMEDWAIT:        "rt_sigtimedwait",
	syscall.SYS_RT_SIGQUEUEINFO:        "rt_sigqueueinfo",
	syscall.SYS_RT_SIGSUSPEND:          "rt_sigsuspend",
	syscall.SYS_PREAD64:                "pread64",
	syscall.SYS_PWRITE64:               "pwrite64",
	syscall.SYS_CHOWN:                  "chown",
	syscall.SYS_GETCWD:                 "getcwd",
	syscall.SYS_CAPGET:                 "capget",
	syscall.SYS_CAPSET:                 "capset",
	syscall.SYS_SIGALTSTACK:            "sigaltstack",
	syscall.SYS_SENDFILE:               "sendfile",
	syscall.SYS_GETPMSG:                "getpmsg",
	syscall.SYS_PUTPMSG:                "putpmsg",
	syscall.SYS_VFORK:                  "vfork",
	syscall.SYS_UGETRLIMIT:             "ugetrlimit",
	syscall.SYS_MMAP2:                  "mmap2",
	syscall.SYS_TRUNCATE64:             "truncate64",
	syscall.SYS_FTRUNCATE64:            "ftruncate64",
	syscall.SYS_STAT64:                 "stat64",
	syscall.SYS_LSTAT64:                "lstat64",
	syscall.SYS_FSTAT64:                "fstat64",
	syscall.SYS_LCHOWN32:               "lchown32",
	syscall.SYS_GETUID32:               "getuid32",
	syscall.SYS_GETGID32:               "getgid32",
	syscall.SYS_GETEUID32:              "geteuid32",
	syscall.SYS_GETEGID32:              "getegid32",
	syscall.SYS_SETREUID32:             "setreuid32",
	syscall.SYS_SETREGID32:             "setregid32",
	syscall.SYS_GETGROUPS32:            "getgroups32",
	syscall.SYS_SETGROUPS32:            "setgroups32",
	syscall.SYS_FCHOWN32:               "fchown32",
	syscall.SYS_SETRESUID32:            "setresuid32",
	syscall.SYS_GETRESUID32:            "getresuid32",
	syscall.SYS_SETRESGID32:            "setresgid32",
	syscall.SYS_GETRESGID32:            "getresgid32",
	syscall.SYS_CHOWN32:                "chown32",
	syscall.SYS_SETUID32:               "setuid32",
	syscall.SYS_SETGID32:               "setgid32",
	syscall.SYS_SETFSUID32:             "setfsuid32",
	syscall.SYS_SETFSGID32:             "setfsgid32",
	syscall.SYS_PIVOT_ROOT:             "pivot_root",
	syscall.SYS_MINCORE:                "mincore",
	syscall.SYS_MADVISE:                "madvise",
	syscall.SYS_MADVISE1:               "madvise1",
	syscall.SYS_GETDENTS64:             "getdents64",
	syscall.SYS_FCNTL64:                "fcntl64",
	syscall.SYS_GETTID:                 "gettid",
	syscall.SYS_READAHEAD:              "readahead",
	syscall.SYS_SETXATTR:               "setxattr",
	syscall.SYS_LSETXATTR:              "lsetxattr",
	syscall.SYS_FSETXATTR:              "fsetxattr",
	syscall.SYS_GETXATTR:               "getxattr",
	syscall.SYS_LGETXATTR:              "lgetxattr",
	syscall.SYS_FGETXATTR:              "fgetxattr",
	syscall.SYS_LISTXATTR:              "listxattr",
	syscall.SYS_LLISTXATTR:             "llistxattr",
	syscall.SYS_FLISTXATTR:             "flistxattr",
	syscall.SYS_REMOVEXATTR:            "removexattr",
	syscall.SYS_LREMOVEXATTR:           "lremovexattr",
	syscall.SYS_FREMOVEXATTR:           "fremovexattr",
	syscall.SYS_TKILL:                  "tkill",
	syscall.SYS_SENDFILE64:             "sendfile64",
	syscall.SYS_FUTEX:                  "futex",
	syscall.SYS_SCHED_SETAFFINITY:      "sched_setaffinity",
	syscall.SYS_SCHED_GETAFFINITY:      "sched_getaffinity",
	syscall.SYS_SET_THREAD_AREA:        "set_thread_area",
	syscall.SYS_GET_THREAD_AREA:        "get_thread_area",
	syscall.SYS_IO_SETUP:               "io_setup",
	syscall.SYS_IO_DESTROY:             "io_destroy",
	syscall.SYS_IO_GETEVENTS:           "io_getevents",
	syscall.SYS_IO_SUBMIT:              "io_submit",
	syscall.SYS_IO_CANCEL:              "io_cancel",
	syscall.SYS_FADVISE64:              "fadvise64",
	syscall.SYS_EXIT_GROUP:             "exit_group",
	syscall.SYS_LOOKUP_DCOOKIE:         "lookup_dcookie",
	syscall.SYS_EPOLL_CREATE:           "epoll_create",
	syscall.SYS_EPOLL_CTL:              "epoll_ctl",
	syscall.SYS_EPOLL_WAIT:             "epoll_wait",
	syscall.SYS_REMAP_FILE_PAGES:       "remap_file_pages",
	syscall.SYS_SET_TID_ADDRESS:        "set_tid_address",
	syscall.SYS_TIMER_CREATE:           "timer_create",
	syscall.SYS_TIMER_SETTIME:          "timer_settime",
	syscall.SYS_TIMER_GETTIME:          "timer_gettime",
	syscall.SYS_TIMER_GETOVERRUN:       "timer_getoverrun",
	syscall.SYS_TIMER_DELETE:           "timer_delete",
	syscall.SYS_CLOCK_SETTIME:          "clock_settime",
	syscall.SYS_CLOCK_GETTIME:          "clock_gettime",
	syscall.SYS_CLOCK_GETRES:           "clock_getres",
	syscall.SYS_CLOCK_NANOSLEEP:        "clock_nanosleep",
	syscall.SYS_STATFS64:               "statfs64",
	syscall.SYS_FSTATFS64:              "fstatfs64",
	syscall.SYS_TGKILL:                 "tgkill",
	syscall.SYS_UTIMES:                 "utimes",
	syscall.SYS_FADVISE64_64:           "fadvise64_64",
	syscall.SYS_VSERVER:                "vserver",
	syscall.SYS_MBIND:                  "mbind",
	syscall.SYS_GET_MEMPOLICY:          "get_mempolicy",
	syscall.SYS_SET_MEMPOLICY:          "set_mempolicy",
	syscall.SYS_MQ_OPEN:                "mq_open",
	syscall.SYS_MQ_UNLINK:              "mq_unlink",
	syscall.SYS_MQ_TIMEDSEND:           "mq_timedsend",
	syscall.SYS_MQ_TIMEDRECEIVE:        "mq_timedreceive",
	syscall.SYS_MQ_NOTIFY:              "mq_notify",
	syscall.SYS_MQ_GETSETATTR:          "mq_getsetattr",
	syscall.SYS_KEXEC_LOAD:             "kexec_load",
	syscall.SYS_WAITID:                 "waitid",
	syscall.SYS_ADD_KEY:                "add_key",
	syscall.SYS_REQUEST_KEY:            "request_key",
	syscall.SYS_KEYCTL:                 "keyctl",
	syscall.SYS_IOPRIO_SET:             "ioprio_set",
	syscall.SYS_IOPRIO_GET:             "ioprio_get",
	syscall.SYS_INOTIFY_INIT:           "inotify_init",
	syscall.SYS_INOTIFY_ADD_WATCH:      "inotify_add_watch",
	syscall.SYS_INOTIFY_RM_WATCH:       "inotify_rm_watch",
	syscall.SYS_MIGRATE_PAGES:          "migrate_pages",
	syscall.SYS_OPENAT:                 "openat",
	syscall.SYS_MKDIRAT:                "mkdirat",
	syscall.SYS_MKNODAT:                "mknodat",
	syscall.SYS_FCHOWNAT:               "fchownat",
	syscall.SYS_FUTIMESAT:              "futimesat",
	syscall.SYS_FSTATAT64:              "fstatat64",
	syscall.SYS_UNLINKAT:               "unlinkat",
	syscall.SYS_RENAMEAT:               "renameat",
	syscall.SYS_LINKAT:                 "linkat",
	syscall.SYS_SYMLINKAT:              "symlinkat",
	syscall.SYS_READLINKAT:             "readlinkat",
	syscall.SYS_FCHMODAT:               "fchmodat",
	syscall.SYS_FACCESSAT:              "faccessat",
	syscall.SYS_PSELECT6:               "pselect6",
	syscall.SYS_PPOLL:                  "ppoll",
	syscall.SYS_UNSHARE:                "unshare",
	syscall.SYS_SET_ROBUST_LIST:        "set_robust_list",
	syscall.SYS_GET_ROBUST_LIST:        "get_robust_list",
	syscall.SYS_SPLICE:                 "splice",
	syscall.SYS_SYNC_FILE_RANGE:        "sync_file_range",
	syscall.SYS_TEE:                    "tee",
	syscall.SYS_VMSPLICE:               "vmsplice",
	syscall.SYS_MOVE_PAGES:             "move_pages",
	syscall.SYS_GETCPU:                 "getcpu",
	syscall.SYS_EPOLL_PWAIT:            "epoll_pwait",
	syscall.SYS_UTIMENSAT:              "utimensat",
	syscall.SYS_SIGNALFD:               "signalfd",
	syscall.SYS_TIMERFD_CREATE:         "timerfd_create",
	syscall.SYS_EVENTFD:                "eventfd",
	syscall.SYS_FALLOCATE:              "fallocate",
	syscall.SYS_TIMERFD_SETTIME:        "timerfd_settime",
	syscall.SYS_TIMERFD_GETTIME:        "timerfd_gettime",
	syscall.SYS_SIGNALFD4:              "signalfd4",
	syscall.SYS_EVENTFD2:               "eventfd2",
	syscall.SYS_EPOLL_CREATE1:          "epoll_create1",
	syscall.SYS_DUP3:                   "dup3",
	syscall.SYS_PIPE2:                  "pipe2",
	syscall.SYS_INOTIFY_INIT1:          "inotify_init1",
	syscall.SYS_PREADV:                 "preadv",
	syscall.SYS_PWRITEV:                "pwritev",
	syscall.SYS_RT_TGSIGQUEUEINFO:      "rt_tgsigqueueinfo",
	syscall.SYS_PERF_EVENT_OPEN:        "perf_event_open",
	syscall.SYS_RECVMMSG:               "recvmmsg",
	syscall.SYS_FANOTIFY_INIT:          "fanotify_init",
	syscall.SYS_FANOTIFY_MARK:          "fanotify_mark",
	syscall.SYS_PRLIMIT64:              "prlimit64",
}
