package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17_Enum_n5  uper.Enumerated = 0
	BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17_Enum_n15 uper.Enumerated = 1
	BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17_Enum_n25 uper.Enumerated = 2
	BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17_Enum_n32 uper.Enumerated = 3
	BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17_Enum_n35 uper.Enumerated = 4
	BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17_Enum_n45 uper.Enumerated = 5
	BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17_Enum_n50 uper.Enumerated = 6
	BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17_Enum_n64 uper.Enumerated = 7
)

type BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17", err)
	}
	return nil
}

func (ie *BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode BandSidelinkPC5_r16_rx_IUC_Scheme2_Mode2Sidelink_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
