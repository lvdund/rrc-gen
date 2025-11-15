package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_AperiodicTriggerState struct {
	associatedReportConfigInfoList []CSI_AssociatedReportConfigInfo                       `lb:1,ub:maxNrofReportConfigPerAperiodicTrigger,madatory`
	ap_CSI_MultiplexingMode_r17    *CSI_AperiodicTriggerState_ap_CSI_MultiplexingMode_r17 `optional,ext-1`
}

func (ie *CSI_AperiodicTriggerState) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.ap_CSI_MultiplexingMode_r17 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_associatedReportConfigInfoList := utils.NewSequence[*CSI_AssociatedReportConfigInfo]([]*CSI_AssociatedReportConfigInfo{}, uper.Constraint{Lb: 1, Ub: maxNrofReportConfigPerAperiodicTrigger}, false)
	for _, i := range ie.associatedReportConfigInfoList {
		tmp_associatedReportConfigInfoList.Value = append(tmp_associatedReportConfigInfoList.Value, &i)
	}
	if err = tmp_associatedReportConfigInfoList.Encode(w); err != nil {
		return utils.WrapError("Encode associatedReportConfigInfoList", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.ap_CSI_MultiplexingMode_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CSI_AperiodicTriggerState", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.ap_CSI_MultiplexingMode_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ap_CSI_MultiplexingMode_r17 optional
			if ie.ap_CSI_MultiplexingMode_r17 != nil {
				if err = ie.ap_CSI_MultiplexingMode_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ap_CSI_MultiplexingMode_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *CSI_AperiodicTriggerState) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_associatedReportConfigInfoList := utils.NewSequence[*CSI_AssociatedReportConfigInfo]([]*CSI_AssociatedReportConfigInfo{}, uper.Constraint{Lb: 1, Ub: maxNrofReportConfigPerAperiodicTrigger}, false)
	fn_associatedReportConfigInfoList := func() *CSI_AssociatedReportConfigInfo {
		return new(CSI_AssociatedReportConfigInfo)
	}
	if err = tmp_associatedReportConfigInfoList.Decode(r, fn_associatedReportConfigInfoList); err != nil {
		return utils.WrapError("Decode associatedReportConfigInfoList", err)
	}
	ie.associatedReportConfigInfoList = []CSI_AssociatedReportConfigInfo{}
	for _, i := range tmp_associatedReportConfigInfoList.Value {
		ie.associatedReportConfigInfoList = append(ie.associatedReportConfigInfoList, *i)
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ap_CSI_MultiplexingMode_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ap_CSI_MultiplexingMode_r17 optional
			if ap_CSI_MultiplexingMode_r17Present {
				ie.ap_CSI_MultiplexingMode_r17 = new(CSI_AperiodicTriggerState_ap_CSI_MultiplexingMode_r17)
				if err = ie.ap_CSI_MultiplexingMode_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ap_CSI_MultiplexingMode_r17", err)
				}
			}
		}
	}
	return nil
}
