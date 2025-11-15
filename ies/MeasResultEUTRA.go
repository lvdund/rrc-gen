package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultEUTRA struct {
	eutra_PhysCellId PhysCellId               `madatory`
	measResult       MeasQuantityResultsEUTRA `madatory`
	cgi_Info         *CGI_InfoEUTRA           `optional`
}

func (ie *MeasResultEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.cgi_Info != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.eutra_PhysCellId.Encode(w); err != nil {
		return utils.WrapError("Encode eutra_PhysCellId", err)
	}
	if err = ie.measResult.Encode(w); err != nil {
		return utils.WrapError("Encode measResult", err)
	}
	if ie.cgi_Info != nil {
		if err = ie.cgi_Info.Encode(w); err != nil {
			return utils.WrapError("Encode cgi_Info", err)
		}
	}
	return nil
}

func (ie *MeasResultEUTRA) Decode(r *uper.UperReader) error {
	var err error
	var cgi_InfoPresent bool
	if cgi_InfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.eutra_PhysCellId.Decode(r); err != nil {
		return utils.WrapError("Decode eutra_PhysCellId", err)
	}
	if err = ie.measResult.Decode(r); err != nil {
		return utils.WrapError("Decode measResult", err)
	}
	if cgi_InfoPresent {
		ie.cgi_Info = new(CGI_InfoEUTRA)
		if err = ie.cgi_Info.Decode(r); err != nil {
			return utils.WrapError("Decode cgi_Info", err)
		}
	}
	return nil
}
