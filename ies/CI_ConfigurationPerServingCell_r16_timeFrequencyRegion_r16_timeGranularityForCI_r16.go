package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16_timeGranularityForCI_r16_Enum_n1  uper.Enumerated = 0
	CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16_timeGranularityForCI_r16_Enum_n2  uper.Enumerated = 1
	CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16_timeGranularityForCI_r16_Enum_n4  uper.Enumerated = 2
	CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16_timeGranularityForCI_r16_Enum_n7  uper.Enumerated = 3
	CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16_timeGranularityForCI_r16_Enum_n14 uper.Enumerated = 4
	CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16_timeGranularityForCI_r16_Enum_n28 uper.Enumerated = 5
)

type CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16_timeGranularityForCI_r16 struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16_timeGranularityForCI_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16_timeGranularityForCI_r16", err)
	}
	return nil
}

func (ie *CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16_timeGranularityForCI_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16_timeGranularityForCI_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
