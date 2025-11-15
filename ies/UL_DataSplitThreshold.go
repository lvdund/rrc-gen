package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UL_DataSplitThreshold_Enum_b0       uper.Enumerated = 0
	UL_DataSplitThreshold_Enum_b100     uper.Enumerated = 1
	UL_DataSplitThreshold_Enum_b200     uper.Enumerated = 2
	UL_DataSplitThreshold_Enum_b400     uper.Enumerated = 3
	UL_DataSplitThreshold_Enum_b800     uper.Enumerated = 4
	UL_DataSplitThreshold_Enum_b1600    uper.Enumerated = 5
	UL_DataSplitThreshold_Enum_b3200    uper.Enumerated = 6
	UL_DataSplitThreshold_Enum_b6400    uper.Enumerated = 7
	UL_DataSplitThreshold_Enum_b12800   uper.Enumerated = 8
	UL_DataSplitThreshold_Enum_b25600   uper.Enumerated = 9
	UL_DataSplitThreshold_Enum_b51200   uper.Enumerated = 10
	UL_DataSplitThreshold_Enum_b102400  uper.Enumerated = 11
	UL_DataSplitThreshold_Enum_b204800  uper.Enumerated = 12
	UL_DataSplitThreshold_Enum_b409600  uper.Enumerated = 13
	UL_DataSplitThreshold_Enum_b819200  uper.Enumerated = 14
	UL_DataSplitThreshold_Enum_b1228800 uper.Enumerated = 15
	UL_DataSplitThreshold_Enum_b1638400 uper.Enumerated = 16
	UL_DataSplitThreshold_Enum_b2457600 uper.Enumerated = 17
	UL_DataSplitThreshold_Enum_b3276800 uper.Enumerated = 18
	UL_DataSplitThreshold_Enum_b4096000 uper.Enumerated = 19
	UL_DataSplitThreshold_Enum_b4915200 uper.Enumerated = 20
	UL_DataSplitThreshold_Enum_b5734400 uper.Enumerated = 21
	UL_DataSplitThreshold_Enum_b6553600 uper.Enumerated = 22
	UL_DataSplitThreshold_Enum_infinity uper.Enumerated = 23
	UL_DataSplitThreshold_Enum_spare8   uper.Enumerated = 24
	UL_DataSplitThreshold_Enum_spare7   uper.Enumerated = 25
	UL_DataSplitThreshold_Enum_spare6   uper.Enumerated = 26
	UL_DataSplitThreshold_Enum_spare5   uper.Enumerated = 27
	UL_DataSplitThreshold_Enum_spare4   uper.Enumerated = 28
	UL_DataSplitThreshold_Enum_spare3   uper.Enumerated = 29
	UL_DataSplitThreshold_Enum_spare2   uper.Enumerated = 30
	UL_DataSplitThreshold_Enum_spare1   uper.Enumerated = 31
)

type UL_DataSplitThreshold struct {
	Value uper.Enumerated `lb:0,ub:31,madatory`
}

func (ie *UL_DataSplitThreshold) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode UL_DataSplitThreshold", err)
	}
	return nil
}

func (ie *UL_DataSplitThreshold) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode UL_DataSplitThreshold", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
