package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DMRS_BundlingPUCCH_Config_r17 struct {
	pucch_DMRS_Bundling_r17            *DMRS_BundlingPUCCH_Config_r17_pucch_DMRS_Bundling_r17            `optional`
	pucch_TimeDomainWindowLength_r17   *int64                                                            `lb:2,ub:8,optional`
	pucch_WindowRestart_r17            *DMRS_BundlingPUCCH_Config_r17_pucch_WindowRestart_r17            `optional`
	pucch_FrequencyHoppingInterval_r17 *DMRS_BundlingPUCCH_Config_r17_pucch_FrequencyHoppingInterval_r17 `optional`
}

func (ie *DMRS_BundlingPUCCH_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pucch_DMRS_Bundling_r17 != nil, ie.pucch_TimeDomainWindowLength_r17 != nil, ie.pucch_WindowRestart_r17 != nil, ie.pucch_FrequencyHoppingInterval_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pucch_DMRS_Bundling_r17 != nil {
		if err = ie.pucch_DMRS_Bundling_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_DMRS_Bundling_r17", err)
		}
	}
	if ie.pucch_TimeDomainWindowLength_r17 != nil {
		if err = w.WriteInteger(*ie.pucch_TimeDomainWindowLength_r17, &uper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode pucch_TimeDomainWindowLength_r17", err)
		}
	}
	if ie.pucch_WindowRestart_r17 != nil {
		if err = ie.pucch_WindowRestart_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_WindowRestart_r17", err)
		}
	}
	if ie.pucch_FrequencyHoppingInterval_r17 != nil {
		if err = ie.pucch_FrequencyHoppingInterval_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_FrequencyHoppingInterval_r17", err)
		}
	}
	return nil
}

func (ie *DMRS_BundlingPUCCH_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var pucch_DMRS_Bundling_r17Present bool
	if pucch_DMRS_Bundling_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_TimeDomainWindowLength_r17Present bool
	if pucch_TimeDomainWindowLength_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_WindowRestart_r17Present bool
	if pucch_WindowRestart_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_FrequencyHoppingInterval_r17Present bool
	if pucch_FrequencyHoppingInterval_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if pucch_DMRS_Bundling_r17Present {
		ie.pucch_DMRS_Bundling_r17 = new(DMRS_BundlingPUCCH_Config_r17_pucch_DMRS_Bundling_r17)
		if err = ie.pucch_DMRS_Bundling_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_DMRS_Bundling_r17", err)
		}
	}
	if pucch_TimeDomainWindowLength_r17Present {
		var tmp_int_pucch_TimeDomainWindowLength_r17 int64
		if tmp_int_pucch_TimeDomainWindowLength_r17, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode pucch_TimeDomainWindowLength_r17", err)
		}
		ie.pucch_TimeDomainWindowLength_r17 = &tmp_int_pucch_TimeDomainWindowLength_r17
	}
	if pucch_WindowRestart_r17Present {
		ie.pucch_WindowRestart_r17 = new(DMRS_BundlingPUCCH_Config_r17_pucch_WindowRestart_r17)
		if err = ie.pucch_WindowRestart_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_WindowRestart_r17", err)
		}
	}
	if pucch_FrequencyHoppingInterval_r17Present {
		ie.pucch_FrequencyHoppingInterval_r17 = new(DMRS_BundlingPUCCH_Config_r17_pucch_FrequencyHoppingInterval_r17)
		if err = ie.pucch_FrequencyHoppingInterval_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_FrequencyHoppingInterval_r17", err)
		}
	}
	return nil
}
