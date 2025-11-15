package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParametersSidelinkEUTRA_NR_v1630_nr struct {
	tx_Sidelink_r16               *BandParametersSidelinkEUTRA_NR_v1630_nr_tx_Sidelink_r16               `optional`
	rx_Sidelink_r16               *BandParametersSidelinkEUTRA_NR_v1630_nr_rx_Sidelink_r16               `optional`
	sl_CrossCarrierScheduling_r16 *BandParametersSidelinkEUTRA_NR_v1630_nr_sl_CrossCarrierScheduling_r16 `optional`
}

func (ie *BandParametersSidelinkEUTRA_NR_v1630_nr) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.tx_Sidelink_r16 != nil, ie.rx_Sidelink_r16 != nil, ie.sl_CrossCarrierScheduling_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.tx_Sidelink_r16 != nil {
		if err = ie.tx_Sidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode tx_Sidelink_r16", err)
		}
	}
	if ie.rx_Sidelink_r16 != nil {
		if err = ie.rx_Sidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode rx_Sidelink_r16", err)
		}
	}
	if ie.sl_CrossCarrierScheduling_r16 != nil {
		if err = ie.sl_CrossCarrierScheduling_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_CrossCarrierScheduling_r16", err)
		}
	}
	return nil
}

func (ie *BandParametersSidelinkEUTRA_NR_v1630_nr) Decode(r *uper.UperReader) error {
	var err error
	var tx_Sidelink_r16Present bool
	if tx_Sidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rx_Sidelink_r16Present bool
	if rx_Sidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_CrossCarrierScheduling_r16Present bool
	if sl_CrossCarrierScheduling_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if tx_Sidelink_r16Present {
		ie.tx_Sidelink_r16 = new(BandParametersSidelinkEUTRA_NR_v1630_nr_tx_Sidelink_r16)
		if err = ie.tx_Sidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode tx_Sidelink_r16", err)
		}
	}
	if rx_Sidelink_r16Present {
		ie.rx_Sidelink_r16 = new(BandParametersSidelinkEUTRA_NR_v1630_nr_rx_Sidelink_r16)
		if err = ie.rx_Sidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rx_Sidelink_r16", err)
		}
	}
	if sl_CrossCarrierScheduling_r16Present {
		ie.sl_CrossCarrierScheduling_r16 = new(BandParametersSidelinkEUTRA_NR_v1630_nr_sl_CrossCarrierScheduling_r16)
		if err = ie.sl_CrossCarrierScheduling_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_CrossCarrierScheduling_r16", err)
		}
	}
	return nil
}
