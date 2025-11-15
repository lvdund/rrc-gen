package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NTN_Config_r17 struct {
	epochTime_r17                  *EpochTime_r17                                 `optional`
	ntn_UlSyncValidityDuration_r17 *NTN_Config_r17_ntn_UlSyncValidityDuration_r17 `optional`
	cellSpecificKoffset_r17        *int64                                         `lb:1,ub:1023,optional`
	kmac_r17                       *int64                                         `lb:1,ub:512,optional`
	ta_Info_r17                    *TA_Info_r17                                   `optional`
	ntn_PolarizationDL_r17         *NTN_Config_r17_ntn_PolarizationDL_r17         `optional`
	ntn_PolarizationUL_r17         *NTN_Config_r17_ntn_PolarizationUL_r17         `optional`
	ephemerisInfo_r17              *EphemerisInfo_r17                             `optional`
	ta_Report_r17                  *NTN_Config_r17_ta_Report_r17                  `optional`
}

func (ie *NTN_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.epochTime_r17 != nil, ie.ntn_UlSyncValidityDuration_r17 != nil, ie.cellSpecificKoffset_r17 != nil, ie.kmac_r17 != nil, ie.ta_Info_r17 != nil, ie.ntn_PolarizationDL_r17 != nil, ie.ntn_PolarizationUL_r17 != nil, ie.ephemerisInfo_r17 != nil, ie.ta_Report_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.epochTime_r17 != nil {
		if err = ie.epochTime_r17.Encode(w); err != nil {
			return utils.WrapError("Encode epochTime_r17", err)
		}
	}
	if ie.ntn_UlSyncValidityDuration_r17 != nil {
		if err = ie.ntn_UlSyncValidityDuration_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ntn_UlSyncValidityDuration_r17", err)
		}
	}
	if ie.cellSpecificKoffset_r17 != nil {
		if err = w.WriteInteger(*ie.cellSpecificKoffset_r17, &uper.Constraint{Lb: 1, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode cellSpecificKoffset_r17", err)
		}
	}
	if ie.kmac_r17 != nil {
		if err = w.WriteInteger(*ie.kmac_r17, &uper.Constraint{Lb: 1, Ub: 512}, false); err != nil {
			return utils.WrapError("Encode kmac_r17", err)
		}
	}
	if ie.ta_Info_r17 != nil {
		if err = ie.ta_Info_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ta_Info_r17", err)
		}
	}
	if ie.ntn_PolarizationDL_r17 != nil {
		if err = ie.ntn_PolarizationDL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ntn_PolarizationDL_r17", err)
		}
	}
	if ie.ntn_PolarizationUL_r17 != nil {
		if err = ie.ntn_PolarizationUL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ntn_PolarizationUL_r17", err)
		}
	}
	if ie.ephemerisInfo_r17 != nil {
		if err = ie.ephemerisInfo_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ephemerisInfo_r17", err)
		}
	}
	if ie.ta_Report_r17 != nil {
		if err = ie.ta_Report_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ta_Report_r17", err)
		}
	}
	return nil
}

func (ie *NTN_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var epochTime_r17Present bool
	if epochTime_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ntn_UlSyncValidityDuration_r17Present bool
	if ntn_UlSyncValidityDuration_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cellSpecificKoffset_r17Present bool
	if cellSpecificKoffset_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var kmac_r17Present bool
	if kmac_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ta_Info_r17Present bool
	if ta_Info_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ntn_PolarizationDL_r17Present bool
	if ntn_PolarizationDL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ntn_PolarizationUL_r17Present bool
	if ntn_PolarizationUL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ephemerisInfo_r17Present bool
	if ephemerisInfo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ta_Report_r17Present bool
	if ta_Report_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if epochTime_r17Present {
		ie.epochTime_r17 = new(EpochTime_r17)
		if err = ie.epochTime_r17.Decode(r); err != nil {
			return utils.WrapError("Decode epochTime_r17", err)
		}
	}
	if ntn_UlSyncValidityDuration_r17Present {
		ie.ntn_UlSyncValidityDuration_r17 = new(NTN_Config_r17_ntn_UlSyncValidityDuration_r17)
		if err = ie.ntn_UlSyncValidityDuration_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ntn_UlSyncValidityDuration_r17", err)
		}
	}
	if cellSpecificKoffset_r17Present {
		var tmp_int_cellSpecificKoffset_r17 int64
		if tmp_int_cellSpecificKoffset_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode cellSpecificKoffset_r17", err)
		}
		ie.cellSpecificKoffset_r17 = &tmp_int_cellSpecificKoffset_r17
	}
	if kmac_r17Present {
		var tmp_int_kmac_r17 int64
		if tmp_int_kmac_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 512}, false); err != nil {
			return utils.WrapError("Decode kmac_r17", err)
		}
		ie.kmac_r17 = &tmp_int_kmac_r17
	}
	if ta_Info_r17Present {
		ie.ta_Info_r17 = new(TA_Info_r17)
		if err = ie.ta_Info_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ta_Info_r17", err)
		}
	}
	if ntn_PolarizationDL_r17Present {
		ie.ntn_PolarizationDL_r17 = new(NTN_Config_r17_ntn_PolarizationDL_r17)
		if err = ie.ntn_PolarizationDL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ntn_PolarizationDL_r17", err)
		}
	}
	if ntn_PolarizationUL_r17Present {
		ie.ntn_PolarizationUL_r17 = new(NTN_Config_r17_ntn_PolarizationUL_r17)
		if err = ie.ntn_PolarizationUL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ntn_PolarizationUL_r17", err)
		}
	}
	if ephemerisInfo_r17Present {
		ie.ephemerisInfo_r17 = new(EphemerisInfo_r17)
		if err = ie.ephemerisInfo_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ephemerisInfo_r17", err)
		}
	}
	if ta_Report_r17Present {
		ie.ta_Report_r17 = new(NTN_Config_r17_ta_Report_r17)
		if err = ie.ta_Report_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ta_Report_r17", err)
		}
	}
	return nil
}
