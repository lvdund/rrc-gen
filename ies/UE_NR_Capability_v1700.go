package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v1700 struct {
	inactiveStatePO_Determination_r17       *UE_NR_Capability_v1700_inactiveStatePO_Determination_r17       `optional`
	highSpeedParameters_v1700               *HighSpeedParameters_v1700                                      `optional`
	powSav_Parameters_v1700                 *PowSav_Parameters_v1700                                        `optional`
	mac_Parameters_v1700                    *MAC_Parameters_v1700                                           `optional`
	ims_Parameters_v1700                    *IMS_Parameters_v1700                                           `optional`
	measAndMobParameters_v1700              MeasAndMobParameters_v1700                                      `madatory`
	appLayerMeasParameters_r17              *AppLayerMeasParameters_r17                                     `optional`
	redCapParameters_r17                    *RedCapParameters_r17                                           `optional`
	ra_SDT_r17                              *UE_NR_Capability_v1700_ra_SDT_r17                              `optional`
	srb_SDT_r17                             *UE_NR_Capability_v1700_srb_SDT_r17                             `optional`
	gNB_SideRTT_BasedPDC_r17                *UE_NR_Capability_v1700_gNB_SideRTT_BasedPDC_r17                `optional`
	bh_RLF_DetectionRecovery_Indication_r17 *UE_NR_Capability_v1700_bh_RLF_DetectionRecovery_Indication_r17 `optional`
	nrdc_Parameters_v1700                   *NRDC_Parameters_v1700                                          `optional`
	bap_Parameters_v1700                    *BAP_Parameters_v1700                                           `optional`
	musim_GapPreference_r17                 *UE_NR_Capability_v1700_musim_GapPreference_r17                 `optional`
	musimLeaveConnected_r17                 *UE_NR_Capability_v1700_musimLeaveConnected_r17                 `optional`
	mbs_Parameters_r17                      MBS_Parameters_r17                                              `madatory`
	nonTerrestrialNetwork_r17               *UE_NR_Capability_v1700_nonTerrestrialNetwork_r17               `optional`
	ntn_ScenarioSupport_r17                 *UE_NR_Capability_v1700_ntn_ScenarioSupport_r17                 `optional`
	sliceInfoforCellReselection_r17         *UE_NR_Capability_v1700_sliceInfoforCellReselection_r17         `optional`
	ue_RadioPagingInfo_r17                  *UE_RadioPagingInfo_r17                                         `optional`
	ul_GapFR2_Pattern_r17                   *uper.BitString                                                 `lb:4,ub:4,optional`
	ntn_Parameters_r17                      *NTN_Parameters_r17                                             `optional`
	nonCriticalExtension                    interface{}                                                     `optional`
}

func (ie *UE_NR_Capability_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.inactiveStatePO_Determination_r17 != nil, ie.highSpeedParameters_v1700 != nil, ie.powSav_Parameters_v1700 != nil, ie.mac_Parameters_v1700 != nil, ie.ims_Parameters_v1700 != nil, ie.appLayerMeasParameters_r17 != nil, ie.redCapParameters_r17 != nil, ie.ra_SDT_r17 != nil, ie.srb_SDT_r17 != nil, ie.gNB_SideRTT_BasedPDC_r17 != nil, ie.bh_RLF_DetectionRecovery_Indication_r17 != nil, ie.nrdc_Parameters_v1700 != nil, ie.bap_Parameters_v1700 != nil, ie.musim_GapPreference_r17 != nil, ie.musimLeaveConnected_r17 != nil, ie.nonTerrestrialNetwork_r17 != nil, ie.ntn_ScenarioSupport_r17 != nil, ie.sliceInfoforCellReselection_r17 != nil, ie.ue_RadioPagingInfo_r17 != nil, ie.ul_GapFR2_Pattern_r17 != nil, ie.ntn_Parameters_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.inactiveStatePO_Determination_r17 != nil {
		if err = ie.inactiveStatePO_Determination_r17.Encode(w); err != nil {
			return utils.WrapError("Encode inactiveStatePO_Determination_r17", err)
		}
	}
	if ie.highSpeedParameters_v1700 != nil {
		if err = ie.highSpeedParameters_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode highSpeedParameters_v1700", err)
		}
	}
	if ie.powSav_Parameters_v1700 != nil {
		if err = ie.powSav_Parameters_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode powSav_Parameters_v1700", err)
		}
	}
	if ie.mac_Parameters_v1700 != nil {
		if err = ie.mac_Parameters_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode mac_Parameters_v1700", err)
		}
	}
	if ie.ims_Parameters_v1700 != nil {
		if err = ie.ims_Parameters_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode ims_Parameters_v1700", err)
		}
	}
	if ie.appLayerMeasParameters_r17 != nil {
		if err = ie.appLayerMeasParameters_r17.Encode(w); err != nil {
			return utils.WrapError("Encode appLayerMeasParameters_r17", err)
		}
	}
	if ie.redCapParameters_r17 != nil {
		if err = ie.redCapParameters_r17.Encode(w); err != nil {
			return utils.WrapError("Encode redCapParameters_r17", err)
		}
	}
	if ie.ra_SDT_r17 != nil {
		if err = ie.ra_SDT_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ra_SDT_r17", err)
		}
	}
	if ie.srb_SDT_r17 != nil {
		if err = ie.srb_SDT_r17.Encode(w); err != nil {
			return utils.WrapError("Encode srb_SDT_r17", err)
		}
	}
	if ie.gNB_SideRTT_BasedPDC_r17 != nil {
		if err = ie.gNB_SideRTT_BasedPDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode gNB_SideRTT_BasedPDC_r17", err)
		}
	}
	if ie.bh_RLF_DetectionRecovery_Indication_r17 != nil {
		if err = ie.bh_RLF_DetectionRecovery_Indication_r17.Encode(w); err != nil {
			return utils.WrapError("Encode bh_RLF_DetectionRecovery_Indication_r17", err)
		}
	}
	if ie.nrdc_Parameters_v1700 != nil {
		if err = ie.nrdc_Parameters_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode nrdc_Parameters_v1700", err)
		}
	}
	if ie.bap_Parameters_v1700 != nil {
		if err = ie.bap_Parameters_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode bap_Parameters_v1700", err)
		}
	}
	if ie.musim_GapPreference_r17 != nil {
		if err = ie.musim_GapPreference_r17.Encode(w); err != nil {
			return utils.WrapError("Encode musim_GapPreference_r17", err)
		}
	}
	if ie.musimLeaveConnected_r17 != nil {
		if err = ie.musimLeaveConnected_r17.Encode(w); err != nil {
			return utils.WrapError("Encode musimLeaveConnected_r17", err)
		}
	}
	if ie.nonTerrestrialNetwork_r17 != nil {
		if err = ie.nonTerrestrialNetwork_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nonTerrestrialNetwork_r17", err)
		}
	}
	if ie.ntn_ScenarioSupport_r17 != nil {
		if err = ie.ntn_ScenarioSupport_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ntn_ScenarioSupport_r17", err)
		}
	}
	if ie.sliceInfoforCellReselection_r17 != nil {
		if err = ie.sliceInfoforCellReselection_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sliceInfoforCellReselection_r17", err)
		}
	}
	if ie.ue_RadioPagingInfo_r17 != nil {
		if err = ie.ue_RadioPagingInfo_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ue_RadioPagingInfo_r17", err)
		}
	}
	if ie.ul_GapFR2_Pattern_r17 != nil {
		if err = w.WriteBitString(ie.ul_GapFR2_Pattern_r17.Bytes, uint(ie.ul_GapFR2_Pattern_r17.NumBits), &uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Encode ul_GapFR2_Pattern_r17", err)
		}
	}
	if ie.ntn_Parameters_r17 != nil {
		if err = ie.ntn_Parameters_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ntn_Parameters_r17", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v1700) Decode(r *uper.UperReader) error {
	var err error
	var inactiveStatePO_Determination_r17Present bool
	if inactiveStatePO_Determination_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var highSpeedParameters_v1700Present bool
	if highSpeedParameters_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var powSav_Parameters_v1700Present bool
	if powSav_Parameters_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mac_Parameters_v1700Present bool
	if mac_Parameters_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ims_Parameters_v1700Present bool
	if ims_Parameters_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var appLayerMeasParameters_r17Present bool
	if appLayerMeasParameters_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var redCapParameters_r17Present bool
	if redCapParameters_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ra_SDT_r17Present bool
	if ra_SDT_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var srb_SDT_r17Present bool
	if srb_SDT_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var gNB_SideRTT_BasedPDC_r17Present bool
	if gNB_SideRTT_BasedPDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bh_RLF_DetectionRecovery_Indication_r17Present bool
	if bh_RLF_DetectionRecovery_Indication_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nrdc_Parameters_v1700Present bool
	if nrdc_Parameters_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bap_Parameters_v1700Present bool
	if bap_Parameters_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var musim_GapPreference_r17Present bool
	if musim_GapPreference_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var musimLeaveConnected_r17Present bool
	if musimLeaveConnected_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonTerrestrialNetwork_r17Present bool
	if nonTerrestrialNetwork_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ntn_ScenarioSupport_r17Present bool
	if ntn_ScenarioSupport_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sliceInfoforCellReselection_r17Present bool
	if sliceInfoforCellReselection_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ue_RadioPagingInfo_r17Present bool
	if ue_RadioPagingInfo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_GapFR2_Pattern_r17Present bool
	if ul_GapFR2_Pattern_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ntn_Parameters_r17Present bool
	if ntn_Parameters_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if inactiveStatePO_Determination_r17Present {
		ie.inactiveStatePO_Determination_r17 = new(UE_NR_Capability_v1700_inactiveStatePO_Determination_r17)
		if err = ie.inactiveStatePO_Determination_r17.Decode(r); err != nil {
			return utils.WrapError("Decode inactiveStatePO_Determination_r17", err)
		}
	}
	if highSpeedParameters_v1700Present {
		ie.highSpeedParameters_v1700 = new(HighSpeedParameters_v1700)
		if err = ie.highSpeedParameters_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode highSpeedParameters_v1700", err)
		}
	}
	if powSav_Parameters_v1700Present {
		ie.powSav_Parameters_v1700 = new(PowSav_Parameters_v1700)
		if err = ie.powSav_Parameters_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode powSav_Parameters_v1700", err)
		}
	}
	if mac_Parameters_v1700Present {
		ie.mac_Parameters_v1700 = new(MAC_Parameters_v1700)
		if err = ie.mac_Parameters_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode mac_Parameters_v1700", err)
		}
	}
	if ims_Parameters_v1700Present {
		ie.ims_Parameters_v1700 = new(IMS_Parameters_v1700)
		if err = ie.ims_Parameters_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode ims_Parameters_v1700", err)
		}
	}
	if appLayerMeasParameters_r17Present {
		ie.appLayerMeasParameters_r17 = new(AppLayerMeasParameters_r17)
		if err = ie.appLayerMeasParameters_r17.Decode(r); err != nil {
			return utils.WrapError("Decode appLayerMeasParameters_r17", err)
		}
	}
	if redCapParameters_r17Present {
		ie.redCapParameters_r17 = new(RedCapParameters_r17)
		if err = ie.redCapParameters_r17.Decode(r); err != nil {
			return utils.WrapError("Decode redCapParameters_r17", err)
		}
	}
	if ra_SDT_r17Present {
		ie.ra_SDT_r17 = new(UE_NR_Capability_v1700_ra_SDT_r17)
		if err = ie.ra_SDT_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ra_SDT_r17", err)
		}
	}
	if srb_SDT_r17Present {
		ie.srb_SDT_r17 = new(UE_NR_Capability_v1700_srb_SDT_r17)
		if err = ie.srb_SDT_r17.Decode(r); err != nil {
			return utils.WrapError("Decode srb_SDT_r17", err)
		}
	}
	if gNB_SideRTT_BasedPDC_r17Present {
		ie.gNB_SideRTT_BasedPDC_r17 = new(UE_NR_Capability_v1700_gNB_SideRTT_BasedPDC_r17)
		if err = ie.gNB_SideRTT_BasedPDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode gNB_SideRTT_BasedPDC_r17", err)
		}
	}
	if bh_RLF_DetectionRecovery_Indication_r17Present {
		ie.bh_RLF_DetectionRecovery_Indication_r17 = new(UE_NR_Capability_v1700_bh_RLF_DetectionRecovery_Indication_r17)
		if err = ie.bh_RLF_DetectionRecovery_Indication_r17.Decode(r); err != nil {
			return utils.WrapError("Decode bh_RLF_DetectionRecovery_Indication_r17", err)
		}
	}
	if nrdc_Parameters_v1700Present {
		ie.nrdc_Parameters_v1700 = new(NRDC_Parameters_v1700)
		if err = ie.nrdc_Parameters_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode nrdc_Parameters_v1700", err)
		}
	}
	if bap_Parameters_v1700Present {
		ie.bap_Parameters_v1700 = new(BAP_Parameters_v1700)
		if err = ie.bap_Parameters_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode bap_Parameters_v1700", err)
		}
	}
	if musim_GapPreference_r17Present {
		ie.musim_GapPreference_r17 = new(UE_NR_Capability_v1700_musim_GapPreference_r17)
		if err = ie.musim_GapPreference_r17.Decode(r); err != nil {
			return utils.WrapError("Decode musim_GapPreference_r17", err)
		}
	}
	if musimLeaveConnected_r17Present {
		ie.musimLeaveConnected_r17 = new(UE_NR_Capability_v1700_musimLeaveConnected_r17)
		if err = ie.musimLeaveConnected_r17.Decode(r); err != nil {
			return utils.WrapError("Decode musimLeaveConnected_r17", err)
		}
	}
	if nonTerrestrialNetwork_r17Present {
		ie.nonTerrestrialNetwork_r17 = new(UE_NR_Capability_v1700_nonTerrestrialNetwork_r17)
		if err = ie.nonTerrestrialNetwork_r17.Decode(r); err != nil {
			return utils.WrapError("Decode nonTerrestrialNetwork_r17", err)
		}
	}
	if ntn_ScenarioSupport_r17Present {
		ie.ntn_ScenarioSupport_r17 = new(UE_NR_Capability_v1700_ntn_ScenarioSupport_r17)
		if err = ie.ntn_ScenarioSupport_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ntn_ScenarioSupport_r17", err)
		}
	}
	if sliceInfoforCellReselection_r17Present {
		ie.sliceInfoforCellReselection_r17 = new(UE_NR_Capability_v1700_sliceInfoforCellReselection_r17)
		if err = ie.sliceInfoforCellReselection_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sliceInfoforCellReselection_r17", err)
		}
	}
	if ue_RadioPagingInfo_r17Present {
		ie.ue_RadioPagingInfo_r17 = new(UE_RadioPagingInfo_r17)
		if err = ie.ue_RadioPagingInfo_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ue_RadioPagingInfo_r17", err)
		}
	}
	if ul_GapFR2_Pattern_r17Present {
		var tmp_bs_ul_GapFR2_Pattern_r17 []byte
		var n_ul_GapFR2_Pattern_r17 uint
		if tmp_bs_ul_GapFR2_Pattern_r17, n_ul_GapFR2_Pattern_r17, err = r.ReadBitString(&uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode ul_GapFR2_Pattern_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_ul_GapFR2_Pattern_r17,
			NumBits: uint64(n_ul_GapFR2_Pattern_r17),
		}
		ie.ul_GapFR2_Pattern_r17 = &tmp_bitstring
	}
	if ntn_Parameters_r17Present {
		ie.ntn_Parameters_r17 = new(NTN_Parameters_r17)
		if err = ie.ntn_Parameters_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ntn_Parameters_r17", err)
		}
	}
	return nil
}
