package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SuccessHO_Config_r17 struct {
	thresholdPercentageT304_r17     *SuccessHO_Config_r17_thresholdPercentageT304_r17     `optional`
	thresholdPercentageT310_r17     *SuccessHO_Config_r17_thresholdPercentageT310_r17     `optional`
	thresholdPercentageT312_r17     *SuccessHO_Config_r17_thresholdPercentageT312_r17     `optional`
	sourceDAPS_FailureReporting_r17 *SuccessHO_Config_r17_sourceDAPS_FailureReporting_r17 `optional`
}

func (ie *SuccessHO_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.thresholdPercentageT304_r17 != nil, ie.thresholdPercentageT310_r17 != nil, ie.thresholdPercentageT312_r17 != nil, ie.sourceDAPS_FailureReporting_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.thresholdPercentageT304_r17 != nil {
		if err = ie.thresholdPercentageT304_r17.Encode(w); err != nil {
			return utils.WrapError("Encode thresholdPercentageT304_r17", err)
		}
	}
	if ie.thresholdPercentageT310_r17 != nil {
		if err = ie.thresholdPercentageT310_r17.Encode(w); err != nil {
			return utils.WrapError("Encode thresholdPercentageT310_r17", err)
		}
	}
	if ie.thresholdPercentageT312_r17 != nil {
		if err = ie.thresholdPercentageT312_r17.Encode(w); err != nil {
			return utils.WrapError("Encode thresholdPercentageT312_r17", err)
		}
	}
	if ie.sourceDAPS_FailureReporting_r17 != nil {
		if err = ie.sourceDAPS_FailureReporting_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sourceDAPS_FailureReporting_r17", err)
		}
	}
	return nil
}

func (ie *SuccessHO_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var thresholdPercentageT304_r17Present bool
	if thresholdPercentageT304_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var thresholdPercentageT310_r17Present bool
	if thresholdPercentageT310_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var thresholdPercentageT312_r17Present bool
	if thresholdPercentageT312_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sourceDAPS_FailureReporting_r17Present bool
	if sourceDAPS_FailureReporting_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if thresholdPercentageT304_r17Present {
		ie.thresholdPercentageT304_r17 = new(SuccessHO_Config_r17_thresholdPercentageT304_r17)
		if err = ie.thresholdPercentageT304_r17.Decode(r); err != nil {
			return utils.WrapError("Decode thresholdPercentageT304_r17", err)
		}
	}
	if thresholdPercentageT310_r17Present {
		ie.thresholdPercentageT310_r17 = new(SuccessHO_Config_r17_thresholdPercentageT310_r17)
		if err = ie.thresholdPercentageT310_r17.Decode(r); err != nil {
			return utils.WrapError("Decode thresholdPercentageT310_r17", err)
		}
	}
	if thresholdPercentageT312_r17Present {
		ie.thresholdPercentageT312_r17 = new(SuccessHO_Config_r17_thresholdPercentageT312_r17)
		if err = ie.thresholdPercentageT312_r17.Decode(r); err != nil {
			return utils.WrapError("Decode thresholdPercentageT312_r17", err)
		}
	}
	if sourceDAPS_FailureReporting_r17Present {
		ie.sourceDAPS_FailureReporting_r17 = new(SuccessHO_Config_r17_sourceDAPS_FailureReporting_r17)
		if err = ie.sourceDAPS_FailureReporting_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sourceDAPS_FailureReporting_r17", err)
		}
	}
	return nil
}
