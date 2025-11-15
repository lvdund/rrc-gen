package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_TxResourceReqCommRelayInfo_r17 struct {
	sl_RelayDRXConfig_r17         *SL_TxResourceReq_v1700       `optional`
	sl_TxResourceReqCommRelay_r17 SL_TxResourceReqCommRelay_r17 `madatory`
}

func (ie *SL_TxResourceReqCommRelayInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_RelayDRXConfig_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_RelayDRXConfig_r17 != nil {
		if err = ie.sl_RelayDRXConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RelayDRXConfig_r17", err)
		}
	}
	if err = ie.sl_TxResourceReqCommRelay_r17.Encode(w); err != nil {
		return utils.WrapError("Encode sl_TxResourceReqCommRelay_r17", err)
	}
	return nil
}

func (ie *SL_TxResourceReqCommRelayInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var sl_RelayDRXConfig_r17Present bool
	if sl_RelayDRXConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_RelayDRXConfig_r17Present {
		ie.sl_RelayDRXConfig_r17 = new(SL_TxResourceReq_v1700)
		if err = ie.sl_RelayDRXConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_RelayDRXConfig_r17", err)
		}
	}
	if err = ie.sl_TxResourceReqCommRelay_r17.Decode(r); err != nil {
		return utils.WrapError("Decode sl_TxResourceReqCommRelay_r17", err)
	}
	return nil
}
