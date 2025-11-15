package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DMRS_BundlingPUSCH_Config_r17 struct {
	pusch_DMRS_Bundling_r17            *DMRS_BundlingPUSCH_Config_r17_pusch_DMRS_Bundling_r17            `optional`
	pusch_TimeDomainWindowLength_r17   *int64                                                            `lb:2,ub:32,optional`
	pusch_WindowRestart_r17            *DMRS_BundlingPUSCH_Config_r17_pusch_WindowRestart_r17            `optional`
	pusch_FrequencyHoppingInterval_r17 *DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17 `optional`
}

func (ie *DMRS_BundlingPUSCH_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pusch_DMRS_Bundling_r17 != nil, ie.pusch_TimeDomainWindowLength_r17 != nil, ie.pusch_WindowRestart_r17 != nil, ie.pusch_FrequencyHoppingInterval_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pusch_DMRS_Bundling_r17 != nil {
		if err = ie.pusch_DMRS_Bundling_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_DMRS_Bundling_r17", err)
		}
	}
	if ie.pusch_TimeDomainWindowLength_r17 != nil {
		if err = w.WriteInteger(*ie.pusch_TimeDomainWindowLength_r17, &uper.Constraint{Lb: 2, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode pusch_TimeDomainWindowLength_r17", err)
		}
	}
	if ie.pusch_WindowRestart_r17 != nil {
		if err = ie.pusch_WindowRestart_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_WindowRestart_r17", err)
		}
	}
	if ie.pusch_FrequencyHoppingInterval_r17 != nil {
		if err = ie.pusch_FrequencyHoppingInterval_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_FrequencyHoppingInterval_r17", err)
		}
	}
	return nil
}

func (ie *DMRS_BundlingPUSCH_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var pusch_DMRS_Bundling_r17Present bool
	if pusch_DMRS_Bundling_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_TimeDomainWindowLength_r17Present bool
	if pusch_TimeDomainWindowLength_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_WindowRestart_r17Present bool
	if pusch_WindowRestart_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_FrequencyHoppingInterval_r17Present bool
	if pusch_FrequencyHoppingInterval_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if pusch_DMRS_Bundling_r17Present {
		ie.pusch_DMRS_Bundling_r17 = new(DMRS_BundlingPUSCH_Config_r17_pusch_DMRS_Bundling_r17)
		if err = ie.pusch_DMRS_Bundling_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_DMRS_Bundling_r17", err)
		}
	}
	if pusch_TimeDomainWindowLength_r17Present {
		var tmp_int_pusch_TimeDomainWindowLength_r17 int64
		if tmp_int_pusch_TimeDomainWindowLength_r17, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode pusch_TimeDomainWindowLength_r17", err)
		}
		ie.pusch_TimeDomainWindowLength_r17 = &tmp_int_pusch_TimeDomainWindowLength_r17
	}
	if pusch_WindowRestart_r17Present {
		ie.pusch_WindowRestart_r17 = new(DMRS_BundlingPUSCH_Config_r17_pusch_WindowRestart_r17)
		if err = ie.pusch_WindowRestart_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_WindowRestart_r17", err)
		}
	}
	if pusch_FrequencyHoppingInterval_r17Present {
		ie.pusch_FrequencyHoppingInterval_r17 = new(DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17)
		if err = ie.pusch_FrequencyHoppingInterval_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_FrequencyHoppingInterval_r17", err)
		}
	}
	return nil
}
