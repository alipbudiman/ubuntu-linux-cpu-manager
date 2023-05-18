package function

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	history              = map[time.Time]string{}
	intervalSleep int    = 3
	totalloop     int    = 0
	version       string = "1.0 beta"
	overload      int    = 0
)

func Credit() {
	fmt.Println("\033[31m \n                                                                                 \n _|_|_|_|  _|      _|    _|_|_|              _|_|_|    _|_|_|_|  _|      _|      \n _|          _|  _|    _|                    _|    _|  _|        _|      _|      \n _|_|_|        _|      _|  _|_|  _|_|_|_|_|  _|    _|  _|_|_|    _|      _|      \n _|          _|  _|    _|    _|              _|    _|  _|          _|  _|        \n _|        _|      _|    _|_|_|              _|_|_|    _|_|_|_|      _|          \n                                                                                 \n \033[0m")
	fmt.Print("\nThis is simple code for keep you CPU running without overload")
	fmt.Print("\nThis program made by FXG-Development | Alif Budiman | Ver " + version)
	fmt.Print("\nrunning properly on linux ubuntu 18")
	fmt.Print("\nVisit my github on: https://github.com/alipbudiman")
}

func LocalTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	return now
}

func SetInterval(time int) {
	intervalS := time
	for {
		if intervalS >= 3 {
			intervalSleep = time
			break
		} else {
			var clintv int
			fmt.Print("\nInterval must more or equal than 3")
			fmt.Print("\nEnter Commands: ")
			fmt.Scanln(&clintv)
			intervalS = clintv
		}
	}
}
func GetIP() string {
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			return ip.String()
		}
	}
	return ""
}

func IsCPUManager(mainsystem string) bool {
	out, err := exec.Command("screen", "-ls").Output()
	if err != nil {
		log.Fatal(err)
	}
	Xreplace := strings.Replace(string(out), "There are screens on:", "", -1)
	Xsplit := strings.Split(string(Xreplace), "\n")
	for i := range Xsplit {
		res1 := strings.Split(Xsplit[i], "\t")
		for i, x := range res1 {
			if i == 1 {
				res2 := strings.Split(x, ".")
				if res2[1] == mainsystem {
					return true
				}
			}
		}
	}
	return false

}

func Clearcache() {
	exec.Command("sync;", "echo", "1", ">", "/proc/sys/vm/drop_caches").Run()
	exec.Command("sync;", "echo", "2", ">", "/proc/sys/vm/drop_caches").Run()
	exec.Command("sync;", "echo", "3", ">", "/proc/sys/vm/drop_caches").Run()
	exec.Command("apt-get", "clean").Run()
}

func isOverload() bool {
	TotalCPU := runtime.NumCPU()
	out, err := exec.Command("uptime").Output()
	if err != nil {
		return false
	} else {
		Xread := string(out)
		Xcpu := strings.Split(Xread, "load average: ")
		Ucpu := strings.Split(Xcpu[1], ",")
		CPUused, err := strconv.ParseFloat(Ucpu[0], 64)
		if err != nil {
			return false
		}
		if CPUused > float64(TotalCPU) {
			return true
		} else {
			return false
		}

	}
}

func TeraceMyCPU(cpumanagerskip string) bool {
	totalloop += 1
	TotalCPU := runtime.NumCPU()
	out, err := exec.Command("uptime").Output()
	if err != nil {
		log.Fatal(err)
	} else {
		Xread := string(out)
		Xcpu := strings.Split(Xread, "load average: ")
		Ucpu := strings.Split(Xcpu[1], ",")
		CPUused, err := strconv.ParseFloat(Ucpu[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		timess := LocalTime()
		history[timess] = "Total Cpu: " + strconv.Itoa(TotalCPU) + " | CPU Used: " + Ucpu[0]
		if CPUused > float64(TotalCPU) {
			overload += 1
			out, err := exec.Command("screen", "-ls").Output()
			if err != nil {
				log.Fatal(err)
			}
			Xreplace := strings.Replace(string(out), "There are screens on:", "", -1)
			Xsplit := strings.Split(string(Xreplace), "\n")
			for i := range Xsplit {
				res1 := strings.Split(Xsplit[i], "\t")
				for i, x := range res1 {
					if i == 1 {
						res2 := strings.Split(x, ".")
						if res2[1] != cpumanagerskip {
							exec.Command("screen", "-S", string(x), "-X", "kill").Run()
						}
					}
				}
			}
		}

	}
}

func CallCPUReport() {
	fmt.Println("-------------------------------------")
	for key, val := range history {
		fmt.Println(key, "\t:", val)
	}
	var cl string
	fmt.Print("\nType any for Exit")
	fmt.Print("\nEnter Commands: ")
	fmt.Scanln(&cl)

}

func CallProfile() {
	fmt.Println("-------------------------------------")
	out, err := exec.Command("uptime").Output()
	if err != nil {
		log.Fatal(err)
	} else {
		Xread := string(out)
		Xcpu := strings.Split(Xread, "load average: ")
		Ucpu := strings.Split(Xcpu[1], ",")
		res := "\n[ Profile ]"
		res += "\nIP address: " + GetIP()
		res += "\nTotal CPU: " + strconv.Itoa(runtime.NumCPU())
		res += "\nCPU used: " + Ucpu[0]
		res += "\nLoop: " + strconv.Itoa(totalloop) + " times"
		res += "\nOverload: " + strconv.Itoa(overload) + " times"
		res += "\nApp version: " + version
		fmt.Println(res)
		var cl string
		fmt.Print("\nType any for Exit")
		fmt.Print("\nEnter Commands: ")
		fmt.Scanln(&cl)
	}
}

func CallSettings() {
	for {
		fmt.Println("-------------------------------------")
		res := "\n[ Settings ]"
		res += "\nInterval: " + strconv.Itoa(intervalSleep) + "/hour"
		res += "\n\n[ Set commands ]"
		res += "\n>> [1]. Set interval"
		res += "\n>> [2]. Exit"
		fmt.Print(res)
		var cl int
		fmt.Print("\nEnter Commands: ")
		fmt.Scanln(&cl)
		if cl == 1 {
			fmt.Print("\nEnter interval (in number): ")
			var interval_req int
			fmt.Scanln(&interval_req)
			SetInterval(interval_req)
			fmt.Println("\nInterval set to: " + strconv.Itoa(interval_req))
		} else if cl == 2 {
			break
		} else {
			time.Sleep(2 * time.Second)
			fmt.Println("\nInvalid command [2]")
		}
	}
}
