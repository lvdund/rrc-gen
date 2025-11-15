package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n1   uper.Enumerated = 0
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n2   uper.Enumerated = 1
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n4   uper.Enumerated = 2
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n5   uper.Enumerated = 3
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n7   uper.Enumerated = 4
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n8   uper.Enumerated = 5
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n10  uper.Enumerated = 6
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n14  uper.Enumerated = 7
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n16  uper.Enumerated = 8
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n20  uper.Enumerated = 9
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n28  uper.Enumerated = 10
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n32  uper.Enumerated = 11
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n35  uper.Enumerated = 12
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n42  uper.Enumerated = 13
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n56  uper.Enumerated = 14
	CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16_Enum_n112 uper.Enumerated = 15
)

type CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16", err)
	}
	return nil
}

func (ie *CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode CI_ConfigurationPerServingCell_r16_ci_PayloadSize_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
