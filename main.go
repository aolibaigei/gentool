package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Hostinfo struct {
	ID       int    `json:"id"`
	Hostname string `json:"hostname"`
	Ipaddr   string `json:"ipaddr"`
	Arch     string `json:"arch"`
	OSs      string `json:"os"`
}

type Containerinfo struct {
	Containerid string `json:"cid"`
	Image       string `json:"image"`
	Command     string `json:"command"`
	Created     string `json:"created"`
	Status      string `json:"status"`
	Port        string `json:"port"`
	Name        string `json:"name"`
}

type Imageinfo struct {
	Image   string `json:"image"`
	Imageid string `json:"imageid"`
	Created string `json:"created"`
	Size    string `json:"size"`
}

func genMachineID() string {

	u1, _ := uuid.NewV4()

	return strings.ReplaceAll(u1.String(), "-", "")

}

func genIpaddr() string {

	// rand.Seed(time.Now().Unix())
	rand.New(rand.NewSource(time.Now().UnixNano()))

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

func genContainerName() string {

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
	timedir[2] = "Up 10 months"

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

func genImageSize() string {
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

func getHost(num int) {

	host := &Hostinfo{}

	for i := 1; i < num+1; i++ {
		host.ID = i
		host.Hostname = genHostName()
		host.Ipaddr = genIpaddr()
		host.OSs = genOSVer()
		host.Arch = genHostArch()

		jsons, errs := json.Marshal(host)
		if errs != nil {
			fmt.Println(errs.Error())
		}
		fmt.Println(string(jsons) + ",")

	}

}

func getContanier(num int) {
	container := &Containerinfo{}
	for i := 0; i < num; i++ {
		container.Containerid = genContainerID()
		container.Image = genContainerImageName()
		container.Port = genContainerPort()
		container.Created = genCreateTime()
		container.Status = genContainerStatus()
		container.Name = genContainerName()
		container.Command = "/bin/bash"

		jsons, errs := json.Marshal(container)
		if errs != nil {
			fmt.Println(errs.Error())
		}
		fmt.Println(string(jsons) + ",")

	}
}

func getImage(num int) {
	image := &Imageinfo{}
	for i := 0; i < num; i++ {
		image.Image = genImageName()
		image.Imageid = genImageID()
		image.Created = genCreateTime()
		image.Size = genImageSize()

		jsons, errs := json.Marshal(image)
		if errs != nil {
			fmt.Println(errs.Error())
		}
		fmt.Println(string(jsons) + ",")

	}
}

func main() {
	// fmt.Println(genMachineID())
	// fmt.Println(genIpaddr())
	// fmt.Println(genHostName())
	// fmt.Println(genHostArch())
	// fmt.Println(genOSVer())

	// getHost(10)
	// getContanier(10)
	getImage(4)

}
