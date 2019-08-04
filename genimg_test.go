// Copyright (c) 2019 邢子文(XING-ZIWEN) <Rick.Xing@Nachmath.com>
// STAMP|42704

package mnist

import (
	"fmt"
	"os"
	"testing"
)

func Test_GenImages(t *testing.T) {

}

const TrainSetFileName = "usable/train-images.idx3-ubyte"

func Test_OpenFile(t *testing.T) {
	file, err := os.Open(TrainSetFileName)
	if err != nil {
	}
	defer file.Close()

	magicNumber := make([]byte, 4)
	// reader := bufio.NewReader(file)
	n, err := file.Read(magicNumber)
	if err != nil {
	}
	fmt.Println(n)
	fmt.Println(magicNumber)

	numberOfImages := make([]byte, 4)
	n, err = file.Read(numberOfImages)
	if err == nil {
	}
	fmt.Println(n)
	fmt.Println(numberOfImages)
	var numberOfImagesI int32 = binary.

}
