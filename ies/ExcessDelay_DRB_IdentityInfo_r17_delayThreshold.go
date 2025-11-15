package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms0dot25 uper.Enumerated = 0
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms0dot5  uper.Enumerated = 1
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms1      uper.Enumerated = 2
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms2      uper.Enumerated = 3
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms4      uper.Enumerated = 4
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms5      uper.Enumerated = 5
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms10     uper.Enumerated = 6
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms20     uper.Enumerated = 7
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms30     uper.Enumerated = 8
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms40     uper.Enumerated = 9
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms50     uper.Enumerated = 10
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms60     uper.Enumerated = 11
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms70     uper.Enumerated = 12
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms80     uper.Enumerated = 13
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms90     uper.Enumerated = 14
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms100    uper.Enumerated = 15
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms150    uper.Enumerated = 16
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms300    uper.Enumerated = 17
	ExcessDelay_DRB_IdentityInfo_r17_delayThreshold_Enum_ms500    uper.Enumerated = 18
)

type ExcessDelay_DRB_IdentityInfo_r17_delayThreshold struct {
	Value uper.Enumerated `lb:0,ub:18,madatory`
}

func (ie *ExcessDelay_DRB_IdentityInfo_r17_delayThreshold) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 18}, false); err != nil {
		return utils.WrapError("Encode ExcessDelay_DRB_IdentityInfo_r17_delayThreshold", err)
	}
	return nil
}

func (ie *ExcessDelay_DRB_IdentityInfo_r17_delayThreshold) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 18}, false); err != nil {
		return utils.WrapError("Decode ExcessDelay_DRB_IdentityInfo_r17_delayThreshold", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
