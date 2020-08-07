package settings

var Graphics = initGraphics()

func initGraphics() *graphics {
	return &graphics{
		Width:        1920,
		Height:       1080,
		WindowWidth:  1280,
		WindowHeight: 720,
		Fullscreen:   false,
		VSync:        false,
		FPSCap:       1000,
		MSAA:         16,
	}
}

type graphics struct {
	Width        int64
	Height       int64
	WindowWidth  int64
	WindowHeight int64
	Fullscreen   bool  //true
	VSync        bool  //false
	FPSCap       int64 //1000
	MSAA         int32 //16
}

func (gr graphics) GetSize() (int64, int64) {
	if gr.Fullscreen {
		return gr.Width, gr.Height
	}
	return gr.WindowWidth, gr.WindowHeight
}

func (gr graphics) GetSizeF() (float64, float64) {
	if gr.Fullscreen {
		return float64(gr.Width), float64(gr.Height)
	}
	return float64(gr.WindowWidth), float64(gr.WindowHeight)
}

func (gr graphics) GetWidth() int64 {
	if gr.Fullscreen {
		return gr.Width
	}
	return gr.WindowWidth
}

func (gr graphics) GetWidthF() float64 {
	if gr.Fullscreen {
		return float64(gr.Width)
	}
	return float64(gr.WindowWidth)
}

func (gr graphics) GetHeight() int64 {
	if gr.Fullscreen {
		return gr.Height
	}
	return gr.WindowHeight
}

func (gr graphics) GetHeightF() float64 {
	if gr.Fullscreen {
		return float64(gr.Height)
	}
	return float64(gr.WindowHeight)
}

func (gr graphics) GetAspectRatio() float64 {
	if gr.Fullscreen {
		return float64(gr.Width) / float64(gr.Height)
	}
	return float64(gr.WindowWidth) / float64(gr.WindowHeight)
}
