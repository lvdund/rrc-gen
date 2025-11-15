package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ChannelAccessConfig_r16 struct {
	energyDetectionConfig_r16           *ChannelAccessConfig_r16_energyDetectionConfig_r16       `lb:-85,ub:-52,optional`
	ul_toDL_COT_SharingED_Threshold_r16 *int64                                                   `lb:-85,ub:-52,optional`
	absenceOfAnyOtherTechnology_r16     *ChannelAccessConfig_r16_absenceOfAnyOtherTechnology_r16 `optional`
}

func (ie *ChannelAccessConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.energyDetectionConfig_r16 != nil, ie.ul_toDL_COT_SharingED_Threshold_r16 != nil, ie.absenceOfAnyOtherTechnology_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.energyDetectionConfig_r16 != nil {
		if err = ie.energyDetectionConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode energyDetectionConfig_r16", err)
		}
	}
	if ie.ul_toDL_COT_SharingED_Threshold_r16 != nil {
		if err = w.WriteInteger(*ie.ul_toDL_COT_SharingED_Threshold_r16, &uper.Constraint{Lb: -85, Ub: -52}, false); err != nil {
			return utils.WrapError("Encode ul_toDL_COT_SharingED_Threshold_r16", err)
		}
	}
	if ie.absenceOfAnyOtherTechnology_r16 != nil {
		if err = ie.absenceOfAnyOtherTechnology_r16.Encode(w); err != nil {
			return utils.WrapError("Encode absenceOfAnyOtherTechnology_r16", err)
		}
	}
	return nil
}

func (ie *ChannelAccessConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var energyDetectionConfig_r16Present bool
	if energyDetectionConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_toDL_COT_SharingED_Threshold_r16Present bool
	if ul_toDL_COT_SharingED_Threshold_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var absenceOfAnyOtherTechnology_r16Present bool
	if absenceOfAnyOtherTechnology_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if energyDetectionConfig_r16Present {
		ie.energyDetectionConfig_r16 = new(ChannelAccessConfig_r16_energyDetectionConfig_r16)
		if err = ie.energyDetectionConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode energyDetectionConfig_r16", err)
		}
	}
	if ul_toDL_COT_SharingED_Threshold_r16Present {
		var tmp_int_ul_toDL_COT_SharingED_Threshold_r16 int64
		if tmp_int_ul_toDL_COT_SharingED_Threshold_r16, err = r.ReadInteger(&uper.Constraint{Lb: -85, Ub: -52}, false); err != nil {
			return utils.WrapError("Decode ul_toDL_COT_SharingED_Threshold_r16", err)
		}
		ie.ul_toDL_COT_SharingED_Threshold_r16 = &tmp_int_ul_toDL_COT_SharingED_Threshold_r16
	}
	if absenceOfAnyOtherTechnology_r16Present {
		ie.absenceOfAnyOtherTechnology_r16 = new(ChannelAccessConfig_r16_absenceOfAnyOtherTechnology_r16)
		if err = ie.absenceOfAnyOtherTechnology_r16.Decode(r); err != nil {
			return utils.WrapError("Decode absenceOfAnyOtherTechnology_r16", err)
		}
	}
	return nil
}
