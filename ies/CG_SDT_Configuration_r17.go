package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_SDT_Configuration_r17 struct {
	cg_SDT_RetransmissionTimer *int64                                            `lb:1,ub:64,optional`
	sdt_SSB_Subset_r17         *CG_SDT_Configuration_r17_sdt_SSB_Subset_r17      `lb:4,ub:4,optional`
	sdt_SSB_PerCG_PUSCH_r17    *CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17 `optional`
	sdt_P0_PUSCH_r17           *int64                                            `lb:-16,ub:15,optional`
	sdt_Alpha_r17              *CG_SDT_Configuration_r17_sdt_Alpha_r17           `optional`
	sdt_DMRS_Ports_r17         *CG_SDT_Configuration_r17_sdt_DMRS_Ports_r17      `lb:8,ub:8,optional`
	sdt_NrofDMRS_Sequences_r17 *int64                                            `lb:1,ub:2,optional`
}

func (ie *CG_SDT_Configuration_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.cg_SDT_RetransmissionTimer != nil, ie.sdt_SSB_Subset_r17 != nil, ie.sdt_SSB_PerCG_PUSCH_r17 != nil, ie.sdt_P0_PUSCH_r17 != nil, ie.sdt_Alpha_r17 != nil, ie.sdt_DMRS_Ports_r17 != nil, ie.sdt_NrofDMRS_Sequences_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.cg_SDT_RetransmissionTimer != nil {
		if err = w.WriteInteger(*ie.cg_SDT_RetransmissionTimer, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
			return utils.WrapError("Encode cg_SDT_RetransmissionTimer", err)
		}
	}
	if ie.sdt_SSB_Subset_r17 != nil {
		if err = ie.sdt_SSB_Subset_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sdt_SSB_Subset_r17", err)
		}
	}
	if ie.sdt_SSB_PerCG_PUSCH_r17 != nil {
		if err = ie.sdt_SSB_PerCG_PUSCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sdt_SSB_PerCG_PUSCH_r17", err)
		}
	}
	if ie.sdt_P0_PUSCH_r17 != nil {
		if err = w.WriteInteger(*ie.sdt_P0_PUSCH_r17, &uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode sdt_P0_PUSCH_r17", err)
		}
	}
	if ie.sdt_Alpha_r17 != nil {
		if err = ie.sdt_Alpha_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sdt_Alpha_r17", err)
		}
	}
	if ie.sdt_DMRS_Ports_r17 != nil {
		if err = ie.sdt_DMRS_Ports_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sdt_DMRS_Ports_r17", err)
		}
	}
	if ie.sdt_NrofDMRS_Sequences_r17 != nil {
		if err = w.WriteInteger(*ie.sdt_NrofDMRS_Sequences_r17, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode sdt_NrofDMRS_Sequences_r17", err)
		}
	}
	return nil
}

func (ie *CG_SDT_Configuration_r17) Decode(r *uper.UperReader) error {
	var err error
	var cg_SDT_RetransmissionTimerPresent bool
	if cg_SDT_RetransmissionTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sdt_SSB_Subset_r17Present bool
	if sdt_SSB_Subset_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sdt_SSB_PerCG_PUSCH_r17Present bool
	if sdt_SSB_PerCG_PUSCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sdt_P0_PUSCH_r17Present bool
	if sdt_P0_PUSCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sdt_Alpha_r17Present bool
	if sdt_Alpha_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sdt_DMRS_Ports_r17Present bool
	if sdt_DMRS_Ports_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sdt_NrofDMRS_Sequences_r17Present bool
	if sdt_NrofDMRS_Sequences_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if cg_SDT_RetransmissionTimerPresent {
		var tmp_int_cg_SDT_RetransmissionTimer int64
		if tmp_int_cg_SDT_RetransmissionTimer, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode cg_SDT_RetransmissionTimer", err)
		}
		ie.cg_SDT_RetransmissionTimer = &tmp_int_cg_SDT_RetransmissionTimer
	}
	if sdt_SSB_Subset_r17Present {
		ie.sdt_SSB_Subset_r17 = new(CG_SDT_Configuration_r17_sdt_SSB_Subset_r17)
		if err = ie.sdt_SSB_Subset_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sdt_SSB_Subset_r17", err)
		}
	}
	if sdt_SSB_PerCG_PUSCH_r17Present {
		ie.sdt_SSB_PerCG_PUSCH_r17 = new(CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17)
		if err = ie.sdt_SSB_PerCG_PUSCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sdt_SSB_PerCG_PUSCH_r17", err)
		}
	}
	if sdt_P0_PUSCH_r17Present {
		var tmp_int_sdt_P0_PUSCH_r17 int64
		if tmp_int_sdt_P0_PUSCH_r17, err = r.ReadInteger(&uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode sdt_P0_PUSCH_r17", err)
		}
		ie.sdt_P0_PUSCH_r17 = &tmp_int_sdt_P0_PUSCH_r17
	}
	if sdt_Alpha_r17Present {
		ie.sdt_Alpha_r17 = new(CG_SDT_Configuration_r17_sdt_Alpha_r17)
		if err = ie.sdt_Alpha_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sdt_Alpha_r17", err)
		}
	}
	if sdt_DMRS_Ports_r17Present {
		ie.sdt_DMRS_Ports_r17 = new(CG_SDT_Configuration_r17_sdt_DMRS_Ports_r17)
		if err = ie.sdt_DMRS_Ports_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sdt_DMRS_Ports_r17", err)
		}
	}
	if sdt_NrofDMRS_Sequences_r17Present {
		var tmp_int_sdt_NrofDMRS_Sequences_r17 int64
		if tmp_int_sdt_NrofDMRS_Sequences_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode sdt_NrofDMRS_Sequences_r17", err)
		}
		ie.sdt_NrofDMRS_Sequences_r17 = &tmp_int_sdt_NrofDMRS_Sequences_r17
	}
	return nil
}
