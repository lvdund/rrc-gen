package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_PosRRC_InactiveConfig_r17 struct {
	srs_PosConfigNUL_r17                    *SRS_PosConfig_r17        `optional`
	srs_PosConfigSUL_r17                    *SRS_PosConfig_r17        `optional`
	bwp_NUL_r17                             *BWP                      `optional`
	bwp_SUL_r17                             *BWP                      `optional`
	inactivePosSRS_TimeAlignmentTimer_r17   *TimeAlignmentTimer       `optional`
	inactivePosSRS_RSRP_ChangeThreshold_r17 *RSRP_ChangeThreshold_r17 `optional`
}

func (ie *SRS_PosRRC_InactiveConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.srs_PosConfigNUL_r17 != nil, ie.srs_PosConfigSUL_r17 != nil, ie.bwp_NUL_r17 != nil, ie.bwp_SUL_r17 != nil, ie.inactivePosSRS_TimeAlignmentTimer_r17 != nil, ie.inactivePosSRS_RSRP_ChangeThreshold_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.srs_PosConfigNUL_r17 != nil {
		if err = ie.srs_PosConfigNUL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode srs_PosConfigNUL_r17", err)
		}
	}
	if ie.srs_PosConfigSUL_r17 != nil {
		if err = ie.srs_PosConfigSUL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode srs_PosConfigSUL_r17", err)
		}
	}
	if ie.bwp_NUL_r17 != nil {
		if err = ie.bwp_NUL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode bwp_NUL_r17", err)
		}
	}
	if ie.bwp_SUL_r17 != nil {
		if err = ie.bwp_SUL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode bwp_SUL_r17", err)
		}
	}
	if ie.inactivePosSRS_TimeAlignmentTimer_r17 != nil {
		if err = ie.inactivePosSRS_TimeAlignmentTimer_r17.Encode(w); err != nil {
			return utils.WrapError("Encode inactivePosSRS_TimeAlignmentTimer_r17", err)
		}
	}
	if ie.inactivePosSRS_RSRP_ChangeThreshold_r17 != nil {
		if err = ie.inactivePosSRS_RSRP_ChangeThreshold_r17.Encode(w); err != nil {
			return utils.WrapError("Encode inactivePosSRS_RSRP_ChangeThreshold_r17", err)
		}
	}
	return nil
}

func (ie *SRS_PosRRC_InactiveConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var srs_PosConfigNUL_r17Present bool
	if srs_PosConfigNUL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var srs_PosConfigSUL_r17Present bool
	if srs_PosConfigSUL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bwp_NUL_r17Present bool
	if bwp_NUL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bwp_SUL_r17Present bool
	if bwp_SUL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var inactivePosSRS_TimeAlignmentTimer_r17Present bool
	if inactivePosSRS_TimeAlignmentTimer_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var inactivePosSRS_RSRP_ChangeThreshold_r17Present bool
	if inactivePosSRS_RSRP_ChangeThreshold_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if srs_PosConfigNUL_r17Present {
		ie.srs_PosConfigNUL_r17 = new(SRS_PosConfig_r17)
		if err = ie.srs_PosConfigNUL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode srs_PosConfigNUL_r17", err)
		}
	}
	if srs_PosConfigSUL_r17Present {
		ie.srs_PosConfigSUL_r17 = new(SRS_PosConfig_r17)
		if err = ie.srs_PosConfigSUL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode srs_PosConfigSUL_r17", err)
		}
	}
	if bwp_NUL_r17Present {
		ie.bwp_NUL_r17 = new(BWP)
		if err = ie.bwp_NUL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode bwp_NUL_r17", err)
		}
	}
	if bwp_SUL_r17Present {
		ie.bwp_SUL_r17 = new(BWP)
		if err = ie.bwp_SUL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode bwp_SUL_r17", err)
		}
	}
	if inactivePosSRS_TimeAlignmentTimer_r17Present {
		ie.inactivePosSRS_TimeAlignmentTimer_r17 = new(TimeAlignmentTimer)
		if err = ie.inactivePosSRS_TimeAlignmentTimer_r17.Decode(r); err != nil {
			return utils.WrapError("Decode inactivePosSRS_TimeAlignmentTimer_r17", err)
		}
	}
	if inactivePosSRS_RSRP_ChangeThreshold_r17Present {
		ie.inactivePosSRS_RSRP_ChangeThreshold_r17 = new(RSRP_ChangeThreshold_r17)
		if err = ie.inactivePosSRS_RSRP_ChangeThreshold_r17.Decode(r); err != nil {
			return utils.WrapError("Decode inactivePosSRS_RSRP_ChangeThreshold_r17", err)
		}
	}
	return nil
}
