package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NZP_CSI_RS_ResourceSet struct {
	nzp_CSI_ResourceSetId           NZP_CSI_RS_ResourceSetId             `madatory`
	nzp_CSI_RS_Resources            []NZP_CSI_RS_ResourceId              `lb:1,ub:maxNrofNZP_CSI_RS_ResourcesPerSet,madatory`
	repetition                      *NZP_CSI_RS_ResourceSet_repetition   `optional`
	aperiodicTriggeringOffset       *int64                               `lb:0,ub:6,optional`
	trs_Info                        *NZP_CSI_RS_ResourceSet_trs_Info     `optional`
	aperiodicTriggeringOffset_r16   *int64                               `lb:0,ub:31,optional,ext-1`
	pdc_Info_r17                    *NZP_CSI_RS_ResourceSet_pdc_Info_r17 `optional,ext-2`
	cmrGroupingAndPairing_r17       *CMRGroupingAndPairing_r17           `optional,ext-2`
	aperiodicTriggeringOffset_r17   *int64                               `lb:0,ub:124,optional,ext-2`
	aperiodicTriggeringOffsetL2_r17 *int64                               `lb:0,ub:31,optional,ext-2`
}

func (ie *NZP_CSI_RS_ResourceSet) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.aperiodicTriggeringOffset_r16 != nil || ie.pdc_Info_r17 != nil || ie.cmrGroupingAndPairing_r17 != nil || ie.aperiodicTriggeringOffset_r17 != nil || ie.aperiodicTriggeringOffsetL2_r17 != nil
	preambleBits := []bool{hasExtensions, ie.repetition != nil, ie.aperiodicTriggeringOffset != nil, ie.trs_Info != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.nzp_CSI_ResourceSetId.Encode(w); err != nil {
		return utils.WrapError("Encode nzp_CSI_ResourceSetId", err)
	}
	tmp_nzp_CSI_RS_Resources := utils.NewSequence[*NZP_CSI_RS_ResourceId]([]*NZP_CSI_RS_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourcesPerSet}, false)
	for _, i := range ie.nzp_CSI_RS_Resources {
		tmp_nzp_CSI_RS_Resources.Value = append(tmp_nzp_CSI_RS_Resources.Value, &i)
	}
	if err = tmp_nzp_CSI_RS_Resources.Encode(w); err != nil {
		return utils.WrapError("Encode nzp_CSI_RS_Resources", err)
	}
	if ie.repetition != nil {
		if err = ie.repetition.Encode(w); err != nil {
			return utils.WrapError("Encode repetition", err)
		}
	}
	if ie.aperiodicTriggeringOffset != nil {
		if err = w.WriteInteger(*ie.aperiodicTriggeringOffset, &uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
			return utils.WrapError("Encode aperiodicTriggeringOffset", err)
		}
	}
	if ie.trs_Info != nil {
		if err = ie.trs_Info.Encode(w); err != nil {
			return utils.WrapError("Encode trs_Info", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.aperiodicTriggeringOffset_r16 != nil, ie.pdc_Info_r17 != nil || ie.cmrGroupingAndPairing_r17 != nil || ie.aperiodicTriggeringOffset_r17 != nil || ie.aperiodicTriggeringOffsetL2_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap NZP_CSI_RS_ResourceSet", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.aperiodicTriggeringOffset_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode aperiodicTriggeringOffset_r16 optional
			if ie.aperiodicTriggeringOffset_r16 != nil {
				if err = extWriter.WriteInteger(*ie.aperiodicTriggeringOffset_r16, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
					return utils.WrapError("Encode aperiodicTriggeringOffset_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.pdc_Info_r17 != nil, ie.cmrGroupingAndPairing_r17 != nil, ie.aperiodicTriggeringOffset_r17 != nil, ie.aperiodicTriggeringOffsetL2_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode pdc_Info_r17 optional
			if ie.pdc_Info_r17 != nil {
				if err = ie.pdc_Info_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdc_Info_r17", err)
				}
			}
			// encode cmrGroupingAndPairing_r17 optional
			if ie.cmrGroupingAndPairing_r17 != nil {
				if err = ie.cmrGroupingAndPairing_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cmrGroupingAndPairing_r17", err)
				}
			}
			// encode aperiodicTriggeringOffset_r17 optional
			if ie.aperiodicTriggeringOffset_r17 != nil {
				if err = extWriter.WriteInteger(*ie.aperiodicTriggeringOffset_r17, &uper.Constraint{Lb: 0, Ub: 124}, false); err != nil {
					return utils.WrapError("Encode aperiodicTriggeringOffset_r17", err)
				}
			}
			// encode aperiodicTriggeringOffsetL2_r17 optional
			if ie.aperiodicTriggeringOffsetL2_r17 != nil {
				if err = extWriter.WriteInteger(*ie.aperiodicTriggeringOffsetL2_r17, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
					return utils.WrapError("Encode aperiodicTriggeringOffsetL2_r17", err)
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

func (ie *NZP_CSI_RS_ResourceSet) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var repetitionPresent bool
	if repetitionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var aperiodicTriggeringOffsetPresent bool
	if aperiodicTriggeringOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var trs_InfoPresent bool
	if trs_InfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.nzp_CSI_ResourceSetId.Decode(r); err != nil {
		return utils.WrapError("Decode nzp_CSI_ResourceSetId", err)
	}
	tmp_nzp_CSI_RS_Resources := utils.NewSequence[*NZP_CSI_RS_ResourceId]([]*NZP_CSI_RS_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourcesPerSet}, false)
	fn_nzp_CSI_RS_Resources := func() *NZP_CSI_RS_ResourceId {
		return new(NZP_CSI_RS_ResourceId)
	}
	if err = tmp_nzp_CSI_RS_Resources.Decode(r, fn_nzp_CSI_RS_Resources); err != nil {
		return utils.WrapError("Decode nzp_CSI_RS_Resources", err)
	}
	ie.nzp_CSI_RS_Resources = []NZP_CSI_RS_ResourceId{}
	for _, i := range tmp_nzp_CSI_RS_Resources.Value {
		ie.nzp_CSI_RS_Resources = append(ie.nzp_CSI_RS_Resources, *i)
	}
	if repetitionPresent {
		ie.repetition = new(NZP_CSI_RS_ResourceSet_repetition)
		if err = ie.repetition.Decode(r); err != nil {
			return utils.WrapError("Decode repetition", err)
		}
	}
	if aperiodicTriggeringOffsetPresent {
		var tmp_int_aperiodicTriggeringOffset int64
		if tmp_int_aperiodicTriggeringOffset, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode aperiodicTriggeringOffset", err)
		}
		ie.aperiodicTriggeringOffset = &tmp_int_aperiodicTriggeringOffset
	}
	if trs_InfoPresent {
		ie.trs_Info = new(NZP_CSI_RS_ResourceSet_trs_Info)
		if err = ie.trs_Info.Decode(r); err != nil {
			return utils.WrapError("Decode trs_Info", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			aperiodicTriggeringOffset_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode aperiodicTriggeringOffset_r16 optional
			if aperiodicTriggeringOffset_r16Present {
				var tmp_int_aperiodicTriggeringOffset_r16 int64
				if tmp_int_aperiodicTriggeringOffset_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
					return utils.WrapError("Decode aperiodicTriggeringOffset_r16", err)
				}
				ie.aperiodicTriggeringOffset_r16 = &tmp_int_aperiodicTriggeringOffset_r16
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			pdc_Info_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cmrGroupingAndPairing_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			aperiodicTriggeringOffset_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			aperiodicTriggeringOffsetL2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode pdc_Info_r17 optional
			if pdc_Info_r17Present {
				ie.pdc_Info_r17 = new(NZP_CSI_RS_ResourceSet_pdc_Info_r17)
				if err = ie.pdc_Info_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdc_Info_r17", err)
				}
			}
			// decode cmrGroupingAndPairing_r17 optional
			if cmrGroupingAndPairing_r17Present {
				ie.cmrGroupingAndPairing_r17 = new(CMRGroupingAndPairing_r17)
				if err = ie.cmrGroupingAndPairing_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode cmrGroupingAndPairing_r17", err)
				}
			}
			// decode aperiodicTriggeringOffset_r17 optional
			if aperiodicTriggeringOffset_r17Present {
				var tmp_int_aperiodicTriggeringOffset_r17 int64
				if tmp_int_aperiodicTriggeringOffset_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 124}, false); err != nil {
					return utils.WrapError("Decode aperiodicTriggeringOffset_r17", err)
				}
				ie.aperiodicTriggeringOffset_r17 = &tmp_int_aperiodicTriggeringOffset_r17
			}
			// decode aperiodicTriggeringOffsetL2_r17 optional
			if aperiodicTriggeringOffsetL2_r17Present {
				var tmp_int_aperiodicTriggeringOffsetL2_r17 int64
				if tmp_int_aperiodicTriggeringOffsetL2_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
					return utils.WrapError("Decode aperiodicTriggeringOffsetL2_r17", err)
				}
				ie.aperiodicTriggeringOffsetL2_r17 = &tmp_int_aperiodicTriggeringOffsetL2_r17
			}
		}
	}
	return nil
}
