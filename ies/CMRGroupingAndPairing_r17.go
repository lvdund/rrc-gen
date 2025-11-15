package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CMRGroupingAndPairing_r17 struct {
	nrofResourcesGroup1_r17 int64                   `lb:1,ub:7,madatory`
	pair1OfNZP_CSI_RS_r17   *NZP_CSI_RS_Pairing_r17 `optional`
	pair2OfNZP_CSI_RS_r17   *NZP_CSI_RS_Pairing_r17 `optional`
}

func (ie *CMRGroupingAndPairing_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pair1OfNZP_CSI_RS_r17 != nil, ie.pair2OfNZP_CSI_RS_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.nrofResourcesGroup1_r17, &uper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteInteger nrofResourcesGroup1_r17", err)
	}
	if ie.pair1OfNZP_CSI_RS_r17 != nil {
		if err = ie.pair1OfNZP_CSI_RS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pair1OfNZP_CSI_RS_r17", err)
		}
	}
	if ie.pair2OfNZP_CSI_RS_r17 != nil {
		if err = ie.pair2OfNZP_CSI_RS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pair2OfNZP_CSI_RS_r17", err)
		}
	}
	return nil
}

func (ie *CMRGroupingAndPairing_r17) Decode(r *uper.UperReader) error {
	var err error
	var pair1OfNZP_CSI_RS_r17Present bool
	if pair1OfNZP_CSI_RS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pair2OfNZP_CSI_RS_r17Present bool
	if pair2OfNZP_CSI_RS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_nrofResourcesGroup1_r17 int64
	if tmp_int_nrofResourcesGroup1_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadInteger nrofResourcesGroup1_r17", err)
	}
	ie.nrofResourcesGroup1_r17 = tmp_int_nrofResourcesGroup1_r17
	if pair1OfNZP_CSI_RS_r17Present {
		ie.pair1OfNZP_CSI_RS_r17 = new(NZP_CSI_RS_Pairing_r17)
		if err = ie.pair1OfNZP_CSI_RS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pair1OfNZP_CSI_RS_r17", err)
		}
	}
	if pair2OfNZP_CSI_RS_r17Present {
		ie.pair2OfNZP_CSI_RS_r17 = new(NZP_CSI_RS_Pairing_r17)
		if err = ie.pair2OfNZP_CSI_RS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pair2OfNZP_CSI_RS_r17", err)
		}
	}
	return nil
}
