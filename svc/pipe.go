package svc

import "io"

type PipeRW struct {
	r *io.PipeReader
	w *io.PipeWriter
}

func (p *PipeRW) Read(data []byte) (n int, err error) {
	return p.r.Read(data)
}

func (p *PipeRW) Write(data []byte) (n int, err error) {
	return p.w.Write(data)
}

func (p *PipeRW) Close() error {
	if err := p.r.Close(); err != nil {
		return err
	}
	if err := p.w.Close(); err != nil {
		return err
	}
	return nil
}

func (p *PipeRW) CloseReader() error {
	return p.r.Close()
}

func (p *PipeRW) CloseWriter() error {
	return p.w.Close()
}

func (p *PipeRW) CloseWithError(err error) error {
	if err := p.r.CloseWithError(err); err != nil {
		return err
	}
	if err := p.w.CloseWithError(err); err != nil {
		return err
	}
	return nil
}

func (p *PipeRW) CloseReaderWithError(err error) error {
	return p.r.CloseWithError(err)
}

func (p *PipeRW) CloseWriterWithError(err error) error {
	return p.w.CloseWithError(err)
}

func NewPipeRW() *PipeRW {
	r, w := io.Pipe()
	return &PipeRW{r: r, w: w}
}
