package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17_Enum_oneSeventh      uper.Enumerated = 0
	CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17_Enum_threeFourteenth uper.Enumerated = 1
	CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17_Enum_twoSeventh      uper.Enumerated = 2
	CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17_Enum_threeSeventh    uper.Enumerated = 3
	CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17_Enum_oneHalf         uper.Enumerated = 4
	CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17_Enum_fourSeventh     uper.Enumerated = 5
	CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17_Enum_fiveSeventh     uper.Enumerated = 6
	CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17_Enum_spare1          uper.Enumerated = 7
)

type CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17", err)
	}
	return nil
}

func (ie *CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode CrossCarrierSchedulingConfig_ccs_BlindDetectionSplit_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
