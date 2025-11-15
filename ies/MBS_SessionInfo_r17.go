package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MBS_SessionInfo_r17 struct {
	mbs_SessionId_r17               TMGI_r17                         `madatory`
	g_RNTI_r17                      RNTI_Value                       `madatory`
	mrb_ListBroadcast_r17           MRB_ListBroadcast_r17            `madatory`
	mtch_SchedulingInfo_r17         *DRX_ConfigPTM_Index_r17         `optional`
	mtch_NeighbourCell_r17          *uper.BitString                  `lb:maxNeighCellMBS_r17,ub:maxNeighCellMBS_r17,optional`
	pdsch_ConfigIndex_r17           *PDSCH_ConfigIndex_r17           `optional`
	mtch_SSB_MappingWindowIndex_r17 *MTCH_SSB_MappingWindowIndex_r17 `optional`
}

func (ie *MBS_SessionInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.mtch_SchedulingInfo_r17 != nil, ie.mtch_NeighbourCell_r17 != nil, ie.pdsch_ConfigIndex_r17 != nil, ie.mtch_SSB_MappingWindowIndex_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.mbs_SessionId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode mbs_SessionId_r17", err)
	}
	if err = ie.g_RNTI_r17.Encode(w); err != nil {
		return utils.WrapError("Encode g_RNTI_r17", err)
	}
	if err = ie.mrb_ListBroadcast_r17.Encode(w); err != nil {
		return utils.WrapError("Encode mrb_ListBroadcast_r17", err)
	}
	if ie.mtch_SchedulingInfo_r17 != nil {
		if err = ie.mtch_SchedulingInfo_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mtch_SchedulingInfo_r17", err)
		}
	}
	if ie.mtch_NeighbourCell_r17 != nil {
		if err = w.WriteBitString(ie.mtch_NeighbourCell_r17.Bytes, uint(ie.mtch_NeighbourCell_r17.NumBits), &uper.Constraint{Lb: maxNeighCellMBS_r17, Ub: maxNeighCellMBS_r17}, false); err != nil {
			return utils.WrapError("Encode mtch_NeighbourCell_r17", err)
		}
	}
	if ie.pdsch_ConfigIndex_r17 != nil {
		if err = ie.pdsch_ConfigIndex_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_ConfigIndex_r17", err)
		}
	}
	if ie.mtch_SSB_MappingWindowIndex_r17 != nil {
		if err = ie.mtch_SSB_MappingWindowIndex_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mtch_SSB_MappingWindowIndex_r17", err)
		}
	}
	return nil
}

func (ie *MBS_SessionInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var mtch_SchedulingInfo_r17Present bool
	if mtch_SchedulingInfo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mtch_NeighbourCell_r17Present bool
	if mtch_NeighbourCell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_ConfigIndex_r17Present bool
	if pdsch_ConfigIndex_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mtch_SSB_MappingWindowIndex_r17Present bool
	if mtch_SSB_MappingWindowIndex_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.mbs_SessionId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode mbs_SessionId_r17", err)
	}
	if err = ie.g_RNTI_r17.Decode(r); err != nil {
		return utils.WrapError("Decode g_RNTI_r17", err)
	}
	if err = ie.mrb_ListBroadcast_r17.Decode(r); err != nil {
		return utils.WrapError("Decode mrb_ListBroadcast_r17", err)
	}
	if mtch_SchedulingInfo_r17Present {
		ie.mtch_SchedulingInfo_r17 = new(DRX_ConfigPTM_Index_r17)
		if err = ie.mtch_SchedulingInfo_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mtch_SchedulingInfo_r17", err)
		}
	}
	if mtch_NeighbourCell_r17Present {
		var tmp_bs_mtch_NeighbourCell_r17 []byte
		var n_mtch_NeighbourCell_r17 uint
		if tmp_bs_mtch_NeighbourCell_r17, n_mtch_NeighbourCell_r17, err = r.ReadBitString(&uper.Constraint{Lb: maxNeighCellMBS_r17, Ub: maxNeighCellMBS_r17}, false); err != nil {
			return utils.WrapError("Decode mtch_NeighbourCell_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_mtch_NeighbourCell_r17,
			NumBits: uint64(n_mtch_NeighbourCell_r17),
		}
		ie.mtch_NeighbourCell_r17 = &tmp_bitstring
	}
	if pdsch_ConfigIndex_r17Present {
		ie.pdsch_ConfigIndex_r17 = new(PDSCH_ConfigIndex_r17)
		if err = ie.pdsch_ConfigIndex_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_ConfigIndex_r17", err)
		}
	}
	if mtch_SSB_MappingWindowIndex_r17Present {
		ie.mtch_SSB_MappingWindowIndex_r17 = new(MTCH_SSB_MappingWindowIndex_r17)
		if err = ie.mtch_SSB_MappingWindowIndex_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mtch_SSB_MappingWindowIndex_r17", err)
		}
	}
	return nil
}
