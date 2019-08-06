// Copyright (c) 2019 邢子文(XING-ZIWEN) <Rick.Xing@Nachmath.com>
// STAMP|42705|42706

package mnist

import (
	"os"

	"github.com/nachmath/godatx/datapump"
)

var (
	_p  *SamplePump
	_   datapump.BinPumper = _p
	_p2 *LabelPump
	_   datapump.BinPumper = _p2
)

// Size
const (
	SampleFileHeaderSize = 16
	LabelFileHeaderSize  = 8
	SampleSize           = W * H
	LabelSize            = 1
)

//---/---/---/---/---/---/---/---/---/---/---/---//

// SamplePump ...
type SamplePump struct {
	F *os.File
}

// Open ...
func (p *SamplePump) Open(idx3 string) error {
	file, err := os.Open(idx3)
	if err != nil {
		return err
	}
	var bytes4 = make([]byte, 4)
	_, err = file.Read(bytes4)
	_, err = file.Read(bytes4)
	// fmt.Println(binary.BigEndian.Uint32(bytes4))
	_, err = file.Read(bytes4)
	// fmt.Println(binary.BigEndian.Uint32(bytes4))
	_, err = file.Read(bytes4)
	// fmt.Println(binary.BigEndian.Uint32(bytes4))
	if err != nil {
		return err
	}
	p.F = file
	return nil
}

// Close ...
func (p *SamplePump) Close() {
	p.F.Close()
}

// One ...
func (p *SamplePump) One(i int) []byte {
	_, err := p.F.Seek(int64(SampleFileHeaderSize+i*SampleSize), 0)
	if err != nil {
		return nil
	}
	var bytes = make([]byte, SampleSize)
	n, _ := p.F.Read(bytes)
	if n != SampleSize {
		return nil
	}
	return bytes
}

// Some ...
func (p *SamplePump) Some(begin, end int) [][]byte {
	_, err := p.F.Seek(int64(SampleFileHeaderSize+begin*SampleSize), 0)
	if err != nil {
		return nil
	}
	somebytes := make([][]byte, end-begin)
	count := 0
	for i := 0; i < len(somebytes); i++ {
		somebytes[i] = make([]byte, SampleSize)
		n, _ := p.F.Read(somebytes[i])
		if n != SampleSize {
			break
		}
		count++
	}
	return somebytes[:count]
}

//---/---/---/---/---/---/---/---/---/---/---/---//

// LabelPump ...
type LabelPump struct {
	F *os.File
}

// Open ...
func (p *LabelPump) Open(idx1 string) error {
	file, err := os.Open(idx1)
	if err != nil {
		return err
	}
	var bytes4 = make([]byte, 4)
	_, err = file.Read(bytes4)
	_, err = file.Read(bytes4)
	// fmt.Println(binary.BigEndian.Uint32(bytes4))
	if err != nil {
		return err
	}
	p.F = file
	return nil
}

// Close ...
func (p *LabelPump) Close() {
	p.F.Close()
}

// One ...
func (p *LabelPump) One(i int) []byte {
	_, err := p.F.Seek(int64(LabelFileHeaderSize+i*LabelSize), 0)
	if err != nil {
		return nil
	}
	var bytes = make([]byte, LabelSize)
	n, _ := p.F.Read(bytes)
	if n != LabelSize {
		return nil
	}
	return bytes
}

// Some ...
func (p *LabelPump) Some(begin, end int) [][]byte {
	_, err := p.F.Seek(int64(LabelFileHeaderSize+begin*LabelSize), 0)
	if err != nil {
		return nil
	}
	somebytes := make([][]byte, end-begin)
	count := 0
	for i := 0; i < len(somebytes); i++ {
		somebytes[i] = make([]byte, LabelSize)
		n, _ := p.F.Read(somebytes[i])
		if n != LabelSize {
			break
		}
		count++
	}
	return somebytes[:count]
}
