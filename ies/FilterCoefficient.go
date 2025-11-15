package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FilterCoefficient_Enum_fc0    uper.Enumerated = 0
	FilterCoefficient_Enum_fc1    uper.Enumerated = 1
	FilterCoefficient_Enum_fc2    uper.Enumerated = 2
	FilterCoefficient_Enum_fc3    uper.Enumerated = 3
	FilterCoefficient_Enum_fc4    uper.Enumerated = 4
	FilterCoefficient_Enum_fc5    uper.Enumerated = 5
	FilterCoefficient_Enum_fc6    uper.Enumerated = 6
	FilterCoefficient_Enum_fc7    uper.Enumerated = 7
	FilterCoefficient_Enum_fc8    uper.Enumerated = 8
	FilterCoefficient_Enum_fc9    uper.Enumerated = 9
	FilterCoefficient_Enum_fc11   uper.Enumerated = 10
	FilterCoefficient_Enum_fc13   uper.Enumerated = 11
	FilterCoefficient_Enum_fc15   uper.Enumerated = 12
	FilterCoefficient_Enum_fc17   uper.Enumerated = 13
	FilterCoefficient_Enum_fc19   uper.Enumerated = 14
	FilterCoefficient_Enum_spare1 uper.Enumerated = 15
)

type FilterCoefficient struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *FilterCoefficient) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode FilterCoefficient", err)
	}
	return nil
}

func (ie *FilterCoefficient) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode FilterCoefficient", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
