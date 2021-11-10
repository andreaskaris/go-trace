# go-trace

Usage:
~~~
[akaris@linux go-trace (master)]$ make build
go build -o bin/trace cmd/trace/trace.go
[akaris@linux go-trace (master)]$ bin/trace -p $(pidof testcmd) -f
2021/11/10 16:06:14 syscall.PtraceAttach(34870)
2021/11/10 16:06:14 syscall.PtraceAttach(34867)
2021-11-10 16:06:14.258172924 +0100 CET m=+0.000568318 PID: 34870, Syscall: sys_epoll_pwait
2021/11/10 16:06:14 syscall.PtraceAttach(34866)
2021/11/10 16:06:14 syscall.PtraceAttach(34868)
2021-11-10 16:06:14.258230659 +0100 CET m=+0.000626083 PID: 34867, Syscall: sys_restart_syscall
2021/11/10 16:06:14 syscall.PtraceAttach(34869)
2021-11-10 16:06:14.258290641 +0100 CET m=+0.000686095 PID: 34866, Syscall: sys_futex
2021-11-10 16:06:14.258322864 +0100 CET m=+0.000718270 PID: 34868, Syscall: sys_futex
2021-11-10 16:06:14.258348944 +0100 CET m=+0.000744348 PID: 34869, Syscall: sys_futex
2021-11-10 16:06:14.343712792 +0100 CET m=+0.086108344 PID: 34870, Syscall: sys_futex
2021-11-10 16:06:14.343906287 +0100 CET m=+0.086301837 PID: 34867, Syscall: sys_nanosleep
2021-11-10 16:06:14.344090188 +0100 CET m=+0.086485844 PID: 34870, Syscall: sys_write
2021-11-10 16:06:14.344096686 +0100 CET m=+0.086492265 PID: 34866, Syscall: sys_futex
2021-11-10 16:06:14.344355733 +0100 CET m=+0.086751238 PID: 34867, Syscall: sys_futex
2021-11-10 16:06:14.344526857 +0100 CET m=+0.086922402 PID: 34866, Syscall: sys_epoll_pwait
2021-11-10 16:06:14.344557261 +0100 CET m=+0.086952830 PID: 34870, Syscall: sys_futex
2021-11-10 16:06:14.34458538 +0100 CET m=+0.086980927 PID: 34867, Syscall: sys_nanosleep
2021-11-10 16:06:14.344869485 +0100 CET m=+0.087264969 PID: 34867, Syscall: sys_nanosleep
2021-11-10 16:06:14.344900483 +0100 CET m=+0.087295953 PID: 34870, Syscall: sys_futex
2021-11-10 16:06:14.345126467 +0100 CET m=+0.087521920 PID: 34867, Syscall: sys_futex
2021-11-10 16:06:14.344901745 +0100 CET m=+0.087297211 PID: 34868, Syscall: sys_futex
(...)
~~~
