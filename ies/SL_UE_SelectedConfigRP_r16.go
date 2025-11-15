package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_UE_SelectedConfigRP_r16 struct {
	sl_CBR_PriorityTxConfigList_r16   *SL_CBR_PriorityTxConfigList_r16                        `optional`
	sl_Thres_RSRP_List_r16            *SL_Thres_RSRP_List_r16                                 `optional`
	sl_MultiReserveResource_r16       *SL_UE_SelectedConfigRP_r16_sl_MultiReserveResource_r16 `optional`
	sl_MaxNumPerReserve_r16           *SL_UE_SelectedConfigRP_r16_sl_MaxNumPerReserve_r16     `optional`
	sl_SensingWindow_r16              *SL_UE_SelectedConfigRP_r16_sl_SensingWindow_r16        `optional`
	sl_SelectionWindowList_r16        *SL_SelectionWindowList_r16                             `optional`
	sl_ResourceReservePeriodList_r16  []SL_ResourceReservePeriod_r16                          `lb:1,ub:16,optional`
	sl_RS_ForSensing_r16              SL_UE_SelectedConfigRP_r16_sl_RS_ForSensing_r16         `madatory`
	sl_CBR_PriorityTxConfigList_v1650 *SL_CBR_PriorityTxConfigList_v1650                      `optional,ext-1`
}

func (ie *SL_UE_SelectedConfigRP_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.sl_CBR_PriorityTxConfigList_v1650 != nil
	preambleBits := []bool{hasExtensions, ie.sl_CBR_PriorityTxConfigList_r16 != nil, ie.sl_Thres_RSRP_List_r16 != nil, ie.sl_MultiReserveResource_r16 != nil, ie.sl_MaxNumPerReserve_r16 != nil, ie.sl_SensingWindow_r16 != nil, ie.sl_SelectionWindowList_r16 != nil, len(ie.sl_ResourceReservePeriodList_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_CBR_PriorityTxConfigList_r16 != nil {
		if err = ie.sl_CBR_PriorityTxConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_CBR_PriorityTxConfigList_r16", err)
		}
	}
	if ie.sl_Thres_RSRP_List_r16 != nil {
		if err = ie.sl_Thres_RSRP_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_Thres_RSRP_List_r16", err)
		}
	}
	if ie.sl_MultiReserveResource_r16 != nil {
		if err = ie.sl_MultiReserveResource_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MultiReserveResource_r16", err)
		}
	}
	if ie.sl_MaxNumPerReserve_r16 != nil {
		if err = ie.sl_MaxNumPerReserve_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MaxNumPerReserve_r16", err)
		}
	}
	if ie.sl_SensingWindow_r16 != nil {
		if err = ie.sl_SensingWindow_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_SensingWindow_r16", err)
		}
	}
	if ie.sl_SelectionWindowList_r16 != nil {
		if err = ie.sl_SelectionWindowList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_SelectionWindowList_r16", err)
		}
	}
	if len(ie.sl_ResourceReservePeriodList_r16) > 0 {
		tmp_sl_ResourceReservePeriodList_r16 := utils.NewSequence[*SL_ResourceReservePeriod_r16]([]*SL_ResourceReservePeriod_r16{}, uper.Constraint{Lb: 1, Ub: 16}, false)
		for _, i := range ie.sl_ResourceReservePeriodList_r16 {
			tmp_sl_ResourceReservePeriodList_r16.Value = append(tmp_sl_ResourceReservePeriodList_r16.Value, &i)
		}
		if err = tmp_sl_ResourceReservePeriodList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ResourceReservePeriodList_r16", err)
		}
	}
	if err = ie.sl_RS_ForSensing_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_RS_ForSensing_r16", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.sl_CBR_PriorityTxConfigList_v1650 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SL_UE_SelectedConfigRP_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.sl_CBR_PriorityTxConfigList_v1650 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sl_CBR_PriorityTxConfigList_v1650 optional
			if ie.sl_CBR_PriorityTxConfigList_v1650 != nil {
				if err = ie.sl_CBR_PriorityTxConfigList_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_CBR_PriorityTxConfigList_v1650", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *SL_UE_SelectedConfigRP_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_CBR_PriorityTxConfigList_r16Present bool
	if sl_CBR_PriorityTxConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_Thres_RSRP_List_r16Present bool
	if sl_Thres_RSRP_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MultiReserveResource_r16Present bool
	if sl_MultiReserveResource_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MaxNumPerReserve_r16Present bool
	if sl_MaxNumPerReserve_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_SensingWindow_r16Present bool
	if sl_SensingWindow_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_SelectionWindowList_r16Present bool
	if sl_SelectionWindowList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ResourceReservePeriodList_r16Present bool
	if sl_ResourceReservePeriodList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_CBR_PriorityTxConfigList_r16Present {
		ie.sl_CBR_PriorityTxConfigList_r16 = new(SL_CBR_PriorityTxConfigList_r16)
		if err = ie.sl_CBR_PriorityTxConfigList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_CBR_PriorityTxConfigList_r16", err)
		}
	}
	if sl_Thres_RSRP_List_r16Present {
		ie.sl_Thres_RSRP_List_r16 = new(SL_Thres_RSRP_List_r16)
		if err = ie.sl_Thres_RSRP_List_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_Thres_RSRP_List_r16", err)
		}
	}
	if sl_MultiReserveResource_r16Present {
		ie.sl_MultiReserveResource_r16 = new(SL_UE_SelectedConfigRP_r16_sl_MultiReserveResource_r16)
		if err = ie.sl_MultiReserveResource_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MultiReserveResource_r16", err)
		}
	}
	if sl_MaxNumPerReserve_r16Present {
		ie.sl_MaxNumPerReserve_r16 = new(SL_UE_SelectedConfigRP_r16_sl_MaxNumPerReserve_r16)
		if err = ie.sl_MaxNumPerReserve_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MaxNumPerReserve_r16", err)
		}
	}
	if sl_SensingWindow_r16Present {
		ie.sl_SensingWindow_r16 = new(SL_UE_SelectedConfigRP_r16_sl_SensingWindow_r16)
		if err = ie.sl_SensingWindow_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_SensingWindow_r16", err)
		}
	}
	if sl_SelectionWindowList_r16Present {
		ie.sl_SelectionWindowList_r16 = new(SL_SelectionWindowList_r16)
		if err = ie.sl_SelectionWindowList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_SelectionWindowList_r16", err)
		}
	}
	if sl_ResourceReservePeriodList_r16Present {
		tmp_sl_ResourceReservePeriodList_r16 := utils.NewSequence[*SL_ResourceReservePeriod_r16]([]*SL_ResourceReservePeriod_r16{}, uper.Constraint{Lb: 1, Ub: 16}, false)
		fn_sl_ResourceReservePeriodList_r16 := func() *SL_ResourceReservePeriod_r16 {
			return new(SL_ResourceReservePeriod_r16)
		}
		if err = tmp_sl_ResourceReservePeriodList_r16.Decode(r, fn_sl_ResourceReservePeriodList_r16); err != nil {
			return utils.WrapError("Decode sl_ResourceReservePeriodList_r16", err)
		}
		ie.sl_ResourceReservePeriodList_r16 = []SL_ResourceReservePeriod_r16{}
		for _, i := range tmp_sl_ResourceReservePeriodList_r16.Value {
			ie.sl_ResourceReservePeriodList_r16 = append(ie.sl_ResourceReservePeriodList_r16, *i)
		}
	}
	if err = ie.sl_RS_ForSensing_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_RS_ForSensing_r16", err)
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			sl_CBR_PriorityTxConfigList_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sl_CBR_PriorityTxConfigList_v1650 optional
			if sl_CBR_PriorityTxConfigList_v1650Present {
				ie.sl_CBR_PriorityTxConfigList_v1650 = new(SL_CBR_PriorityTxConfigList_v1650)
				if err = ie.sl_CBR_PriorityTxConfigList_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_CBR_PriorityTxConfigList_v1650", err)
				}
			}
		}
	}
	return nil
}
