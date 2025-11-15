package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_ConfigCommon struct {
	pdsch_TimeDomainAllocationList *PDSCH_TimeDomainResourceAllocationList `optional`
}

func (ie *PDSCH_ConfigCommon) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pdsch_TimeDomainAllocationList != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pdsch_TimeDomainAllocationList != nil {
		if err = ie.pdsch_TimeDomainAllocationList.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_TimeDomainAllocationList", err)
		}
	}
	return nil
}

func (ie *PDSCH_ConfigCommon) Decode(r *uper.UperReader) error {
	var err error
	var pdsch_TimeDomainAllocationListPresent bool
	if pdsch_TimeDomainAllocationListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if pdsch_TimeDomainAllocationListPresent {
		ie.pdsch_TimeDomainAllocationList = new(PDSCH_TimeDomainResourceAllocationList)
		if err = ie.pdsch_TimeDomainAllocationList.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_TimeDomainAllocationList", err)
		}
	}
	return nil
}
