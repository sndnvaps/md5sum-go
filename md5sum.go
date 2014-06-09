//create file md5sum.md5
//check file's md5sum from md5sum.md5

package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
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

	Buf, buferr := ioutil.ReadFile(fileName)
	if buferr != nil {
		fmt.Println(fileName, buferr)
	}

	temp := hex.EncodeToString(byte2string(md5.Sum(Buf))) + " " + "*" + path.Base(fileName)
	_, err = fout.WriteString(temp)
	if err != nil {
		fmt.Println(md5file, err)
	}
	fmt.Printf("md5sum = %s", temp)

	return temp
}

func Verifymd5sum(file string, md5file string) bool {
	if len(md5file) == 0 {
		f, e := os.Stat(file + ".md5")
		if e != nil {
			fmt.Println(file, ".md5", "Doesn't exits")
			return false
		}
		md5file = f.Name()

	}
	md5sumfile, err := ioutil.ReadFile(md5file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(md5sumfile))

	md5sum := string(md5sumfile)
	str := strings.Split(md5sum, " ")
	//fmt.Println(len(str))
	var md5fileName string
	md5string := str[0] //get the md5sum string from file
	//fmt.Println("str[1][:1] = ", str[1][:1])
	if str[1][:1] == string("*") {

		md5fileName = str[1][1:len(str[1])] // get the filename from md5sum.md5 file
	} else {
		md5fileName = str[1]
	}

	//check filename  with md5sum.md5
	if path.Base(md5fileName) == path.Base(file) {
		fmt.Println("FileName has compact")
	} else {
		fmt.Println("Not the same FileName")
	}

	fmt.Println("md5fileName =", md5fileName[:len(md5fileName)])

	h, _ := os.Open(file)
	buf := md5.New()
	io.Copy(buf, h)

	md5fromfile := hex.EncodeToString(buf.Sum(nil))

	if md5fromfile == md5string {
		return true
	}

	//fmt.Printf("Name = %s md5sum = %s\n", md5string, md5fileName)
	return false
}

func byte2string(in [16]byte) []byte {
	return in[:16]
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

	if Type == "C" /*&& len(md5in) != 0  */ {
		T := Verifymd5sum(F, md5in)
		if T {
			fmt.Println("verifiy success ")
		} else {
			fmt.Println("verify failed ")
		}
	}

}
