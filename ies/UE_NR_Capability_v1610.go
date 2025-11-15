package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v1610 struct {
	inDeviceCoexInd_r16                 *UE_NR_Capability_v1610_inDeviceCoexInd_r16                 `optional`
	dl_DedicatedMessageSegmentation_r16 *UE_NR_Capability_v1610_dl_DedicatedMessageSegmentation_r16 `optional`
	nrdc_Parameters_v1610               *NRDC_Parameters_v1610                                      `optional`
	powSav_Parameters_r16               *PowSav_Parameters_r16                                      `optional`
	fr1_Add_UE_NR_Capabilities_v1610    *UE_NR_CapabilityAddFRX_Mode_v1610                          `optional`
	fr2_Add_UE_NR_Capabilities_v1610    *UE_NR_CapabilityAddFRX_Mode_v1610                          `optional`
	bh_RLF_Indication_r16               *UE_NR_Capability_v1610_bh_RLF_Indication_r16               `optional`
	directSN_AdditionFirstRRC_IAB_r16   *UE_NR_Capability_v1610_directSN_AdditionFirstRRC_IAB_r16   `optional`
	bap_Parameters_r16                  *BAP_Parameters_r16                                         `optional`
	referenceTimeProvision_r16          *UE_NR_Capability_v1610_referenceTimeProvision_r16          `optional`
	sidelinkParameters_r16              *SidelinkParameters_r16                                     `optional`
	highSpeedParameters_r16             *HighSpeedParameters_r16                                    `optional`
	mac_Parameters_v1610                *MAC_Parameters_v1610                                       `optional`
	mcgRLF_RecoveryViaSCG_r16           *UE_NR_Capability_v1610_mcgRLF_RecoveryViaSCG_r16           `optional`
	resumeWithStoredMCG_SCells_r16      *UE_NR_Capability_v1610_resumeWithStoredMCG_SCells_r16      `optional`
	resumeWithStoredSCG_r16             *UE_NR_Capability_v1610_resumeWithStoredSCG_r16             `optional`
	resumeWithSCG_Config_r16            *UE_NR_Capability_v1610_resumeWithSCG_Config_r16            `optional`
	ue_BasedPerfMeas_Parameters_r16     *UE_BasedPerfMeas_Parameters_r16                            `optional`
	son_Parameters_r16                  *SON_Parameters_r16                                         `optional`
	onDemandSIB_Connected_r16           *UE_NR_Capability_v1610_onDemandSIB_Connected_r16           `optional`
	nonCriticalExtension                *UE_NR_Capability_v1640                                     `optional`
}

func (ie *UE_NR_Capability_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.inDeviceCoexInd_r16 != nil, ie.dl_DedicatedMessageSegmentation_r16 != nil, ie.nrdc_Parameters_v1610 != nil, ie.powSav_Parameters_r16 != nil, ie.fr1_Add_UE_NR_Capabilities_v1610 != nil, ie.fr2_Add_UE_NR_Capabilities_v1610 != nil, ie.bh_RLF_Indication_r16 != nil, ie.directSN_AdditionFirstRRC_IAB_r16 != nil, ie.bap_Parameters_r16 != nil, ie.referenceTimeProvision_r16 != nil, ie.sidelinkParameters_r16 != nil, ie.highSpeedParameters_r16 != nil, ie.mac_Parameters_v1610 != nil, ie.mcgRLF_RecoveryViaSCG_r16 != nil, ie.resumeWithStoredMCG_SCells_r16 != nil, ie.resumeWithStoredSCG_r16 != nil, ie.resumeWithSCG_Config_r16 != nil, ie.ue_BasedPerfMeas_Parameters_r16 != nil, ie.son_Parameters_r16 != nil, ie.onDemandSIB_Connected_r16 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.inDeviceCoexInd_r16 != nil {
		if err = ie.inDeviceCoexInd_r16.Encode(w); err != nil {
			return utils.WrapError("Encode inDeviceCoexInd_r16", err)
		}
	}
	if ie.dl_DedicatedMessageSegmentation_r16 != nil {
		if err = ie.dl_DedicatedMessageSegmentation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dl_DedicatedMessageSegmentation_r16", err)
		}
	}
	if ie.nrdc_Parameters_v1610 != nil {
		if err = ie.nrdc_Parameters_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode nrdc_Parameters_v1610", err)
		}
	}
	if ie.powSav_Parameters_r16 != nil {
		if err = ie.powSav_Parameters_r16.Encode(w); err != nil {
			return utils.WrapError("Encode powSav_Parameters_r16", err)
		}
	}
	if ie.fr1_Add_UE_NR_Capabilities_v1610 != nil {
		if err = ie.fr1_Add_UE_NR_Capabilities_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode fr1_Add_UE_NR_Capabilities_v1610", err)
		}
	}
	if ie.fr2_Add_UE_NR_Capabilities_v1610 != nil {
		if err = ie.fr2_Add_UE_NR_Capabilities_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode fr2_Add_UE_NR_Capabilities_v1610", err)
		}
	}
	if ie.bh_RLF_Indication_r16 != nil {
		if err = ie.bh_RLF_Indication_r16.Encode(w); err != nil {
			return utils.WrapError("Encode bh_RLF_Indication_r16", err)
		}
	}
	if ie.directSN_AdditionFirstRRC_IAB_r16 != nil {
		if err = ie.directSN_AdditionFirstRRC_IAB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode directSN_AdditionFirstRRC_IAB_r16", err)
		}
	}
	if ie.bap_Parameters_r16 != nil {
		if err = ie.bap_Parameters_r16.Encode(w); err != nil {
			return utils.WrapError("Encode bap_Parameters_r16", err)
		}
	}
	if ie.referenceTimeProvision_r16 != nil {
		if err = ie.referenceTimeProvision_r16.Encode(w); err != nil {
			return utils.WrapError("Encode referenceTimeProvision_r16", err)
		}
	}
	if ie.sidelinkParameters_r16 != nil {
		if err = ie.sidelinkParameters_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sidelinkParameters_r16", err)
		}
	}
	if ie.highSpeedParameters_r16 != nil {
		if err = ie.highSpeedParameters_r16.Encode(w); err != nil {
			return utils.WrapError("Encode highSpeedParameters_r16", err)
		}
	}
	if ie.mac_Parameters_v1610 != nil {
		if err = ie.mac_Parameters_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode mac_Parameters_v1610", err)
		}
	}
	if ie.mcgRLF_RecoveryViaSCG_r16 != nil {
		if err = ie.mcgRLF_RecoveryViaSCG_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mcgRLF_RecoveryViaSCG_r16", err)
		}
	}
	if ie.resumeWithStoredMCG_SCells_r16 != nil {
		if err = ie.resumeWithStoredMCG_SCells_r16.Encode(w); err != nil {
			return utils.WrapError("Encode resumeWithStoredMCG_SCells_r16", err)
		}
	}
	if ie.resumeWithStoredSCG_r16 != nil {
		if err = ie.resumeWithStoredSCG_r16.Encode(w); err != nil {
			return utils.WrapError("Encode resumeWithStoredSCG_r16", err)
		}
	}
	if ie.resumeWithSCG_Config_r16 != nil {
		if err = ie.resumeWithSCG_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode resumeWithSCG_Config_r16", err)
		}
	}
	if ie.ue_BasedPerfMeas_Parameters_r16 != nil {
		if err = ie.ue_BasedPerfMeas_Parameters_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ue_BasedPerfMeas_Parameters_r16", err)
		}
	}
	if ie.son_Parameters_r16 != nil {
		if err = ie.son_Parameters_r16.Encode(w); err != nil {
			return utils.WrapError("Encode son_Parameters_r16", err)
		}
	}
	if ie.onDemandSIB_Connected_r16 != nil {
		if err = ie.onDemandSIB_Connected_r16.Encode(w); err != nil {
			return utils.WrapError("Encode onDemandSIB_Connected_r16", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v1610) Decode(r *uper.UperReader) error {
	var err error
	var inDeviceCoexInd_r16Present bool
	if inDeviceCoexInd_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_DedicatedMessageSegmentation_r16Present bool
	if dl_DedicatedMessageSegmentation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nrdc_Parameters_v1610Present bool
	if nrdc_Parameters_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var powSav_Parameters_r16Present bool
	if powSav_Parameters_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fr1_Add_UE_NR_Capabilities_v1610Present bool
	if fr1_Add_UE_NR_Capabilities_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fr2_Add_UE_NR_Capabilities_v1610Present bool
	if fr2_Add_UE_NR_Capabilities_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bh_RLF_Indication_r16Present bool
	if bh_RLF_Indication_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var directSN_AdditionFirstRRC_IAB_r16Present bool
	if directSN_AdditionFirstRRC_IAB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bap_Parameters_r16Present bool
	if bap_Parameters_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var referenceTimeProvision_r16Present bool
	if referenceTimeProvision_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sidelinkParameters_r16Present bool
	if sidelinkParameters_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var highSpeedParameters_r16Present bool
	if highSpeedParameters_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mac_Parameters_v1610Present bool
	if mac_Parameters_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mcgRLF_RecoveryViaSCG_r16Present bool
	if mcgRLF_RecoveryViaSCG_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var resumeWithStoredMCG_SCells_r16Present bool
	if resumeWithStoredMCG_SCells_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var resumeWithStoredSCG_r16Present bool
	if resumeWithStoredSCG_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var resumeWithSCG_Config_r16Present bool
	if resumeWithSCG_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ue_BasedPerfMeas_Parameters_r16Present bool
	if ue_BasedPerfMeas_Parameters_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var son_Parameters_r16Present bool
	if son_Parameters_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var onDemandSIB_Connected_r16Present bool
	if onDemandSIB_Connected_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if inDeviceCoexInd_r16Present {
		ie.inDeviceCoexInd_r16 = new(UE_NR_Capability_v1610_inDeviceCoexInd_r16)
		if err = ie.inDeviceCoexInd_r16.Decode(r); err != nil {
			return utils.WrapError("Decode inDeviceCoexInd_r16", err)
		}
	}
	if dl_DedicatedMessageSegmentation_r16Present {
		ie.dl_DedicatedMessageSegmentation_r16 = new(UE_NR_Capability_v1610_dl_DedicatedMessageSegmentation_r16)
		if err = ie.dl_DedicatedMessageSegmentation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dl_DedicatedMessageSegmentation_r16", err)
		}
	}
	if nrdc_Parameters_v1610Present {
		ie.nrdc_Parameters_v1610 = new(NRDC_Parameters_v1610)
		if err = ie.nrdc_Parameters_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode nrdc_Parameters_v1610", err)
		}
	}
	if powSav_Parameters_r16Present {
		ie.powSav_Parameters_r16 = new(PowSav_Parameters_r16)
		if err = ie.powSav_Parameters_r16.Decode(r); err != nil {
			return utils.WrapError("Decode powSav_Parameters_r16", err)
		}
	}
	if fr1_Add_UE_NR_Capabilities_v1610Present {
		ie.fr1_Add_UE_NR_Capabilities_v1610 = new(UE_NR_CapabilityAddFRX_Mode_v1610)
		if err = ie.fr1_Add_UE_NR_Capabilities_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode fr1_Add_UE_NR_Capabilities_v1610", err)
		}
	}
	if fr2_Add_UE_NR_Capabilities_v1610Present {
		ie.fr2_Add_UE_NR_Capabilities_v1610 = new(UE_NR_CapabilityAddFRX_Mode_v1610)
		if err = ie.fr2_Add_UE_NR_Capabilities_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode fr2_Add_UE_NR_Capabilities_v1610", err)
		}
	}
	if bh_RLF_Indication_r16Present {
		ie.bh_RLF_Indication_r16 = new(UE_NR_Capability_v1610_bh_RLF_Indication_r16)
		if err = ie.bh_RLF_Indication_r16.Decode(r); err != nil {
			return utils.WrapError("Decode bh_RLF_Indication_r16", err)
		}
	}
	if directSN_AdditionFirstRRC_IAB_r16Present {
		ie.directSN_AdditionFirstRRC_IAB_r16 = new(UE_NR_Capability_v1610_directSN_AdditionFirstRRC_IAB_r16)
		if err = ie.directSN_AdditionFirstRRC_IAB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode directSN_AdditionFirstRRC_IAB_r16", err)
		}
	}
	if bap_Parameters_r16Present {
		ie.bap_Parameters_r16 = new(BAP_Parameters_r16)
		if err = ie.bap_Parameters_r16.Decode(r); err != nil {
			return utils.WrapError("Decode bap_Parameters_r16", err)
		}
	}
	if referenceTimeProvision_r16Present {
		ie.referenceTimeProvision_r16 = new(UE_NR_Capability_v1610_referenceTimeProvision_r16)
		if err = ie.referenceTimeProvision_r16.Decode(r); err != nil {
			return utils.WrapError("Decode referenceTimeProvision_r16", err)
		}
	}
	if sidelinkParameters_r16Present {
		ie.sidelinkParameters_r16 = new(SidelinkParameters_r16)
		if err = ie.sidelinkParameters_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sidelinkParameters_r16", err)
		}
	}
	if highSpeedParameters_r16Present {
		ie.highSpeedParameters_r16 = new(HighSpeedParameters_r16)
		if err = ie.highSpeedParameters_r16.Decode(r); err != nil {
			return utils.WrapError("Decode highSpeedParameters_r16", err)
		}
	}
	if mac_Parameters_v1610Present {
		ie.mac_Parameters_v1610 = new(MAC_Parameters_v1610)
		if err = ie.mac_Parameters_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode mac_Parameters_v1610", err)
		}
	}
	if mcgRLF_RecoveryViaSCG_r16Present {
		ie.mcgRLF_RecoveryViaSCG_r16 = new(UE_NR_Capability_v1610_mcgRLF_RecoveryViaSCG_r16)
		if err = ie.mcgRLF_RecoveryViaSCG_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mcgRLF_RecoveryViaSCG_r16", err)
		}
	}
	if resumeWithStoredMCG_SCells_r16Present {
		ie.resumeWithStoredMCG_SCells_r16 = new(UE_NR_Capability_v1610_resumeWithStoredMCG_SCells_r16)
		if err = ie.resumeWithStoredMCG_SCells_r16.Decode(r); err != nil {
			return utils.WrapError("Decode resumeWithStoredMCG_SCells_r16", err)
		}
	}
	if resumeWithStoredSCG_r16Present {
		ie.resumeWithStoredSCG_r16 = new(UE_NR_Capability_v1610_resumeWithStoredSCG_r16)
		if err = ie.resumeWithStoredSCG_r16.Decode(r); err != nil {
			return utils.WrapError("Decode resumeWithStoredSCG_r16", err)
		}
	}
	if resumeWithSCG_Config_r16Present {
		ie.resumeWithSCG_Config_r16 = new(UE_NR_Capability_v1610_resumeWithSCG_Config_r16)
		if err = ie.resumeWithSCG_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode resumeWithSCG_Config_r16", err)
		}
	}
	if ue_BasedPerfMeas_Parameters_r16Present {
		ie.ue_BasedPerfMeas_Parameters_r16 = new(UE_BasedPerfMeas_Parameters_r16)
		if err = ie.ue_BasedPerfMeas_Parameters_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ue_BasedPerfMeas_Parameters_r16", err)
		}
	}
	if son_Parameters_r16Present {
		ie.son_Parameters_r16 = new(SON_Parameters_r16)
		if err = ie.son_Parameters_r16.Decode(r); err != nil {
			return utils.WrapError("Decode son_Parameters_r16", err)
		}
	}
	if onDemandSIB_Connected_r16Present {
		ie.onDemandSIB_Connected_r16 = new(UE_NR_Capability_v1610_onDemandSIB_Connected_r16)
		if err = ie.onDemandSIB_Connected_r16.Decode(r); err != nil {
			return utils.WrapError("Decode onDemandSIB_Connected_r16", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UE_NR_Capability_v1640)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
