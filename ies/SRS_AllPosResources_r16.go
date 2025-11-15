package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_AllPosResources_r16 struct {
	srs_PosResources_r16  SRS_PosResources_r16   `madatory`
	srs_PosResourceAP_r16 *SRS_PosResourceAP_r16 `optional`
	srs_PosResourceSP_r16 *SRS_PosResourceSP_r16 `optional`
}

func (ie *SRS_AllPosResources_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.srs_PosResourceAP_r16 != nil, ie.srs_PosResourceSP_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.srs_PosResources_r16.Encode(w); err != nil {
		return utils.WrapError("Encode srs_PosResources_r16", err)
	}
	if ie.srs_PosResourceAP_r16 != nil {
		if err = ie.srs_PosResourceAP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode srs_PosResourceAP_r16", err)
		}
	}
	if ie.srs_PosResourceSP_r16 != nil {
		if err = ie.srs_PosResourceSP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode srs_PosResourceSP_r16", err)
		}
	}
	return nil
}

func (ie *SRS_AllPosResources_r16) Decode(r *uper.UperReader) error {
	var err error
	var srs_PosResourceAP_r16Present bool
	if srs_PosResourceAP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var srs_PosResourceSP_r16Present bool
	if srs_PosResourceSP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.srs_PosResources_r16.Decode(r); err != nil {
		return utils.WrapError("Decode srs_PosResources_r16", err)
	}
	if srs_PosResourceAP_r16Present {
		ie.srs_PosResourceAP_r16 = new(SRS_PosResourceAP_r16)
		if err = ie.srs_PosResourceAP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode srs_PosResourceAP_r16", err)
		}
	}
	if srs_PosResourceSP_r16Present {
		ie.srs_PosResourceSP_r16 = new(SRS_PosResourceSP_r16)
		if err = ie.srs_PosResourceSP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode srs_PosResourceSP_r16", err)
		}
	}
	return nil
}
