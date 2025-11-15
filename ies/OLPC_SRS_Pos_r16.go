package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type OLPC_SRS_Pos_r16 struct {
	olpc_SRS_PosBasedOnPRS_Serving_r16      *OLPC_SRS_Pos_r16_olpc_SRS_PosBasedOnPRS_Serving_r16      `optional`
	olpc_SRS_PosBasedOnSSB_Neigh_r16        *OLPC_SRS_Pos_r16_olpc_SRS_PosBasedOnSSB_Neigh_r16        `optional`
	olpc_SRS_PosBasedOnPRS_Neigh_r16        *OLPC_SRS_Pos_r16_olpc_SRS_PosBasedOnPRS_Neigh_r16        `optional`
	maxNumberPathLossEstimatePerServing_r16 *OLPC_SRS_Pos_r16_maxNumberPathLossEstimatePerServing_r16 `optional`
}

func (ie *OLPC_SRS_Pos_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.olpc_SRS_PosBasedOnPRS_Serving_r16 != nil, ie.olpc_SRS_PosBasedOnSSB_Neigh_r16 != nil, ie.olpc_SRS_PosBasedOnPRS_Neigh_r16 != nil, ie.maxNumberPathLossEstimatePerServing_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.olpc_SRS_PosBasedOnPRS_Serving_r16 != nil {
		if err = ie.olpc_SRS_PosBasedOnPRS_Serving_r16.Encode(w); err != nil {
			return utils.WrapError("Encode olpc_SRS_PosBasedOnPRS_Serving_r16", err)
		}
	}
	if ie.olpc_SRS_PosBasedOnSSB_Neigh_r16 != nil {
		if err = ie.olpc_SRS_PosBasedOnSSB_Neigh_r16.Encode(w); err != nil {
			return utils.WrapError("Encode olpc_SRS_PosBasedOnSSB_Neigh_r16", err)
		}
	}
	if ie.olpc_SRS_PosBasedOnPRS_Neigh_r16 != nil {
		if err = ie.olpc_SRS_PosBasedOnPRS_Neigh_r16.Encode(w); err != nil {
			return utils.WrapError("Encode olpc_SRS_PosBasedOnPRS_Neigh_r16", err)
		}
	}
	if ie.maxNumberPathLossEstimatePerServing_r16 != nil {
		if err = ie.maxNumberPathLossEstimatePerServing_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumberPathLossEstimatePerServing_r16", err)
		}
	}
	return nil
}

func (ie *OLPC_SRS_Pos_r16) Decode(r *uper.UperReader) error {
	var err error
	var olpc_SRS_PosBasedOnPRS_Serving_r16Present bool
	if olpc_SRS_PosBasedOnPRS_Serving_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var olpc_SRS_PosBasedOnSSB_Neigh_r16Present bool
	if olpc_SRS_PosBasedOnSSB_Neigh_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var olpc_SRS_PosBasedOnPRS_Neigh_r16Present bool
	if olpc_SRS_PosBasedOnPRS_Neigh_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumberPathLossEstimatePerServing_r16Present bool
	if maxNumberPathLossEstimatePerServing_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if olpc_SRS_PosBasedOnPRS_Serving_r16Present {
		ie.olpc_SRS_PosBasedOnPRS_Serving_r16 = new(OLPC_SRS_Pos_r16_olpc_SRS_PosBasedOnPRS_Serving_r16)
		if err = ie.olpc_SRS_PosBasedOnPRS_Serving_r16.Decode(r); err != nil {
			return utils.WrapError("Decode olpc_SRS_PosBasedOnPRS_Serving_r16", err)
		}
	}
	if olpc_SRS_PosBasedOnSSB_Neigh_r16Present {
		ie.olpc_SRS_PosBasedOnSSB_Neigh_r16 = new(OLPC_SRS_Pos_r16_olpc_SRS_PosBasedOnSSB_Neigh_r16)
		if err = ie.olpc_SRS_PosBasedOnSSB_Neigh_r16.Decode(r); err != nil {
			return utils.WrapError("Decode olpc_SRS_PosBasedOnSSB_Neigh_r16", err)
		}
	}
	if olpc_SRS_PosBasedOnPRS_Neigh_r16Present {
		ie.olpc_SRS_PosBasedOnPRS_Neigh_r16 = new(OLPC_SRS_Pos_r16_olpc_SRS_PosBasedOnPRS_Neigh_r16)
		if err = ie.olpc_SRS_PosBasedOnPRS_Neigh_r16.Decode(r); err != nil {
			return utils.WrapError("Decode olpc_SRS_PosBasedOnPRS_Neigh_r16", err)
		}
	}
	if maxNumberPathLossEstimatePerServing_r16Present {
		ie.maxNumberPathLossEstimatePerServing_r16 = new(OLPC_SRS_Pos_r16_maxNumberPathLossEstimatePerServing_r16)
		if err = ie.maxNumberPathLossEstimatePerServing_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumberPathLossEstimatePerServing_r16", err)
		}
	}
	return nil
}
