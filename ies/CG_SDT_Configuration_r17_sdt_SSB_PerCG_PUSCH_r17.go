package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17_Enum_oneEighth uper.Enumerated = 0
	CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17_Enum_oneFourth uper.Enumerated = 1
	CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17_Enum_half      uper.Enumerated = 2
	CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17_Enum_one       uper.Enumerated = 3
	CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17_Enum_two       uper.Enumerated = 4
	CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17_Enum_four      uper.Enumerated = 5
	CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17_Enum_eight     uper.Enumerated = 6
	CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17_Enum_sixteen   uper.Enumerated = 7
)

type CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17", err)
	}
	return nil
}

func (ie *CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
