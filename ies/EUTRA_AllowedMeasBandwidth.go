package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	EUTRA_AllowedMeasBandwidth_Enum_mbw6   uper.Enumerated = 0
	EUTRA_AllowedMeasBandwidth_Enum_mbw15  uper.Enumerated = 1
	EUTRA_AllowedMeasBandwidth_Enum_mbw25  uper.Enumerated = 2
	EUTRA_AllowedMeasBandwidth_Enum_mbw50  uper.Enumerated = 3
	EUTRA_AllowedMeasBandwidth_Enum_mbw75  uper.Enumerated = 4
	EUTRA_AllowedMeasBandwidth_Enum_mbw100 uper.Enumerated = 5
)

type EUTRA_AllowedMeasBandwidth struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *EUTRA_AllowedMeasBandwidth) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode EUTRA_AllowedMeasBandwidth", err)
	}
	return nil
}

func (ie *EUTRA_AllowedMeasBandwidth) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode EUTRA_AllowedMeasBandwidth", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
