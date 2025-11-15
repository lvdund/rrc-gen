package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SidelinkUEInformationNR_v1700_IEs struct {
	sl_TxResourceReqList_v1700        *SL_TxResourceReqList_v1700                    `optional`
	sl_RxDRX_ReportList_v1700         *SL_RxDRX_ReportList_v1700                     `optional`
	sl_RxInterestedGC_BC_DestList_r17 *SL_RxInterestedGC_BC_DestList_r17             `optional`
	sl_RxInterestedFreqListDisc_r17   *SL_InterestedFreqList_r16                     `optional`
	sl_TxResourceReqListDisc_r17      *SL_TxResourceReqListDisc_r17                  `optional`
	sl_TxResourceReqListCommRelay_r17 *SL_TxResourceReqListCommRelay_r17             `optional`
	ue_Type_r17                       *SidelinkUEInformationNR_v1700_IEs_ue_Type_r17 `optional`
	sl_SourceIdentityRemoteUE_r17     *SL_SourceIdentity_r17                         `optional`
	nonCriticalExtension              interface{}                                    `optional`
}

func (ie *SidelinkUEInformationNR_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_TxResourceReqList_v1700 != nil, ie.sl_RxDRX_ReportList_v1700 != nil, ie.sl_RxInterestedGC_BC_DestList_r17 != nil, ie.sl_RxInterestedFreqListDisc_r17 != nil, ie.sl_TxResourceReqListDisc_r17 != nil, ie.sl_TxResourceReqListCommRelay_r17 != nil, ie.ue_Type_r17 != nil, ie.sl_SourceIdentityRemoteUE_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_TxResourceReqList_v1700 != nil {
		if err = ie.sl_TxResourceReqList_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TxResourceReqList_v1700", err)
		}
	}
	if ie.sl_RxDRX_ReportList_v1700 != nil {
		if err = ie.sl_RxDRX_ReportList_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RxDRX_ReportList_v1700", err)
		}
	}
	if ie.sl_RxInterestedGC_BC_DestList_r17 != nil {
		if err = ie.sl_RxInterestedGC_BC_DestList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RxInterestedGC_BC_DestList_r17", err)
		}
	}
	if ie.sl_RxInterestedFreqListDisc_r17 != nil {
		if err = ie.sl_RxInterestedFreqListDisc_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RxInterestedFreqListDisc_r17", err)
		}
	}
	if ie.sl_TxResourceReqListDisc_r17 != nil {
		if err = ie.sl_TxResourceReqListDisc_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TxResourceReqListDisc_r17", err)
		}
	}
	if ie.sl_TxResourceReqListCommRelay_r17 != nil {
		if err = ie.sl_TxResourceReqListCommRelay_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TxResourceReqListCommRelay_r17", err)
		}
	}
	if ie.ue_Type_r17 != nil {
		if err = ie.ue_Type_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ue_Type_r17", err)
		}
	}
	if ie.sl_SourceIdentityRemoteUE_r17 != nil {
		if err = ie.sl_SourceIdentityRemoteUE_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_SourceIdentityRemoteUE_r17", err)
		}
	}
	return nil
}

func (ie *SidelinkUEInformationNR_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var sl_TxResourceReqList_v1700Present bool
	if sl_TxResourceReqList_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_RxDRX_ReportList_v1700Present bool
	if sl_RxDRX_ReportList_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_RxInterestedGC_BC_DestList_r17Present bool
	if sl_RxInterestedGC_BC_DestList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_RxInterestedFreqListDisc_r17Present bool
	if sl_RxInterestedFreqListDisc_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TxResourceReqListDisc_r17Present bool
	if sl_TxResourceReqListDisc_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TxResourceReqListCommRelay_r17Present bool
	if sl_TxResourceReqListCommRelay_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ue_Type_r17Present bool
	if ue_Type_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_SourceIdentityRemoteUE_r17Present bool
	if sl_SourceIdentityRemoteUE_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_TxResourceReqList_v1700Present {
		ie.sl_TxResourceReqList_v1700 = new(SL_TxResourceReqList_v1700)
		if err = ie.sl_TxResourceReqList_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TxResourceReqList_v1700", err)
		}
	}
	if sl_RxDRX_ReportList_v1700Present {
		ie.sl_RxDRX_ReportList_v1700 = new(SL_RxDRX_ReportList_v1700)
		if err = ie.sl_RxDRX_ReportList_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode sl_RxDRX_ReportList_v1700", err)
		}
	}
	if sl_RxInterestedGC_BC_DestList_r17Present {
		ie.sl_RxInterestedGC_BC_DestList_r17 = new(SL_RxInterestedGC_BC_DestList_r17)
		if err = ie.sl_RxInterestedGC_BC_DestList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_RxInterestedGC_BC_DestList_r17", err)
		}
	}
	if sl_RxInterestedFreqListDisc_r17Present {
		ie.sl_RxInterestedFreqListDisc_r17 = new(SL_InterestedFreqList_r16)
		if err = ie.sl_RxInterestedFreqListDisc_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_RxInterestedFreqListDisc_r17", err)
		}
	}
	if sl_TxResourceReqListDisc_r17Present {
		ie.sl_TxResourceReqListDisc_r17 = new(SL_TxResourceReqListDisc_r17)
		if err = ie.sl_TxResourceReqListDisc_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TxResourceReqListDisc_r17", err)
		}
	}
	if sl_TxResourceReqListCommRelay_r17Present {
		ie.sl_TxResourceReqListCommRelay_r17 = new(SL_TxResourceReqListCommRelay_r17)
		if err = ie.sl_TxResourceReqListCommRelay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TxResourceReqListCommRelay_r17", err)
		}
	}
	if ue_Type_r17Present {
		ie.ue_Type_r17 = new(SidelinkUEInformationNR_v1700_IEs_ue_Type_r17)
		if err = ie.ue_Type_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ue_Type_r17", err)
		}
	}
	if sl_SourceIdentityRemoteUE_r17Present {
		ie.sl_SourceIdentityRemoteUE_r17 = new(SL_SourceIdentity_r17)
		if err = ie.sl_SourceIdentityRemoteUE_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_SourceIdentityRemoteUE_r17", err)
		}
	}
	return nil
}
