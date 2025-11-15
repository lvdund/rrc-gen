package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCSetupComplete_IEs struct {
	selectedPLMN_Identity    int64                                    `lb:1,ub:maxPLMN,madatory`
	registeredAMF            *RegisteredAMF                           `optional`
	guami_Type               *RRCSetupComplete_IEs_guami_Type         `optional`
	s_NSSAI_List             []S_NSSAI                                `lb:1,ub:maxNrofS_NSSAI,optional`
	dedicatedNAS_Message     DedicatedNAS_Message                     `madatory`
	ng_5G_S_TMSI_Value       *RRCSetupComplete_IEs_ng_5G_S_TMSI_Value `lb:9,ub:9,optional`
	lateNonCriticalExtension *[]byte                                  `optional`
	nonCriticalExtension     *RRCSetupComplete_v1610_IEs              `optional`
}

func (ie *RRCSetupComplete_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.registeredAMF != nil, ie.guami_Type != nil, len(ie.s_NSSAI_List) > 0, ie.ng_5G_S_TMSI_Value != nil, ie.lateNonCriticalExtension != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.selectedPLMN_Identity, &uper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
		return utils.WrapError("WriteInteger selectedPLMN_Identity", err)
	}
	if ie.registeredAMF != nil {
		if err = ie.registeredAMF.Encode(w); err != nil {
			return utils.WrapError("Encode registeredAMF", err)
		}
	}
	if ie.guami_Type != nil {
		if err = ie.guami_Type.Encode(w); err != nil {
			return utils.WrapError("Encode guami_Type", err)
		}
	}
	if len(ie.s_NSSAI_List) > 0 {
		tmp_s_NSSAI_List := utils.NewSequence[*S_NSSAI]([]*S_NSSAI{}, uper.Constraint{Lb: 1, Ub: maxNrofS_NSSAI}, false)
		for _, i := range ie.s_NSSAI_List {
			tmp_s_NSSAI_List.Value = append(tmp_s_NSSAI_List.Value, &i)
		}
		if err = tmp_s_NSSAI_List.Encode(w); err != nil {
			return utils.WrapError("Encode s_NSSAI_List", err)
		}
	}
	if err = ie.dedicatedNAS_Message.Encode(w); err != nil {
		return utils.WrapError("Encode dedicatedNAS_Message", err)
	}
	if ie.ng_5G_S_TMSI_Value != nil {
		if err = ie.ng_5G_S_TMSI_Value.Encode(w); err != nil {
			return utils.WrapError("Encode ng_5G_S_TMSI_Value", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCSetupComplete_IEs) Decode(r *uper.UperReader) error {
	var err error
	var registeredAMFPresent bool
	if registeredAMFPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var guami_TypePresent bool
	if guami_TypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var s_NSSAI_ListPresent bool
	if s_NSSAI_ListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ng_5G_S_TMSI_ValuePresent bool
	if ng_5G_S_TMSI_ValuePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_selectedPLMN_Identity int64
	if tmp_int_selectedPLMN_Identity, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
		return utils.WrapError("ReadInteger selectedPLMN_Identity", err)
	}
	ie.selectedPLMN_Identity = tmp_int_selectedPLMN_Identity
	if registeredAMFPresent {
		ie.registeredAMF = new(RegisteredAMF)
		if err = ie.registeredAMF.Decode(r); err != nil {
			return utils.WrapError("Decode registeredAMF", err)
		}
	}
	if guami_TypePresent {
		ie.guami_Type = new(RRCSetupComplete_IEs_guami_Type)
		if err = ie.guami_Type.Decode(r); err != nil {
			return utils.WrapError("Decode guami_Type", err)
		}
	}
	if s_NSSAI_ListPresent {
		tmp_s_NSSAI_List := utils.NewSequence[*S_NSSAI]([]*S_NSSAI{}, uper.Constraint{Lb: 1, Ub: maxNrofS_NSSAI}, false)
		fn_s_NSSAI_List := func() *S_NSSAI {
			return new(S_NSSAI)
		}
		if err = tmp_s_NSSAI_List.Decode(r, fn_s_NSSAI_List); err != nil {
			return utils.WrapError("Decode s_NSSAI_List", err)
		}
		ie.s_NSSAI_List = []S_NSSAI{}
		for _, i := range tmp_s_NSSAI_List.Value {
			ie.s_NSSAI_List = append(ie.s_NSSAI_List, *i)
		}
	}
	if err = ie.dedicatedNAS_Message.Decode(r); err != nil {
		return utils.WrapError("Decode dedicatedNAS_Message", err)
	}
	if ng_5G_S_TMSI_ValuePresent {
		ie.ng_5G_S_TMSI_Value = new(RRCSetupComplete_IEs_ng_5G_S_TMSI_Value)
		if err = ie.ng_5G_S_TMSI_Value.Decode(r); err != nil {
			return utils.WrapError("Decode ng_5G_S_TMSI_Value", err)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCSetupComplete_v1610_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
