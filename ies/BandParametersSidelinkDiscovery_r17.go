package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParametersSidelinkDiscovery_r17 struct {
	sl_CrossCarrierScheduling_r17           *BandParametersSidelinkDiscovery_r17_sl_CrossCarrierScheduling_r17           `optional`
	sl_TransmissionMode2_PartialSensing_r17 *BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17 `optional`
	tx_IUC_Scheme1_Mode2Sidelink_r17        *BandParametersSidelinkDiscovery_r17_tx_IUC_Scheme1_Mode2Sidelink_r17        `optional`
}

func (ie *BandParametersSidelinkDiscovery_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_CrossCarrierScheduling_r17 != nil, ie.sl_TransmissionMode2_PartialSensing_r17 != nil, ie.tx_IUC_Scheme1_Mode2Sidelink_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_CrossCarrierScheduling_r17 != nil {
		if err = ie.sl_CrossCarrierScheduling_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_CrossCarrierScheduling_r17", err)
		}
	}
	if ie.sl_TransmissionMode2_PartialSensing_r17 != nil {
		if err = ie.sl_TransmissionMode2_PartialSensing_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TransmissionMode2_PartialSensing_r17", err)
		}
	}
	if ie.tx_IUC_Scheme1_Mode2Sidelink_r17 != nil {
		if err = ie.tx_IUC_Scheme1_Mode2Sidelink_r17.Encode(w); err != nil {
			return utils.WrapError("Encode tx_IUC_Scheme1_Mode2Sidelink_r17", err)
		}
	}
	return nil
}

func (ie *BandParametersSidelinkDiscovery_r17) Decode(r *uper.UperReader) error {
	var err error
	var sl_CrossCarrierScheduling_r17Present bool
	if sl_CrossCarrierScheduling_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TransmissionMode2_PartialSensing_r17Present bool
	if sl_TransmissionMode2_PartialSensing_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tx_IUC_Scheme1_Mode2Sidelink_r17Present bool
	if tx_IUC_Scheme1_Mode2Sidelink_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_CrossCarrierScheduling_r17Present {
		ie.sl_CrossCarrierScheduling_r17 = new(BandParametersSidelinkDiscovery_r17_sl_CrossCarrierScheduling_r17)
		if err = ie.sl_CrossCarrierScheduling_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_CrossCarrierScheduling_r17", err)
		}
	}
	if sl_TransmissionMode2_PartialSensing_r17Present {
		ie.sl_TransmissionMode2_PartialSensing_r17 = new(BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17)
		if err = ie.sl_TransmissionMode2_PartialSensing_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TransmissionMode2_PartialSensing_r17", err)
		}
	}
	if tx_IUC_Scheme1_Mode2Sidelink_r17Present {
		ie.tx_IUC_Scheme1_Mode2Sidelink_r17 = new(BandParametersSidelinkDiscovery_r17_tx_IUC_Scheme1_Mode2Sidelink_r17)
		if err = ie.tx_IUC_Scheme1_Mode2Sidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode tx_IUC_Scheme1_Mode2Sidelink_r17", err)
		}
	}
	return nil
}
