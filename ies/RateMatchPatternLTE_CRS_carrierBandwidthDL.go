package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RateMatchPatternLTE_CRS_carrierBandwidthDL_Enum_n6     uper.Enumerated = 0
	RateMatchPatternLTE_CRS_carrierBandwidthDL_Enum_n15    uper.Enumerated = 1
	RateMatchPatternLTE_CRS_carrierBandwidthDL_Enum_n25    uper.Enumerated = 2
	RateMatchPatternLTE_CRS_carrierBandwidthDL_Enum_n50    uper.Enumerated = 3
	RateMatchPatternLTE_CRS_carrierBandwidthDL_Enum_n75    uper.Enumerated = 4
	RateMatchPatternLTE_CRS_carrierBandwidthDL_Enum_n100   uper.Enumerated = 5
	RateMatchPatternLTE_CRS_carrierBandwidthDL_Enum_spare2 uper.Enumerated = 6
	RateMatchPatternLTE_CRS_carrierBandwidthDL_Enum_spare1 uper.Enumerated = 7
)

type RateMatchPatternLTE_CRS_carrierBandwidthDL struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RateMatchPatternLTE_CRS_carrierBandwidthDL) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RateMatchPatternLTE_CRS_carrierBandwidthDL", err)
	}
	return nil
}

func (ie *RateMatchPatternLTE_CRS_carrierBandwidthDL) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RateMatchPatternLTE_CRS_carrierBandwidthDL", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
