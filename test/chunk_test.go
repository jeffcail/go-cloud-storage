package test

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"testing"
)

// 50M
const chunkSize = 50 * 1024 * 1024

func TestGenerateChunkFile(t *testing.T) {
	fileInfo, err := os.Stat("./music/zhoujielun.mp4")
	if err != nil {
		t.Fatal(err)
	}
	chunNum := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))
	openFile, err := os.OpenFile("./music/zhoujielun.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b := make([]byte, chunkSize)

	for i := 0; i < int(chunNum); i++ {
		openFile.Seek(int64(i*chunkSize), 0)
		if chunkSize > fileInfo.Size()-int64(i*chunkSize) {
			b = make([]byte, fileInfo.Size()-int64(i*chunkSize))
		}

		openFile.Read(b)
		err = ioutil.WriteFile(fmt.Sprintf("chunk_%d.mp4", i), b, 0666)
		if err != nil {
			t.Fatal(err)
		}
		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		f.Write(b)
		f.Close()
	}
	openFile.Close()
}

func TestMergeChunkFile(t *testing.T) {
	file, err := os.OpenFile("./zhoujielun2.mp4", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	fileInfo, err := os.Stat("./music/zhoujielun.mp4")
	if err != nil {
		t.Fatal(err)
	}
	chunNum := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))
	for i := 0; i < int(chunNum); i++ {
		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_RDONLY, os.ModePerm)
		if err != nil {
			t.Fatal()
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}

		file.Write(b)
		f.Close()
	}
	file.Close()
}

func TestCheckFileHash(t *testing.T) {
	file1, err := os.OpenFile("./music/zhoujielun.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b1, err := ioutil.ReadAll(file1)
	if err != nil {
		t.Fatal(err)
	}

	file2, err := os.OpenFile("./zhoujielun2.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := ioutil.ReadAll(file2)
	if err != nil {
		t.Fatal(err)
	}

	s1 := fmt.Sprintf("%x", md5.Sum(b1))
	s2 := fmt.Sprintf("%x", md5.Sum(b2))
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 == s2)
}
