package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PHR_Config_phr_Tx_PowerFactorChange_Enum_dB1      uper.Enumerated = 0
	PHR_Config_phr_Tx_PowerFactorChange_Enum_dB3      uper.Enumerated = 1
	PHR_Config_phr_Tx_PowerFactorChange_Enum_dB6      uper.Enumerated = 2
	PHR_Config_phr_Tx_PowerFactorChange_Enum_infinity uper.Enumerated = 3
)

type PHR_Config_phr_Tx_PowerFactorChange struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *PHR_Config_phr_Tx_PowerFactorChange) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode PHR_Config_phr_Tx_PowerFactorChange", err)
	}
	return nil
}

func (ie *PHR_Config_phr_Tx_PowerFactorChange) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode PHR_Config_phr_Tx_PowerFactorChange", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
