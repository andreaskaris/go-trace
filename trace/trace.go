package trace

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	syscall "golang.org/x/sys/unix"
)

func listChildren(pid int) ([]int, error) {
	var children []int

	files, err := os.ReadDir(
		fmt.Sprintf("/proc/%d/task", pid),
	)
	if err != nil {
		return nil, fmt.Errorf("Could not read dir: %s", err.Error())
	}
	for _, file := range files {
		if file.IsDir() {
			i, err := strconv.Atoi(file.Name())
			if err != nil {
				return nil, fmt.Errorf("Could not convert name to int: %s", err.Error())
			}
			children = append(children, i)
		}
	}
	return children, nil
}

func strace(pid int, follow bool) (err error, detachErr error) {
	// must lock OSThread, otherwise we'll get no such process errors
	// https://github.com/golang/go/issues/43685
	// https://github.com/golang/go/issues/7699
	// https://pkg.go.dev/runtime#LockOSThread
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	var regs syscall.PtraceRegs = syscall.PtraceRegs{}
	var exit bool

	// attach
	log.Printf("syscall.PtraceAttach(%d)", pid)
	err = syscall.PtraceAttach(pid)
	if err != nil {
		err = fmt.Errorf("Error in syscall.PtraceAttach(%d): %s", pid, err.Error())
		return
	}

	defer func() {
		log.Printf("Detaching from pid: %d", pid)
		detachErr = syscall.PtraceDetach(pid)
	}()

	// wait for process
	_, err = syscall.Wait4(pid, nil, 0, nil)
	if err != nil {
		return
	}

	// set ptrace options
	options := syscall.PTRACE_O_TRACESYSGOOD | syscall.PTRACE_O_TRACEEXEC | syscall.PTRACE_O_TRACEEXIT
	if follow {
		options = syscall.PTRACE_O_TRACESYSGOOD | syscall.PTRACE_O_TRACEFORK | syscall.PTRACE_O_TRACEVFORK | syscall.PTRACE_O_TRACECLONE | syscall.PTRACE_O_TRACEEXEC | syscall.PTRACE_O_TRACEEXIT
	}

	err = syscall.PtraceSetOptions(pid, options)
	if err != nil {
		err = fmt.Errorf("Failed at syscall.PtraceSetOptions(%d): %s", pid, err.Error())
		return
	}
	for {
		if exit {
			// read the regs
			err = syscall.PtraceGetRegs(pid, &regs)
			if err != nil {
				err = fmt.Errorf("Failed at syscall.PtraceGetRegs(%d): %s", pid, err.Error())
				return
			}
			//fmt.Printf("%v\n", regs)
			fmt.Printf("%v PID: %d, Syscall: %s\n", time.Now(), pid, SyscallName[regs.Orig_rax])
		}
		exit = !exit

		// wait for the next syscall according to options
		err = syscall.PtraceSyscall(pid, 0)
		if err != nil {
			err = fmt.Errorf("Failed at syscall.PtraceSyscall(%d): %s", pid, err.Error())
			return
		}
		// wait for process
		_, err = syscall.Wait4(pid, nil, 0, nil)
		if err != nil {
			return
		}
	}
}

func Strace(pid int, follow bool) error {
	var err error
	attachTo := []int{pid}
	if follow {
		attachTo, err = listChildren(pid)
		if err != nil {
			return err
		}
	}

	wg := sync.WaitGroup{}
	for _, pid := range attachTo {
		wg.Add(1)
		go func(pid int) {
			defer wg.Done()
			err, detachErr := strace(pid, follow)
			if err != nil {
				log.Print(err)
			}
			if detachErr != nil {
				log.Print(detachErr)
			}
		}(pid)
	}
	wg.Wait()

	return nil
}
