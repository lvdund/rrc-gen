package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReestablishmentInfo struct {
	sourcePhysCellId          PhysCellId            `madatory`
	targetCellShortMAC_I      ShortMAC_I            `madatory`
	additionalReestabInfoList *ReestabNCellInfoList `optional`
}

func (ie *ReestablishmentInfo) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.additionalReestabInfoList != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.sourcePhysCellId.Encode(w); err != nil {
		return utils.WrapError("Encode sourcePhysCellId", err)
	}
	if err = ie.targetCellShortMAC_I.Encode(w); err != nil {
		return utils.WrapError("Encode targetCellShortMAC_I", err)
	}
	if ie.additionalReestabInfoList != nil {
		if err = ie.additionalReestabInfoList.Encode(w); err != nil {
			return utils.WrapError("Encode additionalReestabInfoList", err)
		}
	}
	return nil
}

func (ie *ReestablishmentInfo) Decode(r *uper.UperReader) error {
	var err error
	var additionalReestabInfoListPresent bool
	if additionalReestabInfoListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.sourcePhysCellId.Decode(r); err != nil {
		return utils.WrapError("Decode sourcePhysCellId", err)
	}
	if err = ie.targetCellShortMAC_I.Decode(r); err != nil {
		return utils.WrapError("Decode targetCellShortMAC_I", err)
	}
	if additionalReestabInfoListPresent {
		ie.additionalReestabInfoList = new(ReestabNCellInfoList)
		if err = ie.additionalReestabInfoList.Decode(r); err != nil {
			return utils.WrapError("Decode additionalReestabInfoList", err)
		}
	}
	return nil
}
