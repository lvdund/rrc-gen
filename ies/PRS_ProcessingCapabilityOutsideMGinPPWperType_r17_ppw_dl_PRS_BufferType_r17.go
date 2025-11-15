package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_dl_PRS_BufferType_r17_Enum_type1 uper.Enumerated = 0
	PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_dl_PRS_BufferType_r17_Enum_type2 uper.Enumerated = 1
)

type PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_dl_PRS_BufferType_r17 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_dl_PRS_BufferType_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_dl_PRS_BufferType_r17", err)
	}
	return nil
}

func (ie *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_dl_PRS_BufferType_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_dl_PRS_BufferType_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
