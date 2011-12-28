// Copyright 2011 Xing Xing <mikespook@gmail.com> All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gearman

import (
	"bytes"
	"errors"
	//    "log"
)

// Client side job
type ClientJob struct {
	Data                []byte
	Handle, UniqueId    string
	magicCode, DataType uint32
}

// Create a new job
func NewClientJob(magiccode, datatype uint32, data []byte) (job *ClientJob) {
	return &ClientJob{magicCode: magiccode,
		DataType: datatype,
		Data:     data}
}

// Decode a job from byte slice
func DecodeClientJob(data []byte) (job *ClientJob, err error) {
	if len(data) < 12 {
		err = errors.New("Data length is too small.")
		return
	}
	datatype := byteToUint32([4]byte{data[4], data[5], data[6], data[7]})
	l := byteToUint32([4]byte{data[8], data[9], data[10], data[11]})
	if len(data[12:]) != int(l) {
		err = errors.New("Invalid data length.")
		return
	}
	data = data[12:]
	job = NewClientJob(RES, datatype, data)
	return
}

// Encode a job to byte slice
func (job *ClientJob) Encode() (data []byte) {
	magiccode := uint32ToByte(job.magicCode)
	datatype := uint32ToByte(job.DataType)
	data = make([]byte, 0, 1024*64)
	data = append(data, magiccode[:]...)
	data = append(data, datatype[:]...)
	l := len(job.Data)
	datalength := uint32ToByte(uint32(l))
	data = append(data, datalength[:]...)
	data = append(data, job.Data...)
	return
}

// Extract the job's result.
func (job *ClientJob) Result() (data []byte, err error) {
	switch job.DataType {
	case WORK_FAIL:
		job.Handle = string(job.Data)
		err = errors.New("Work fail.")
		return
	case WORK_EXCEPTION:
		err = errors.New("Work exception.")
		fallthrough
	case WORK_COMPLETE:
		s := bytes.SplitN(job.Data, []byte{'\x00'}, 2)
		if len(s) != 2 {
			err = errors.New("Invalid data.")
			return
		}
		job.Handle = string(s[0])
		data = s[1]
	default:
		err = errors.New("The job is not a result.")
	}
	return
}

// Extract the job's update
func (job *ClientJob) Update() (data []byte, err error) {
	if job.DataType != WORK_DATA && job.DataType != WORK_WARNING {
		err = errors.New("The job is not a update.")
		return
	}
	s := bytes.SplitN(job.Data, []byte{'\x00'}, 2)
	if len(s) != 2 {
		err = errors.New("Invalid data.")
		return
	}
	if job.DataType == WORK_WARNING {
		err = errors.New("Work warning")
	}
	job.Handle = string(s[0])
	data = s[1]
	return
}
