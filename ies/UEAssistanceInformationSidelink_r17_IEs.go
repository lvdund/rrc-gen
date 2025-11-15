package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UEAssistanceInformationSidelink_r17_IEs struct {
	sl_PreferredDRX_ConfigList_r17 []SL_DRX_ConfigUC_SemiStatic_r17 `lb:1,ub:maxNrofSL_RxInfoSet_r17,optional`
	lateNonCriticalExtension       *[]byte                          `optional`
	nonCriticalExtension           interface{}                      `optional`
}

func (ie *UEAssistanceInformationSidelink_r17_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.sl_PreferredDRX_ConfigList_r17) > 0, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.sl_PreferredDRX_ConfigList_r17) > 0 {
		tmp_sl_PreferredDRX_ConfigList_r17 := utils.NewSequence[*SL_DRX_ConfigUC_SemiStatic_r17]([]*SL_DRX_ConfigUC_SemiStatic_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_RxInfoSet_r17}, false)
		for _, i := range ie.sl_PreferredDRX_ConfigList_r17 {
			tmp_sl_PreferredDRX_ConfigList_r17.Value = append(tmp_sl_PreferredDRX_ConfigList_r17.Value, &i)
		}
		if err = tmp_sl_PreferredDRX_ConfigList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PreferredDRX_ConfigList_r17", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UEAssistanceInformationSidelink_r17_IEs) Decode(r *uper.UperReader) error {
	var err error
	var sl_PreferredDRX_ConfigList_r17Present bool
	if sl_PreferredDRX_ConfigList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_PreferredDRX_ConfigList_r17Present {
		tmp_sl_PreferredDRX_ConfigList_r17 := utils.NewSequence[*SL_DRX_ConfigUC_SemiStatic_r17]([]*SL_DRX_ConfigUC_SemiStatic_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_RxInfoSet_r17}, false)
		fn_sl_PreferredDRX_ConfigList_r17 := func() *SL_DRX_ConfigUC_SemiStatic_r17 {
			return new(SL_DRX_ConfigUC_SemiStatic_r17)
		}
		if err = tmp_sl_PreferredDRX_ConfigList_r17.Decode(r, fn_sl_PreferredDRX_ConfigList_r17); err != nil {
			return utils.WrapError("Decode sl_PreferredDRX_ConfigList_r17", err)
		}
		ie.sl_PreferredDRX_ConfigList_r17 = []SL_DRX_ConfigUC_SemiStatic_r17{}
		for _, i := range tmp_sl_PreferredDRX_ConfigList_r17.Value {
			ie.sl_PreferredDRX_ConfigList_r17 = append(ie.sl_PreferredDRX_ConfigList_r17, *i)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	return nil
}
