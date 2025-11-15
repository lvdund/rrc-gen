package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfiguration_v1610_IEs struct {
	otherConfig_v1610                   *OtherConfig_v1610                                   `optional`
	bap_Config_r16                      *BAP_Config_r16                                      `optional,setuprelease`
	iab_IP_AddressConfigurationList_r16 *IAB_IP_AddressConfigurationList_r16                 `optional`
	conditionalReconfiguration_r16      *ConditionalReconfiguration_r16                      `optional`
	daps_SourceRelease_r16              *RRCReconfiguration_v1610_IEs_daps_SourceRelease_r16 `optional`
	t316_r16                            *T316_r16                                            `optional,setuprelease`
	needForGapsConfigNR_r16             *NeedForGapsConfigNR_r16                             `optional,setuprelease`
	onDemandSIB_Request_r16             *OnDemandSIB_Request_r16                             `optional,setuprelease`
	dedicatedPosSysInfoDelivery_r16     *[]byte                                              `optional`
	sl_ConfigDedicatedNR_r16            *SL_ConfigDedicatedNR_r16                            `optional,setuprelease`
	sl_ConfigDedicatedEUTRA_Info_r16    *SL_ConfigDedicatedEUTRA_Info_r16                    `optional,setuprelease`
	targetCellSMTC_SCG_r16              *SSB_MTC                                             `optional`
	nonCriticalExtension                *RRCReconfiguration_v1700_IEs                        `optional`
}

func (ie *RRCReconfiguration_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.otherConfig_v1610 != nil, ie.bap_Config_r16 != nil, ie.iab_IP_AddressConfigurationList_r16 != nil, ie.conditionalReconfiguration_r16 != nil, ie.daps_SourceRelease_r16 != nil, ie.t316_r16 != nil, ie.needForGapsConfigNR_r16 != nil, ie.onDemandSIB_Request_r16 != nil, ie.dedicatedPosSysInfoDelivery_r16 != nil, ie.sl_ConfigDedicatedNR_r16 != nil, ie.sl_ConfigDedicatedEUTRA_Info_r16 != nil, ie.targetCellSMTC_SCG_r16 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.otherConfig_v1610 != nil {
		if err = ie.otherConfig_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode otherConfig_v1610", err)
		}
	}
	if ie.bap_Config_r16 != nil {
		tmp_bap_Config_r16 := utils.SetupRelease[*BAP_Config_r16]{
			Setup: ie.bap_Config_r16,
		}
		if err = tmp_bap_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode bap_Config_r16", err)
		}
	}
	if ie.iab_IP_AddressConfigurationList_r16 != nil {
		if err = ie.iab_IP_AddressConfigurationList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode iab_IP_AddressConfigurationList_r16", err)
		}
	}
	if ie.conditionalReconfiguration_r16 != nil {
		if err = ie.conditionalReconfiguration_r16.Encode(w); err != nil {
			return utils.WrapError("Encode conditionalReconfiguration_r16", err)
		}
	}
	if ie.daps_SourceRelease_r16 != nil {
		if err = ie.daps_SourceRelease_r16.Encode(w); err != nil {
			return utils.WrapError("Encode daps_SourceRelease_r16", err)
		}
	}
	if ie.t316_r16 != nil {
		tmp_t316_r16 := utils.SetupRelease[*T316_r16]{
			Setup: ie.t316_r16,
		}
		if err = tmp_t316_r16.Encode(w); err != nil {
			return utils.WrapError("Encode t316_r16", err)
		}
	}
	if ie.needForGapsConfigNR_r16 != nil {
		tmp_needForGapsConfigNR_r16 := utils.SetupRelease[*NeedForGapsConfigNR_r16]{
			Setup: ie.needForGapsConfigNR_r16,
		}
		if err = tmp_needForGapsConfigNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode needForGapsConfigNR_r16", err)
		}
	}
	if ie.onDemandSIB_Request_r16 != nil {
		tmp_onDemandSIB_Request_r16 := utils.SetupRelease[*OnDemandSIB_Request_r16]{
			Setup: ie.onDemandSIB_Request_r16,
		}
		if err = tmp_onDemandSIB_Request_r16.Encode(w); err != nil {
			return utils.WrapError("Encode onDemandSIB_Request_r16", err)
		}
	}
	if ie.dedicatedPosSysInfoDelivery_r16 != nil {
		if err = w.WriteOctetString(*ie.dedicatedPosSysInfoDelivery_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode dedicatedPosSysInfoDelivery_r16", err)
		}
	}
	if ie.sl_ConfigDedicatedNR_r16 != nil {
		tmp_sl_ConfigDedicatedNR_r16 := utils.SetupRelease[*SL_ConfigDedicatedNR_r16]{
			Setup: ie.sl_ConfigDedicatedNR_r16,
		}
		if err = tmp_sl_ConfigDedicatedNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ConfigDedicatedNR_r16", err)
		}
	}
	if ie.sl_ConfigDedicatedEUTRA_Info_r16 != nil {
		tmp_sl_ConfigDedicatedEUTRA_Info_r16 := utils.SetupRelease[*SL_ConfigDedicatedEUTRA_Info_r16]{
			Setup: ie.sl_ConfigDedicatedEUTRA_Info_r16,
		}
		if err = tmp_sl_ConfigDedicatedEUTRA_Info_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ConfigDedicatedEUTRA_Info_r16", err)
		}
	}
	if ie.targetCellSMTC_SCG_r16 != nil {
		if err = ie.targetCellSMTC_SCG_r16.Encode(w); err != nil {
			return utils.WrapError("Encode targetCellSMTC_SCG_r16", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfiguration_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var otherConfig_v1610Present bool
	if otherConfig_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bap_Config_r16Present bool
	if bap_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var iab_IP_AddressConfigurationList_r16Present bool
	if iab_IP_AddressConfigurationList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var conditionalReconfiguration_r16Present bool
	if conditionalReconfiguration_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var daps_SourceRelease_r16Present bool
	if daps_SourceRelease_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var t316_r16Present bool
	if t316_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var needForGapsConfigNR_r16Present bool
	if needForGapsConfigNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var onDemandSIB_Request_r16Present bool
	if onDemandSIB_Request_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dedicatedPosSysInfoDelivery_r16Present bool
	if dedicatedPosSysInfoDelivery_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ConfigDedicatedNR_r16Present bool
	if sl_ConfigDedicatedNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ConfigDedicatedEUTRA_Info_r16Present bool
	if sl_ConfigDedicatedEUTRA_Info_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var targetCellSMTC_SCG_r16Present bool
	if targetCellSMTC_SCG_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if otherConfig_v1610Present {
		ie.otherConfig_v1610 = new(OtherConfig_v1610)
		if err = ie.otherConfig_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode otherConfig_v1610", err)
		}
	}
	if bap_Config_r16Present {
		tmp_bap_Config_r16 := utils.SetupRelease[*BAP_Config_r16]{}
		if err = tmp_bap_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode bap_Config_r16", err)
		}
		ie.bap_Config_r16 = tmp_bap_Config_r16.Setup
	}
	if iab_IP_AddressConfigurationList_r16Present {
		ie.iab_IP_AddressConfigurationList_r16 = new(IAB_IP_AddressConfigurationList_r16)
		if err = ie.iab_IP_AddressConfigurationList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode iab_IP_AddressConfigurationList_r16", err)
		}
	}
	if conditionalReconfiguration_r16Present {
		ie.conditionalReconfiguration_r16 = new(ConditionalReconfiguration_r16)
		if err = ie.conditionalReconfiguration_r16.Decode(r); err != nil {
			return utils.WrapError("Decode conditionalReconfiguration_r16", err)
		}
	}
	if daps_SourceRelease_r16Present {
		ie.daps_SourceRelease_r16 = new(RRCReconfiguration_v1610_IEs_daps_SourceRelease_r16)
		if err = ie.daps_SourceRelease_r16.Decode(r); err != nil {
			return utils.WrapError("Decode daps_SourceRelease_r16", err)
		}
	}
	if t316_r16Present {
		tmp_t316_r16 := utils.SetupRelease[*T316_r16]{}
		if err = tmp_t316_r16.Decode(r); err != nil {
			return utils.WrapError("Decode t316_r16", err)
		}
		ie.t316_r16 = tmp_t316_r16.Setup
	}
	if needForGapsConfigNR_r16Present {
		tmp_needForGapsConfigNR_r16 := utils.SetupRelease[*NeedForGapsConfigNR_r16]{}
		if err = tmp_needForGapsConfigNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode needForGapsConfigNR_r16", err)
		}
		ie.needForGapsConfigNR_r16 = tmp_needForGapsConfigNR_r16.Setup
	}
	if onDemandSIB_Request_r16Present {
		tmp_onDemandSIB_Request_r16 := utils.SetupRelease[*OnDemandSIB_Request_r16]{}
		if err = tmp_onDemandSIB_Request_r16.Decode(r); err != nil {
			return utils.WrapError("Decode onDemandSIB_Request_r16", err)
		}
		ie.onDemandSIB_Request_r16 = tmp_onDemandSIB_Request_r16.Setup
	}
	if dedicatedPosSysInfoDelivery_r16Present {
		var tmp_os_dedicatedPosSysInfoDelivery_r16 []byte
		if tmp_os_dedicatedPosSysInfoDelivery_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode dedicatedPosSysInfoDelivery_r16", err)
		}
		ie.dedicatedPosSysInfoDelivery_r16 = &tmp_os_dedicatedPosSysInfoDelivery_r16
	}
	if sl_ConfigDedicatedNR_r16Present {
		tmp_sl_ConfigDedicatedNR_r16 := utils.SetupRelease[*SL_ConfigDedicatedNR_r16]{}
		if err = tmp_sl_ConfigDedicatedNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ConfigDedicatedNR_r16", err)
		}
		ie.sl_ConfigDedicatedNR_r16 = tmp_sl_ConfigDedicatedNR_r16.Setup
	}
	if sl_ConfigDedicatedEUTRA_Info_r16Present {
		tmp_sl_ConfigDedicatedEUTRA_Info_r16 := utils.SetupRelease[*SL_ConfigDedicatedEUTRA_Info_r16]{}
		if err = tmp_sl_ConfigDedicatedEUTRA_Info_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ConfigDedicatedEUTRA_Info_r16", err)
		}
		ie.sl_ConfigDedicatedEUTRA_Info_r16 = tmp_sl_ConfigDedicatedEUTRA_Info_r16.Setup
	}
	if targetCellSMTC_SCG_r16Present {
		ie.targetCellSMTC_SCG_r16 = new(SSB_MTC)
		if err = ie.targetCellSMTC_SCG_r16.Decode(r); err != nil {
			return utils.WrapError("Decode targetCellSMTC_SCG_r16", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCReconfiguration_v1700_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
