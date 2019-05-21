package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func genMachineID() string {

	u1, _ := uuid.NewV4()

	return strings.ReplaceAll(u1.String(), "-", "")

}

func genIpaddr() string {

	rand.Seed(time.Now().Unix())

	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))

	return ip
}

func genHostName() string {

	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	hostname := []byte{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 12; i++ {
		hostname = append(hostname, bytes[r.Intn(len(bytes))])
	}

	return string(hostname)
}

func genHostArch() string {

	archdir := make(map[int]string)
	archdir[0] = "arch64"
	archdir[1] = "x86_64"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	arch := archdir[r.Intn(2)]

	return arch
}

func genOSVer() string {

	osdir := make(map[int]string)
	osdir[0] = "Ubuntu"
	osdir[1] = "CentOS"
	osdir[2] = "KylinOS"
	osdir[3] = "DeepinOS"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	osver := osdir[r.Intn(len(osdir))]

	return osver
}

func genKernelVer() string {

	kerneldir := make(map[int]string)
	kerneldir[0] = "4.16.0-201.fc18.x86_64"
	kerneldir[1] = "4.18.0-20-generic"
	kerneldir[2] = "kernel-2.6.18-53.el5"
	kerneldir[3] = "4.12.14-25.22.1"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	kernelver := kerneldir[r.Intn(len(kerneldir))]

	return kernelver
}

// func genHostInfo() string {
// 	return "test"
// }

// func genContainerInfo() string {
// 	return "test"
// }

// func genImageIsnfo() string {
// 	return "test"
// }

// func genExecEvent() string {
// 	return "test"
// }

// func genconninfo() string {
// 	return "test"
// }

// func gencveinfo() string {
// 	return "test"
// }

func main() {
	// fmt.Println(genMachineID())
	// fmt.Println(genIpaddr())
	// fmt.Println(genHostName())
	// fmt.Println(genHostArch())
	// fmt.Println(genOSVer())
	fmt.Println(genKernelVer())

}
