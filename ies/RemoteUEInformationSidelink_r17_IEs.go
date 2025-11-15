package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RemoteUEInformationSidelink_r17_IEs struct {
	sl_RequestedSIB_List_r17   *SL_RequestedSIB_List_r17   `optional,setuprelease`
	sl_PagingInfo_RemoteUE_r17 *SL_PagingInfo_RemoteUE_r17 `optional,setuprelease`
	lateNonCriticalExtension   *[]byte                     `optional`
	nonCriticalExtension       interface{}                 `optional`
}

func (ie *RemoteUEInformationSidelink_r17_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_RequestedSIB_List_r17 != nil, ie.sl_PagingInfo_RemoteUE_r17 != nil, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_RequestedSIB_List_r17 != nil {
		tmp_sl_RequestedSIB_List_r17 := utils.SetupRelease[*SL_RequestedSIB_List_r17]{
			Setup: ie.sl_RequestedSIB_List_r17,
		}
		if err = tmp_sl_RequestedSIB_List_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RequestedSIB_List_r17", err)
		}
	}
	if ie.sl_PagingInfo_RemoteUE_r17 != nil {
		tmp_sl_PagingInfo_RemoteUE_r17 := utils.SetupRelease[*SL_PagingInfo_RemoteUE_r17]{
			Setup: ie.sl_PagingInfo_RemoteUE_r17,
		}
		if err = tmp_sl_PagingInfo_RemoteUE_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PagingInfo_RemoteUE_r17", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RemoteUEInformationSidelink_r17_IEs) Decode(r *uper.UperReader) error {
	var err error
	var sl_RequestedSIB_List_r17Present bool
	if sl_RequestedSIB_List_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PagingInfo_RemoteUE_r17Present bool
	if sl_PagingInfo_RemoteUE_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_RequestedSIB_List_r17Present {
		tmp_sl_RequestedSIB_List_r17 := utils.SetupRelease[*SL_RequestedSIB_List_r17]{}
		if err = tmp_sl_RequestedSIB_List_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_RequestedSIB_List_r17", err)
		}
		ie.sl_RequestedSIB_List_r17 = tmp_sl_RequestedSIB_List_r17.Setup
	}
	if sl_PagingInfo_RemoteUE_r17Present {
		tmp_sl_PagingInfo_RemoteUE_r17 := utils.SetupRelease[*SL_PagingInfo_RemoteUE_r17]{}
		if err = tmp_sl_PagingInfo_RemoteUE_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PagingInfo_RemoteUE_r17", err)
		}
		ie.sl_PagingInfo_RemoteUE_r17 = tmp_sl_PagingInfo_RemoteUE_r17.Setup
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
