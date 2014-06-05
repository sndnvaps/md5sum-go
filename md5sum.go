//create file md5sum.md5
//check file's md5sum from md5sum.md5

package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	//"path"
	"strings"
)

var (
	F      string // file need to calculate md5sum
	md5out string //write md5sum to file
	Type   string // 'C', "M"
	md5in  string // md5sum in file need to verify
	help   string // show the help info
)

func init() {
	flag.StringVar(&F, "file", "", "File need to calculate md5sum")
	flag.StringVar(&md5out, "md5out", "", "output md5sum to file")
	flag.StringVar(&Type, "Type", "", "C = Check md5sum, M =  Make md5sum ")
	flag.StringVar(&md5in, "md5in", "", "Input md5sum file to check")
	flag.StringVar(&help, "help", "", "print the help info")

	flag.Usage = func() {
		fmt.Printf("Usage: %s -options=value ...\n", os.Args[0])
		flag.PrintDefaults()

	}

}

func MakeMd5() string {
	fileName := F
	var md5file string
	if len(md5out) == 0 {
		md5file = F + ".md5"
	} else {
		md5file = md5out
	}
	//fmt.Println("md5file = ", md5file)
	//md5file := "filelist.go.text.md5"
	//bufs := make([]byte, 1024)

	fin, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(fileName, err)
	}
	defer fin.Close()

	fout, foute := os.Create(md5file)
	if foute != nil {
		fmt.Println(md5file, foute)
	}
	defer fout.Close()
	/*
		fSize, _ := os.Stat(fileName)
		size := fSize.Size()
		filebuf := make([]byte, int(size))

		//md5string := md5.New()

		for size > 0 {

			nsize, e := fin.Read(bufs)
			if e != nil {
				fmt.Println(e)
			}

				_, md5e := md5string.Write(bufs)
				if md5e != nil {
					fmt.Println(md5e)
				}
				filebuf = md5string.Sum(bufs)

			filebuf = append(filebuf, bufs...)

			size = size - int64(nsize)

			if size == 0 {
				break
			}

		}
	*/

	//fmt.Printf("md5.sum = %s\n", hex.EncodeToString(filebuf))
	//fmt.Println("md5.sum= ", filebuf)

	//md5string.Sum(nil)
	Buf, buferr := ioutil.ReadFile(fileName)
	if buferr != nil {
		fmt.Println(fileName, buferr)
	}
	/*
		fmt.Println("md5.sum1 = %s\n", len(hex.EncodeToString(byte2string(md5.Sum(Buf)))), hex.EncodeToString(byte2string(md5.Sum(Buf))))
		md5string.Write(Buf)
	*/
	//fmt.Println("md5.sum2 = %s", len(hex.EncodeToString(md5string.Sum(nil))), hex.EncodeToString(md5string.Sum(nil)))
	temp := hex.EncodeToString(byte2string(md5.Sum(Buf))) + " " + fileName
	_, err = fout.WriteString(temp)
	if err != nil {
		fmt.Println(md5file, err)
	}
	fmt.Printf("md5sum = %s", temp)

	return temp
}

func Verifymd5sum(file string, md5file string) bool {
	md5sumfile, err := ioutil.ReadFile(md5file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(md5sumfile))
	md5sum := string(md5sumfile)
	str := strings.Split(md5sum, " ")
	fmt.Println(len(str))
	md5string := str[0]
	//md5fileName := str[1]

	//check md5sum from file
	Buf, buferr := ioutil.ReadFile(file)
	if buferr != nil {
		fmt.Println(buferr)
	}

	md5fromfile := hex.EncodeToString(byte2string(md5.Sum(Buf)))

	if md5fromfile == md5string {
		return true
	}

	//fmt.Printf("Name = %s md5sum = %s\n", md5string, md5fileName)
	return false
}

func byte2string(in [16]byte) []byte {
	tmp := make([]byte, 16)
	for _, value := range in {
		tmp = append(tmp, value)
	}

	return tmp[16:]
}

func main() {

	flag.Parse()

	if len(Type) == 0 || len(help) != 0 {
		flag.Usage()
	}

	if Type == "M" {
		MakeMd5()
		return
	}

	if Type == "C" {
		T := Verifymd5sum(F, md5in)
		if T {
			fmt.Println("verifiy success ")
		} else {
			fmt.Println("verify failed ")
		}
		return

	}

	/*
		MakeMd5()
		T := Verifymd5sum("filelist.go.text", "filelist.go.text.md5")
		if T {
			fmt.Println("verifiy success ")
		} else {
			fmt.Println("verify failed ")
		}
	*/

}
