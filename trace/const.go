package trace

import (
	syscall "golang.org/x/sys/unix"
)

var SyscallName map[uint64]string = map[uint64]string{
	syscall.SYS_READ:                   "sys_read",
	syscall.SYS_WRITE:                  "sys_write",
	syscall.SYS_OPEN:                   "sys_open",
	syscall.SYS_CLOSE:                  "sys_close",
	syscall.SYS_STAT:                   "sys_stat",
	syscall.SYS_FSTAT:                  "sys_fstat",
	syscall.SYS_LSTAT:                  "sys_lstat",
	syscall.SYS_POLL:                   "sys_poll",
	syscall.SYS_LSEEK:                  "sys_lseek",
	syscall.SYS_MMAP:                   "sys_mmap",
	syscall.SYS_MPROTECT:               "sys_mprotect",
	syscall.SYS_MUNMAP:                 "sys_munmap",
	syscall.SYS_BRK:                    "sys_brk",
	syscall.SYS_RT_SIGACTION:           "sys_rt_sigaction",
	syscall.SYS_RT_SIGPROCMASK:         "sys_rt_sigprocmask",
	syscall.SYS_RT_SIGRETURN:           "sys_rt_sigreturn",
	syscall.SYS_IOCTL:                  "sys_ioctl",
	syscall.SYS_PREAD64:                "sys_pread64",
	syscall.SYS_PWRITE64:               "sys_pwrite64",
	syscall.SYS_READV:                  "sys_readv",
	syscall.SYS_WRITEV:                 "sys_writev",
	syscall.SYS_ACCESS:                 "sys_access",
	syscall.SYS_PIPE:                   "sys_pipe",
	syscall.SYS_SELECT:                 "sys_select",
	syscall.SYS_SCHED_YIELD:            "sys_sched_yield",
	syscall.SYS_MREMAP:                 "sys_mremap",
	syscall.SYS_MSYNC:                  "sys_msync",
	syscall.SYS_MINCORE:                "sys_mincore",
	syscall.SYS_MADVISE:                "sys_madvise",
	syscall.SYS_SHMGET:                 "sys_shmget",
	syscall.SYS_SHMAT:                  "sys_shmat",
	syscall.SYS_SHMCTL:                 "sys_shmctl",
	syscall.SYS_DUP:                    "sys_dup",
	syscall.SYS_DUP2:                   "sys_dup2",
	syscall.SYS_PAUSE:                  "sys_pause",
	syscall.SYS_NANOSLEEP:              "sys_nanosleep",
	syscall.SYS_GETITIMER:              "sys_getitimer",
	syscall.SYS_ALARM:                  "sys_alarm",
	syscall.SYS_SETITIMER:              "sys_setitimer",
	syscall.SYS_GETPID:                 "sys_getpid",
	syscall.SYS_SENDFILE:               "sys_sendfile",
	syscall.SYS_SOCKET:                 "sys_socket",
	syscall.SYS_CONNECT:                "sys_connect",
	syscall.SYS_ACCEPT:                 "sys_accept",
	syscall.SYS_SENDTO:                 "sys_sendto",
	syscall.SYS_RECVFROM:               "sys_recvfrom",
	syscall.SYS_SENDMSG:                "sys_sendmsg",
	syscall.SYS_RECVMSG:                "sys_recvmsg",
	syscall.SYS_SHUTDOWN:               "sys_shutdown",
	syscall.SYS_BIND:                   "sys_bind",
	syscall.SYS_LISTEN:                 "sys_listen",
	syscall.SYS_GETSOCKNAME:            "sys_getsockname",
	syscall.SYS_GETPEERNAME:            "sys_getpeername",
	syscall.SYS_SOCKETPAIR:             "sys_socketpair",
	syscall.SYS_SETSOCKOPT:             "sys_setsockopt",
	syscall.SYS_GETSOCKOPT:             "sys_getsockopt",
	syscall.SYS_CLONE:                  "sys_clone",
	syscall.SYS_FORK:                   "sys_fork",
	syscall.SYS_VFORK:                  "sys_vfork",
	syscall.SYS_EXECVE:                 "sys_execve",
	syscall.SYS_EXIT:                   "sys_exit",
	syscall.SYS_WAIT4:                  "sys_wait4",
	syscall.SYS_KILL:                   "sys_kill",
	syscall.SYS_UNAME:                  "sys_uname",
	syscall.SYS_SEMGET:                 "sys_semget",
	syscall.SYS_SEMOP:                  "sys_semop",
	syscall.SYS_SEMCTL:                 "sys_semctl",
	syscall.SYS_SHMDT:                  "sys_shmdt",
	syscall.SYS_MSGGET:                 "sys_msgget",
	syscall.SYS_MSGSND:                 "sys_msgsnd",
	syscall.SYS_MSGRCV:                 "sys_msgrcv",
	syscall.SYS_MSGCTL:                 "sys_msgctl",
	syscall.SYS_FCNTL:                  "sys_fcntl",
	syscall.SYS_FLOCK:                  "sys_flock",
	syscall.SYS_FSYNC:                  "sys_fsync",
	syscall.SYS_FDATASYNC:              "sys_fdatasync",
	syscall.SYS_TRUNCATE:               "sys_truncate",
	syscall.SYS_FTRUNCATE:              "sys_ftruncate",
	syscall.SYS_GETDENTS:               "sys_getdents",
	syscall.SYS_GETCWD:                 "sys_getcwd",
	syscall.SYS_CHDIR:                  "sys_chdir",
	syscall.SYS_FCHDIR:                 "sys_fchdir",
	syscall.SYS_RENAME:                 "sys_rename",
	syscall.SYS_MKDIR:                  "sys_mkdir",
	syscall.SYS_RMDIR:                  "sys_rmdir",
	syscall.SYS_CREAT:                  "sys_creat",
	syscall.SYS_LINK:                   "sys_link",
	syscall.SYS_UNLINK:                 "sys_unlink",
	syscall.SYS_SYMLINK:                "sys_symlink",
	syscall.SYS_READLINK:               "sys_readlink",
	syscall.SYS_CHMOD:                  "sys_chmod",
	syscall.SYS_FCHMOD:                 "sys_fchmod",
	syscall.SYS_CHOWN:                  "sys_chown",
	syscall.SYS_FCHOWN:                 "sys_fchown",
	syscall.SYS_LCHOWN:                 "sys_lchown",
	syscall.SYS_UMASK:                  "sys_umask",
	syscall.SYS_GETTIMEOFDAY:           "sys_gettimeofday",
	syscall.SYS_GETRLIMIT:              "sys_getrlimit",
	syscall.SYS_GETRUSAGE:              "sys_getrusage",
	syscall.SYS_SYSINFO:                "sys_sysinfo",
	syscall.SYS_TIMES:                  "sys_times",
	syscall.SYS_PTRACE:                 "sys_ptrace",
	syscall.SYS_GETUID:                 "sys_getuid",
	syscall.SYS_SYSLOG:                 "sys_syslog",
	syscall.SYS_GETGID:                 "sys_getgid",
	syscall.SYS_SETUID:                 "sys_setuid",
	syscall.SYS_SETGID:                 "sys_setgid",
	syscall.SYS_GETEUID:                "sys_geteuid",
	syscall.SYS_GETEGID:                "sys_getegid",
	syscall.SYS_SETPGID:                "sys_setpgid",
	syscall.SYS_GETPPID:                "sys_getppid",
	syscall.SYS_GETPGRP:                "sys_getpgrp",
	syscall.SYS_SETSID:                 "sys_setsid",
	syscall.SYS_SETREUID:               "sys_setreuid",
	syscall.SYS_SETREGID:               "sys_setregid",
	syscall.SYS_GETGROUPS:              "sys_getgroups",
	syscall.SYS_SETGROUPS:              "sys_setgroups",
	syscall.SYS_SETRESUID:              "sys_setresuid",
	syscall.SYS_GETRESUID:              "sys_getresuid",
	syscall.SYS_SETRESGID:              "sys_setresgid",
	syscall.SYS_GETRESGID:              "sys_getresgid",
	syscall.SYS_GETPGID:                "sys_getpgid",
	syscall.SYS_SETFSUID:               "sys_setfsuid",
	syscall.SYS_SETFSGID:               "sys_setfsgid",
	syscall.SYS_GETSID:                 "sys_getsid",
	syscall.SYS_CAPGET:                 "sys_capget",
	syscall.SYS_CAPSET:                 "sys_capset",
	syscall.SYS_RT_SIGPENDING:          "sys_rt_sigpending",
	syscall.SYS_RT_SIGTIMEDWAIT:        "sys_rt_sigtimedwait",
	syscall.SYS_RT_SIGQUEUEINFO:        "sys_rt_sigqueueinfo",
	syscall.SYS_RT_SIGSUSPEND:          "sys_rt_sigsuspend",
	syscall.SYS_SIGALTSTACK:            "sys_sigaltstack",
	syscall.SYS_UTIME:                  "sys_utime",
	syscall.SYS_MKNOD:                  "sys_mknod",
	syscall.SYS_USELIB:                 "sys_uselib",
	syscall.SYS_PERSONALITY:            "sys_personality",
	syscall.SYS_USTAT:                  "sys_ustat",
	syscall.SYS_STATFS:                 "sys_statfs",
	syscall.SYS_FSTATFS:                "sys_fstatfs",
	syscall.SYS_SYSFS:                  "sys_sysfs",
	syscall.SYS_GETPRIORITY:            "sys_getpriority",
	syscall.SYS_SETPRIORITY:            "sys_setpriority",
	syscall.SYS_SCHED_SETPARAM:         "sys_sched_setparam",
	syscall.SYS_SCHED_GETPARAM:         "sys_sched_getparam",
	syscall.SYS_SCHED_SETSCHEDULER:     "sys_sched_setscheduler",
	syscall.SYS_SCHED_GETSCHEDULER:     "sys_sched_getscheduler",
	syscall.SYS_SCHED_GET_PRIORITY_MAX: "sys_sched_get_priority_max",
	syscall.SYS_SCHED_GET_PRIORITY_MIN: "sys_sched_get_priority_min",
	syscall.SYS_SCHED_RR_GET_INTERVAL:  "sys_sched_rr_get_interval",
	syscall.SYS_MLOCK:                  "sys_mlock",
	syscall.SYS_MUNLOCK:                "sys_munlock",
	syscall.SYS_MLOCKALL:               "sys_mlockall",
	syscall.SYS_MUNLOCKALL:             "sys_munlockall",
	syscall.SYS_VHANGUP:                "sys_vhangup",
	syscall.SYS_MODIFY_LDT:             "sys_modify_ldt",
	syscall.SYS_PIVOT_ROOT:             "sys_pivot_root",
	syscall.SYS__SYSCTL:                "sys__sysctl",
	syscall.SYS_PRCTL:                  "sys_prctl",
	syscall.SYS_ARCH_PRCTL:             "sys_arch_prctl",
	syscall.SYS_ADJTIMEX:               "sys_adjtimex",
	syscall.SYS_SETRLIMIT:              "sys_setrlimit",
	syscall.SYS_CHROOT:                 "sys_chroot",
	syscall.SYS_SYNC:                   "sys_sync",
	syscall.SYS_ACCT:                   "sys_acct",
	syscall.SYS_SETTIMEOFDAY:           "sys_settimeofday",
	syscall.SYS_MOUNT:                  "sys_mount",
	syscall.SYS_UMOUNT2:                "sys_umount2",
	syscall.SYS_SWAPON:                 "sys_swapon",
	syscall.SYS_SWAPOFF:                "sys_swapoff",
	syscall.SYS_REBOOT:                 "sys_reboot",
	syscall.SYS_SETHOSTNAME:            "sys_sethostname",
	syscall.SYS_SETDOMAINNAME:          "sys_setdomainname",
	syscall.SYS_IOPL:                   "sys_iopl",
	syscall.SYS_IOPERM:                 "sys_ioperm",
	syscall.SYS_CREATE_MODULE:          "sys_create_module",
	syscall.SYS_INIT_MODULE:            "sys_init_module",
	syscall.SYS_DELETE_MODULE:          "sys_delete_module",
	syscall.SYS_GET_KERNEL_SYMS:        "sys_get_kernel_syms",
	syscall.SYS_QUERY_MODULE:           "sys_query_module",
	syscall.SYS_QUOTACTL:               "sys_quotactl",
	syscall.SYS_NFSSERVCTL:             "sys_nfsservctl",
	syscall.SYS_GETPMSG:                "sys_getpmsg",
	syscall.SYS_PUTPMSG:                "sys_putpmsg",
	syscall.SYS_AFS_SYSCALL:            "sys_afs_syscall",
	syscall.SYS_TUXCALL:                "sys_tuxcall",
	syscall.SYS_SECURITY:               "sys_security",
	syscall.SYS_GETTID:                 "sys_gettid",
	syscall.SYS_READAHEAD:              "sys_readahead",
	syscall.SYS_SETXATTR:               "sys_setxattr",
	syscall.SYS_LSETXATTR:              "sys_lsetxattr",
	syscall.SYS_FSETXATTR:              "sys_fsetxattr",
	syscall.SYS_GETXATTR:               "sys_getxattr",
	syscall.SYS_LGETXATTR:              "sys_lgetxattr",
	syscall.SYS_FGETXATTR:              "sys_fgetxattr",
	syscall.SYS_LISTXATTR:              "sys_listxattr",
	syscall.SYS_LLISTXATTR:             "sys_llistxattr",
	syscall.SYS_FLISTXATTR:             "sys_flistxattr",
	syscall.SYS_REMOVEXATTR:            "sys_removexattr",
	syscall.SYS_LREMOVEXATTR:           "sys_lremovexattr",
	syscall.SYS_FREMOVEXATTR:           "sys_fremovexattr",
	syscall.SYS_TKILL:                  "sys_tkill",
	syscall.SYS_TIME:                   "sys_time",
	syscall.SYS_FUTEX:                  "sys_futex",
	syscall.SYS_SCHED_SETAFFINITY:      "sys_sched_setaffinity",
	syscall.SYS_SCHED_GETAFFINITY:      "sys_sched_getaffinity",
	syscall.SYS_SET_THREAD_AREA:        "sys_set_thread_area",
	syscall.SYS_IO_SETUP:               "sys_io_setup",
	syscall.SYS_IO_DESTROY:             "sys_io_destroy",
	syscall.SYS_IO_GETEVENTS:           "sys_io_getevents",
	syscall.SYS_IO_SUBMIT:              "sys_io_submit",
	syscall.SYS_IO_CANCEL:              "sys_io_cancel",
	syscall.SYS_GET_THREAD_AREA:        "sys_get_thread_area",
	syscall.SYS_LOOKUP_DCOOKIE:         "sys_lookup_dcookie",
	syscall.SYS_EPOLL_CREATE:           "sys_epoll_create",
	syscall.SYS_EPOLL_CTL_OLD:          "sys_epoll_ctl_old",
	syscall.SYS_EPOLL_WAIT_OLD:         "sys_epoll_wait_old",
	syscall.SYS_REMAP_FILE_PAGES:       "sys_remap_file_pages",
	syscall.SYS_GETDENTS64:             "sys_getdents64",
	syscall.SYS_SET_TID_ADDRESS:        "sys_set_tid_address",
	syscall.SYS_RESTART_SYSCALL:        "sys_restart_syscall",
	syscall.SYS_SEMTIMEDOP:             "sys_semtimedop",
	syscall.SYS_FADVISE64:              "sys_fadvise64",
	syscall.SYS_TIMER_CREATE:           "sys_timer_create",
	syscall.SYS_TIMER_SETTIME:          "sys_timer_settime",
	syscall.SYS_TIMER_GETTIME:          "sys_timer_gettime",
	syscall.SYS_TIMER_GETOVERRUN:       "sys_timer_getoverrun",
	syscall.SYS_TIMER_DELETE:           "sys_timer_delete",
	syscall.SYS_CLOCK_SETTIME:          "sys_clock_settime",
	syscall.SYS_CLOCK_GETTIME:          "sys_clock_gettime",
	syscall.SYS_CLOCK_GETRES:           "sys_clock_getres",
	syscall.SYS_CLOCK_NANOSLEEP:        "sys_clock_nanosleep",
	syscall.SYS_EXIT_GROUP:             "sys_exit_group",
	syscall.SYS_EPOLL_WAIT:             "sys_epoll_wait",
	syscall.SYS_EPOLL_CTL:              "sys_epoll_ctl",
	syscall.SYS_TGKILL:                 "sys_tgkill",
	syscall.SYS_UTIMES:                 "sys_utimes",
	syscall.SYS_VSERVER:                "sys_vserver",
	syscall.SYS_MBIND:                  "sys_mbind",
	syscall.SYS_SET_MEMPOLICY:          "sys_set_mempolicy",
	syscall.SYS_GET_MEMPOLICY:          "sys_get_mempolicy",
	syscall.SYS_MQ_OPEN:                "sys_mq_open",
	syscall.SYS_MQ_UNLINK:              "sys_mq_unlink",
	syscall.SYS_MQ_TIMEDSEND:           "sys_mq_timedsend",
	syscall.SYS_MQ_TIMEDRECEIVE:        "sys_mq_timedreceive",
	syscall.SYS_MQ_NOTIFY:              "sys_mq_notify",
	syscall.SYS_MQ_GETSETATTR:          "sys_mq_getsetattr",
	syscall.SYS_KEXEC_LOAD:             "sys_kexec_load",
	syscall.SYS_WAITID:                 "sys_waitid",
	syscall.SYS_ADD_KEY:                "sys_add_key",
	syscall.SYS_REQUEST_KEY:            "sys_request_key",
	syscall.SYS_KEYCTL:                 "sys_keyctl",
	syscall.SYS_IOPRIO_SET:             "sys_ioprio_set",
	syscall.SYS_IOPRIO_GET:             "sys_ioprio_get",
	syscall.SYS_INOTIFY_INIT:           "sys_inotify_init",
	syscall.SYS_INOTIFY_ADD_WATCH:      "sys_inotify_add_watch",
	syscall.SYS_INOTIFY_RM_WATCH:       "sys_inotify_rm_watch",
	syscall.SYS_MIGRATE_PAGES:          "sys_migrate_pages",
	syscall.SYS_OPENAT:                 "sys_openat",
	syscall.SYS_MKDIRAT:                "sys_mkdirat",
	syscall.SYS_MKNODAT:                "sys_mknodat",
	syscall.SYS_FCHOWNAT:               "sys_fchownat",
	syscall.SYS_FUTIMESAT:              "sys_futimesat",
	syscall.SYS_NEWFSTATAT:             "sys_newfstatat",
	syscall.SYS_UNLINKAT:               "sys_unlinkat",
	syscall.SYS_RENAMEAT:               "sys_renameat",
	syscall.SYS_LINKAT:                 "sys_linkat",
	syscall.SYS_SYMLINKAT:              "sys_symlinkat",
	syscall.SYS_READLINKAT:             "sys_readlinkat",
	syscall.SYS_FCHMODAT:               "sys_fchmodat",
	syscall.SYS_FACCESSAT:              "sys_faccessat",
	syscall.SYS_PSELECT6:               "sys_pselect6",
	syscall.SYS_PPOLL:                  "sys_ppoll",
	syscall.SYS_UNSHARE:                "sys_unshare",
	syscall.SYS_SET_ROBUST_LIST:        "sys_set_robust_list",
	syscall.SYS_GET_ROBUST_LIST:        "sys_get_robust_list",
	syscall.SYS_SPLICE:                 "sys_splice",
	syscall.SYS_TEE:                    "sys_tee",
	syscall.SYS_SYNC_FILE_RANGE:        "sys_sync_file_range",
	syscall.SYS_VMSPLICE:               "sys_vmsplice",
	syscall.SYS_MOVE_PAGES:             "sys_move_pages",
	syscall.SYS_UTIMENSAT:              "sys_utimensat",
	syscall.SYS_EPOLL_PWAIT:            "sys_epoll_pwait",
	syscall.SYS_SIGNALFD:               "sys_signalfd",
	syscall.SYS_TIMERFD_CREATE:         "sys_timerfd_create",
	syscall.SYS_EVENTFD:                "sys_eventfd",
	syscall.SYS_FALLOCATE:              "sys_fallocate",
	syscall.SYS_TIMERFD_SETTIME:        "sys_timerfd_settime",
	syscall.SYS_TIMERFD_GETTIME:        "sys_timerfd_gettime",
	syscall.SYS_ACCEPT4:                "sys_accept4",
	syscall.SYS_SIGNALFD4:              "sys_signalfd4",
	syscall.SYS_EVENTFD2:               "sys_eventfd2",
	syscall.SYS_EPOLL_CREATE1:          "sys_epoll_create1",
	syscall.SYS_DUP3:                   "sys_dup3",
	syscall.SYS_PIPE2:                  "sys_pipe2",
	syscall.SYS_INOTIFY_INIT1:          "sys_inotify_init1",
	syscall.SYS_PREADV:                 "sys_preadv",
	syscall.SYS_PWRITEV:                "sys_pwritev",
	syscall.SYS_RT_TGSIGQUEUEINFO:      "sys_rt_tgsigqueueinfo",
	syscall.SYS_PERF_EVENT_OPEN:        "sys_perf_event_open",
	syscall.SYS_RECVMMSG:               "sys_recvmmsg",
	syscall.SYS_FANOTIFY_INIT:          "sys_fanotify_init",
	syscall.SYS_FANOTIFY_MARK:          "sys_fanotify_mark",
	syscall.SYS_PRLIMIT64:              "sys_prlimit64",
	syscall.SYS_NAME_TO_HANDLE_AT:      "sys_name_to_handle_at",
	syscall.SYS_OPEN_BY_HANDLE_AT:      "sys_open_by_handle_at",
	syscall.SYS_CLOCK_ADJTIME:          "sys_clock_adjtime",
	syscall.SYS_SYNCFS:                 "sys_syncfs",
	syscall.SYS_SENDMMSG:               "sys_sendmmsg",
	syscall.SYS_SETNS:                  "sys_setns",
	syscall.SYS_GETCPU:                 "sys_getcpu",
	syscall.SYS_PROCESS_VM_READV:       "sys_process_vm_readv",
	syscall.SYS_PROCESS_VM_WRITEV:      "sys_process_vm_writev",
	syscall.SYS_KCMP:                   "sys_kcmp",
	syscall.SYS_FINIT_MODULE:           "sys_finit_module",
	syscall.SYS_SCHED_SETATTR:          "sys_sched_setattr",
	syscall.SYS_SCHED_GETATTR:          "sys_sched_getattr",
	syscall.SYS_RENAMEAT2:              "sys_renameat2",
	syscall.SYS_SECCOMP:                "sys_seccomp",
	syscall.SYS_GETRANDOM:              "sys_getrandom",
	syscall.SYS_MEMFD_CREATE:           "sys_memfd_create",
	syscall.SYS_KEXEC_FILE_LOAD:        "sys_kexec_file_load",
	syscall.SYS_BPF:                    "sys_bpf",
	syscall.SYS_EXECVEAT:               "sys_execveat",
	syscall.SYS_USERFAULTFD:            "sys_userfaultfd",
	syscall.SYS_MEMBARRIER:             "sys_membarrier",
	syscall.SYS_MLOCK2:                 "sys_mlock2",
	syscall.SYS_COPY_FILE_RANGE:        "sys_copy_file_range",
	syscall.SYS_PREADV2:                "sys_preadv2",
	syscall.SYS_PWRITEV2:               "sys_pwritev2",
	syscall.SYS_PKEY_MPROTECT:          "sys_pkey_mprotect",
	syscall.SYS_PKEY_ALLOC:             "sys_pkey_alloc",
	syscall.SYS_PKEY_FREE:              "sys_pkey_free",
	syscall.SYS_STATX:                  "sys_statx",
	syscall.SYS_IO_PGETEVENTS:          "sys_io_pgetevents",
	syscall.SYS_RSEQ:                   "sys_rseq",
	syscall.SYS_PIDFD_SEND_SIGNAL:      "sys_pidfd_send_signal",
	syscall.SYS_IO_URING_SETUP:         "sys_io_uring_setup",
	syscall.SYS_IO_URING_ENTER:         "sys_io_uring_enter",
	syscall.SYS_IO_URING_REGISTER:      "sys_io_uring_register",
	syscall.SYS_OPEN_TREE:              "sys_open_tree",
	syscall.SYS_MOVE_MOUNT:             "sys_move_mount",
	syscall.SYS_FSOPEN:                 "sys_fsopen",
	syscall.SYS_FSCONFIG:               "sys_fsconfig",
	syscall.SYS_FSMOUNT:                "sys_fsmount",
	syscall.SYS_FSPICK:                 "sys_fspick",
	syscall.SYS_PIDFD_OPEN:             "sys_pidfd_open",
	syscall.SYS_CLONE3:                 "sys_clone3",
	syscall.SYS_OPENAT2:                "sys_openat2",
	syscall.SYS_PIDFD_GETFD:            "sys_pidfd_getfd",
}
