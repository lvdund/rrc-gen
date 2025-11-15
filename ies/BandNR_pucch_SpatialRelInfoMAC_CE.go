package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_pucch_SpatialRelInfoMAC_CE_Enum_supported uper.Enumerated = 0
)

type BandNR_pucch_SpatialRelInfoMAC_CE struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *BandNR_pucch_SpatialRelInfoMAC_CE) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode BandNR_pucch_SpatialRelInfoMAC_CE", err)
	}
	return nil
}

func (ie *BandNR_pucch_SpatialRelInfoMAC_CE) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode BandNR_pucch_SpatialRelInfoMAC_CE", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
