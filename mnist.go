// Copyright (c) 2019 邢子文(XING-ZIWEN) <Rick.Xing@Nachmath.com>
// STAMP|42706

package mnist

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

// Data file name
const (
	TrainSetSampleFileName = "usable/train-images.idx3-ubyte"
	TrainSetLabelFileName  = "usable/train-labels.idx1-ubyte"
)

// Image size
const (
	W = 28
	H = 28
)

// Example ...
type Example struct {
	index  int
	sample []byte
	label  byte
}

// ExamplePump ...
type ExamplePump struct {
	sp SamplePump
	lp LabelPump
}

// Open ...
func (p *ExamplePump) Open(idx3, idx1 string) error {
	err1 := p.sp.Open(idx3)
	err2 := p.lp.Open(idx1)
	if err1 != nil || err2 != nil {
		p.sp.Close()
		p.lp.Close()
		if err1 != nil {
			return err1
		}
		return err2
	}
	return nil
}

// Close ...
func (p *ExamplePump) Close() {
	p.sp.Close()
	p.lp.Close()
}

// One ...
func (p *ExamplePump) One(i int) Example {
	return Example{
		index:  i,
		sample: p.sp.One(i),
		label:  p.lp.One(i)[0],
	}
}

// GenExampleImage 根据Example生成对应的图片文件
// 图片文件名: example_images/`index`_`label`.png
// 背景色为白色, 前景色为黑色, 注:灰度字节值越大表现为颜色越黑
func GenExampleImage(e Example) bool {
	if len(e.sample) != W*H {
		return false
	}
	filename := "example_images/" +
		strconv.Itoa(e.index) + "_" +
		strconv.Itoa(int(e.label)) + ".jpg"
	file, err := os.Create(filename)
	if err != nil {
		return false
	}
	defer file.Close()
	img := image.NewRGBA(image.Rect(0, 0, W, H))
	for ri := 0; ri < H; ri++ {
		for cj := 0; cj < W; cj++ {
			gray := e.sample[ri*W+cj]
			if gray == 0 {
				img.SetRGBA(cj, ri, color.RGBA{255, 255, 255, 0xff})
			} else {
				img.SetRGBA(cj, ri, color.RGBA{255 - gray, 255 - gray, 255 - gray, 0xff})
			}
		}
	}
	if png.Encode(file, img) != nil {
		return false
	}
	return true
}

// PrintExample 打印Example效果
func PrintExample(e Example) {
	fmt.Println("Index:", e.index)
	fmt.Println("Label:", e.label)
	for i, v := range e.sample {
		fmt.Printf("%2x", v)
		if i%W == W-1 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n")
}
