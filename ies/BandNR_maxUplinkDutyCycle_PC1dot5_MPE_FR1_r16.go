package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_Enum_n10  uper.Enumerated = 0
	BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_Enum_n15  uper.Enumerated = 1
	BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_Enum_n20  uper.Enumerated = 2
	BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_Enum_n25  uper.Enumerated = 3
	BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_Enum_n30  uper.Enumerated = 4
	BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_Enum_n40  uper.Enumerated = 5
	BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_Enum_n50  uper.Enumerated = 6
	BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_Enum_n60  uper.Enumerated = 7
	BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_Enum_n70  uper.Enumerated = 8
	BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_Enum_n80  uper.Enumerated = 9
	BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_Enum_n90  uper.Enumerated = 10
	BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_Enum_n100 uper.Enumerated = 11
)

type BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16 struct {
	Value uper.Enumerated `lb:0,ub:11,madatory`
}

func (ie *BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("Encode BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16", err)
	}
	return nil
}

func (ie *BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("Decode BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
