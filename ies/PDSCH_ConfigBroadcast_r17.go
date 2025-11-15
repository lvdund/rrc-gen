package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_ConfigBroadcast_r17 struct {
	pdschConfigList_r17                []PDSCH_ConfigPTM_r17                       `lb:1,ub:maxNrofPDSCH_ConfigPTM_r17,madatory`
	pdsch_TimeDomainAllocationList_r17 *PDSCH_TimeDomainResourceAllocationList_r16 `optional`
	rateMatchPatternToAddModList_r17   []RateMatchPattern                          `lb:1,ub:maxNrofRateMatchPatterns,optional`
	lte_CRS_ToMatchAround_r17          *RateMatchPatternLTE_CRS                    `optional`
	mcs_Table_r17                      *PDSCH_ConfigBroadcast_r17_mcs_Table_r17    `optional`
	xOverhead_r17                      *PDSCH_ConfigBroadcast_r17_xOverhead_r17    `optional`
}

func (ie *PDSCH_ConfigBroadcast_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pdsch_TimeDomainAllocationList_r17 != nil, len(ie.rateMatchPatternToAddModList_r17) > 0, ie.lte_CRS_ToMatchAround_r17 != nil, ie.mcs_Table_r17 != nil, ie.xOverhead_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_pdschConfigList_r17 := utils.NewSequence[*PDSCH_ConfigPTM_r17]([]*PDSCH_ConfigPTM_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPDSCH_ConfigPTM_r17}, false)
	for _, i := range ie.pdschConfigList_r17 {
		tmp_pdschConfigList_r17.Value = append(tmp_pdschConfigList_r17.Value, &i)
	}
	if err = tmp_pdschConfigList_r17.Encode(w); err != nil {
		return utils.WrapError("Encode pdschConfigList_r17", err)
	}
	if ie.pdsch_TimeDomainAllocationList_r17 != nil {
		if err = ie.pdsch_TimeDomainAllocationList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_TimeDomainAllocationList_r17", err)
		}
	}
	if len(ie.rateMatchPatternToAddModList_r17) > 0 {
		tmp_rateMatchPatternToAddModList_r17 := utils.NewSequence[*RateMatchPattern]([]*RateMatchPattern{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
		for _, i := range ie.rateMatchPatternToAddModList_r17 {
			tmp_rateMatchPatternToAddModList_r17.Value = append(tmp_rateMatchPatternToAddModList_r17.Value, &i)
		}
		if err = tmp_rateMatchPatternToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode rateMatchPatternToAddModList_r17", err)
		}
	}
	if ie.lte_CRS_ToMatchAround_r17 != nil {
		if err = ie.lte_CRS_ToMatchAround_r17.Encode(w); err != nil {
			return utils.WrapError("Encode lte_CRS_ToMatchAround_r17", err)
		}
	}
	if ie.mcs_Table_r17 != nil {
		if err = ie.mcs_Table_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mcs_Table_r17", err)
		}
	}
	if ie.xOverhead_r17 != nil {
		if err = ie.xOverhead_r17.Encode(w); err != nil {
			return utils.WrapError("Encode xOverhead_r17", err)
		}
	}
	return nil
}

func (ie *PDSCH_ConfigBroadcast_r17) Decode(r *uper.UperReader) error {
	var err error
	var pdsch_TimeDomainAllocationList_r17Present bool
	if pdsch_TimeDomainAllocationList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rateMatchPatternToAddModList_r17Present bool
	if rateMatchPatternToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lte_CRS_ToMatchAround_r17Present bool
	if lte_CRS_ToMatchAround_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mcs_Table_r17Present bool
	if mcs_Table_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var xOverhead_r17Present bool
	if xOverhead_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_pdschConfigList_r17 := utils.NewSequence[*PDSCH_ConfigPTM_r17]([]*PDSCH_ConfigPTM_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPDSCH_ConfigPTM_r17}, false)
	fn_pdschConfigList_r17 := func() *PDSCH_ConfigPTM_r17 {
		return new(PDSCH_ConfigPTM_r17)
	}
	if err = tmp_pdschConfigList_r17.Decode(r, fn_pdschConfigList_r17); err != nil {
		return utils.WrapError("Decode pdschConfigList_r17", err)
	}
	ie.pdschConfigList_r17 = []PDSCH_ConfigPTM_r17{}
	for _, i := range tmp_pdschConfigList_r17.Value {
		ie.pdschConfigList_r17 = append(ie.pdschConfigList_r17, *i)
	}
	if pdsch_TimeDomainAllocationList_r17Present {
		ie.pdsch_TimeDomainAllocationList_r17 = new(PDSCH_TimeDomainResourceAllocationList_r16)
		if err = ie.pdsch_TimeDomainAllocationList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_TimeDomainAllocationList_r17", err)
		}
	}
	if rateMatchPatternToAddModList_r17Present {
		tmp_rateMatchPatternToAddModList_r17 := utils.NewSequence[*RateMatchPattern]([]*RateMatchPattern{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
		fn_rateMatchPatternToAddModList_r17 := func() *RateMatchPattern {
			return new(RateMatchPattern)
		}
		if err = tmp_rateMatchPatternToAddModList_r17.Decode(r, fn_rateMatchPatternToAddModList_r17); err != nil {
			return utils.WrapError("Decode rateMatchPatternToAddModList_r17", err)
		}
		ie.rateMatchPatternToAddModList_r17 = []RateMatchPattern{}
		for _, i := range tmp_rateMatchPatternToAddModList_r17.Value {
			ie.rateMatchPatternToAddModList_r17 = append(ie.rateMatchPatternToAddModList_r17, *i)
		}
	}
	if lte_CRS_ToMatchAround_r17Present {
		ie.lte_CRS_ToMatchAround_r17 = new(RateMatchPatternLTE_CRS)
		if err = ie.lte_CRS_ToMatchAround_r17.Decode(r); err != nil {
			return utils.WrapError("Decode lte_CRS_ToMatchAround_r17", err)
		}
	}
	if mcs_Table_r17Present {
		ie.mcs_Table_r17 = new(PDSCH_ConfigBroadcast_r17_mcs_Table_r17)
		if err = ie.mcs_Table_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mcs_Table_r17", err)
		}
	}
	if xOverhead_r17Present {
		ie.xOverhead_r17 = new(PDSCH_ConfigBroadcast_r17_xOverhead_r17)
		if err = ie.xOverhead_r17.Decode(r); err != nil {
			return utils.WrapError("Decode xOverhead_r17", err)
		}
	}
	return nil
}
