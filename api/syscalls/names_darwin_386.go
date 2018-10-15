package syscalls

import "syscall"

var _ = syscall.Exit

var Name = map[uint64]string{
	syscall.SYS_SYSCALL:                        "syscall",
	syscall.SYS_EXIT:                           "exit",
	syscall.SYS_FORK:                           "fork",
	syscall.SYS_READ:                           "read",
	syscall.SYS_WRITE:                          "write",
	syscall.SYS_OPEN:                           "open",
	syscall.SYS_CLOSE:                          "close",
	syscall.SYS_WAIT4:                          "wait4",
	syscall.SYS_LINK:                           "link",
	syscall.SYS_UNLINK:                         "unlink",
	syscall.SYS_CHDIR:                          "chdir",
	syscall.SYS_FCHDIR:                         "fchdir",
	syscall.SYS_MKNOD:                          "mknod",
	syscall.SYS_CHMOD:                          "chmod",
	syscall.SYS_CHOWN:                          "chown",
	syscall.SYS_GETFSSTAT:                      "getfsstat",
	syscall.SYS_GETPID:                         "getpid",
	syscall.SYS_SETUID:                         "setuid",
	syscall.SYS_GETUID:                         "getuid",
	syscall.SYS_GETEUID:                        "geteuid",
	syscall.SYS_PTRACE:                         "ptrace",
	syscall.SYS_RECVMSG:                        "recvmsg",
	syscall.SYS_SENDMSG:                        "sendmsg",
	syscall.SYS_RECVFROM:                       "recvfrom",
	syscall.SYS_ACCEPT:                         "accept",
	syscall.SYS_GETPEERNAME:                    "getpeername",
	syscall.SYS_GETSOCKNAME:                    "getsockname",
	syscall.SYS_ACCESS:                         "access",
	syscall.SYS_CHFLAGS:                        "chflags",
	syscall.SYS_FCHFLAGS:                       "fchflags",
	syscall.SYS_SYNC:                           "sync",
	syscall.SYS_KILL:                           "kill",
	syscall.SYS_GETPPID:                        "getppid",
	syscall.SYS_DUP:                            "dup",
	syscall.SYS_PIPE:                           "pipe",
	syscall.SYS_GETEGID:                        "getegid",
	syscall.SYS_PROFIL:                         "profil",
	syscall.SYS_SIGACTION:                      "sigaction",
	syscall.SYS_GETGID:                         "getgid",
	syscall.SYS_SIGPROCMASK:                    "sigprocmask",
	syscall.SYS_GETLOGIN:                       "getlogin",
	syscall.SYS_SETLOGIN:                       "setlogin",
	syscall.SYS_ACCT:                           "acct",
	syscall.SYS_SIGPENDING:                     "sigpending",
	syscall.SYS_SIGALTSTACK:                    "sigaltstack",
	syscall.SYS_IOCTL:                          "ioctl",
	syscall.SYS_REBOOT:                         "reboot",
	syscall.SYS_REVOKE:                         "revoke",
	syscall.SYS_SYMLINK:                        "symlink",
	syscall.SYS_READLINK:                       "readlink",
	syscall.SYS_EXECVE:                         "execve",
	syscall.SYS_UMASK:                          "umask",
	syscall.SYS_CHROOT:                         "chroot",
	syscall.SYS_MSYNC:                          "msync",
	syscall.SYS_VFORK:                          "vfork",
	syscall.SYS_MUNMAP:                         "munmap",
	syscall.SYS_MPROTECT:                       "mprotect",
	syscall.SYS_MADVISE:                        "madvise",
	syscall.SYS_MINCORE:                        "mincore",
	syscall.SYS_GETGROUPS:                      "getgroups",
	syscall.SYS_SETGROUPS:                      "setgroups",
	syscall.SYS_GETPGRP:                        "getpgrp",
	syscall.SYS_SETPGID:                        "setpgid",
	syscall.SYS_SETITIMER:                      "setitimer",
	syscall.SYS_SWAPON:                         "swapon",
	syscall.SYS_GETITIMER:                      "getitimer",
	syscall.SYS_GETDTABLESIZE:                  "getdtablesize",
	syscall.SYS_DUP2:                           "dup2",
	syscall.SYS_FCNTL:                          "fcntl",
	syscall.SYS_SELECT:                         "select",
	syscall.SYS_FSYNC:                          "fsync",
	syscall.SYS_SETPRIORITY:                    "setpriority",
	syscall.SYS_SOCKET:                         "socket",
	syscall.SYS_CONNECT:                        "connect",
	syscall.SYS_GETPRIORITY:                    "getpriority",
	syscall.SYS_BIND:                           "bind",
	syscall.SYS_SETSOCKOPT:                     "setsockopt",
	syscall.SYS_LISTEN:                         "listen",
	syscall.SYS_SIGSUSPEND:                     "sigsuspend",
	syscall.SYS_GETTIMEOFDAY:                   "gettimeofday",
	syscall.SYS_GETRUSAGE:                      "getrusage",
	syscall.SYS_GETSOCKOPT:                     "getsockopt",
	syscall.SYS_READV:                          "readv",
	syscall.SYS_WRITEV:                         "writev",
	syscall.SYS_SETTIMEOFDAY:                   "settimeofday",
	syscall.SYS_FCHOWN:                         "fchown",
	syscall.SYS_FCHMOD:                         "fchmod",
	syscall.SYS_SETREUID:                       "setreuid",
	syscall.SYS_SETREGID:                       "setregid",
	syscall.SYS_RENAME:                         "rename",
	syscall.SYS_FLOCK:                          "flock",
	syscall.SYS_MKFIFO:                         "mkfifo",
	syscall.SYS_SENDTO:                         "sendto",
	syscall.SYS_SHUTDOWN:                       "shutdown",
	syscall.SYS_SOCKETPAIR:                     "socketpair",
	syscall.SYS_MKDIR:                          "mkdir",
	syscall.SYS_RMDIR:                          "rmdir",
	syscall.SYS_UTIMES:                         "utimes",
	syscall.SYS_FUTIMES:                        "futimes",
	syscall.SYS_ADJTIME:                        "adjtime",
	syscall.SYS_GETHOSTUUID:                    "gethostuuid",
	syscall.SYS_SETSID:                         "setsid",
	syscall.SYS_GETPGID:                        "getpgid",
	syscall.SYS_SETPRIVEXEC:                    "setprivexec",
	syscall.SYS_PREAD:                          "pread",
	syscall.SYS_PWRITE:                         "pwrite",
	syscall.SYS_NFSSVC:                         "nfssvc",
	syscall.SYS_STATFS:                         "statfs",
	syscall.SYS_FSTATFS:                        "fstatfs",
	syscall.SYS_UNMOUNT:                        "unmount",
	syscall.SYS_GETFH:                          "getfh",
	syscall.SYS_QUOTACTL:                       "quotactl",
	syscall.SYS_MOUNT:                          "mount",
	syscall.SYS_CSOPS:                          "csops",
	syscall.SYS_WAITID:                         "waitid",
	syscall.SYS_ADD_PROFIL:                     "add_profil",
	syscall.SYS_KDEBUG_TRACE:                   "kdebug_trace",
	syscall.SYS_SETGID:                         "setgid",
	syscall.SYS_SETEGID:                        "setegid",
	syscall.SYS_SETEUID:                        "seteuid",
	syscall.SYS_SIGRETURN:                      "sigreturn",
	syscall.SYS_CHUD:                           "chud",
	syscall.SYS_FDATASYNC:                      "fdatasync",
	syscall.SYS_STAT:                           "stat",
	syscall.SYS_FSTAT:                          "fstat",
	syscall.SYS_LSTAT:                          "lstat",
	syscall.SYS_PATHCONF:                       "pathconf",
	syscall.SYS_FPATHCONF:                      "fpathconf",
	syscall.SYS_GETRLIMIT:                      "getrlimit",
	syscall.SYS_SETRLIMIT:                      "setrlimit",
	syscall.SYS_GETDIRENTRIES:                  "getdirentries",
	syscall.SYS_MMAP:                           "mmap",
	syscall.SYS_LSEEK:                          "lseek",
	syscall.SYS_TRUNCATE:                       "truncate",
	syscall.SYS_FTRUNCATE:                      "ftruncate",
	syscall.SYS___SYSCTL:                       "__sysctl",
	syscall.SYS_MLOCK:                          "mlock",
	syscall.SYS_MUNLOCK:                        "munlock",
	syscall.SYS_UNDELETE:                       "undelete",
	syscall.SYS_ATSOCKET:                       "atsocket",
	syscall.SYS_ATGETMSG:                       "atgetmsg",
	syscall.SYS_ATPUTMSG:                       "atputmsg",
	syscall.SYS_ATPSNDREQ:                      "atpsndreq",
	syscall.SYS_ATPSNDRSP:                      "atpsndrsp",
	syscall.SYS_ATPGETREQ:                      "atpgetreq",
	syscall.SYS_ATPGETRSP:                      "atpgetrsp",
	syscall.SYS_MKCOMPLEX:                      "mkcomplex",
	syscall.SYS_STATV:                          "statv",
	syscall.SYS_LSTATV:                         "lstatv",
	syscall.SYS_FSTATV:                         "fstatv",
	syscall.SYS_GETATTRLIST:                    "getattrlist",
	syscall.SYS_SETATTRLIST:                    "setattrlist",
	syscall.SYS_GETDIRENTRIESATTR:              "getdirentriesattr",
	syscall.SYS_EXCHANGEDATA:                   "exchangedata",
	syscall.SYS_SEARCHFS:                       "searchfs",
	syscall.SYS_DELETE:                         "delete",
	syscall.SYS_COPYFILE:                       "copyfile",
	syscall.SYS_FGETATTRLIST:                   "fgetattrlist",
	syscall.SYS_FSETATTRLIST:                   "fsetattrlist",
	syscall.SYS_POLL:                           "poll",
	syscall.SYS_WATCHEVENT:                     "watchevent",
	syscall.SYS_WAITEVENT:                      "waitevent",
	syscall.SYS_MODWATCH:                       "modwatch",
	syscall.SYS_GETXATTR:                       "getxattr",
	syscall.SYS_FGETXATTR:                      "fgetxattr",
	syscall.SYS_SETXATTR:                       "setxattr",
	syscall.SYS_FSETXATTR:                      "fsetxattr",
	syscall.SYS_REMOVEXATTR:                    "removexattr",
	syscall.SYS_FREMOVEXATTR:                   "fremovexattr",
	syscall.SYS_LISTXATTR:                      "listxattr",
	syscall.SYS_FLISTXATTR:                     "flistxattr",
	syscall.SYS_FSCTL:                          "fsctl",
	syscall.SYS_INITGROUPS:                     "initgroups",
	syscall.SYS_POSIX_SPAWN:                    "posix_spawn",
	syscall.SYS_FFSCTL:                         "ffsctl",
	syscall.SYS_NFSCLNT:                        "nfsclnt",
	syscall.SYS_FHOPEN:                         "fhopen",
	syscall.SYS_MINHERIT:                       "minherit",
	syscall.SYS_SEMSYS:                         "semsys",
	syscall.SYS_MSGSYS:                         "msgsys",
	syscall.SYS_SHMSYS:                         "shmsys",
	syscall.SYS_SEMCTL:                         "semctl",
	syscall.SYS_SEMGET:                         "semget",
	syscall.SYS_SEMOP:                          "semop",
	syscall.SYS_MSGCTL:                         "msgctl",
	syscall.SYS_MSGGET:                         "msgget",
	syscall.SYS_MSGSND:                         "msgsnd",
	syscall.SYS_MSGRCV:                         "msgrcv",
	syscall.SYS_SHMAT:                          "shmat",
	syscall.SYS_SHMCTL:                         "shmctl",
	syscall.SYS_SHMDT:                          "shmdt",
	syscall.SYS_SHMGET:                         "shmget",
	syscall.SYS_SHM_OPEN:                       "shm_open",
	syscall.SYS_SHM_UNLINK:                     "shm_unlink",
	syscall.SYS_SEM_OPEN:                       "sem_open",
	syscall.SYS_SEM_CLOSE:                      "sem_close",
	syscall.SYS_SEM_UNLINK:                     "sem_unlink",
	syscall.SYS_SEM_WAIT:                       "sem_wait",
	syscall.SYS_SEM_TRYWAIT:                    "sem_trywait",
	syscall.SYS_SEM_POST:                       "sem_post",
	syscall.SYS_SEM_GETVALUE:                   "sem_getvalue",
	syscall.SYS_SEM_INIT:                       "sem_init",
	syscall.SYS_SEM_DESTROY:                    "sem_destroy",
	syscall.SYS_OPEN_EXTENDED:                  "open_extended",
	syscall.SYS_UMASK_EXTENDED:                 "umask_extended",
	syscall.SYS_STAT_EXTENDED:                  "stat_extended",
	syscall.SYS_LSTAT_EXTENDED:                 "lstat_extended",
	syscall.SYS_FSTAT_EXTENDED:                 "fstat_extended",
	syscall.SYS_CHMOD_EXTENDED:                 "chmod_extended",
	syscall.SYS_FCHMOD_EXTENDED:                "fchmod_extended",
	syscall.SYS_ACCESS_EXTENDED:                "access_extended",
	syscall.SYS_SETTID:                         "settid",
	syscall.SYS_GETTID:                         "gettid",
	syscall.SYS_SETSGROUPS:                     "setsgroups",
	syscall.SYS_GETSGROUPS:                     "getsgroups",
	syscall.SYS_SETWGROUPS:                     "setwgroups",
	syscall.SYS_GETWGROUPS:                     "getwgroups",
	syscall.SYS_MKFIFO_EXTENDED:                "mkfifo_extended",
	syscall.SYS_MKDIR_EXTENDED:                 "mkdir_extended",
	syscall.SYS_IDENTITYSVC:                    "identitysvc",
	syscall.SYS_SHARED_REGION_CHECK_NP:         "shared_region_check_np",
	syscall.SYS_VM_PRESSURE_MONITOR:            "vm_pressure_monitor",
	syscall.SYS_PSYNCH_RW_LONGRDLOCK:           "psynch_rw_longrdlock",
	syscall.SYS_PSYNCH_RW_YIELDWRLOCK:          "psynch_rw_yieldwrlock",
	syscall.SYS_PSYNCH_RW_DOWNGRADE:            "psynch_rw_downgrade",
	syscall.SYS_PSYNCH_RW_UPGRADE:              "psynch_rw_upgrade",
	syscall.SYS_PSYNCH_MUTEXWAIT:               "psynch_mutexwait",
	syscall.SYS_PSYNCH_MUTEXDROP:               "psynch_mutexdrop",
	syscall.SYS_PSYNCH_CVBROAD:                 "psynch_cvbroad",
	syscall.SYS_PSYNCH_CVSIGNAL:                "psynch_cvsignal",
	syscall.SYS_PSYNCH_CVWAIT:                  "psynch_cvwait",
	syscall.SYS_PSYNCH_RW_RDLOCK:               "psynch_rw_rdlock",
	syscall.SYS_PSYNCH_RW_WRLOCK:               "psynch_rw_wrlock",
	syscall.SYS_PSYNCH_RW_UNLOCK:               "psynch_rw_unlock",
	syscall.SYS_PSYNCH_RW_UNLOCK2:              "psynch_rw_unlock2",
	syscall.SYS_GETSID:                         "getsid",
	syscall.SYS_SETTID_WITH_PID:                "settid_with_pid",
	syscall.SYS_PSYNCH_CVCLRPREPOST:            "psynch_cvclrprepost",
	syscall.SYS_AIO_FSYNC:                      "aio_fsync",
	syscall.SYS_AIO_RETURN:                     "aio_return",
	syscall.SYS_AIO_SUSPEND:                    "aio_suspend",
	syscall.SYS_AIO_CANCEL:                     "aio_cancel",
	syscall.SYS_AIO_ERROR:                      "aio_error",
	syscall.SYS_AIO_READ:                       "aio_read",
	syscall.SYS_AIO_WRITE:                      "aio_write",
	syscall.SYS_LIO_LISTIO:                     "lio_listio",
	syscall.SYS_IOPOLICYSYS:                    "iopolicysys",
	syscall.SYS_PROCESS_POLICY:                 "process_policy",
	syscall.SYS_MLOCKALL:                       "mlockall",
	syscall.SYS_MUNLOCKALL:                     "munlockall",
	syscall.SYS_ISSETUGID:                      "issetugid",
	syscall.SYS___PTHREAD_KILL:                 "__pthread_kill",
	syscall.SYS___PTHREAD_SIGMASK:              "__pthread_sigmask",
	syscall.SYS___SIGWAIT:                      "__sigwait",
	syscall.SYS___DISABLE_THREADSIGNAL:         "__disable_threadsignal",
	syscall.SYS___PTHREAD_MARKCANCEL:           "__pthread_markcancel",
	syscall.SYS___PTHREAD_CANCELED:             "__pthread_canceled",
	syscall.SYS___SEMWAIT_SIGNAL:               "__semwait_signal",
	syscall.SYS_PROC_INFO:                      "proc_info",
	syscall.SYS_SENDFILE:                       "sendfile",
	syscall.SYS_STAT64:                         "stat64",
	syscall.SYS_FSTAT64:                        "fstat64",
	syscall.SYS_LSTAT64:                        "lstat64",
	syscall.SYS_STAT64_EXTENDED:                "stat64_extended",
	syscall.SYS_LSTAT64_EXTENDED:               "lstat64_extended",
	syscall.SYS_FSTAT64_EXTENDED:               "fstat64_extended",
	syscall.SYS_GETDIRENTRIES64:                "getdirentries64",
	syscall.SYS_STATFS64:                       "statfs64",
	syscall.SYS_FSTATFS64:                      "fstatfs64",
	syscall.SYS_GETFSSTAT64:                    "getfsstat64",
	syscall.SYS___PTHREAD_CHDIR:                "__pthread_chdir",
	syscall.SYS___PTHREAD_FCHDIR:               "__pthread_fchdir",
	syscall.SYS_AUDIT:                          "audit",
	syscall.SYS_AUDITON:                        "auditon",
	syscall.SYS_GETAUID:                        "getauid",
	syscall.SYS_SETAUID:                        "setauid",
	syscall.SYS_GETAUDIT:                       "getaudit",
	syscall.SYS_SETAUDIT:                       "setaudit",
	syscall.SYS_GETAUDIT_ADDR:                  "getaudit_addr",
	syscall.SYS_SETAUDIT_ADDR:                  "setaudit_addr",
	syscall.SYS_AUDITCTL:                       "auditctl",
	syscall.SYS_BSDTHREAD_CREATE:               "bsdthread_create",
	syscall.SYS_BSDTHREAD_TERMINATE:            "bsdthread_terminate",
	syscall.SYS_KQUEUE:                         "kqueue",
	syscall.SYS_KEVENT:                         "kevent",
	syscall.SYS_LCHOWN:                         "lchown",
	syscall.SYS_STACK_SNAPSHOT:                 "stack_snapshot",
	syscall.SYS_BSDTHREAD_REGISTER:             "bsdthread_register",
	syscall.SYS_WORKQ_OPEN:                     "workq_open",
	syscall.SYS_WORKQ_KERNRETURN:               "workq_kernreturn",
	syscall.SYS_KEVENT64:                       "kevent64",
	syscall.SYS___OLD_SEMWAIT_SIGNAL:           "__old_semwait_signal",
	syscall.SYS___OLD_SEMWAIT_SIGNAL_NOCANCEL:  "__old_semwait_signal_nocancel",
	syscall.SYS_THREAD_SELFID:                  "thread_selfid",
	syscall.SYS___MAC_EXECVE:                   "__mac_execve",
	syscall.SYS___MAC_SYSCALL:                  "__mac_syscall",
	syscall.SYS___MAC_GET_FILE:                 "__mac_get_file",
	syscall.SYS___MAC_SET_FILE:                 "__mac_set_file",
	syscall.SYS___MAC_GET_LINK:                 "__mac_get_link",
	syscall.SYS___MAC_SET_LINK:                 "__mac_set_link",
	syscall.SYS___MAC_GET_PROC:                 "__mac_get_proc",
	syscall.SYS___MAC_SET_PROC:                 "__mac_set_proc",
	syscall.SYS___MAC_GET_FD:                   "__mac_get_fd",
	syscall.SYS___MAC_SET_FD:                   "__mac_set_fd",
	syscall.SYS___MAC_GET_PID:                  "__mac_get_pid",
	syscall.SYS___MAC_GET_LCID:                 "__mac_get_lcid",
	syscall.SYS___MAC_GET_LCTX:                 "__mac_get_lctx",
	syscall.SYS___MAC_SET_LCTX:                 "__mac_set_lctx",
	syscall.SYS_SETLCID:                        "setlcid",
	syscall.SYS_GETLCID:                        "getlcid",
	syscall.SYS_READ_NOCANCEL:                  "read_nocancel",
	syscall.SYS_WRITE_NOCANCEL:                 "write_nocancel",
	syscall.SYS_OPEN_NOCANCEL:                  "open_nocancel",
	syscall.SYS_CLOSE_NOCANCEL:                 "close_nocancel",
	syscall.SYS_WAIT4_NOCANCEL:                 "wait4_nocancel",
	syscall.SYS_RECVMSG_NOCANCEL:               "recvmsg_nocancel",
	syscall.SYS_SENDMSG_NOCANCEL:               "sendmsg_nocancel",
	syscall.SYS_RECVFROM_NOCANCEL:              "recvfrom_nocancel",
	syscall.SYS_ACCEPT_NOCANCEL:                "accept_nocancel",
	syscall.SYS_MSYNC_NOCANCEL:                 "msync_nocancel",
	syscall.SYS_FCNTL_NOCANCEL:                 "fcntl_nocancel",
	syscall.SYS_SELECT_NOCANCEL:                "select_nocancel",
	syscall.SYS_FSYNC_NOCANCEL:                 "fsync_nocancel",
	syscall.SYS_CONNECT_NOCANCEL:               "connect_nocancel",
	syscall.SYS_SIGSUSPEND_NOCANCEL:            "sigsuspend_nocancel",
	syscall.SYS_READV_NOCANCEL:                 "readv_nocancel",
	syscall.SYS_WRITEV_NOCANCEL:                "writev_nocancel",
	syscall.SYS_SENDTO_NOCANCEL:                "sendto_nocancel",
	syscall.SYS_PREAD_NOCANCEL:                 "pread_nocancel",
	syscall.SYS_PWRITE_NOCANCEL:                "pwrite_nocancel",
	syscall.SYS_WAITID_NOCANCEL:                "waitid_nocancel",
	syscall.SYS_POLL_NOCANCEL:                  "poll_nocancel",
	syscall.SYS_MSGSND_NOCANCEL:                "msgsnd_nocancel",
	syscall.SYS_MSGRCV_NOCANCEL:                "msgrcv_nocancel",
	syscall.SYS_SEM_WAIT_NOCANCEL:              "sem_wait_nocancel",
	syscall.SYS_AIO_SUSPEND_NOCANCEL:           "aio_suspend_nocancel",
	syscall.SYS___SIGWAIT_NOCANCEL:             "__sigwait_nocancel",
	syscall.SYS___SEMWAIT_SIGNAL_NOCANCEL:      "__semwait_signal_nocancel",
	syscall.SYS___MAC_MOUNT:                    "__mac_mount",
	syscall.SYS___MAC_GET_MOUNT:                "__mac_get_mount",
	syscall.SYS___MAC_GETFSSTAT:                "__mac_getfsstat",
	syscall.SYS_FSGETPATH:                      "fsgetpath",
	syscall.SYS_AUDIT_SESSION_SELF:             "audit_session_self",
	syscall.SYS_AUDIT_SESSION_JOIN:             "audit_session_join",
	syscall.SYS_FILEPORT_MAKEPORT:              "fileport_makeport",
	syscall.SYS_FILEPORT_MAKEFD:                "fileport_makefd",
	syscall.SYS_AUDIT_SESSION_PORT:             "audit_session_port",
	syscall.SYS_PID_SUSPEND:                    "pid_suspend",
	syscall.SYS_PID_RESUME:                     "pid_resume",
	syscall.SYS_PID_HIBERNATE:                  "pid_hibernate",
	syscall.SYS_PID_SHUTDOWN_SOCKETS:           "pid_shutdown_sockets",
	syscall.SYS_SHARED_REGION_MAP_AND_SLIDE_NP: "shared_region_map_and_slide_np",
	syscall.SYS_MAXSYSCALL:                     "maxsyscall",
}
