package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16_bundleSizeSet1_r16_Enum_n4          uper.Enumerated = 0
	PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16_bundleSizeSet1_r16_Enum_wideband    uper.Enumerated = 1
	PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16_bundleSizeSet1_r16_Enum_n2_wideband uper.Enumerated = 2
	PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16_bundleSizeSet1_r16_Enum_n4_wideband uper.Enumerated = 3
)

type PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16_bundleSizeSet1_r16 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16_bundleSizeSet1_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16_bundleSizeSet1_r16", err)
	}
	return nil
}

func (ie *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16_bundleSizeSet1_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode PDSCH_Config_prb_BundlingTypeDCI_1_2_r16_dynamicBundling_r16_bundleSizeSet1_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
