package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SuccessHO_Report_r17 struct {
	sourceCellInfo_r17         *SuccessHO_Report_r17_sourceCellInfo_r17       `optional`
	targetCellInfo_r17         *SuccessHO_Report_r17_targetCellInfo_r17       `optional`
	measResultNeighCells_r17   *SuccessHO_Report_r17_measResultNeighCells_r17 `optional`
	locationInfo_r17           *LocationInfo_r16                              `optional`
	timeSinceCHO_Reconfig_r17  *TimeSinceCHO_Reconfig_r17                     `optional`
	shr_Cause_r17              *SHR_Cause_r17                                 `optional`
	ra_InformationCommon_r17   *RA_InformationCommon_r16                      `optional`
	upInterruptionTimeAtHO_r17 *UPInterruptionTimeAtHO_r17                    `optional`
	c_RNTI_r17                 *RNTI_Value                                    `optional`
}

func (ie *SuccessHO_Report_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sourceCellInfo_r17 != nil, ie.targetCellInfo_r17 != nil, ie.measResultNeighCells_r17 != nil, ie.locationInfo_r17 != nil, ie.timeSinceCHO_Reconfig_r17 != nil, ie.shr_Cause_r17 != nil, ie.ra_InformationCommon_r17 != nil, ie.upInterruptionTimeAtHO_r17 != nil, ie.c_RNTI_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sourceCellInfo_r17 != nil {
		if err = ie.sourceCellInfo_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sourceCellInfo_r17", err)
		}
	}
	if ie.targetCellInfo_r17 != nil {
		if err = ie.targetCellInfo_r17.Encode(w); err != nil {
			return utils.WrapError("Encode targetCellInfo_r17", err)
		}
	}
	if ie.measResultNeighCells_r17 != nil {
		if err = ie.measResultNeighCells_r17.Encode(w); err != nil {
			return utils.WrapError("Encode measResultNeighCells_r17", err)
		}
	}
	if ie.locationInfo_r17 != nil {
		if err = ie.locationInfo_r17.Encode(w); err != nil {
			return utils.WrapError("Encode locationInfo_r17", err)
		}
	}
	if ie.timeSinceCHO_Reconfig_r17 != nil {
		if err = ie.timeSinceCHO_Reconfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode timeSinceCHO_Reconfig_r17", err)
		}
	}
	if ie.shr_Cause_r17 != nil {
		if err = ie.shr_Cause_r17.Encode(w); err != nil {
			return utils.WrapError("Encode shr_Cause_r17", err)
		}
	}
	if ie.ra_InformationCommon_r17 != nil {
		if err = ie.ra_InformationCommon_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ra_InformationCommon_r17", err)
		}
	}
	if ie.upInterruptionTimeAtHO_r17 != nil {
		if err = ie.upInterruptionTimeAtHO_r17.Encode(w); err != nil {
			return utils.WrapError("Encode upInterruptionTimeAtHO_r17", err)
		}
	}
	if ie.c_RNTI_r17 != nil {
		if err = ie.c_RNTI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode c_RNTI_r17", err)
		}
	}
	return nil
}

func (ie *SuccessHO_Report_r17) Decode(r *uper.UperReader) error {
	var err error
	var sourceCellInfo_r17Present bool
	if sourceCellInfo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var targetCellInfo_r17Present bool
	if targetCellInfo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultNeighCells_r17Present bool
	if measResultNeighCells_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var locationInfo_r17Present bool
	if locationInfo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var timeSinceCHO_Reconfig_r17Present bool
	if timeSinceCHO_Reconfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var shr_Cause_r17Present bool
	if shr_Cause_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ra_InformationCommon_r17Present bool
	if ra_InformationCommon_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var upInterruptionTimeAtHO_r17Present bool
	if upInterruptionTimeAtHO_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var c_RNTI_r17Present bool
	if c_RNTI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sourceCellInfo_r17Present {
		ie.sourceCellInfo_r17 = new(SuccessHO_Report_r17_sourceCellInfo_r17)
		if err = ie.sourceCellInfo_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sourceCellInfo_r17", err)
		}
	}
	if targetCellInfo_r17Present {
		ie.targetCellInfo_r17 = new(SuccessHO_Report_r17_targetCellInfo_r17)
		if err = ie.targetCellInfo_r17.Decode(r); err != nil {
			return utils.WrapError("Decode targetCellInfo_r17", err)
		}
	}
	if measResultNeighCells_r17Present {
		ie.measResultNeighCells_r17 = new(SuccessHO_Report_r17_measResultNeighCells_r17)
		if err = ie.measResultNeighCells_r17.Decode(r); err != nil {
			return utils.WrapError("Decode measResultNeighCells_r17", err)
		}
	}
	if locationInfo_r17Present {
		ie.locationInfo_r17 = new(LocationInfo_r16)
		if err = ie.locationInfo_r17.Decode(r); err != nil {
			return utils.WrapError("Decode locationInfo_r17", err)
		}
	}
	if timeSinceCHO_Reconfig_r17Present {
		ie.timeSinceCHO_Reconfig_r17 = new(TimeSinceCHO_Reconfig_r17)
		if err = ie.timeSinceCHO_Reconfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode timeSinceCHO_Reconfig_r17", err)
		}
	}
	if shr_Cause_r17Present {
		ie.shr_Cause_r17 = new(SHR_Cause_r17)
		if err = ie.shr_Cause_r17.Decode(r); err != nil {
			return utils.WrapError("Decode shr_Cause_r17", err)
		}
	}
	if ra_InformationCommon_r17Present {
		ie.ra_InformationCommon_r17 = new(RA_InformationCommon_r16)
		if err = ie.ra_InformationCommon_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ra_InformationCommon_r17", err)
		}
	}
	if upInterruptionTimeAtHO_r17Present {
		ie.upInterruptionTimeAtHO_r17 = new(UPInterruptionTimeAtHO_r17)
		if err = ie.upInterruptionTimeAtHO_r17.Decode(r); err != nil {
			return utils.WrapError("Decode upInterruptionTimeAtHO_r17", err)
		}
	}
	if c_RNTI_r17Present {
		ie.c_RNTI_r17 = new(RNTI_Value)
		if err = ie.c_RNTI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode c_RNTI_r17", err)
		}
	}
	return nil
}
