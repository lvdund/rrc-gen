package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqNeighCellInfo_v1710 struct {
	ssb_PositionQCL_r17 *SSB_PositionQCL_Relation_r17 `optional`
}

func (ie *InterFreqNeighCellInfo_v1710) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ssb_PositionQCL_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ssb_PositionQCL_r17 != nil {
		if err = ie.ssb_PositionQCL_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_PositionQCL_r17", err)
		}
	}
	return nil
}

func (ie *InterFreqNeighCellInfo_v1710) Decode(r *uper.UperReader) error {
	var err error
	var ssb_PositionQCL_r17Present bool
	if ssb_PositionQCL_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ssb_PositionQCL_r17Present {
		ie.ssb_PositionQCL_r17 = new(SSB_PositionQCL_Relation_r17)
		if err = ie.ssb_PositionQCL_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_PositionQCL_r17", err)
		}
	}
	return nil
}
