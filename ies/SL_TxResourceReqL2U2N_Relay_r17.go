package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_TxResourceReqL2U2N_Relay_r17 struct {
	sl_DestinationIdentityL2U2N_r17      *SL_DestinationIdentity_r16                             `optional`
	sl_TxInterestedFreqListL2U2N_r17     SL_TxInterestedFreqList_r16                             `madatory`
	sl_TypeTxSyncListL2U2N_r17           []SL_TypeTxSync_r16                                     `lb:1,ub:maxNrofFreqSL_r16,madatory`
	sl_LocalID_Request_r17               *SL_TxResourceReqL2U2N_Relay_r17_sl_LocalID_Request_r17 `optional`
	sl_PagingIdentityRemoteUE_r17        *SL_PagingIdentityRemoteUE_r17                          `optional`
	sl_CapabilityInformationSidelink_r17 *[]byte                                                 `optional`
}

func (ie *SL_TxResourceReqL2U2N_Relay_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_DestinationIdentityL2U2N_r17 != nil, ie.sl_LocalID_Request_r17 != nil, ie.sl_PagingIdentityRemoteUE_r17 != nil, ie.sl_CapabilityInformationSidelink_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_DestinationIdentityL2U2N_r17 != nil {
		if err = ie.sl_DestinationIdentityL2U2N_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_DestinationIdentityL2U2N_r17", err)
		}
	}
	if err = ie.sl_TxInterestedFreqListL2U2N_r17.Encode(w); err != nil {
		return utils.WrapError("Encode sl_TxInterestedFreqListL2U2N_r17", err)
	}
	tmp_sl_TypeTxSyncListL2U2N_r17 := utils.NewSequence[*SL_TypeTxSync_r16]([]*SL_TypeTxSync_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
	for _, i := range ie.sl_TypeTxSyncListL2U2N_r17 {
		tmp_sl_TypeTxSyncListL2U2N_r17.Value = append(tmp_sl_TypeTxSyncListL2U2N_r17.Value, &i)
	}
	if err = tmp_sl_TypeTxSyncListL2U2N_r17.Encode(w); err != nil {
		return utils.WrapError("Encode sl_TypeTxSyncListL2U2N_r17", err)
	}
	if ie.sl_LocalID_Request_r17 != nil {
		if err = ie.sl_LocalID_Request_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_LocalID_Request_r17", err)
		}
	}
	if ie.sl_PagingIdentityRemoteUE_r17 != nil {
		if err = ie.sl_PagingIdentityRemoteUE_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PagingIdentityRemoteUE_r17", err)
		}
	}
	if ie.sl_CapabilityInformationSidelink_r17 != nil {
		if err = w.WriteOctetString(*ie.sl_CapabilityInformationSidelink_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode sl_CapabilityInformationSidelink_r17", err)
		}
	}
	return nil
}

func (ie *SL_TxResourceReqL2U2N_Relay_r17) Decode(r *uper.UperReader) error {
	var err error
	var sl_DestinationIdentityL2U2N_r17Present bool
	if sl_DestinationIdentityL2U2N_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_LocalID_Request_r17Present bool
	if sl_LocalID_Request_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PagingIdentityRemoteUE_r17Present bool
	if sl_PagingIdentityRemoteUE_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_CapabilityInformationSidelink_r17Present bool
	if sl_CapabilityInformationSidelink_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_DestinationIdentityL2U2N_r17Present {
		ie.sl_DestinationIdentityL2U2N_r17 = new(SL_DestinationIdentity_r16)
		if err = ie.sl_DestinationIdentityL2U2N_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_DestinationIdentityL2U2N_r17", err)
		}
	}
	if err = ie.sl_TxInterestedFreqListL2U2N_r17.Decode(r); err != nil {
		return utils.WrapError("Decode sl_TxInterestedFreqListL2U2N_r17", err)
	}
	tmp_sl_TypeTxSyncListL2U2N_r17 := utils.NewSequence[*SL_TypeTxSync_r16]([]*SL_TypeTxSync_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
	fn_sl_TypeTxSyncListL2U2N_r17 := func() *SL_TypeTxSync_r16 {
		return new(SL_TypeTxSync_r16)
	}
	if err = tmp_sl_TypeTxSyncListL2U2N_r17.Decode(r, fn_sl_TypeTxSyncListL2U2N_r17); err != nil {
		return utils.WrapError("Decode sl_TypeTxSyncListL2U2N_r17", err)
	}
	ie.sl_TypeTxSyncListL2U2N_r17 = []SL_TypeTxSync_r16{}
	for _, i := range tmp_sl_TypeTxSyncListL2U2N_r17.Value {
		ie.sl_TypeTxSyncListL2U2N_r17 = append(ie.sl_TypeTxSyncListL2U2N_r17, *i)
	}
	if sl_LocalID_Request_r17Present {
		ie.sl_LocalID_Request_r17 = new(SL_TxResourceReqL2U2N_Relay_r17_sl_LocalID_Request_r17)
		if err = ie.sl_LocalID_Request_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_LocalID_Request_r17", err)
		}
	}
	if sl_PagingIdentityRemoteUE_r17Present {
		ie.sl_PagingIdentityRemoteUE_r17 = new(SL_PagingIdentityRemoteUE_r17)
		if err = ie.sl_PagingIdentityRemoteUE_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PagingIdentityRemoteUE_r17", err)
		}
	}
	if sl_CapabilityInformationSidelink_r17Present {
		var tmp_os_sl_CapabilityInformationSidelink_r17 []byte
		if tmp_os_sl_CapabilityInformationSidelink_r17, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode sl_CapabilityInformationSidelink_r17", err)
		}
		ie.sl_CapabilityInformationSidelink_r17 = &tmp_os_sl_CapabilityInformationSidelink_r17
	}
	return nil
}
