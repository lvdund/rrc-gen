package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17_Enum_n2  uper.Enumerated = 0
	CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17_Enum_n4  uper.Enumerated = 1
	CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17_Enum_n8  uper.Enumerated = 2
	CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17_Enum_n12 uper.Enumerated = 3
	CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17_Enum_n16 uper.Enumerated = 4
	CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17_Enum_n24 uper.Enumerated = 5
	CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17_Enum_n32 uper.Enumerated = 6
)

type CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17 struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17", err)
	}
	return nil
}

func (ie *CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode CSI_MultiTRP_SupportedCombinations_r17_maxNumTx_Ports_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
