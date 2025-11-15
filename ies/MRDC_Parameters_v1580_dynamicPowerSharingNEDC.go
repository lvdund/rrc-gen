package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MRDC_Parameters_v1580_dynamicPowerSharingNEDC_Enum_supported uper.Enumerated = 0
)

type MRDC_Parameters_v1580_dynamicPowerSharingNEDC struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *MRDC_Parameters_v1580_dynamicPowerSharingNEDC) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode MRDC_Parameters_v1580_dynamicPowerSharingNEDC", err)
	}
	return nil
}

func (ie *MRDC_Parameters_v1580_dynamicPowerSharingNEDC) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode MRDC_Parameters_v1580_dynamicPowerSharingNEDC", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
