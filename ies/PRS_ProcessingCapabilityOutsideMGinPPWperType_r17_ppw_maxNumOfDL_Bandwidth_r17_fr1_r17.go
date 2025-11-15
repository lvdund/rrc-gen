package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_Bandwidth_r17_fr1_r17_Enum_mhz5   uper.Enumerated = 0
	PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_Bandwidth_r17_fr1_r17_Enum_mhz10  uper.Enumerated = 1
	PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_Bandwidth_r17_fr1_r17_Enum_mhz20  uper.Enumerated = 2
	PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_Bandwidth_r17_fr1_r17_Enum_mhz40  uper.Enumerated = 3
	PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_Bandwidth_r17_fr1_r17_Enum_mhz50  uper.Enumerated = 4
	PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_Bandwidth_r17_fr1_r17_Enum_mhz80  uper.Enumerated = 5
	PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_Bandwidth_r17_fr1_r17_Enum_mhz100 uper.Enumerated = 6
)

type PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_Bandwidth_r17_fr1_r17 struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_Bandwidth_r17_fr1_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_Bandwidth_r17_fr1_r17", err)
	}
	return nil
}

func (ie *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_Bandwidth_r17_fr1_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_Bandwidth_r17_fr1_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
