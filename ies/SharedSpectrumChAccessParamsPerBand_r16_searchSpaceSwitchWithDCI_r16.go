package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SharedSpectrumChAccessParamsPerBand_r16_searchSpaceSwitchWithDCI_r16_Enum_supported uper.Enumerated = 0
)

type SharedSpectrumChAccessParamsPerBand_r16_searchSpaceSwitchWithDCI_r16 struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *SharedSpectrumChAccessParamsPerBand_r16_searchSpaceSwitchWithDCI_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode SharedSpectrumChAccessParamsPerBand_r16_searchSpaceSwitchWithDCI_r16", err)
	}
	return nil
}

func (ie *SharedSpectrumChAccessParamsPerBand_r16_searchSpaceSwitchWithDCI_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode SharedSpectrumChAccessParamsPerBand_r16_searchSpaceSwitchWithDCI_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
