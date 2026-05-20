package outliers

import (
	"encoding/json"
	"errors"
	"io"
	"os/exec"
)

type Outliers struct {
	cmd *exec.Cmd
	w   io.WriteCloser
	r   io.ReadCloser
	enc *json.Encoder
	dec *json.Decoder
}

func New() (*Outliers, error) {
	// -u = unbuffered output
	cmd := exec.Command("python", "-u", "outliers.py")
	w, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	r, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	o := Outliers{
		cmd: cmd,
		w:   w,
		enc: json.NewEncoder(w),
		r:   r,
		dec: json.NewDecoder(r),
	}

	return &o, nil
}

func (o *Outliers) Close() error {
	return errors.Join(o.w.Close(), o.r.Close(), o.cmd.Process.Kill())
}

func (o *Outliers) Detect(data []float64) ([]int, error) {
	if err := o.enc.Encode(data); err != nil {
		return nil, err
	}

	var out []int
	if err := o.dec.Decode(&out); err != nil {
		return nil, err
	}

	return out, nil
}
