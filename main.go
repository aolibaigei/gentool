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

func genContainerID() string {

	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	containerid := []byte{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 12; i++ {
		containerid = append(containerid, bytes[r.Intn(len(bytes))])
	}

	return string(containerid)
}

func genImageID() string {

	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	imageid := []byte{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 12; i++ {
		imageid = append(imageid, bytes[r.Intn(len(bytes))])
	}

	return string(imageid)
}

func genContainerImageName() string {
	imagedir := make(map[int]string)
	imagedir[0] = "ubuntu"
	imagedir[1] = "redis"
	imagedir[2] = "mongo"
	imagedir[3] = "nginx"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	image := imagedir[r.Intn(len(imagedir))]

	return image
}

func genImageName() string {
	imagedir := make(map[int]string)
	imagedir[0] = "ubuntu:latest"
	imagedir[1] = "redis:1.4"
	imagedir[2] = "mongo:2.5.8"
	imagedir[3] = "nginx:1.15.2"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	image := imagedir[r.Intn(len(imagedir))]

	return image
}

func genCreateTime() string {
	timedir := make(map[int]string)
	timedir[0] = "18 seconds ago"
	timedir[1] = "2 months ago"
	timedir[2] = "5 days ago"
	timedir[3] = "15 days ago"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	time := timedir[r.Intn(len(timedir))]

	return time
}

func genContainerStatus() string {
	timedir := make(map[int]string)
	timedir[0] = "Up 14 seconds"
	timedir[1] = "Up 5 days"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	time := timedir[r.Intn(len(timedir))]

	return time
}

func genContainerPort() string {
	portdir := make(map[int]string)
	portdir[0] = "80"
	portdir[1] = "8080"
	portdir[2] = "8443"
	portdir[3] = "5432"
	portdir[4] = "3306"
	portdir[5] = "11211"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	port := portdir[r.Intn(len(portdir))]

	return port
}

func genContainerSize() string {
	sizedir := make(map[int]string)
	sizedir[0] = "80MB"
	sizedir[1] = "65.5MB"
	sizedir[2] = "165.5MB"
	sizedir[3] = "202MB"
	sizedir[4] = "104MB"
	sizedir[5] = "130MB"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	size := sizedir[r.Intn(len(sizedir))]

	return size
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
	fmt.Println(genContainerSize())

}
