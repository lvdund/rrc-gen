package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PBPS_CPS_Config_r17 struct {
	sl_AllowedResourceSelectionConfig_r17 *SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17 `optional`
	sl_MinNumCandidateSlotsPeriodic_r17   *int64                                                        `lb:1,ub:32,optional`
	sl_PBPS_OccasionReservePeriodList_r17 []int64                                                       `lb:1,ub:16,e_lb:1,e_ub:16,optional`
	sl_Additional_PBPS_Occasion_r17       *SL_PBPS_CPS_Config_r17_sl_Additional_PBPS_Occasion_r17       `optional`
	sl_CPS_WindowPeriodic_r17             *int64                                                        `lb:5,ub:30,optional`
	sl_MinNumCandidateSlotsAperiodic_r17  *int64                                                        `lb:1,ub:32,optional`
	sl_MinNumRssiMeasurementSlots_r17     *int64                                                        `lb:1,ub:800,optional`
	sl_DefaultCBR_RandomSelection_r17     *int64                                                        `lb:0,ub:100,optional`
	sl_DefaultCBR_PartialSensing_r17      *int64                                                        `lb:0,ub:100,optional`
	sl_CPS_WindowAperiodic_r17            *int64                                                        `lb:0,ub:30,optional`
	sl_PartialSensingInactiveTime_r17     *SL_PBPS_CPS_Config_r17_sl_PartialSensingInactiveTime_r17     `optional`
}

func (ie *SL_PBPS_CPS_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_AllowedResourceSelectionConfig_r17 != nil, ie.sl_MinNumCandidateSlotsPeriodic_r17 != nil, len(ie.sl_PBPS_OccasionReservePeriodList_r17) > 0, ie.sl_Additional_PBPS_Occasion_r17 != nil, ie.sl_CPS_WindowPeriodic_r17 != nil, ie.sl_MinNumCandidateSlotsAperiodic_r17 != nil, ie.sl_MinNumRssiMeasurementSlots_r17 != nil, ie.sl_DefaultCBR_RandomSelection_r17 != nil, ie.sl_DefaultCBR_PartialSensing_r17 != nil, ie.sl_CPS_WindowAperiodic_r17 != nil, ie.sl_PartialSensingInactiveTime_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_AllowedResourceSelectionConfig_r17 != nil {
		if err = ie.sl_AllowedResourceSelectionConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_AllowedResourceSelectionConfig_r17", err)
		}
	}
	if ie.sl_MinNumCandidateSlotsPeriodic_r17 != nil {
		if err = w.WriteInteger(*ie.sl_MinNumCandidateSlotsPeriodic_r17, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode sl_MinNumCandidateSlotsPeriodic_r17", err)
		}
	}
	if len(ie.sl_PBPS_OccasionReservePeriodList_r17) > 0 {
		tmp_sl_PBPS_OccasionReservePeriodList_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 16}, false)
		for _, i := range ie.sl_PBPS_OccasionReservePeriodList_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 1, Ub: 16}, false)
			tmp_sl_PBPS_OccasionReservePeriodList_r17.Value = append(tmp_sl_PBPS_OccasionReservePeriodList_r17.Value, &tmp_ie)
		}
		if err = tmp_sl_PBPS_OccasionReservePeriodList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PBPS_OccasionReservePeriodList_r17", err)
		}
	}
	if ie.sl_Additional_PBPS_Occasion_r17 != nil {
		if err = ie.sl_Additional_PBPS_Occasion_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_Additional_PBPS_Occasion_r17", err)
		}
	}
	if ie.sl_CPS_WindowPeriodic_r17 != nil {
		if err = w.WriteInteger(*ie.sl_CPS_WindowPeriodic_r17, &uper.Constraint{Lb: 5, Ub: 30}, false); err != nil {
			return utils.WrapError("Encode sl_CPS_WindowPeriodic_r17", err)
		}
	}
	if ie.sl_MinNumCandidateSlotsAperiodic_r17 != nil {
		if err = w.WriteInteger(*ie.sl_MinNumCandidateSlotsAperiodic_r17, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode sl_MinNumCandidateSlotsAperiodic_r17", err)
		}
	}
	if ie.sl_MinNumRssiMeasurementSlots_r17 != nil {
		if err = w.WriteInteger(*ie.sl_MinNumRssiMeasurementSlots_r17, &uper.Constraint{Lb: 1, Ub: 800}, false); err != nil {
			return utils.WrapError("Encode sl_MinNumRssiMeasurementSlots_r17", err)
		}
	}
	if ie.sl_DefaultCBR_RandomSelection_r17 != nil {
		if err = w.WriteInteger(*ie.sl_DefaultCBR_RandomSelection_r17, &uper.Constraint{Lb: 0, Ub: 100}, false); err != nil {
			return utils.WrapError("Encode sl_DefaultCBR_RandomSelection_r17", err)
		}
	}
	if ie.sl_DefaultCBR_PartialSensing_r17 != nil {
		if err = w.WriteInteger(*ie.sl_DefaultCBR_PartialSensing_r17, &uper.Constraint{Lb: 0, Ub: 100}, false); err != nil {
			return utils.WrapError("Encode sl_DefaultCBR_PartialSensing_r17", err)
		}
	}
	if ie.sl_CPS_WindowAperiodic_r17 != nil {
		if err = w.WriteInteger(*ie.sl_CPS_WindowAperiodic_r17, &uper.Constraint{Lb: 0, Ub: 30}, false); err != nil {
			return utils.WrapError("Encode sl_CPS_WindowAperiodic_r17", err)
		}
	}
	if ie.sl_PartialSensingInactiveTime_r17 != nil {
		if err = ie.sl_PartialSensingInactiveTime_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PartialSensingInactiveTime_r17", err)
		}
	}
	return nil
}

func (ie *SL_PBPS_CPS_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var sl_AllowedResourceSelectionConfig_r17Present bool
	if sl_AllowedResourceSelectionConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MinNumCandidateSlotsPeriodic_r17Present bool
	if sl_MinNumCandidateSlotsPeriodic_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PBPS_OccasionReservePeriodList_r17Present bool
	if sl_PBPS_OccasionReservePeriodList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_Additional_PBPS_Occasion_r17Present bool
	if sl_Additional_PBPS_Occasion_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_CPS_WindowPeriodic_r17Present bool
	if sl_CPS_WindowPeriodic_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MinNumCandidateSlotsAperiodic_r17Present bool
	if sl_MinNumCandidateSlotsAperiodic_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MinNumRssiMeasurementSlots_r17Present bool
	if sl_MinNumRssiMeasurementSlots_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_DefaultCBR_RandomSelection_r17Present bool
	if sl_DefaultCBR_RandomSelection_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_DefaultCBR_PartialSensing_r17Present bool
	if sl_DefaultCBR_PartialSensing_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_CPS_WindowAperiodic_r17Present bool
	if sl_CPS_WindowAperiodic_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PartialSensingInactiveTime_r17Present bool
	if sl_PartialSensingInactiveTime_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_AllowedResourceSelectionConfig_r17Present {
		ie.sl_AllowedResourceSelectionConfig_r17 = new(SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17)
		if err = ie.sl_AllowedResourceSelectionConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_AllowedResourceSelectionConfig_r17", err)
		}
	}
	if sl_MinNumCandidateSlotsPeriodic_r17Present {
		var tmp_int_sl_MinNumCandidateSlotsPeriodic_r17 int64
		if tmp_int_sl_MinNumCandidateSlotsPeriodic_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode sl_MinNumCandidateSlotsPeriodic_r17", err)
		}
		ie.sl_MinNumCandidateSlotsPeriodic_r17 = &tmp_int_sl_MinNumCandidateSlotsPeriodic_r17
	}
	if sl_PBPS_OccasionReservePeriodList_r17Present {
		tmp_sl_PBPS_OccasionReservePeriodList_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 16}, false)
		fn_sl_PBPS_OccasionReservePeriodList_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 1, Ub: 16}, false)
			return &ie
		}
		if err = tmp_sl_PBPS_OccasionReservePeriodList_r17.Decode(r, fn_sl_PBPS_OccasionReservePeriodList_r17); err != nil {
			return utils.WrapError("Decode sl_PBPS_OccasionReservePeriodList_r17", err)
		}
		ie.sl_PBPS_OccasionReservePeriodList_r17 = []int64{}
		for _, i := range tmp_sl_PBPS_OccasionReservePeriodList_r17.Value {
			ie.sl_PBPS_OccasionReservePeriodList_r17 = append(ie.sl_PBPS_OccasionReservePeriodList_r17, int64(i.Value))
		}
	}
	if sl_Additional_PBPS_Occasion_r17Present {
		ie.sl_Additional_PBPS_Occasion_r17 = new(SL_PBPS_CPS_Config_r17_sl_Additional_PBPS_Occasion_r17)
		if err = ie.sl_Additional_PBPS_Occasion_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_Additional_PBPS_Occasion_r17", err)
		}
	}
	if sl_CPS_WindowPeriodic_r17Present {
		var tmp_int_sl_CPS_WindowPeriodic_r17 int64
		if tmp_int_sl_CPS_WindowPeriodic_r17, err = r.ReadInteger(&uper.Constraint{Lb: 5, Ub: 30}, false); err != nil {
			return utils.WrapError("Decode sl_CPS_WindowPeriodic_r17", err)
		}
		ie.sl_CPS_WindowPeriodic_r17 = &tmp_int_sl_CPS_WindowPeriodic_r17
	}
	if sl_MinNumCandidateSlotsAperiodic_r17Present {
		var tmp_int_sl_MinNumCandidateSlotsAperiodic_r17 int64
		if tmp_int_sl_MinNumCandidateSlotsAperiodic_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode sl_MinNumCandidateSlotsAperiodic_r17", err)
		}
		ie.sl_MinNumCandidateSlotsAperiodic_r17 = &tmp_int_sl_MinNumCandidateSlotsAperiodic_r17
	}
	if sl_MinNumRssiMeasurementSlots_r17Present {
		var tmp_int_sl_MinNumRssiMeasurementSlots_r17 int64
		if tmp_int_sl_MinNumRssiMeasurementSlots_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 800}, false); err != nil {
			return utils.WrapError("Decode sl_MinNumRssiMeasurementSlots_r17", err)
		}
		ie.sl_MinNumRssiMeasurementSlots_r17 = &tmp_int_sl_MinNumRssiMeasurementSlots_r17
	}
	if sl_DefaultCBR_RandomSelection_r17Present {
		var tmp_int_sl_DefaultCBR_RandomSelection_r17 int64
		if tmp_int_sl_DefaultCBR_RandomSelection_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 100}, false); err != nil {
			return utils.WrapError("Decode sl_DefaultCBR_RandomSelection_r17", err)
		}
		ie.sl_DefaultCBR_RandomSelection_r17 = &tmp_int_sl_DefaultCBR_RandomSelection_r17
	}
	if sl_DefaultCBR_PartialSensing_r17Present {
		var tmp_int_sl_DefaultCBR_PartialSensing_r17 int64
		if tmp_int_sl_DefaultCBR_PartialSensing_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 100}, false); err != nil {
			return utils.WrapError("Decode sl_DefaultCBR_PartialSensing_r17", err)
		}
		ie.sl_DefaultCBR_PartialSensing_r17 = &tmp_int_sl_DefaultCBR_PartialSensing_r17
	}
	if sl_CPS_WindowAperiodic_r17Present {
		var tmp_int_sl_CPS_WindowAperiodic_r17 int64
		if tmp_int_sl_CPS_WindowAperiodic_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 30}, false); err != nil {
			return utils.WrapError("Decode sl_CPS_WindowAperiodic_r17", err)
		}
		ie.sl_CPS_WindowAperiodic_r17 = &tmp_int_sl_CPS_WindowAperiodic_r17
	}
	if sl_PartialSensingInactiveTime_r17Present {
		ie.sl_PartialSensingInactiveTime_r17 = new(SL_PBPS_CPS_Config_r17_sl_PartialSensingInactiveTime_r17)
		if err = ie.sl_PartialSensingInactiveTime_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PartialSensingInactiveTime_r17", err)
		}
	}
	return nil
}
