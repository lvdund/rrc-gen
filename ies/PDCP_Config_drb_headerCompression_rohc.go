package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_Config_drb_headerCompression_rohc struct {
	maxCID           int64                                                    `lb:1,ub:16383,madatory`
	profiles         PDCP_Config_drb_headerCompression_rohc_profiles          `madatory`
	drb_ContinueROHC *PDCP_Config_drb_headerCompression_rohc_drb_ContinueROHC `optional`
}

func (ie *PDCP_Config_drb_headerCompression_rohc) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.drb_ContinueROHC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.maxCID, &uper.Constraint{Lb: 1, Ub: 16383}, false); err != nil {
		return utils.WrapError("WriteInteger maxCID", err)
	}
	if err = ie.profiles.Encode(w); err != nil {
		return utils.WrapError("Encode profiles", err)
	}
	if ie.drb_ContinueROHC != nil {
		if err = ie.drb_ContinueROHC.Encode(w); err != nil {
			return utils.WrapError("Encode drb_ContinueROHC", err)
		}
	}
	return nil
}

func (ie *PDCP_Config_drb_headerCompression_rohc) Decode(r *uper.UperReader) error {
	var err error
	var drb_ContinueROHCPresent bool
	if drb_ContinueROHCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_maxCID int64
	if tmp_int_maxCID, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16383}, false); err != nil {
		return utils.WrapError("ReadInteger maxCID", err)
	}
	ie.maxCID = tmp_int_maxCID
	if err = ie.profiles.Decode(r); err != nil {
		return utils.WrapError("Decode profiles", err)
	}
	if drb_ContinueROHCPresent {
		ie.drb_ContinueROHC = new(PDCP_Config_drb_headerCompression_rohc_drb_ContinueROHC)
		if err = ie.drb_ContinueROHC.Decode(r); err != nil {
			return utils.WrapError("Decode drb_ContinueROHC", err)
		}
	}
	return nil
}
