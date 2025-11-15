package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB17_IEs_r17 struct {
	trs_ResourceSetConfig_r17 []TRS_ResourceSet_r17               `lb:1,ub:maxNrofTRS_ResourceSets_r17,madatory`
	validityDuration_r17      *SIB17_IEs_r17_validityDuration_r17 `optional`
	lateNonCriticalExtension  *[]byte                             `optional`
}

func (ie *SIB17_IEs_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.validityDuration_r17 != nil, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_trs_ResourceSetConfig_r17 := utils.NewSequence[*TRS_ResourceSet_r17]([]*TRS_ResourceSet_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofTRS_ResourceSets_r17}, false)
	for _, i := range ie.trs_ResourceSetConfig_r17 {
		tmp_trs_ResourceSetConfig_r17.Value = append(tmp_trs_ResourceSetConfig_r17.Value, &i)
	}
	if err = tmp_trs_ResourceSetConfig_r17.Encode(w); err != nil {
		return utils.WrapError("Encode trs_ResourceSetConfig_r17", err)
	}
	if ie.validityDuration_r17 != nil {
		if err = ie.validityDuration_r17.Encode(w); err != nil {
			return utils.WrapError("Encode validityDuration_r17", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB17_IEs_r17) Decode(r *uper.UperReader) error {
	var err error
	var validityDuration_r17Present bool
	if validityDuration_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_trs_ResourceSetConfig_r17 := utils.NewSequence[*TRS_ResourceSet_r17]([]*TRS_ResourceSet_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofTRS_ResourceSets_r17}, false)
	fn_trs_ResourceSetConfig_r17 := func() *TRS_ResourceSet_r17 {
		return new(TRS_ResourceSet_r17)
	}
	if err = tmp_trs_ResourceSetConfig_r17.Decode(r, fn_trs_ResourceSetConfig_r17); err != nil {
		return utils.WrapError("Decode trs_ResourceSetConfig_r17", err)
	}
	ie.trs_ResourceSetConfig_r17 = []TRS_ResourceSet_r17{}
	for _, i := range tmp_trs_ResourceSetConfig_r17.Value {
		ie.trs_ResourceSetConfig_r17 = append(ie.trs_ResourceSetConfig_r17, *i)
	}
	if validityDuration_r17Present {
		ie.validityDuration_r17 = new(SIB17_IEs_r17_validityDuration_r17)
		if err = ie.validityDuration_r17.Decode(r); err != nil {
			return utils.WrapError("Decode validityDuration_r17", err)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	return nil
}
