package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n50   uper.Enumerated = 0
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n75   uper.Enumerated = 1
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n100  uper.Enumerated = 2
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n125  uper.Enumerated = 3
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n150  uper.Enumerated = 4
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n175  uper.Enumerated = 5
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n200  uper.Enumerated = 6
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n225  uper.Enumerated = 7
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n250  uper.Enumerated = 8
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n275  uper.Enumerated = 9
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n300  uper.Enumerated = 10
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n350  uper.Enumerated = 11
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n400  uper.Enumerated = 12
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n450  uper.Enumerated = 13
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_n500  uper.Enumerated = 14
	NAICS_Capability_Entry_numberOfAggregatedPRB_Enum_spare uper.Enumerated = 15
)

type NAICS_Capability_Entry_numberOfAggregatedPRB struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *NAICS_Capability_Entry_numberOfAggregatedPRB) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode NAICS_Capability_Entry_numberOfAggregatedPRB", err)
	}
	return nil
}

func (ie *NAICS_Capability_Entry_numberOfAggregatedPRB) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode NAICS_Capability_Entry_numberOfAggregatedPRB", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
