package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookParameters_type1_multiPanel_nrofPanels_Enum_n2 uper.Enumerated = 0
	CodebookParameters_type1_multiPanel_nrofPanels_Enum_n4 uper.Enumerated = 1
)

type CodebookParameters_type1_multiPanel_nrofPanels struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *CodebookParameters_type1_multiPanel_nrofPanels) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode CodebookParameters_type1_multiPanel_nrofPanels", err)
	}
	return nil
}

func (ie *CodebookParameters_type1_multiPanel_nrofPanels) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode CodebookParameters_type1_multiPanel_nrofPanels", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
