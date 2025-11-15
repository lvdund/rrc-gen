package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_TxResourceReqCommRelay_r17_Choice_nothing uint64 = iota
	SL_TxResourceReqCommRelay_r17_Choice_sl_TxResourceReqL2U2N_Relay_r17
	SL_TxResourceReqCommRelay_r17_Choice_sl_TxResourceReqL3U2N_Relay_r17
)

type SL_TxResourceReqCommRelay_r17 struct {
	Choice                          uint64
	sl_TxResourceReqL2U2N_Relay_r17 *SL_TxResourceReqL2U2N_Relay_r17
	sl_TxResourceReqL3U2N_Relay_r17 *SL_TxResourceReq_r16
}

func (ie *SL_TxResourceReqCommRelay_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_TxResourceReqCommRelay_r17_Choice_sl_TxResourceReqL2U2N_Relay_r17:
		if err = ie.sl_TxResourceReqL2U2N_Relay_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode sl_TxResourceReqL2U2N_Relay_r17", err)
		}
	case SL_TxResourceReqCommRelay_r17_Choice_sl_TxResourceReqL3U2N_Relay_r17:
		if err = ie.sl_TxResourceReqL3U2N_Relay_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode sl_TxResourceReqL3U2N_Relay_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SL_TxResourceReqCommRelay_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_TxResourceReqCommRelay_r17_Choice_sl_TxResourceReqL2U2N_Relay_r17:
		ie.sl_TxResourceReqL2U2N_Relay_r17 = new(SL_TxResourceReqL2U2N_Relay_r17)
		if err = ie.sl_TxResourceReqL2U2N_Relay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TxResourceReqL2U2N_Relay_r17", err)
		}
	case SL_TxResourceReqCommRelay_r17_Choice_sl_TxResourceReqL3U2N_Relay_r17:
		ie.sl_TxResourceReqL3U2N_Relay_r17 = new(SL_TxResourceReq_r16)
		if err = ie.sl_TxResourceReqL3U2N_Relay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TxResourceReqL3U2N_Relay_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
