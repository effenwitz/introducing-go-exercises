package main

import (
	"container/list"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings" //strings package
	//input/output package
	"os" //file and folder operations
)

func main() {
	fmt.Println("Packages in go")
	// ioEx()
	// fileFolderExReadFile()
	// fileFolderExWriteFile()
	//readDirEx()
	//listEx()
	//sortByName()
	//nonCryptoHash()
	cyrptoHashEx()
}

func stringsEx() {
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
}

//main interfaces -Reader and Writer
//for in memory read write
func ioEx() {
	//buffer struct in bytes package
	// var buf bytes.Buffer
	// buf.Write([]bytes("test"))

	r := strings.NewReader("hi")
	http.Post("http://example.com", "text/plain", r)

}

//folder and file operations
func fileFolderExReadFile() {
	fmt.Println("File and Folder Package")
	//open a file
	file, err := os.Open("testRead.txt")
	if err != nil {
		fmt.Println(err)
		//handle error here
		return
	}
	defer file.Close()

	//get file size
	fileInfo, err := file.Stat()
	if err != nil {
		return
	}
	fileSize := fileInfo.Size()
	fmt.Println(fileSize)

	//read the file in one shot
	bs := make([]byte, fileSize)
	_, err2 := file.Read(bs)
	if err2 != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)

	//shorter way to read file
	bs1, err1 := ioutil.ReadFile("test.txt")
	if err1 != nil {
		return
	}
	str1 := string(bs1)
	fmt.Println(str1)

}

//to create and write a file
func fileFolderExWriteFile() {
	file, err := os.Create("testCreate.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err1 := file.WriteString("Kirty")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
}

//to get contents of a directory
func readDirEx() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err1 := dir.Readdir(-1)
	if err1 != nil {
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

//containers and sort
func listEx() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)
	x.PushBack("Kirty")

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}

//Person in a struct
type Person struct {
	Name string
	Age  int
}

//ByName is a Person array
type ByName []Person

//Len is a method of sort.Interface
func (ps ByName) Len() int {
	return len(ps)
}

//Less is a method of sort.Interface
func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

//Swap method from sort Interface
func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func sortByName() {
	people := []Person{
		{"xJack", 10},
		{"Jill", 11},
	}

	sort.Sort(ByName(people))
	fmt.Println(people)
}

//hashes and cryptography
//hash - cryptographic hash and non-cryptographic hash
func nonCryptoHash() {
	//create a hasher
	h := crc32.NewIEEE()
	//write data to it
	h.Write([]byte("kirty"))
	//calculate the crc32 checksum
	v := h.Sum32()
	fmt.Println(v)

}

//compare two files using nonCryptoHash
func getHash(filename string) (uint32, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	//create a hasher
	h := crc32.NewIEEE()
	//copy the file into the hasher
	_, err1 := io.Copy(h, f)
	if err1 != nil {
		return 0, err
	}
	return h.Sum32(), nil
}

//cryptographic hash
func cyrptoHashEx() {
	h := sha1.New()
	h.Write([]byte("kirty"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
	hx := hex.EncodeToString(bs)
	fmt.Println(hx)
}
