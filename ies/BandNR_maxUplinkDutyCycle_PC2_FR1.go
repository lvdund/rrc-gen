package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_maxUplinkDutyCycle_PC2_FR1_Enum_n60  uper.Enumerated = 0
	BandNR_maxUplinkDutyCycle_PC2_FR1_Enum_n70  uper.Enumerated = 1
	BandNR_maxUplinkDutyCycle_PC2_FR1_Enum_n80  uper.Enumerated = 2
	BandNR_maxUplinkDutyCycle_PC2_FR1_Enum_n90  uper.Enumerated = 3
	BandNR_maxUplinkDutyCycle_PC2_FR1_Enum_n100 uper.Enumerated = 4
)

type BandNR_maxUplinkDutyCycle_PC2_FR1 struct {
	Value uper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *BandNR_maxUplinkDutyCycle_PC2_FR1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode BandNR_maxUplinkDutyCycle_PC2_FR1", err)
	}
	return nil
}

func (ie *BandNR_maxUplinkDutyCycle_PC2_FR1) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode BandNR_maxUplinkDutyCycle_PC2_FR1", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
