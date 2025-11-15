package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_TxResourceReqDisc_r17 struct {
	sl_DestinationIdentityDisc_r17  SL_DestinationIdentity_r16                    `madatory`
	sl_SourceIdentityRelayUE_r17    *SL_SourceIdentity_r17                        `optional`
	sl_CastTypeDisc_r17             SL_TxResourceReqDisc_r17_sl_CastTypeDisc_r17  `madatory`
	sl_TxInterestedFreqListDisc_r17 SL_TxInterestedFreqList_r16                   `madatory`
	sl_TypeTxSyncListDisc_r17       []SL_TypeTxSync_r16                           `lb:1,ub:maxNrofFreqSL_r16,madatory`
	sl_DiscoveryType_r17            SL_TxResourceReqDisc_r17_sl_DiscoveryType_r17 `madatory`
}

func (ie *SL_TxResourceReqDisc_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_SourceIdentityRelayUE_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.sl_DestinationIdentityDisc_r17.Encode(w); err != nil {
		return utils.WrapError("Encode sl_DestinationIdentityDisc_r17", err)
	}
	if ie.sl_SourceIdentityRelayUE_r17 != nil {
		if err = ie.sl_SourceIdentityRelayUE_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_SourceIdentityRelayUE_r17", err)
		}
	}
	if err = ie.sl_CastTypeDisc_r17.Encode(w); err != nil {
		return utils.WrapError("Encode sl_CastTypeDisc_r17", err)
	}
	if err = ie.sl_TxInterestedFreqListDisc_r17.Encode(w); err != nil {
		return utils.WrapError("Encode sl_TxInterestedFreqListDisc_r17", err)
	}
	tmp_sl_TypeTxSyncListDisc_r17 := utils.NewSequence[*SL_TypeTxSync_r16]([]*SL_TypeTxSync_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
	for _, i := range ie.sl_TypeTxSyncListDisc_r17 {
		tmp_sl_TypeTxSyncListDisc_r17.Value = append(tmp_sl_TypeTxSyncListDisc_r17.Value, &i)
	}
	if err = tmp_sl_TypeTxSyncListDisc_r17.Encode(w); err != nil {
		return utils.WrapError("Encode sl_TypeTxSyncListDisc_r17", err)
	}
	if err = ie.sl_DiscoveryType_r17.Encode(w); err != nil {
		return utils.WrapError("Encode sl_DiscoveryType_r17", err)
	}
	return nil
}

func (ie *SL_TxResourceReqDisc_r17) Decode(r *uper.UperReader) error {
	var err error
	var sl_SourceIdentityRelayUE_r17Present bool
	if sl_SourceIdentityRelayUE_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.sl_DestinationIdentityDisc_r17.Decode(r); err != nil {
		return utils.WrapError("Decode sl_DestinationIdentityDisc_r17", err)
	}
	if sl_SourceIdentityRelayUE_r17Present {
		ie.sl_SourceIdentityRelayUE_r17 = new(SL_SourceIdentity_r17)
		if err = ie.sl_SourceIdentityRelayUE_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_SourceIdentityRelayUE_r17", err)
		}
	}
	if err = ie.sl_CastTypeDisc_r17.Decode(r); err != nil {
		return utils.WrapError("Decode sl_CastTypeDisc_r17", err)
	}
	if err = ie.sl_TxInterestedFreqListDisc_r17.Decode(r); err != nil {
		return utils.WrapError("Decode sl_TxInterestedFreqListDisc_r17", err)
	}
	tmp_sl_TypeTxSyncListDisc_r17 := utils.NewSequence[*SL_TypeTxSync_r16]([]*SL_TypeTxSync_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
	fn_sl_TypeTxSyncListDisc_r17 := func() *SL_TypeTxSync_r16 {
		return new(SL_TypeTxSync_r16)
	}
	if err = tmp_sl_TypeTxSyncListDisc_r17.Decode(r, fn_sl_TypeTxSyncListDisc_r17); err != nil {
		return utils.WrapError("Decode sl_TypeTxSyncListDisc_r17", err)
	}
	ie.sl_TypeTxSyncListDisc_r17 = []SL_TypeTxSync_r16{}
	for _, i := range tmp_sl_TypeTxSyncListDisc_r17.Value {
		ie.sl_TypeTxSyncListDisc_r17 = append(ie.sl_TypeTxSyncListDisc_r17, *i)
	}
	if err = ie.sl_DiscoveryType_r17.Decode(r); err != nil {
		return utils.WrapError("Decode sl_DiscoveryType_r17", err)
	}
	return nil
}
