package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LTE_NeighCellsCRS_AssistInfo_r17 struct {
	neighCarrierBandwidthDL_r17       *LTE_NeighCellsCRS_AssistInfo_r17_neighCarrierBandwidthDL_r17 `optional`
	neighCarrierFreqDL_r17            *int64                                                        `lb:0,ub:16383,optional`
	neighCellId_r17                   *EUTRA_PhysCellId                                             `optional`
	neighCRS_muting_r17               *LTE_NeighCellsCRS_AssistInfo_r17_neighCRS_muting_r17         `optional`
	neighMBSFN_SubframeConfigList_r17 *EUTRA_MBSFN_SubframeConfigList                               `optional`
	neighNrofCRS_Ports_r17            *LTE_NeighCellsCRS_AssistInfo_r17_neighNrofCRS_Ports_r17      `optional`
	neighV_Shift_r17                  *LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17            `optional`
}

func (ie *LTE_NeighCellsCRS_AssistInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.neighCarrierBandwidthDL_r17 != nil, ie.neighCarrierFreqDL_r17 != nil, ie.neighCellId_r17 != nil, ie.neighCRS_muting_r17 != nil, ie.neighMBSFN_SubframeConfigList_r17 != nil, ie.neighNrofCRS_Ports_r17 != nil, ie.neighV_Shift_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.neighCarrierBandwidthDL_r17 != nil {
		if err = ie.neighCarrierBandwidthDL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode neighCarrierBandwidthDL_r17", err)
		}
	}
	if ie.neighCarrierFreqDL_r17 != nil {
		if err = w.WriteInteger(*ie.neighCarrierFreqDL_r17, &uper.Constraint{Lb: 0, Ub: 16383}, false); err != nil {
			return utils.WrapError("Encode neighCarrierFreqDL_r17", err)
		}
	}
	if ie.neighCellId_r17 != nil {
		if err = ie.neighCellId_r17.Encode(w); err != nil {
			return utils.WrapError("Encode neighCellId_r17", err)
		}
	}
	if ie.neighCRS_muting_r17 != nil {
		if err = ie.neighCRS_muting_r17.Encode(w); err != nil {
			return utils.WrapError("Encode neighCRS_muting_r17", err)
		}
	}
	if ie.neighMBSFN_SubframeConfigList_r17 != nil {
		if err = ie.neighMBSFN_SubframeConfigList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode neighMBSFN_SubframeConfigList_r17", err)
		}
	}
	if ie.neighNrofCRS_Ports_r17 != nil {
		if err = ie.neighNrofCRS_Ports_r17.Encode(w); err != nil {
			return utils.WrapError("Encode neighNrofCRS_Ports_r17", err)
		}
	}
	if ie.neighV_Shift_r17 != nil {
		if err = ie.neighV_Shift_r17.Encode(w); err != nil {
			return utils.WrapError("Encode neighV_Shift_r17", err)
		}
	}
	return nil
}

func (ie *LTE_NeighCellsCRS_AssistInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var neighCarrierBandwidthDL_r17Present bool
	if neighCarrierBandwidthDL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var neighCarrierFreqDL_r17Present bool
	if neighCarrierFreqDL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var neighCellId_r17Present bool
	if neighCellId_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var neighCRS_muting_r17Present bool
	if neighCRS_muting_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var neighMBSFN_SubframeConfigList_r17Present bool
	if neighMBSFN_SubframeConfigList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var neighNrofCRS_Ports_r17Present bool
	if neighNrofCRS_Ports_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var neighV_Shift_r17Present bool
	if neighV_Shift_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if neighCarrierBandwidthDL_r17Present {
		ie.neighCarrierBandwidthDL_r17 = new(LTE_NeighCellsCRS_AssistInfo_r17_neighCarrierBandwidthDL_r17)
		if err = ie.neighCarrierBandwidthDL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode neighCarrierBandwidthDL_r17", err)
		}
	}
	if neighCarrierFreqDL_r17Present {
		var tmp_int_neighCarrierFreqDL_r17 int64
		if tmp_int_neighCarrierFreqDL_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 16383}, false); err != nil {
			return utils.WrapError("Decode neighCarrierFreqDL_r17", err)
		}
		ie.neighCarrierFreqDL_r17 = &tmp_int_neighCarrierFreqDL_r17
	}
	if neighCellId_r17Present {
		ie.neighCellId_r17 = new(EUTRA_PhysCellId)
		if err = ie.neighCellId_r17.Decode(r); err != nil {
			return utils.WrapError("Decode neighCellId_r17", err)
		}
	}
	if neighCRS_muting_r17Present {
		ie.neighCRS_muting_r17 = new(LTE_NeighCellsCRS_AssistInfo_r17_neighCRS_muting_r17)
		if err = ie.neighCRS_muting_r17.Decode(r); err != nil {
			return utils.WrapError("Decode neighCRS_muting_r17", err)
		}
	}
	if neighMBSFN_SubframeConfigList_r17Present {
		ie.neighMBSFN_SubframeConfigList_r17 = new(EUTRA_MBSFN_SubframeConfigList)
		if err = ie.neighMBSFN_SubframeConfigList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode neighMBSFN_SubframeConfigList_r17", err)
		}
	}
	if neighNrofCRS_Ports_r17Present {
		ie.neighNrofCRS_Ports_r17 = new(LTE_NeighCellsCRS_AssistInfo_r17_neighNrofCRS_Ports_r17)
		if err = ie.neighNrofCRS_Ports_r17.Decode(r); err != nil {
			return utils.WrapError("Decode neighNrofCRS_Ports_r17", err)
		}
	}
	if neighV_Shift_r17Present {
		ie.neighV_Shift_r17 = new(LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17)
		if err = ie.neighV_Shift_r17.Decode(r); err != nil {
			return utils.WrapError("Decode neighV_Shift_r17", err)
		}
	}
	return nil
}
