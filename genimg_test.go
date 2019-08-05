// Copyright (c) 2019 邢子文(XING-ZIWEN) <Rick.Xing@Nachmath.com>
// STAMP|42704

package mnist

import (
	"fmt"
	"testing"
)

const TrainSetSampleFileName = "usable/train-images.idx3-ubyte"
const TrainSetLabelFileName = "usable/train-labels.idx1-ubyte"

func Test(t *testing.T) {
	var mnistSamplePump SamplePump
	err := mnistSamplePump.Open(TrainSetSampleFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer mnistSamplePump.Close()

	oneSample := mnistSamplePump.One(0)
	fmt.Println(len(oneSample))

	// someSamples := mnistSamplePump.Some(59998, 60000)
	// fmt.Println(someSamples)

	var mnistLabelPump LabelPump
	err = mnistLabelPump.Open(TrainSetLabelFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer mnistLabelPump.Close()

	oneLabel := mnistLabelPump.One(0)
	fmt.Println(oneLabel)

	someLabels := mnistLabelPump.Some(59998, 60000)
	fmt.Println(someLabels)
}
