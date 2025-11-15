package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_b56    uper.Enumerated = 0
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_b144   uper.Enumerated = 1
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_b208   uper.Enumerated = 2
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_b256   uper.Enumerated = 3
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_b282   uper.Enumerated = 4
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_b480   uper.Enumerated = 5
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_b640   uper.Enumerated = 6
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_b800   uper.Enumerated = 7
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_b1000  uper.Enumerated = 8
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_b72    uper.Enumerated = 9
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_spare6 uper.Enumerated = 10
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_spare5 uper.Enumerated = 11
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_spare4 uper.Enumerated = 12
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_spare3 uper.Enumerated = 13
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_spare2 uper.Enumerated = 14
	FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17_Enum_spare1 uper.Enumerated = 15
)

type FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17", err)
	}
	return nil
}

func (ie *FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode FeatureCombinationPreambles_r17_groupBconfigured_r17_ra_SizeGroupA_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
