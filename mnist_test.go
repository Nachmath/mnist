// Copyright (c) 2019 邢子文(XING-ZIWEN) <Rick.Xing@Nachmath.com>
// STAMP|42704

package mnist

import (
	"fmt"
	"testing"
)

func Test_SamplePump_LabelPump(t *testing.T) {
	var sp SamplePump
	err := sp.Open(TrainSetSampleFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sp.Close()

	var lp LabelPump
	err = lp.Open(TrainSetLabelFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer lp.Close()

	oneSample := sp.One(0)
	fmt.Println(oneSample)
	oneLabel := lp.One(0)
	fmt.Println(oneLabel)

	someSamples := sp.Some(59998, 60000)
	fmt.Println(someSamples)
	someLabels := lp.Some(59998, 60000)
	fmt.Println(someLabels)
}

func Test_ExamplePump(t *testing.T) {
	var ep ExamplePump
	err := ep.Open(TrainSetSampleFileName, TrainSetLabelFileName)
	if err != nil {
		t.Fatal(err)
	}
	defer ep.Close()

	example := ep.One(55)
	fmt.Println(example)
}

func Test_GenImage(t *testing.T) {
	var ep ExamplePump
	err := ep.Open(TrainSetSampleFileName, TrainSetLabelFileName)
	if err != nil {
		t.Fatal(err)
	}
	defer ep.Close()

	example := ep.One(55)
	// fmt.Println(example)

	if !(GenExampleImage(example) == true) {
		t.Fatal()
	}
	PrintExample(example)

	for i := 1500; i < 1600; i++ {
		GenExampleImage(ep.One(i))
	}

}
