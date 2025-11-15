package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17_Enum_n2     uper.Enumerated = 0
	NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17_Enum_n4     uper.Enumerated = 1
	NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17_Enum_n6     uper.Enumerated = 2
	NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17_Enum_n12    uper.Enumerated = 3
	NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17_Enum_spare4 uper.Enumerated = 4
	NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17_Enum_spare3 uper.Enumerated = 5
	NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17_Enum_spare2 uper.Enumerated = 6
	NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17_Enum_spare1 uper.Enumerated = 7
)

type NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17", err)
	}
	return nil
}

func (ie *NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
