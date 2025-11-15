package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRB_PDCP_ConfigBroadcast_r17 struct {
	pdcp_SN_SizeDL_r17    *MRB_PDCP_ConfigBroadcast_r17_pdcp_SN_SizeDL_r17   `optional`
	headerCompression_r17 MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17 `lb:1,ub:16,madatory`
	t_Reordering_r17      *MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17     `optional`
}

func (ie *MRB_PDCP_ConfigBroadcast_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pdcp_SN_SizeDL_r17 != nil, ie.t_Reordering_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pdcp_SN_SizeDL_r17 != nil {
		if err = ie.pdcp_SN_SizeDL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdcp_SN_SizeDL_r17", err)
		}
	}
	if err = ie.headerCompression_r17.Encode(w); err != nil {
		return utils.WrapError("Encode headerCompression_r17", err)
	}
	if ie.t_Reordering_r17 != nil {
		if err = ie.t_Reordering_r17.Encode(w); err != nil {
			return utils.WrapError("Encode t_Reordering_r17", err)
		}
	}
	return nil
}

func (ie *MRB_PDCP_ConfigBroadcast_r17) Decode(r *uper.UperReader) error {
	var err error
	var pdcp_SN_SizeDL_r17Present bool
	if pdcp_SN_SizeDL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var t_Reordering_r17Present bool
	if t_Reordering_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if pdcp_SN_SizeDL_r17Present {
		ie.pdcp_SN_SizeDL_r17 = new(MRB_PDCP_ConfigBroadcast_r17_pdcp_SN_SizeDL_r17)
		if err = ie.pdcp_SN_SizeDL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdcp_SN_SizeDL_r17", err)
		}
	}
	if err = ie.headerCompression_r17.Decode(r); err != nil {
		return utils.WrapError("Decode headerCompression_r17", err)
	}
	if t_Reordering_r17Present {
		ie.t_Reordering_r17 = new(MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17)
		if err = ie.t_Reordering_r17.Decode(r); err != nil {
			return utils.WrapError("Decode t_Reordering_r17", err)
		}
	}
	return nil
}
